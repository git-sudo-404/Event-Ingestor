package com.gitsudo404.event_ingestion.controller;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.asyncDispatch;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.request;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import com.gitsudo404.event_ingestion.dto.VehicleEngineTelemetryEventRequestDTO;
import com.gitsudo404.event_ingestion.middleware.rate_limiter.RateLimiter;
import com.gitsudo404.event_ingestion.service.VehicleEventService;
import java.util.concurrent.CompletableFuture;
import org.junit.jupiter.api.Test;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.validation.beanvalidation.LocalValidatorFactoryBean;

class VehicleEventHttpTest {

  @Test
  void acceptsNumericEngineTelemetryAtItsHttpEndpoint() throws Exception {
    VehicleEventService service = mock(VehicleEventService.class);
    when(service.ingestEngineTelemetryEvent(any(VehicleEngineTelemetryEventRequestDTO.class)))
        .thenReturn(CompletableFuture.completedFuture(null));
    MockMvc mvc = mvc(service);

    MvcResult pending = mvc.perform(post("/api/v1/events/vehicle/engine-telemetry")
        .contentType(MediaType.APPLICATION_JSON)
        .content(engineTelemetryJson(true)))
        .andExpect(request().asyncStarted())
        .andReturn();

    mvc.perform(asyncDispatch(pending)).andExpect(status().isOk());
  }

  @Test
  void rejectsEngineTelemetryWithoutRpm() throws Exception {
    MockMvc mvc = mvc(mock(VehicleEventService.class));

    mvc.perform(post("/api/v1/events/vehicle/engine-telemetry")
        .contentType(MediaType.APPLICATION_JSON)
        .content(engineTelemetryJson(false)))
        .andExpect(status().isBadRequest());
  }

  private MockMvc mvc(VehicleEventService service) {
    RateLimiter allowed = key -> true;
    VehicleEventController controller = new VehicleEventController(
        service, allowed, allowed, allowed, allowed, allowed);
    LocalValidatorFactoryBean validator = new LocalValidatorFactoryBean();
    validator.afterPropertiesSet();
    return MockMvcBuilders.standaloneSetup(controller).setValidator(validator).build();
  }

  private String engineTelemetryJson(boolean includeRpm) {
    String rpm = includeRpm ? "\"rpm\":1200.0," : "";
    return "{" + rpm
        + "\"event_id\":\"event-123\","
        + "\"event_type\":\"vehicle.engine_telemetry\","
        + "\"schema_version\":1,"
        + "\"source\":\"vehicle-simulator\","
        + "\"timestamp\":\"2026-11-05T00:00:00Z\","
        + "\"entity_id\":\"vehicle-123\","
        + "\"coolant_temp_c\":85.0,"
        + "\"engine_load_pct\":35.0}";
  }
}
