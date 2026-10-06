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

type LogLevel string

const (
	LogDebug LogLevel = "DEBUG"
	LogInfo  LogLevel = "INFO"
	LogWarn  LogLevel = "WARN"
	LogError LogLevel = "ERROR"
)

type ApplicationLogEvent struct {
	EventMetaData
	Level      string            `json:"level"` // INFO, WARN, ERROR, DEBUG
	Service    string            `json:"service"`
	Message    string            `json:"message"`
	RequestID  *string           `json:"request_id,omitempty"`
	TraceID    *string           `json:"trace_id,omitempty"`
	Host       *string           `json:"host,omitempty"`
	StatusCode *int              `json:"status_code,omitempty"`
	LatencyMs  *float64          `json:"latency_ms,omitempty"` // request processing time in milliseconds
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func NewApplicationLogEvent() *ApplicationLogEvent {
	return &ApplicationLogEvent{
		EventMetaData: NewEventMetaData(),
		Level:         "",
		Service:       "",
		Message:       "",
	}
}

func (ale *ApplicationLogEvent) SetLevel(Level string) *ApplicationLogEvent {
	ale.Level = Level
	return ale
}

func (ale *ApplicationLogEvent) SetService(Service string) *ApplicationLogEvent {
	ale.Service = Service
	return ale
}
func (ale *ApplicationLogEvent) SetMessage(Message string) *ApplicationLogEvent {
	ale.Message = Message
	return ale
}
func (ale *ApplicationLogEvent) SetRequestID(RequestID string) *ApplicationLogEvent {
	ale.RequestID = &RequestID
	return ale
}
func (ale *ApplicationLogEvent) SetTraceID(TraceID string) *ApplicationLogEvent {
	ale.TraceID = &TraceID
	return ale
}
func (ale *ApplicationLogEvent) SetHost(Host string) *ApplicationLogEvent {
	ale.Host = &Host
	return ale
}
func (ale *ApplicationLogEvent) SetStatusCode(StatusCode int) *ApplicationLogEvent {
	ale.StatusCode = &StatusCode
	return ale
}
func (ale *ApplicationLogEvent) SetLatencyMs(LatencyMs float64) *ApplicationLogEvent {
	ale.LatencyMs = &LatencyMs
	return ale
}

func (ale *ApplicationLogEvent) SetMetadata(Metadata map[string]string) *ApplicationLogEvent {
	ale.Metadata = Metadata
	return ale
}
