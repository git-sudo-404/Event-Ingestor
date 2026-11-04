package com.gitsudo404.event_ingestion.producer;

public enum KafkaTopic {
  VEHICLE_GPS_EVENT("events.vehicle.gps"),
  VEHICLE_ENGINE_TELEMETRY_EVENT("events.vehicle.engine-telemetry"),
  VEHICLE_FUEL_EVENT("events.vehicle.fuel"),
  VEHICLE_BATTERY_EVENT("events.vehicle.battery"),
  VEHICLE_DIAGNOSTIC_EVENT("events.vehicle.diagnostic");

  private final String topicName;

  KafkaTopic(String topicName) {
    this.topicName = topicName;
  }

  public String topicName() {
    return topicName;
  }
}
