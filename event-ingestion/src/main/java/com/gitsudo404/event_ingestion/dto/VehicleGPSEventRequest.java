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

public class VehicleGPSEventRequest extends EventMetaData {

  @JsonProperty("latitude")
  @NotNull
  private Float latitude;

  @JsonProperty("longitude")
  @NotNull
  private Float longitude;

  @JsonProperty("altitude_m")
  private Float altitudeM;

  @JsonProperty("speed_kmh")
  @NotNull
  private Float speedKMH;

  @JsonProperty("heading_dir")
  @NotNull
  private Float headingDir;

  @JsonProperty("gps_accuracy_m")
  private Float gpsAccuracyM;

  @JsonProperty("odometer_km")
  private Float odometerKM;

  public Float getLatitude() {
    return latitude;
  }

  public void setLatitude(Float latitude) {
    this.latitude = latitude;
  }

  public Float getLongitude() {
    return longitude;
  }

  public void setLongitude(Float longitude) {
    this.longitude = longitude;
  }

  public Float getAltitudeM() {
    return altitudeM;
  }

  public void setAltitudeM(Float altitudeM) {
    this.altitudeM = altitudeM;
  }

  public Float getSpeedKMH() {
    return speedKMH;
  }

  public void setSpeedKMH(Float speedKMH) {
    this.speedKMH = speedKMH;
  }

  public Float getHeadingDir() {
    return headingDir;
  }

  public void setHeadingDir(Float headingDir) {
    this.headingDir = headingDir;
  }

  public Float getGpsAccuracyM() {
    return gpsAccuracyM;
  }

  public void setGpsAccuracyM(Float gpsAccuracyM) {
    this.gpsAccuracyM = gpsAccuracyM;
  }

  public Float getOdometerKM() {
    return odometerKM;
  }

  public void setOdometerKM(Float odometerKM) {
    this.odometerKM = odometerKM;
  }
}
