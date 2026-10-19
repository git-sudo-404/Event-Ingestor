# NOTE: This is just for local development, to create all the needed topics and partitions in kafka, DO NOT RUN THIS IN PROD

#!/usr/bin/env bash 


CONTAINER="kafka_broker_1"
BOOTSTRAP_SERVER="localhost:9092"

docker compose exec "$CONTAINER" /opt/kafka/bin/kafka-topics.sh \
  --bootstrap-server "$BOOTSTRAP_SERVER" \
  --create --if-not-exists \
  --topic vehicle-gps-event \
  --partitions 3 \
  --replication-factor 1



