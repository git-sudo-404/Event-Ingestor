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

import "time"

type VehicleTelemetryEventType string

const (
	VechicleGPSEvent            VehicleTelemetryEventType = "vehicle.location"
	VehicleEngineTelemetryEvent VehicleTelemetryEventType = "vehicle.engine_telemetry"
	VehicleDiagnosticEvent      VehicleTelemetryEventType = "vehicle.diagnostic"
	VehicleFuelEvent            VehicleTelemetryEventType = "vehicle.fuel"
	VehicleBatteryEvent         VehicleTelemetryEventType = "vehicle.battery"
)

type DiagnosticSeverity string

const (
	SeverityInfo     DiagnosticSeverity = "INFO"
	SeverityWarning  DiagnosticSeverity = "WARNING"
	SeverityCritical DiagnosticSeverity = "CRITICAL"
)

type EventMetaData struct {
	EventID       string    `json:"event_id"`
	EventType     string    `json:"event_type"`
	SchemaVersion int       `json:"schema_version"`
	Source        string    `json:"source"`
	TimeStamp     time.Time `json:"timestamp"`
	EntityID      string    `json:"entity_id"`
}

//NOTE: Any struct embedding the EventMetaData will also have access to these setter functions since GO can promote these functions to the embedding struct

func (emd *EventMetaData) SetEventID(EventID string) *EventMetaData {
	emd.EventID = EventID
	return emd
}

func (emd *EventMetaData) SetEventType(EventType string) *EventMetaData {
	emd.EventType = EventType
	return emd
}

func (emd *EventMetaData) SetSchemaVersion(SchemaVersion int) *EventMetaData {
	emd.SchemaVersion = SchemaVersion
	return emd
}

func (emd *EventMetaData) SetSource(Source string) *EventMetaData {
	emd.Source = Source
	return emd
}

func (emd *EventMetaData) SetTimeStamp(TimeStamp time.Time) *EventMetaData {
	emd.TimeStamp = TimeStamp
	return emd
}

func (emd *EventMetaData) SetEntityID(EntityID string) *EventMetaData {
	emd.EntityID = EntityID
	return emd
}

type GPSEvent struct {
	EventMetaData
	Latitude     float64  `json:"latitude"`
	Longitude    float64  `json:"longitude"`
	AltitudeM    *float64 `json:"altitude_m,omitempty"`     // height above sea level in meters
	SpeedKMH     float64  `json:"speed_kmh"`                // speed in km/h
	HeadingDir   float64  `json:"heading_dir"`              // direction of travel: 0=N, 90=E, 180=S, 270=W
	GPSAccuracyM *float64 `json:"gps_accuracy_m,omitempty"` // estimated GPS error in meters
	OdometerKM   *float64 `json:"odometer_km,omitempty"`    // total distance travelled in km
}

func NewGPSEvent() *GPSEvent {
	return &GPSEvent{
		EventMetaData: EventMetaData{
			EventID:       "",
			EventType:     "",
			SchemaVersion: 0,
			Source:        "",
			TimeStamp:     time.Now(),
			EntityID:      "",
		},
		Latitude:   0,
		Longitude:  0,
		SpeedKMH:   0,
		HeadingDir: 0,
	}
}

func (gps *GPSEvent) SetLatitude(Latitude float64) *GPSEvent {
	gps.Latitude = Latitude
	return gps
}

func (gps *GPSEvent) SetLongitude(Longitude float64) *GPSEvent {
	gps.Longitude = Longitude
	return gps
}

func (gps *GPSEvent) SetAltitudeM(AltitudeM float64) *GPSEvent {
	gps.AltitudeM = &AltitudeM
	return gps
}

func (gps *GPSEvent) SetSpeedKMH(SpeedKMH float64) *GPSEvent {
	gps.SpeedKMH = SpeedKMH
	return gps
}

func (gps *GPSEvent) SetHeadingDir(HeadingDir float64) *GPSEvent {
	gps.HeadingDir = HeadingDir
	return gps
}

func (gps *GPSEvent) SetGPSAccuracyM(GPSAccuracyM float64) *GPSEvent {
	gps.GPSAccuracyM = &GPSAccuracyM
	return gps
}

func (gps *GPSEvent) SetOdometerKM(OdometerKM float64) *GPSEvent {
	gps.OdometerKM = &OdometerKM
	return gps
}

type EngineTelemetryEvent struct {
	EventMetaData
	RPM           float64 `json:"rpm"`             // engine rotations per minute
	CoolantTempC  float64 `json:"coolant_temp_c"`  // temperature in °C
	EngineLoadPct float64 `json:"engine_load_pct"` // percentage of engine capacity currently being used
}

