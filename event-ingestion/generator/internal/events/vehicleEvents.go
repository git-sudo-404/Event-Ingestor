// MIT License
//
// Copyright (c) 2026 git-sudo-404
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package events

import "time"

type EventMetaData struct {
	EventID       string    `json:"event_id"`
	EventType     string    `json:"event_type"`
	SchemaVersion int       `json:"schema_version"`
	Source        string    `json:"source"`
	TimeStamp     time.Time `json:"timestamp"`
	EntityID      string    `json:"entity_id"`
}

type gpsEvent struct {
	EventMetaData
	Latitude     float64  `json:"latitude"`
	Longitude    float64  `json:"longitude"`
	Altitude     *float64 `json:"altitude,omitempty"`       // height above sea level in meters
	Speed        float64  `json:"speed"`                    // speed in km/h
	HeadingDEG   float64  `json:"heading_deg"`              // direction of travel: 0=N, 90=E, 180=S, 270=W
	GPSAccuracyM *float64 `json:"gps_accuracy_m,omitempty"` // estimated GPS error in meters
	OdometerKM   *float64 `json:"odometer_km,omitempty"`    // total distance travelled in km
}

type engineTelemetryEvent struct {
	EventMetaData
	RPM         float64 `json:"rpm"`          // engine rotations per minute
	CoolantTemp float64 `json:"coolant_temp"` // temperature in °C
	EngineLoad  float64 `json:"engine_load"`  // percentage of engine capacity currently being used
}

type DiagnosticEvent struct {
	EventMetaData
	Code        string `json:"code"`
	Severity    string `json:"severity"` // e.g. info, warning, critical
	Description string `json:"description"`
}

type FuelEvent struct {
	EventMetaData
	FuelLevelPct       float64  `json:"fuel_level_pct"`                  // percentage of fuel remaining in the tank
	FuelRateLPH        *float64 `json:"fuel_rate_lph,omitempty"`         // fuel consumed per hour in litres
	FuelConsumedTotalL *float64 `json:"fuel_consumed_total_l,omitempty"` // total fuel consumed in litres
}

type BatteryEvent struct {
	EventMetaData
	VoltageV         float64  `json:"voltage_v"`                     // battery voltage
	CurrentA         *float64 `json:"current_a,omitempty"`           // current flowing through the battery in amps
	TemperatureC     *float64 `json:"temperature_c,omitempty"`       // battery temperature in °C
	StateOfChargePct *float64 `json:"state_of_charge_pct,omitempty"` // estimated battery charge remaining as a percentage
}
