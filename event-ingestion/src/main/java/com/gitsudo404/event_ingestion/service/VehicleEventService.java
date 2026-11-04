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

package com.gitsudo404.event_ingestion.service;

import com.gitsudo404.event_ingestion.dto.EventMetaData;
import com.gitsudo404.event_ingestion.dto.VehicleBatteryEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleDiagnosticEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleEngineTelemetryEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleFuelEventRequestDTO;
import com.gitsudo404.event_ingestion.dto.VehicleGPSEventRequestDTO;
import com.gitsudo404.event_ingestion.producer.EventPublisher;
import com.gitsudo404.event_ingestion.producer.KafkaTopic;

import java.util.concurrent.CompletableFuture;
import org.springframework.stereotype.Service;
import tools.jackson.databind.ObjectMapper;

@Service
public class VehicleEventService {

  private final EventPublisher eventPublisher;
  private final ObjectMapper objectMapper;

  public VehicleEventService(EventPublisher eventPublisher, ObjectMapper objectMapper) {
    this.eventPublisher = eventPublisher;
    this.objectMapper = objectMapper;
  }

  public CompletableFuture<Void> ingestGPSEvent(VehicleGPSEventRequestDTO request) {
    return publish(request, KafkaTopic.VEHICLE_GPS_EVENT);
  }

  public CompletableFuture<Void> ingestBatteryEvent(VehicleBatteryEventRequestDTO request) {
    return publish(request, KafkaTopic.VEHICLE_BATTERY_EVENT);
  }

  public CompletableFuture<Void> ingestFuelEvent(VehicleFuelEventRequestDTO request) {
    return publish(request, KafkaTopic.VEHICLE_FUEL_EVENT);
  }

  public CompletableFuture<Void> ingestDiagnosticEvent(VehicleDiagnosticEventRequestDTO request) {
    return publish(request, KafkaTopic.VEHICLE_DIAGNOSTIC_EVENT);
  }

  public CompletableFuture<Void> ingestEngineTelemetryEvent(VehicleEngineTelemetryEventRequestDTO request) {
    return publish(request, KafkaTopic.VEHICLE_ENGINE_TELEMETRY_EVENT);
  }

  private CompletableFuture<Void> publish(EventMetaData request, KafkaTopic topic) {
    String requestString = objectMapper.writeValueAsString(request);
    return eventPublisher.publish(topic, request.getEntityId(), requestString)
        .thenAccept(sendResult -> {
        });
  }
}
