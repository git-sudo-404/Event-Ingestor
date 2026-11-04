/*
 * MIT License
 *
 * Copyright (c) 2026 git-sudo-404
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * Of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * Copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all Copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING, BUT NOT LIMITED TO, THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES, OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package com.gitsudo404.event_ingestion.controller;

import com.gitsudo404.event_ingestion.dto.EventResponseDTO;
import com.gitsudo404.event_ingestion.dto.EventMetaData;
import com.gitsudo404.event_ingestion.dto.VehicleBatteryEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleDiagnosticEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleEngineTelemetryEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleFuelEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleGPSEventRequestDTO;
import com.gitsudo404.event_ingestion.middleware.rate_limiter.RateLimiter;
import com.gitsudo404.event_ingestion.model.VehicleEventType;
import com.gitsudo404.event_ingestion.service.VehicleEventService;
import jakarta.validation.Valid;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeUnit;
import java.util.function.Supplier;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/events/vehicle")
public class VehicleEventController {

  private final VehicleEventService vehicleEventService;
  private final RateLimiter gpsRateLimiter;
  private final RateLimiter batteryRateLimiter;
  private final RateLimiter fuelRateLimiter;
  private final RateLimiter diagnosticRateLimiter;
  private final RateLimiter engineTelemetryRateLimiter;

  public VehicleEventController(
      VehicleEventService vehicleEventService,
      @Qualifier("vehicleGPSEventRateLimiter") RateLimiter gpsRateLimiter,
      @Qualifier("vehicleBatteryEventRateLimiter") RateLimiter batteryRateLimiter,
      @Qualifier("vehicleFuelEventRateLimiter") RateLimiter fuelRateLimiter,
      @Qualifier("vehicleDiagnosticEventRateLimiter") RateLimiter diagnosticRateLimiter,
      @Qualifier("vehicleEngineTelemetryEventRateLimiter") RateLimiter engineTelemetryRateLimiter) {
    this.vehicleEventService = vehicleEventService;
    this.gpsRateLimiter = gpsRateLimiter;
    this.batteryRateLimiter = batteryRateLimiter;
    this.fuelRateLimiter = fuelRateLimiter;
    this.diagnosticRateLimiter = diagnosticRateLimiter;
    this.engineTelemetryRateLimiter = engineTelemetryRateLimiter;
  }

  @PostMapping("/gps")
  public CompletableFuture<ResponseEntity<EventResponseDTO>> handleVehicleGPSEvent(
      @Valid @RequestBody VehicleGPSEventRequestDTO request) {
    return handleEvent(request, VehicleEventType.GPS, gpsRateLimiter,
        () -> vehicleEventService.ingestGPSEvent(request));
  }

  @PostMapping("/battery")
  public CompletableFuture<ResponseEntity<EventResponseDTO>> handleVehicleBatteryEvent(
      @Valid @RequestBody VehicleBatteryEventRequestDTO request) {
    return handleEvent(request, VehicleEventType.BATTERY, batteryRateLimiter,
        () -> vehicleEventService.ingestBatteryEvent(request));
  }

  @PostMapping("/fuel")
  public CompletableFuture<ResponseEntity<EventResponseDTO>> handleVehicleFuelEvent(
      @Valid @RequestBody VehicleFuelEventRequestDTO request) {
    return handleEvent(request, VehicleEventType.FUEL, fuelRateLimiter,
        () -> vehicleEventService.ingestFuelEvent(request));
  }

  @PostMapping("/diagnostic")
  public CompletableFuture<ResponseEntity<EventResponseDTO>> handleVehicleDiagnosticEvent(
      @Valid @RequestBody VehicleDiagnosticEventRequestDTO request) {
    return handleEvent(request, VehicleEventType.DIAGNOSTIC, diagnosticRateLimiter,
        () -> vehicleEventService.ingestDiagnosticEvent(request));
  }

  @PostMapping("/engine-telemetry")
  public CompletableFuture<ResponseEntity<EventResponseDTO>> handleVehicleEngineTelemetryEvent(
      @Valid @RequestBody VehicleEngineTelemetryEventRequestDTO request) {
    return handleEvent(request, VehicleEventType.ENGINE_TELEMETRY, engineTelemetryRateLimiter,
        () -> vehicleEventService.ingestEngineTelemetryEvent(request));
  }

  private CompletableFuture<ResponseEntity<EventResponseDTO>> handleEvent(
      EventMetaData request, VehicleEventType expectedEventType, RateLimiter rateLimiter,
      Supplier<CompletableFuture<Void>> ingestion) {
    if (!expectedEventType.matches(request.getEventType())) {
      EventResponseDTO response = new EventResponseDTO()
          .SetEventId(request.getEventId())
          .SetMessage("Event type does not match endpoint")
          .SetStatus("invalid-event-type");
      return CompletableFuture.completedFuture(ResponseEntity.badRequest().body(response));
    }

    if (!rateLimiter.allowRequest(request.getEntityId())) {
      EventResponseDTO response = new EventResponseDTO().SetEventId(request.getEventId())
          .SetMessage("Too many requests")
          .SetStatus("429-rate-limited");
      return CompletableFuture.completedFuture(ResponseEntity.status(HttpStatus.TOO_MANY_REQUESTS).body(response));
    }

    return ingestion.get()
        .orTimeout(10, TimeUnit.SECONDS)
        .handle((ignored, failure) -> {
          if (failure != null) {
            EventResponseDTO response = new EventResponseDTO()
                .SetEventId(request.getEventId())
                .SetMessage("Could not confirm event delivery to Kafka")
                .SetStatus("delivery-failed");
            return ResponseEntity.status(HttpStatus.SERVICE_UNAVAILABLE).body(response);
          }

          EventResponseDTO response = new EventResponseDTO()
              .SetEventId(request.getEventId())
              .SetMessage("Event accepted by Kafka")
              .SetStatus("success");
          return ResponseEntity.ok(response);
        });
  }
}