func NewEngineTelemetryEvent() *EngineTelemetryEvent {
	return &EngineTelemetryEvent{
		EventMetaData: EventMetaData{
			EventID:       "",
			EventType:     "",
			SchemaVersion: 0,
			Source:        "",
			TimeStamp:     time.Now(),
			EntityID:      "",
		},
		RPM:           0,
		CoolantTempC:  0,
		EngineLoadPct: 0,
	}
}

func (ete *EngineTelemetryEvent) SetRPM(RPM float64) *EngineTelemetryEvent {
	ete.RPM = RPM
	return ete
}

func (ete *EngineTelemetryEvent) SetCoolantTempC(CoolantTempC float64) *EngineTelemetryEvent {
	ete.CoolantTempC = CoolantTempC
	return ete
}

func (ete *EngineTelemetryEvent) SetEngineLoadPct(EngineLoadPct float64) *EngineTelemetryEvent {
	ete.EngineLoadPct = EngineLoadPct
	return ete
}

type DiagnosticEvent struct {
	EventMetaData
	Code        string `json:"code"`
	Severity    string `json:"severity"` // e.g. info, warning, critical
	Description string `json:"description"`
}

func NewDiagnosticEvent() *DiagnosticEvent {
	return &DiagnosticEvent{
		EventMetaData: EventMetaData{
			EventID:       "",
			EventType:     "",
			SchemaVersion: 0,
			Source:        "",
			TimeStamp:     time.Now(),
			EntityID:      "",
		},
		Code:        "",
		Severity:    "",
		Description: "",
	}
}

func (de *DiagnosticEvent) SetCode(Code string) *DiagnosticEvent {
	de.Code = Code
	return de
}

func (de *DiagnosticEvent) SetSeverity(Severity string) *DiagnosticEvent {
	de.Severity = Severity
	return de
}

func (de *DiagnosticEvent) SetDescription(Description string) *DiagnosticEvent {
	de.Description = Description
	return de
}

type FuelEvent struct {
	EventMetaData
	FuelLevelPct       float64  `json:"fuel_level_pct"`                  // percentage of fuel remaining in the tank
	FuelRateLPH        *float64 `json:"fuel_rate_lph,omitempty"`         // fuel consumed per hour in litres
	FuelConsumedTotalL *float64 `json:"fuel_consumed_total_l,omitempty"` // total fuel consumed in litres
}

func NewFuelEvent() *FuelEvent {
	return &FuelEvent{
		EventMetaData: EventMetaData{
			EventID:       "",
			EventType:     "",
			SchemaVersion: 0,
			Source:        "",
			TimeStamp:     time.Now(),
			EntityID:      "",
		},
		FuelLevelPct: 0,
	}
}

func (fe *FuelEvent) SetFuelLevelPct(FuelLevelPct float64) *FuelEvent {
	fe.FuelLevelPct = FuelLevelPct
	return fe
}
func (fe *FuelEvent) SetFuelRateLPH(FuelRateLPH float64) *FuelEvent {
	fe.FuelRateLPH = &FuelRateLPH
	return fe
}
func (fe *FuelEvent) SetFuelConsumedTotalL(FuelConsumedTotalL float64) *FuelEvent {
	fe.FuelConsumedTotalL = &FuelConsumedTotalL
	return fe
}

type BatteryEvent struct {
	EventMetaData
	VoltageV         float64  `json:"voltage_v"`                     // battery voltage
	CurrentA         *float64 `json:"current_a,omitempty"`           // current flowing through the battery in amps
	TemperatureC     *float64 `json:"temperature_c,omitempty"`       // battery temperature in °C
	StateOfChargePct *float64 `json:"state_of_charge_pct,omitempty"` // estimated battery charge remaining as a percentage
}

func NewBatteryEvent() *BatteryEvent {
	return &BatteryEvent{
		EventMetaData: EventMetaData{
			EventID:       "",
			EventType:     "",
			SchemaVersion: 0,
			Source:        "",
			TimeStamp:     time.Now(),
			EntityID:      "",
		},
		VoltageV: 0,
	}
}

func (be *BatteryEvent) SetVolateV(VoltageV float64) *BatteryEvent {
	be.VoltageV = VoltageV
	return be
}
func (be *BatteryEvent) SetCurrentA(CurrentA float64) *BatteryEvent {
	be.CurrentA = &CurrentA
	return be
}
func (be *BatteryEvent) SetTemperatureC(TemperatureC float64) *BatteryEvent {
	be.TemperatureC = &TemperatureC
	return be
}
func (be *BatteryEvent) SetStateOfChargePct(StateOfChargePct float64) *BatteryEvent {
	be.StateOfChargePct = &StateOfChargePct
	return be
}
