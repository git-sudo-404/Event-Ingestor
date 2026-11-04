package com.gitsudo404.event_ingestion.controller;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

import com.gitsudo404.event_ingestion.dto.EventResponseDTO;
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
    when(service.ingest(request)).thenReturn(delivery);

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
    when(service.ingest(request)).thenReturn(CompletableFuture.failedFuture(new RuntimeException("Kafka unavailable")));

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

  private VehicleEventController controller(VehicleEventService service, boolean allowed) {
    RateLimiter limiter = key -> allowed;
    return new VehicleEventController(service, limiter, limiter, limiter, limiter, limiter);
  }

  private VehicleGPSEventRequestDTO request() {
    VehicleGPSEventRequestDTO request = new VehicleGPSEventRequestDTO();
    request.setEventId("event-123");
    request.setEntityId("vehicle-123");
    return request;
  }
}
