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

package generator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"generator/internal/events"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func (v *vehicleState) makeVehicleEventRequest(eventJSON []byte) {
	resp, err := v.client.Post(
		cfg.EventIngestionURL+"/vehicle/"+strconv.Itoa(int(v.EventNumber))+"/events",
		"application/json",
		bytes.NewReader(eventJSON),
	)
	if err != nil {
		fmt.Println("[ERROR] Error Occured while sending http event from vehicle")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("ingestion returned %s", resp.Status)
	}
	fmt.Println("[LOG] Sending : ", string(eventJSON))
}

func (v *vehicleState) generateVehicleGPSEvent() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.EventNumber += 1
	GPSEvent := events.NewVehicleGPSEvent()
	GPSEvent.
		SetEventID("vehicle.event." + strconv.Itoa(int(v.EventNumber))).
		SetEventType(string(events.VehicleGPSEventType)).
		SetSchemaVersion(int(v.SchemaVersion)).
		SetSource("vehicle-simulator").
		SetTimeStamp(time.Now()).
		SetEntityID(v.VehicleID)
	GPSEvent.
		SetLatitude(v.Latitude).
		SetLongitude(v.Longitude).
		SetAltitudeM(v.AltitudeM).
		SetSpeedKMH(v.SpeedKMH).
		SetHeadingDir(v.HeadingDir).
		SetGPSAccuracyM(v.GPSAccuracyM).
		SetOdometerKM(v.OdometerKM)
	GPSEventJSON, err := json.Marshal(GPSEvent)
	if err != nil {
		fmt.Println("[ERROR] Error while Marshaling Vehicle GPSEvent to JSON", err)
		return
	}
	v.makeVehicleEventRequest(GPSEventJSON)
}

func (v *vehicleState) generateVehicleFuelEvent() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.EventNumber += 1
	FuelEvent := events.NewVehicleFuelEvent()
	FuelEvent.
		SetEventID("vehicle.event." + strconv.Itoa(int(v.EventNumber))).
		SetEventType(string(events.VehicleFuelEventType)).
		SetSchemaVersion(int(v.SchemaVersion)).
		SetSource("vehicle-simulator").
		SetTimeStamp(time.Now()).
		SetEntityID(v.VehicleID)
	FuelEvent.
		SetFuelLevelPct(v.FuelLevelPct).
		SetFuelRateLPH(v.FuelRateLPH).
		SetFuelConsumedTotalL(v.FuelConsumedTotalL)
	FuelEventJSON, err := json.Marshal(FuelEvent)
	if err != nil {
		fmt.Println("[ERROR] Error marshaling Vehicle Fuel Event to JSON", err)
		return
	}
	v.makeVehicleEventRequest(FuelEventJSON)
}

func (v *vehicleState) generateVehicleBatteryEvent() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.EventNumber += 1
	BatteryEvent := events.NewVehicleBatteryEvent()
	BatteryEvent.
		SetEventID("vehicle.event." + strconv.Itoa(int(v.EventNumber))).
		SetEventType(string(events.VehicleBatteryEventType)).
		SetSchemaVersion(int(v.SchemaVersion)).
		SetSource("vehicle-simulator").
		SetTimeStamp(time.Now()).
		SetEntityID(v.VehicleID)
	BatteryEvent.
		SetVolateV(v.VoltageV).
		SetCurrentA(v.CurrentA).
		SetTemperatureC(v.TemperatureC).
		SetStateOfChargePct(v.StateOfChargePct)
	BatteryEventJSON, err := json.Marshal(BatteryEvent)
	if err != nil {
		fmt.Println("[ERROR] Error marshaling Vehicle Battery Event to JSON", err)
		return
	}
	v.makeVehicleEventRequest(BatteryEventJSON)
}

