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
import com.gitsudo404.event_ingestion.dto.VehicleGPSEventRequestDTO;
import com.gitsudo404.event_ingestion.middleware.rate_limiter.RateLimiter;
import com.gitsudo404.event_ingestion.service.VehicleEventService;
import jakarta.validation.Valid;

import org.springframework.beans.factory.annotation.Qualifier;
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
  // TODO: Here any incoming request whose schema doesn't meet the validation
  // would get a 400 and be rejected , add a failover handling here in v2
  public EventResponseDTO handleVehicleGPSEvent(
      @Valid @RequestBody VehicleGPSEventRequestDTO request) {
    // check authentication -- will be done by mTLS using nginx

    if (!gpsRateLimiter.allowRequest(request.getEntityId())) {
      return new EventResponseDTO().SetEventId(request.getEventId())
          .SetMessage("Too many requests")
          .SetStatus("429-rate-limited"); // TODO: Use some other proper status returning here
    }
    // validate schema
    System.out.println(request);

    // EventResponseDTO response =
    return vehicleEventService.ingest(request);
    // return response;
  }
}
