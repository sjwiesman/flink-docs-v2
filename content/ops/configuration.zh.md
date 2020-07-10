---
title: "Configuration"
draft: false
weight: 2
---

Stateful Functions includes a few SDK specific configurations. These are configured through your job's `flink-conf.yaml`.

| Key                                 | Default                  | Type               | Description                                                                                                                     |
| ----------------------------------- | ------------------------ | ------------------ | ------------------------------------------------------------------------------------------------------------------------------- |
| statefun.module.global-config.<KEY> | (none)                   | String             | Adds the given key/value pair to the Stateful Functions global configuration.                                                   |
| statefun.message.serializer         | WITH_PROTOBUF_PAYLOADS | Message Serializer | he serializer to use for on the wire messages. Options are WITH_PROTOBUF_PAYLOADS, WITH_KRYO_PAYLOADS, WITH_RAW_PAYLOADS. |
| statefun.flink-job-name             | StatefulFunctions        | String             | The name to display in the Flink-UI.                                                                                            |
| statefun.feedback.memory.size       | 32 MB                    | Memory             | TThe number of bytes to use for in memory buffering of the feedback channel, before spilling to disk.                           |

