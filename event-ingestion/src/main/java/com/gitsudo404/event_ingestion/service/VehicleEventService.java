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

import com.gitsudo404.event_ingestion.dto.EventResponseDTO;
import com.gitsudo404.event_ingestion.dto.VehicleGPSEventRequestDTO;
import com.gitsudo404.event_ingestion.producer.EventPublisher;
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

  public EventResponseDTO ingest(VehicleGPSEventRequestDTO request) {
    String requestString = objectMapper.writeValueAsString(request);
    eventPublisher.publish(requestString);
    return new EventResponseDTO()
        .SetMessage("DUMMMY MEssage").SetStatus("success").SetEventId(request.getEventId());
  }
}
