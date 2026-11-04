## Event processing

- Deduplicate retried vehicle events by `event_id` in the consumer. The ingestion API can time out before Kafka confirms a send, so a vehicle retry may produce two records with the same event ID. Use a durable check together with the consumer's business write; the vehicle ID Kafka key provides partition ordering, not deduplication.
