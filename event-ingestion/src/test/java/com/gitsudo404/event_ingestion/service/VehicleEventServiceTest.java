package com.gitsudo404.event_ingestion.service;

import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

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

    new VehicleEventService(publisher, mapper).ingest(request).join();

    verify(publisher).publish(KafkaTopic.VEHICLE_GPS_EVENT, "vehicle-123", message);
  }
}
