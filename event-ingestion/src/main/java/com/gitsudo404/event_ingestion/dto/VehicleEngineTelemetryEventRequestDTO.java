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
import jakarta.validation.constraints.NotNull;

public class VehicleEngineTelemetryEventRequestDTO extends EventMetaData {

  @JsonProperty("rpm")
  @NotNull
  private Float rpm;

  @JsonProperty("coolant_temp_c")
  @NotNull
  private Float coolantTempC;

  @JsonProperty("engine_load_pct")
  @NotNull
  private Float engineLoadPct;

  public Float getRpm() {
    return rpm;
  }

  public void setRpm(Float rpm) {
    this.rpm = rpm;
  }

  public Float getCoolantTempC() {
    return coolantTempC;
  }

  public void setCoolantTempC(Float coolantTempC) {
    this.coolantTempC = coolantTempC;
  }

  public Float getEngineLoadPct() {
    return engineLoadPct;
  }

  public void setEngineLoadPct(Float engineLoadPct) {
    this.engineLoadPct = engineLoadPct;
  }
}
