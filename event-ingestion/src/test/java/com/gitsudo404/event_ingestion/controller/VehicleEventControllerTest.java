package com.gitsudo404.event_ingestion.controller;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

import com.gitsudo404.event_ingestion.dto.EventResponseDTO;
import com.gitsudo404.event_ingestion.dto.VehicleBatteryEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleDiagnosticEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleEngineTelemetryEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleFuelEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleGPSEventRequestDTO;
import com.gitsudo404.event_ingestion.middleware.rate_limiter.RateLimiter;
import com.gitsudo404.event_ingestion.service.VehicleEventService;
import java.util.concurrent.CompletableFuture;
import org.junit.jupiter.api.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

class VehicleEventControllerTest {

  @Test
  void respondsOnlyAfterKafkaAcknowledgesTheEvent() {
    VehicleEventService service = mock(VehicleEventService.class);
    CompletableFuture<Void> delivery = new CompletableFuture<>();
    VehicleGPSEventRequestDTO request = request();
    when(service.ingestGPSEvent(request)).thenReturn(delivery);

    CompletableFuture<ResponseEntity<EventResponseDTO>> response = controller(service, true)
        .handleVehicleGPSEvent(request);

    assertFalse(response.isDone());
    delivery.complete(null);
    assertEquals(HttpStatus.OK, response.join().getStatusCode());
    assertEquals(request.getEventId(), response.join().getBody().GetEventId());
  }

  @Test
  void returnsServiceUnavailableWhenDeliveryFails() {
    VehicleEventService service = mock(VehicleEventService.class);
    VehicleGPSEventRequestDTO request = request();
    when(service.ingestGPSEvent(request)).thenReturn(CompletableFuture.failedFuture(new RuntimeException("Kafka unavailable")));

    ResponseEntity<EventResponseDTO> response = controller(service, true)
        .handleVehicleGPSEvent(request).join();

    assertEquals(HttpStatus.SERVICE_UNAVAILABLE, response.getStatusCode());
    assertEquals(request.getEventId(), response.getBody().GetEventId());
  }

  @Test
  void returnsHttp429WhenRateLimited() {
    VehicleEventService service = mock(VehicleEventService.class);

    ResponseEntity<EventResponseDTO> response = controller(service, false)
        .handleVehicleGPSEvent(request()).join();

    assertEquals(HttpStatus.TOO_MANY_REQUESTS, response.getStatusCode());
  }

  @Test
  void routesOtherVehicleEventsToTheirServices() {
    VehicleEventService service = mock(VehicleEventService.class);
    RateLimiter gps = mock(RateLimiter.class);
    RateLimiter battery = mock(RateLimiter.class);
    RateLimiter fuel = mock(RateLimiter.class);
    RateLimiter diagnostic = mock(RateLimiter.class);
    RateLimiter engineTelemetry = mock(RateLimiter.class);
    when(battery.allowRequest("vehicle-123")).thenReturn(true);
    when(fuel.allowRequest("vehicle-123")).thenReturn(true);
    when(diagnostic.allowRequest("vehicle-123")).thenReturn(true);
    when(engineTelemetry.allowRequest("vehicle-123")).thenReturn(true);

    VehicleEventController controller = new VehicleEventController(
        service, gps, battery, fuel, diagnostic, engineTelemetry);
    VehicleBatteryEventRequestDTO batteryRequest = new VehicleBatteryEventRequestDTO();
    batteryRequest.setEventId("event-123");
    batteryRequest.setEntityId("vehicle-123");
    batteryRequest.setEventType("vehicle.battery");
    VehicleFuelEventRequestDTO fuelRequest = new VehicleFuelEventRequestDTO();
    fuelRequest.setEventId("event-123");
    fuelRequest.setEntityId("vehicle-123");
    fuelRequest.setEventType("vehicle.fuel");
    VehicleDiagnosticEventRequestDTO diagnosticRequest = new VehicleDiagnosticEventRequestDTO();
    diagnosticRequest.setEventId("event-123");
    diagnosticRequest.setEntityId("vehicle-123");
    diagnosticRequest.setEventType("vehicle.diagnostic");
    VehicleEngineTelemetryEventRequestDTO engineRequest = new VehicleEngineTelemetryEventRequestDTO();
    engineRequest.setEventId("event-123");
    engineRequest.setEntityId("vehicle-123");
    engineRequest.setEventType("vehicle.engine_telemetry");
    when(service.ingestBatteryEvent(batteryRequest)).thenReturn(CompletableFuture.completedFuture(null));
    when(service.ingestFuelEvent(fuelRequest)).thenReturn(CompletableFuture.completedFuture(null));
    when(service.ingestDiagnosticEvent(diagnosticRequest)).thenReturn(CompletableFuture.completedFuture(null));
    when(service.ingestEngineTelemetryEvent(engineRequest)).thenReturn(CompletableFuture.completedFuture(null));

    assertEquals(HttpStatus.OK, controller.handleVehicleBatteryEvent(batteryRequest).join().getStatusCode());
    assertEquals(HttpStatus.OK, controller.handleVehicleFuelEvent(fuelRequest).join().getStatusCode());
    assertEquals(HttpStatus.OK, controller.handleVehicleDiagnosticEvent(diagnosticRequest).join().getStatusCode());
    assertEquals(HttpStatus.OK, controller.handleVehicleEngineTelemetryEvent(engineRequest).join().getStatusCode());
    verify(service).ingestBatteryEvent(batteryRequest);
    verify(service).ingestFuelEvent(fuelRequest);
    verify(service).ingestDiagnosticEvent(diagnosticRequest);
    verify(service).ingestEngineTelemetryEvent(engineRequest);
  }

  @Test
  void rejectsEventTypeThatDoesNotMatchTheEndpoint() {
    VehicleEventService service = mock(VehicleEventService.class);
    VehicleGPSEventRequestDTO request = request();
    request.setEventType("vehicle.fuel");

    ResponseEntity<EventResponseDTO> response = controller(service, true)
        .handleVehicleGPSEvent(request).join();

    assertEquals(HttpStatus.BAD_REQUEST, response.getStatusCode());
  }

  private VehicleEventController controller(VehicleEventService service, boolean allowed) {
    RateLimiter limiter = key -> allowed;
    return new VehicleEventController(service, limiter, limiter, limiter, limiter, limiter);
  }

  private VehicleGPSEventRequestDTO request() {
    VehicleGPSEventRequestDTO request = new VehicleGPSEventRequestDTO();
    request.setEventId("event-123");
    request.setEntityId("vehicle-123");
    request.setEventType("vehicle.location");
    return request;
  }
}
