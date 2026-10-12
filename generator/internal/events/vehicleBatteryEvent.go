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

type VehicleBatteryEvent struct {
	EventMetaData
	VoltageV         float64  `json:"voltage_v"`                     // battery voltage
	CurrentA         *float64 `json:"current_a,omitempty"`           // current flowing through the battery in amps
	TemperatureC     *float64 `json:"temperature_c,omitempty"`       // battery temperature in °C
	StateOfChargePct *float64 `json:"state_of_charge_pct,omitempty"` // estimated battery charge remaining as a percentage
}

func NewVehicleBatteryEvent() *VehicleBatteryEvent {
	return &VehicleBatteryEvent{
		EventMetaData: NewEventMetaData(),
		VoltageV:      0,
	}
}

func (be *VehicleBatteryEvent) SetVolateV(VoltageV float64) *VehicleBatteryEvent {
	be.VoltageV = VoltageV
	return be
}
func (be *VehicleBatteryEvent) SetCurrentA(CurrentA float64) *VehicleBatteryEvent {
	be.CurrentA = &CurrentA
	return be
}
func (be *VehicleBatteryEvent) SetTemperatureC(TemperatureC float64) *VehicleBatteryEvent {
	be.TemperatureC = &TemperatureC
	return be
}
func (be *VehicleBatteryEvent) SetStateOfChargePct(StateOfChargePct float64) *VehicleBatteryEvent {
	be.StateOfChargePct = &StateOfChargePct
	return be
}
