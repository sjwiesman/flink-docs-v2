---
title: "Metriken"
draft: falsch
weight: 3
---

Stateful Funktionen enthalten eine Reihe von SDK-spezifischen Metriken. Neben den [Standardmetrischen Skopen](https://ci.apache.org/projects/flink/flink-docs-stable/monitoring/metrics.html#system-scope)unterstützt Stateful Functions `Funktionsumfang` , welche eine Ebene unterhalb des Operatorbereichs liegt.

#### metrics.scope.function

* Standard: &lt;host&gt;.taskmanager.&lt;tm_id&gt;.&lt;job_name&gt;.&lt;operator_name&gt;.&lt;subtask_index&gt;.&lt;function_namespace&gt;.&lt;function_name&gt;
* Wird auf alle Metriken angewendet, die auf eine Funktion gespeichert wurden.

| Metriken              | Bereich  | Beschreibung                                                                                                       | Typ    |
| --------------------- | -------- | ------------------------------------------------------------------------------------------------------------------ | ------ |
| in                    | Funktion | Die Anzahl der eingehenden Nachrichten.                                                                            | Zähler |
| inBewerten            | Funktion | Die durchschnittliche Anzahl eingehender Nachrichten pro Sekunde.                                                  | Meter  |
| außerhalb lokal       | Funktion | Die Anzahl der Nachrichten, die an eine Funktion im gleichen Aufgabenfeld gesendet werden.                         | Zähler |
| außerhalb localRate   | Funktion | Die durchschnittliche Anzahl der Nachrichten, die an eine Funktion im selben Aufgabenfeld gesendet werden.         | Meter  |
| außerhalb             | Funktion | Die Anzahl der Nachrichten, die an eine Funktion auf einem anderen Aufgabenfeld gesendet werden.                   | Zähler |
| ausverkaufte Rate     | Funktion | Die durchschnittliche Anzahl der Nachrichten, die an eine Funktion auf einem anderen Aufgabenfeld gesendet werden. | Meter  |
| äußere Eier           | Funktion | Die Anzahl der Nachrichten, die an einen egres gesendet werden.                                                    | Zähler |
| feedback.produced     | Operator | Die Anzahl der Nachrichten aus dem Feedback-Kanal.                                                                 | Meter  |
| feedback.producedRate | Operator | Die durchschnittliche Anzahl der Nachrichten, die vom Feedback-Kanal pro Sekunde gelesen werden.                   | Meter  |
