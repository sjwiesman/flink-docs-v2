---
title: "Konzepte"
draft: falsch
weight: 2
---

Stateful Functions bietet ein Framework für die Erstellung von Event-Drivent-Anwendungen. Hier erläutern wir wichtige Aspekte der Architektur Stateful Function.

## Event-Einzug

Stateful Function Applikationen sitzen direkt in dem ereignisgetriebenen Raum, so dass der natürliche Startpunkt darin besteht, Ereignisse in das System zu bringen.

![ingresse](/fig/concepts/statefun-app-ingress.svg)

In zustandsfähigen Funktionen nennt man die Komponente, die Datensätze in das System einfügt, Ereigniseingänge. Dies kann alles sein von einem Kafka-Thema bis hin zu einer Nachrichtenwarteschlange, an eine HTTP-Anfrage - alles, was Daten in das System einholen und die intitiellen Funktionen auslösen kann, um mit der Berechnung zu beginnen.

## Zugelassene Funktionen

Im Mittelpunkt des Diagramms stehen die namesake Stateful Funktionen.

![funktionen](/fig/concepts/statefun-app-functions.svg)

Stellen Sie sich diese als die Bausteine für Ihren Service vor. Sie können sich gegenseitig willkürlich vermitteln, was eine Möglichkeit ist, wie sich dieses Framework vom traditionellen Stream-Verarbeitungskonzept der Welt abwendet. Anstatt einen statischen Dataflow DAG aufzubauen, können diese Funktionen auf willkürlichen, potenziell zyklischen und sogar runden Strecken miteinander kommunizieren.

Wenn Sie mit Schauspieler-Programmierung vertraut sind, teilen Sie damit gewisse Ähnlichkeiten in der Fähigkeit, dynamisch zwischen den Komponenten zu kommunizieren. Es gibt jedoch eine Reihe bedeutender Unterschiede.

## Dauerhafte Zustände

Der erste ist, dass alle Funktionen lokal eingebetteten Status haben, bekannt als fortgesetzte Zustände.

![status](/fig/concepts/statefun-app-state.svg)

Eine der KernStärken von Apache Flink ist die Fähigkeit, fehlertolerante Lokalzustände bereitzustellen. Wenn Sie innerhalb einer Funktion eine Berechnung durchführen, arbeiten Sie immer mit lokalem Status in lokalen Variablen.

## Fehlertoleranz

Stateful Function's ist sowohl für Staat als auch für Nachrichten immer noch in der Lage, genau das zu bieten, was die Anwender von einem modernen Datenverarbeitungs-Framework erwarten.

![fehlerhafte Toleranz](/fig/concepts/statefun-app-fault-tolerance.svg)

Im Falle eines Scheiterns wird der gesamte Zustand der Welt (sowohl fortgesetzte Staaten als auch Nachrichten) zurückgeschraubt, um eine völlig fehlerfreie Ausführung zu simulieren.

Diese Garantien werden ohne Datenbank zur Verfügung gestellt, stattdessen nutzt Stateful Function's bewährten Snapshoting-Mechanismus von Apache Flink.

## Event-Egress

Schließlich können Applikationen Daten über Event-egresse an externe Systeme ausgeben.

![egress](/fig/concepts/statefun-app-egress.svg)
=
Natürlich führen Funktionen beliebige Berechnungen durch und können alles tun, was sie wollen. Dazu gehören auch das Ausführen von RPC-Aufrufen und das Verbinden mit anderen Systemen. Durch die Verwendung eines Event-Egress, können Anwendungen vorkompilierte Integrationen nutzen, die am oberen Rand des Apache Flink Connector-Ökosystems gebaut wurden.