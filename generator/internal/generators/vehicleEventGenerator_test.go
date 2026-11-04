package generator

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVehicleEventRequestUsesNamedEndpoint(t *testing.T) {
	paths := make(chan string, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths <- r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	previousURL := cfg.EventIngestionURL
	cfg.EventIngestionURL = server.URL
	defer func() { cfg.EventIngestionURL = previousURL }()

	vehicle := &vehicleState{}
	vehicle.makeVehicleEventRequest("battery", []byte(`{"event_id":"event-123"}`))

	if path := <-paths; path != "/api/v1/events/vehicle/battery" {
		t.Fatalf("unexpected endpoint: %s", path)
	}
}

func TestVehicleEventIDsDifferAcrossVehicles(t *testing.T) {
	first := &vehicleState{VehicleID: "vehicle_1", EventNumber: 1}
	second := &vehicleState{VehicleID: "vehicle_2", EventNumber: 1}
	if first.eventID() == second.eventID() {
		t.Fatal("different vehicles received the same event ID")
	}
	if first.eventID() != first.eventID() {
		t.Fatal("an event ID changed between attempts")
	}
}
