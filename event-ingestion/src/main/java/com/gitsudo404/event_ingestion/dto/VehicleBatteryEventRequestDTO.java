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

public class VehicleBatteryEventRequestDTO extends EventMetaData {

  @JsonProperty("voltage_v")
  @NotNull
  private Float voltageV;

  @JsonProperty("current_a")
  private Float currentA;

  @JsonProperty("temperature_c")
  private Float temperatureC;

  @JsonProperty("state_of_charge_pct")
  private Float stateOfChargePct;

  public Float getVoltageV() {
    return voltageV;
  }

  public void setVoltageV(Float voltageV) {
    this.voltageV = voltageV;
  }

  public Float getCurrentA() {
    return currentA;
  }

  public void setCurrentA(Float currentA) {
    this.currentA = currentA;
  }

  public Float getTemperatureC() {
    return temperatureC;
  }

  public void setTemperatureC(Float temperatureC) {
    this.temperatureC = temperatureC;
  }

  public Float getStateOfChargePct() {
    return stateOfChargePct;
  }

  public void setStateOfChargePct(Float stateOfChargePct) {
    this.stateOfChargePct = stateOfChargePct;
  }
}
