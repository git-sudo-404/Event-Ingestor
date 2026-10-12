# Event-Ingestor

High Throughput event Ingestion

```
.
├── event-ingestion
│   ├── event-ingestion
│   │   └── bin
│   │       ├── src
│   │       │   ├── main
│   │       │   │   ├── java
│   │       │   │   │   └── com
│   │       │   │   │       └── gitsudo404
│   │       │   │   │           └── event_ingestion
│   │       │   │   │               ├── config
│   │       │   │   │               ├── controller
│   │       │   │   │               ├── dto
│   │       │   │   │               ├── exception
│   │       │   │   │               ├── producer
│   │       │   │   │               ├── ratelimit
│   │       │   │   │               ├── service
│   │       │   │   │               └── validation
│   │       │   │   └── resources
│   │       │   │       └── templates
│   │       │   └── test
│   │       │       └── java
│   │       │           └── com
│   │       │               └── gitsudo404
│   │       │                   └── event_ingestion
│   │       └── target
│   │           ├── classes
│   │           │   └── com
│   │           │       └── gitsudo404
│   │           │           └── event_ingestion
│   │           │               ├── config
│   │           │               ├── controller
│   │           │               ├── dto
│   │           │               ├── exception
│   │           │               ├── producer
│   │           │               ├── ratelimit
│   │           │               ├── service
│   │           │               └── validation
│   │           └── test-classes
│   │               └── com
│   │                   └── gitsudo404
│   │                       └── event_ingestion
│   ├── HELP.md
│   ├── mvnw
│   ├── mvnw.cmd
│   ├── pom.xml
│   ├── src
│   │   ├── main
│   │   │   ├── java
│   │   │   │   └── com
│   │   │   │       └── gitsudo404
│   │   │   │           └── event_ingestion
│   │   │   │               ├── config
│   │   │   │               ├── controller
│   │   │   │               │   └── VehicleEventController.java
│   │   │   │               ├── dto
│   │   │   │               │   ├── EventMetaData.java
│   │   │   │               │   ├── VehicleBatteryEventRequest.java
│   │   │   │               │   ├── VehicleDiagnosticEventRequest.java
│   │   │   │               │   ├── VehicleEngineTelemetryEventRequest.java
│   │   │   │               │   ├── VehicleFuelEventRequest.java
│   │   │   │               │   └── VehicleGPSEventRequest.java
│   │   │   │               ├── EventIngestionApplication.java
│   │   │   │               ├── exception
│   │   │   │               ├── producer
│   │   │   │               ├── ratelimit
│   │   │   │               ├── service
│   │   │   │               └── validation
│   │   │   └── resources
│   │   │       ├── application.properties
│   │   │       ├── static
│   │   │       └── templates
│   │   └── test
│   │       └── java
│   │           └── com
│   │               └── gitsudo404
│   │                   └── event_ingestion
│   │                       └── EventIngestionApplicationTests.java
│   └── target
│       ├── classes
│       │   ├── application.properties
│       │   └── com
│       │       └── gitsudo404
│       │           └── event_ingestion
│       │               ├── config
│       │               ├── controller
│       │               │   └── VehicleEventController.class
│       │               ├── dto
│       │               │   ├── EventMetaData.class
│       │               │   ├── VehicleBatteryEventRequest.class
│       │               │   ├── VehicleDiagnosticEventRequest.class
│       │               │   ├── VehicleEngineTelemetryEventRequest.class
│       │               │   ├── VehicleFuelEventRequest.class
│       │               │   └── VehicleGPSEventRequest.class
│       │               ├── EventIngestionApplication.class
│       │               ├── exception
│       │               ├── producer
│       │               ├── ratelimit
│       │               ├── service
│       │               └── validation
│       └── test-classes
│           └── com
│               └── gitsudo404
│                   └── event_ingestion
│                       └── EventIngestionApplicationTests.class
├── event-processing
│   ├── HELP.md
│   ├── mvnw
│   ├── mvnw.cmd
│   ├── pom.xml
│   └── src
│       ├── main
│       │   ├── java
│       │   │   └── com
│       │   │       └── gitsudo404
│       │   │           └── event_processing
│       │   │               └── EventProcessingApplication.java
│       │   └── resources
│       │       └── application.properties
│       └── test
│           └── java
│               └── com
│                   └── gitsudo404
│                       └── event_processing
│                           └── EventProcessingApplicationTests.java
├── generator
│   ├── cmd
│   │   └── main.go
│   ├── config.yaml
│   ├── go.mod
│   ├── go.sum
│   └── internal
│       ├── config
│       │   └── config.go
│       ├── events
│       │   ├── applicationLogEvent.go
│       │   ├── clickstreamEvent.go
│       │   ├── eventMetaData.go
│       │   ├── iotEvents.go
│       │   ├── systemMetricEvents.go
│       │   ├── transactionEvents.go
│       │   ├── vehicleBatteryEvent.go
│       │   ├── vehicleDiagnosticEvent.go
│       │   ├── vehicleEngineTelemetryEvent.go
│       │   ├── vehicleFuelEvent.go
│       │   └── vehicleGPSEvent.go
│       └── generators
│           ├── applicationLogEventGenerator.go
│           ├── generator.go
│           ├── vehicleEventGenerator.go
│           └── vehicleState.go
├── infrastructure
│   └── DB
│       └── clickhouse
│           └── DOCKERFILE
├── LICENSE
└── README.md


```