func (v *vehicleState) generateVehicleEngineTelemetryEvent() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.EventNumber += 1
	EngineTelemetryEvent := events.NewVehicleEngineTelemetryEvent()
	EngineTelemetryEvent.
		SetEventID("vehicle.event." + strconv.Itoa(int(v.EventNumber))).
		SetEventType(string(events.VehicleEngineTelemetryEventType)).
		SetSchemaVersion(int(v.SchemaVersion)).
		SetSource("vehicle-simulator").
		SetTimeStamp(time.Now()).
		SetEntityID(v.VehicleID)
	EngineTelemetryEvent.
		SetRPM(v.RPM).
		SetCoolantTempC(v.CoolantTempC).
		SetEngineLoadPct(v.EngineLoadPct)
	EngineTelemetryEventJSON, err := json.Marshal(EngineTelemetryEvent)
	if err != nil {
		fmt.Println("[ERROR] Error marshaling Vehicle Engine Telemetry Event to JSON", err)
		return
	}
	v.makeVehicleEventRequest(EngineTelemetryEventJSON)
}

func (v *vehicleState) generateVehicleDiagnosticEvent(code string, severity string, description string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.EventNumber += 1
	DiagnosticEvent := events.NewVehicleDiagnosticEvent()
	DiagnosticEvent.
		SetEventID("vehicle.event." + strconv.Itoa(int(v.EventNumber))).
		SetEventType(string(events.VehicleDiagnosticEventType)).
		SetSchemaVersion(int(v.SchemaVersion)).
		SetSource("vehicle-simulator").
		SetTimeStamp(time.Now()).
		SetEntityID(v.VehicleID)
	DiagnosticEvent.
		SetCode(code).
		SetSeverity(severity).
		SetDescription(description)
	DiagnosticEventJSON, err := json.Marshal(DiagnosticEvent)
	if err != nil {
		fmt.Println("[ERROR] Error marshaling Vehicle Diagnostic Event to JSON", err)
		return
	}
	v.makeVehicleEventRequest(DiagnosticEventJSON)
}

func (v *vehicleState) startGeneratingEvents(ctx context.Context, vehicleEventsWG *sync.WaitGroup, interval time.Duration) {
	if interval <= 0 {
		return
	}

	randomWaitTime := rand.Float64() * 10
	time.Sleep(time.Duration(randomWaitTime * float64(time.Second)))

	ticker := time.NewTicker(interval)
	vehicleBatteryEventTicker := time.NewTicker(time.Second * 10)
	vehicleEngineTelemetryEventTicker := time.NewTicker(time.Second * 5)
	defer func() {
		ticker.Stop()
		vehicleEventsWG.Done()
		vehicleBatteryEventTicker.Stop()
		vehicleEngineTelemetryEventTicker.Stop()
	}()
	lastUpdate := time.Now()
	for {
		select {
		case <-ctx.Done():
			return
		case now := <-ticker.C:
			elapsed := now.Sub(lastUpdate)
			lastUpdate = now
			if elapsed <= 0 {
				continue
			}
			v.mu.Lock()
			distanceKM := v.updateGPSStates(elapsed)
			v.updateEngineStates(elapsed)
			v.updateFuelStates(elapsed, distanceKM)
			v.updateBatteryStates(elapsed)
			v.mu.Unlock()
			v.generateVehicleGPSEvent()
		case <-vehicleBatteryEventTicker.C:
			v.generateVehicleBatteryEvent()
			v.generateVehicleFuelEvent()
		case <-vehicleEngineTelemetryEventTicker.C:
			v.generateVehicleEngineTelemetryEvent()
		}
		if v.SpeedKMH >= 60 {
			v.generateVehicleDiagnosticEvent("high.speed", string(events.SeverityWarning), "Vehicle crossing speed limit")
		}
		if v.FuelLevelPct <= 10 {
			v.generateVehicleDiagnosticEvent("low.fuel", string(events.SeverityCritical), "Fuel is low "+strconv.Itoa(int(v.FuelLevelPct))+"%")
		}
		if v.StateOfChargePct <= 10 {
			v.generateVehicleDiagnosticEvent("low.battery", string(events.SeverityCritical), "Battery is low "+strconv.Itoa(int(v.FuelLevelPct))+"%")
		}
		if v.CoolantTempC >= 40 {
			v.generateVehicleDiagnosticEvent("high.coolant.temp", string(events.SeverityInfo), "Coolant Temperature is high "+strconv.Itoa(int(v.CoolantTempC))+"C")
		}
	}
}
