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

type VehicleEngineTelemetryEvent struct {
	EventMetaData
	RPM           float64 `json:"rpm"`             // engine rotations per minute
	CoolantTempC  float64 `json:"coolant_temp_c"`  // temperature in °C
	EngineLoadPct float64 `json:"engine_load_pct"` // percentage of engine capacity currently being used
}

func NewVehicleEngineTelemetryEvent() *VehicleEngineTelemetryEvent {
	return &VehicleEngineTelemetryEvent{
		EventMetaData: NewEventMetaData(),
		RPM:           0,
		CoolantTempC:  0,
		EngineLoadPct: 0,
	}
}

func (ete *VehicleEngineTelemetryEvent) SetRPM(RPM float64) *VehicleEngineTelemetryEvent {
	ete.RPM = RPM
	return ete
}

func (ete *VehicleEngineTelemetryEvent) SetCoolantTempC(CoolantTempC float64) *VehicleEngineTelemetryEvent {
	ete.CoolantTempC = CoolantTempC
	return ete
}

func (ete *VehicleEngineTelemetryEvent) SetEngineLoadPct(EngineLoadPct float64) *VehicleEngineTelemetryEvent {
	ete.EngineLoadPct = EngineLoadPct
	return ete
}
