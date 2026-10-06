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

type SystemMetricType string

const (
	SystemCPU     SystemMetricType = "system.cpu"
	SystemMemory  SystemMetricType = "system.memory"
	SystemDisk    SystemMetricType = "system.disk"
	SystemNetwork SystemMetricType = "system.network"
	SystemLoad    SystemMetricType = "system.load"
	SystemProcess SystemMetricType = "system.process"
)

type SystemMetricEvent struct {
	EventMetaData

	Host       string            `json:"host"`
	MetricType SystemMetricType  `json:"metric_type"`
	Value      float64           `json:"value"`
	Unit       string            `json:"unit"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func (sme *SystemMetricEvent) SetHost(Host string) *SystemMetricEvent {
	sme.Host = Host
	return sme
}

func (sme *SystemMetricEvent) SetMetricType(MetricType SystemMetricType) *SystemMetricEvent {
	sme.MetricType = MetricType
	return sme
}

func (sme *SystemMetricEvent) SetValue(Value float64) *SystemMetricEvent {
	sme.Value = Value
	return sme
}

func (sme *SystemMetricEvent) SetUnit(Unit string) *SystemMetricEvent {
	sme.Unit = Unit
	return sme
}

func (sme *SystemMetricEvent) SetMetadata(Metadata map[string]string) *SystemMetricEvent {
	sme.Metadata = Metadata
	return sme
}
