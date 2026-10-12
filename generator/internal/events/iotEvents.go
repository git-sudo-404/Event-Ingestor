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

type IOTEventType string

const (
	IOTTemperature IOTEventType = "iot.temperature"
	IOTHumidity    IOTEventType = "iot.humidity"
	IOTPressure    IOTEventType = "iot.pressure"
	IOTVibration   IOTEventType = "iot.vibration"
	IOTAirQuality  IOTEventType = "iot.air_quality"
	IOTPower       IOTEventType = "iot.power"
)

type IOTEvent struct {
	EventMetaData
	DeviceID   string            `json:"device_id"`
	SensorType IOTEventType      `json:"sensor_type"`
	Value      float64           `json:"value"`
	Unit       string            `json:"unit"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func (ie *IOTEvent) SetDeviceID(DeviceID string) *IOTEvent {
	ie.DeviceID = DeviceID
	return ie
}

func (ie *IOTEvent) SetSensorType(SensorType IOTEventType) *IOTEvent {
	ie.SensorType = SensorType
	return ie
}

func (ie *IOTEvent) SetValue(Value float64) *IOTEvent {
	ie.Value = Value
	return ie
}

func (ie *IOTEvent) SetUnit(Unit string) *IOTEvent {
	ie.Unit = Unit
	return ie
}

func (ie *IOTEvent) SetMetadata(Metadata map[string]string) *IOTEvent {
	ie.Metadata = Metadata
	return ie
}
