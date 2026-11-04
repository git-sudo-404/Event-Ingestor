package com.gitsudo404.event_ingestion.service;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

import com.gitsudo404.event_ingestion.dto.EventMetaData;
import com.gitsudo404.event_ingestion.dto.VehicleBatteryEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleDiagnosticEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleEngineTelemetryEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleFuelEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleGPSEventRequestDTO;
import com.gitsudo404.event_ingestion.producer.EventPublisher;
import com.gitsudo404.event_ingestion.producer.KafkaTopic;
import java.util.concurrent.CompletableFuture;
import org.junit.jupiter.api.Test;
import tools.jackson.databind.ObjectMapper;

class VehicleEventServiceTest {

  @Test
  void usesVehicleIdAsKafkaKey() {
    EventPublisher publisher = mock(EventPublisher.class);
    ObjectMapper mapper = mock(ObjectMapper.class);
    VehicleGPSEventRequestDTO request = new VehicleGPSEventRequestDTO();
    request.setEntityId("vehicle-123");
    String message = "{\"event_id\":\"event-123\"}";
    when(mapper.writeValueAsString(request)).thenReturn(message);
    when(publisher.publish(KafkaTopic.VEHICLE_GPS_EVENT, "vehicle-123", message))
        .thenReturn(CompletableFuture.completedFuture(null));

    new VehicleEventService(publisher, mapper).ingestGPSEvent(request).join();

    verify(publisher).publish(KafkaTopic.VEHICLE_GPS_EVENT, "vehicle-123", message);
  }

  @Test
  void publishesEachVehicleEventToItsOwnTopicWithTheVehicleKey() {
    EventPublisher publisher = mock(EventPublisher.class);
    ObjectMapper mapper = mock(ObjectMapper.class);
    String message = "{\"entity_id\":\"vehicle-123\"}";
    when(mapper.writeValueAsString(any(EventMetaData.class))).thenReturn(message);
    when(publisher.publish(any(KafkaTopic.class), eq("vehicle-123"), eq(message)))
        .thenReturn(CompletableFuture.completedFuture(null));
    VehicleEventService service = new VehicleEventService(publisher, mapper);

    VehicleBatteryEventRequestDTO battery = new VehicleBatteryEventRequestDTO();
    battery.setEntityId("vehicle-123");
    VehicleFuelEventRequestDTO fuel = new VehicleFuelEventRequestDTO();
    fuel.setEntityId("vehicle-123");
    VehicleDiagnosticEventRequestDTO diagnostic = new VehicleDiagnosticEventRequestDTO();
    diagnostic.setEntityId("vehicle-123");
    VehicleEngineTelemetryEventRequestDTO engine = new VehicleEngineTelemetryEventRequestDTO();
    engine.setEntityId("vehicle-123");

    service.ingestBatteryEvent(battery).join();
    service.ingestFuelEvent(fuel).join();
    service.ingestDiagnosticEvent(diagnostic).join();
    service.ingestEngineTelemetryEvent(engine).join();

    verify(publisher).publish(KafkaTopic.VEHICLE_BATTERY_EVENT, "vehicle-123", message);
    verify(publisher).publish(KafkaTopic.VEHICLE_FUEL_EVENT, "vehicle-123", message);
    verify(publisher).publish(KafkaTopic.VEHICLE_DIAGNOSTIC_EVENT, "vehicle-123", message);
    verify(publisher).publish(KafkaTopic.VEHICLE_ENGINE_TELEMETRY_EVENT, "vehicle-123", message);
  }
}
