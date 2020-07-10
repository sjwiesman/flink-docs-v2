---
title: "Konfiguration"
draft: falsch
weight: 2
---

Stateful Functions enthält einige SDK-spezifische Konfigurationen. Diese werden über den `flink-conf.yaml` Ihres Jobs konfiguriert.

| Schlüssel                           | Standard              | Typ                       | Beschreibung                                                                                                                             |
| ----------------------------------- | --------------------- | ------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| statefun.module.global-config.<KEY> | (keine)               | String                    | Fügt das angegebene Schlüssel/Wert-Paar zur globalen Konfiguration Stateful Functions hinzu.                                             |
| statefun.message.serializer         | Mit PROTOBUF_PAYLOADS | Nachrichten-Serialisierer | er serializer für die Drahtmeldungen. Optionen sind WITH_PROTOBUF_PAYLOADS, MITH_KRYO_PAYLOADS, MITH_RAW_PAYLOADS.                 |
| statefun.flink-Job-Name             | Stateful-Funktionen   | String                    | Der Name der im Flink-UI angezeigt wird.                                                                                                 |
| statefun.feedback.memory.size       | 32 MB                 | Speicher                  | Die Anzahl der Bytes, die für die Speichernutzung des Feedback-Kanals verwendet werden sollen, bevor auf die Festplatte gesprungen wird. |

