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
	"math/rand"
	"strconv"
	"sync"
)

var startingPoints = []struct{ lat, lon, alt float64 }{
	{40.7580, -73.9855, 16},  // New York
	{41.8781, -87.6298, 182}, // Chicago
	{34.0522, -118.2437, 19}, // Los Angeles
	{29.7604, -95.3698, 15},  // Houston
}
var dir = []float64{0, 90, 180, 270}
var accuracyM = []float64{1, 2, 3, 4, 5, 6, 7, 2.4, 2.4, 5.2, 2.5, 10.3, 12.4, 13.6, 20.1, 2.3, 4.5, 5.6, 6.4}
var tempC = []float64{10.0, 15.0, 20.0, 25.0, 30.0, 35.0}

type vehicleState struct {
	VehicleID string

	//GPS states
	Latitude     float64
	Longitude    float64
	AltitudeM    float64
	SpeedKMH     float64
	HeadingDir   float64
	GPSAccuracyM float64
	OdometerKM   float64

	//Fuel states
	FuelLevelPct       float64
	FuelRateLPH        float64
	FuelConsumedTotalL float64

	//Engine states
	RPM           float64
	CoolantTempC  float64
	EngineLoadPct float64

	//Battery states
	VoltageV         float64
	CurrentA         float64
	TemperatureC     float64
	StateOfChargePct float64

	//EventMetaData states
	EventNumber float64
	mu          sync.Mutex
}

func CreateNewVehicle(vehicleNumber int) *vehicleState {

	VehicleID := "vehicle_" + strconv.Itoa(vehicleNumber)

	point := startingPoints[rand.Intn(len(startingPoints))]
	Latitude := point.lat + (rand.Float64()-0.5)*0.01
	Longitude := point.lon + (rand.Float64()-0.5)*0.01
	AltitudeM := point.alt

	SpeedKMH := float64(0)
	HeadingDir := dir[rand.Intn(len(dir))]

	GPSAccuracyM := accuracyM[rand.Intn(len(accuracyM))]
	OdometerKM := rand.Float64() * 100000

	FuelLevelPct := rand.ExpFloat64() * 100
	FuelRateLPH := float64(0)
	FuelConsumedTotalL := OdometerKM * 8.0 / 100.0

	RPM := float64(0)
	CoolantTempC := tempC[rand.Intn(len(tempC))] + rand.Float64() - rand.Float64()
	EngineLoadPct := float64(0)

	VoltageV := 12.6
	CurrentA := 0.0
	TemperatureC := CoolantTempC // both start near ambient temperature
	StateOfChargePct := 80.0 + rand.Float64()*20.0

	EventNumber := float64(0)

	return &vehicleState{
		VehicleID:          VehicleID,
		Latitude:           Latitude,
		Longitude:          Longitude,
		AltitudeM:          AltitudeM,
		SpeedKMH:           SpeedKMH,
		HeadingDir:         HeadingDir,
		GPSAccuracyM:       GPSAccuracyM,
		OdometerKM:         OdometerKM,
		FuelLevelPct:       FuelLevelPct,
		FuelRateLPH:        FuelRateLPH,
		FuelConsumedTotalL: FuelConsumedTotalL,
		RPM:                RPM,
		CoolantTempC:       CoolantTempC,
		EngineLoadPct:      EngineLoadPct,
		VoltageV:           VoltageV,
		CurrentA:           CurrentA,
		TemperatureC:       TemperatureC,
		StateOfChargePct:   StateOfChargePct,
		EventNumber:        EventNumber,
	}
}

func (v *vehicleState) updateGPSStates() {

}
