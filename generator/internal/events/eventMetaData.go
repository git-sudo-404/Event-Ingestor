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

type EventType string

const (
	VehicleGPSEventType             EventType = "vehicle.location"
	VehicleEngineTelemetryEventType EventType = "vehicle.engine_telemetry"
	VehicleDiagnosticEventType      EventType = "vehicle.diagnostic"
	VehicleFuelEventType            EventType = "vehicle.fuel"
	VehicleBatteryEventType         EventType = "vehicle.battery"
)

type EventMetaData struct {
	EventID       string    `json:"event_id"`
	EventType     string    `json:"event_type"`
	SchemaVersion int       `json:"schema_version"`
	Source        string    `json:"source"`
	TimeStamp     time.Time `json:"timestamp"`
	EntityID      string    `json:"entity_id"`
}

func NewEventMetaData() EventMetaData {
	return EventMetaData{
		EventID:       "",
		EventType:     "",
		SchemaVersion: 0,
		Source:        "",
		TimeStamp:     time.Now(),
		EntityID:      "",
	}
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
