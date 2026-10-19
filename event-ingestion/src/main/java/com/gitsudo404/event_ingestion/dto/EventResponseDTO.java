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

package com.gitsudo404.event_ingestion.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class EventResponseDTO {

  @JsonProperty("event_id")
  public String eventId;

  @JsonProperty("status")
  public String status;

  @JsonProperty("message")
  public String message;

  public EventResponseDTO SetEventId(String eventId) {
    this.eventId = eventId;
    return this;
  }

  public EventResponseDTO SetStatus(String status) {
    this.status = status;
    return this;
  }

  public EventResponseDTO SetMessage(String message) {
    this.message = message;
    return this;
  }

  public String GetEventId() {
    return this.eventId;
  }

  public String GetStatus() {
    return this.status;
  }

  public String GetMessage() {
    return this.message;
  }
}
