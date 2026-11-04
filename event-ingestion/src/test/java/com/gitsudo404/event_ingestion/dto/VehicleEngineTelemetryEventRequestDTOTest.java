package com.gitsudo404.event_ingestion.dto;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import jakarta.validation.Validation;
import jakarta.validation.Validator;
import java.time.Instant;
import org.junit.jupiter.api.Test;

class VehicleEngineTelemetryEventRequestDTOTest {

  @Test
  void acceptsNumericTelemetryAndRequiresRpm() {
    Validator validator = Validation.buildDefaultValidatorFactory().getValidator();
    VehicleEngineTelemetryEventRequestDTO request = new VehicleEngineTelemetryEventRequestDTO();
    request.setEventId("event-123");
    request.setEventType("vehicle.engine_telemetry");
    request.setSchemaVersion(1);
    request.setSource("vehicle-simulator");
    request.setTimestamp(Instant.now());
    request.setEntityId("vehicle-123");
    request.setRpm(1200.0f);
    request.setCoolantTempC(85.0f);
    request.setEngineLoadPct(35.0f);

    assertTrue(validator.validate(request).isEmpty());

    request.setRpm(null);
    assertEquals("rpm", validator.validate(request).iterator().next().getPropertyPath().toString());
  }
}
