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
 * The above copyright notice and this permission notice shall be included in all
 * Copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING, BUT NOT LIMITED TO, THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES, OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package events

type VehicleGPSEvent struct {
	EventMetaData
	Latitude     float64  `json:"latitude"`
	Longitude    float64  `json:"longitude"`
	AltitudeM    *float64 `json:"altitude_m,omitempty"`     // height above sea level in meters
	SpeedKMH     float64  `json:"speed_kmh"`                // speed in km/h
	HeadingDir   float64  `json:"heading_dir"`              // direction of travel: 0=N, 90=E, 180=S, 270=W
	GPSAccuracyM *float64 `json:"gps_accuracy_m,omitempty"` // estimated GPS error in meters
	OdometerKM   *float64 `json:"odometer_km,omitempty"`    // total distance travelled in km
}

func NewVehicleGPSEvent() *VehicleGPSEvent {
	return &VehicleGPSEvent{
		EventMetaData: NewEventMetaData(),
		Latitude:      0,
		Longitude:     0,
		SpeedKMH:      0,
		HeadingDir:    0,
	}
}

func (gps *VehicleGPSEvent) SetLatitude(Latitude float64) *VehicleGPSEvent {
	gps.Latitude = Latitude
	return gps
}

func (gps *VehicleGPSEvent) SetLongitude(Longitude float64) *VehicleGPSEvent {
	gps.Longitude = Longitude
	return gps
}

func (gps *VehicleGPSEvent) SetAltitudeM(AltitudeM float64) *VehicleGPSEvent {
	gps.AltitudeM = &AltitudeM
	return gps
}

func (gps *VehicleGPSEvent) SetSpeedKMH(SpeedKMH float64) *VehicleGPSEvent {
	gps.SpeedKMH = SpeedKMH
	return gps
}

func (gps *VehicleGPSEvent) SetHeadingDir(HeadingDir float64) *VehicleGPSEvent {
	gps.HeadingDir = HeadingDir
	return gps
}

func (gps *VehicleGPSEvent) SetGPSAccuracyM(GPSAccuracyM float64) *VehicleGPSEvent {
	gps.GPSAccuracyM = &GPSAccuracyM
	return gps
}

func (gps *VehicleGPSEvent) SetOdometerKM(OdometerKM float64) *VehicleGPSEvent {
	gps.OdometerKM = &OdometerKM
	return gps
}
