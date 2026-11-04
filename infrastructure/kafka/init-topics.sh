#!/usr/bin/env bash
# NOTE: This is just for local development, to create all the needed topics and partitions in kafka, DO NOT RUN THIS IN PROD
set -euo pipefail

CONTAINER="kafka_broker_1"
BOOTSTRAP_SERVER="localhost:9092"

for topic in \
  events.vehicle.gps \
  events.vehicle.battery \
  events.vehicle.fuel \
  events.vehicle.diagnostic \
  events.vehicle.engine-telemetry; do
  docker compose exec "$CONTAINER" /opt/kafka/bin/kafka-topics.sh \
    --bootstrap-server "$BOOTSTRAP_SERVER" \
    --create --if-not-exists \
    --topic "$topic" \
    --partitions 3 \
    --replication-factor 1
done
