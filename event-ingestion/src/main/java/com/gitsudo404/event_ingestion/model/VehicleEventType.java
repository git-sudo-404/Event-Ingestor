package com.gitsudo404.event_ingestion.model;

public enum VehicleEventType {
  GPS("vehicle.location"),
  BATTERY("vehicle.battery"),
  FUEL("vehicle.fuel"),
  DIAGNOSTIC("vehicle.diagnostic"),
  ENGINE_TELEMETRY("vehicle.engine_telemetry");

  private final String value;

  VehicleEventType(String value) {
    this.value = value;
  }

  public boolean matches(String eventType) {
    return value.equals(eventType);
  }
}
