---
title: "Verteilte Architektur"
draft: falsch
weight: 2
---

Ein Stateful Functions Deployment besteht aus wenigen Komponenten, die miteinander interagieren. Hier beschreiben wir diese Stücke und ihre Beziehung zueinander und die Apache Flink Laufzeit.

## Hochstufige Ansicht

Eine *Stateful Functions* Bereitstellung besteht aus einer Reihe von **Apache Flink Stateful Functions** Prozessen und, optional verschiedene Deployments, die entfernte Funktionen ausführen.

![architektur](/fig/concepts/arch_overview.svg)

Die Flink Worker Prozesse (TaskManagers) empfangen die Ereignisse von den ingress Systemen (Kafka, Kinesis, etc.) und leiten sie zu den Zielfunktionen weiter. Sie rufen die Funktionen auf und leiten die resultierenden Meldungen an die nächsten entsprechenden Zielfunktionen. Mitteilungen, die für Eier bestimmt sind, werden in ein Eiersystem geschrieben (wiederum, Kafka, Kinesis, ...).

## Komponenten

Der schwere Hebel wird von den Apache Flink Prozessen ausgeführt, die den Zustand verwalten, die Nachricht handhaben und die zugewiesenen Funktionen aufrufen. Der Flink Cluster besteht typischerweise aus einem Master und mehreren Arbeitern (TaskManager).

![komponenten](fig/concepts/arch_components.svg)

Zusätzlich zu den Apache Flink Prozessen, für eine vollständige Bereitstellung benötigt [ZooKeeper](https://zookeeper.apache.org/) (für [Master Failover](https://ci.apache.org/projects/flink/flink-docs-stable/ops/jobmanager_high_availability.html)) und Massenspeicher (S3, HDFS, NAS, GCS, Azure Blob Store, etc. um Flink's [Checkpoints](https://ci.apache.org/projects/flink/flink-docs-master/concepts/stateful-stream-processing.html#checkpointing) zu speichern. Im Gegenzug erfordert der Einsatz keine Datenbank, und Flink-Prozesse erfordern keine persistenten Volumen.

## Logische Ko-Position, physische Trennung

Ein Kernprinzip vieler Stream-Prozessoren ist, dass Anwendungslogik und Anwendungszustand gemeinsam lokalisiert werden müssen. Dieser Ansatz ist die Grundlage für ihre Widersprüchlichkeit. Stateful Function verfolgt hierzu einen einzigartigen Ansatz, indem *logisch* Zustand und Berechnung kodiert, aber *physisch* trennt.

  - *Logical Co-Location:* Messaging, State Access/Updates und Funktionsaufrufe werden wie in der DataStream API von Flink eng miteinander verwaltet. Der Status wird über den Schlüssel geteilt und Nachrichten per Schlüssel an den Staat weitergeleitet. Es gibt einen einzigen Schriftsteller pro Taste zu einer Zeit, auch die Planung der Funktionsanweisungen.

  - *Physikalische Trennung:* Funktionen können aus der Ferne ausgeführt werden, wobei der Zugriff auf Nachrichten und Zustände als Teil des Anrufs zur Verfügung gestellt wird. seine Art und Weise können Funktionen wie zustandslose Prozesse unabhängig verwaltet werden.


## Deployment Styles für Funktionen

Die zustandsfähigen Funktionen selbst können auf verschiedene Weise eingesetzt werden, die bestimmte Eigenschaften miteinander abdecken: Lose Kopplung und unabhängige Skalierung auf der einen Seite mit Performance-Overhead auf der anderen Seite. Jedes Modul von Funktionen kann von einer anderen Art sein, so dass einige Funktionen entfernt laufen können, während andere einbettet werden könnten.

#### Remote-Funktionen

*Remote-Funktionen* verwenden das oben genannte Prinzip *physikalische Trennung* unter Beibehaltung *logischer Co-Position*. Die Status/Messaging-Ebene (d.h. die Flink-Prozesse) und die Funktionsebene werden unabhängig voneinander eingesetzt, verwaltet und skaliert.

Funktionsaufrufe erfolgen über ein HTTP-/gRPC-Protokoll und gehen durch einen Dienst, der Anfragen an jeden verfügbaren Endpunkt weiterleitet zum Beispiel ein Kubernetes (load-balancing) Dienst, das AWS Request Gateway für Lambda, etc. Da die Anrufe sich selbst enthalten (Meldung, Zustand, Zugriff auf Timer usw.) können die Zielfunktionen wie jede zustandslose Anwendung behandelt werden.

![fernbedienung](/fig/concepts/arch_funs_remote.svg)


Weitere Informationen finden Sie in der Dokumentation zum [Python SDK](/sdk/python/) und [Remote-Module](/sdk/#remote-module).

#### Co-lokalisierte Funktionen

Eine alternative Möglichkeit, Funktionen bereitzustellen, ist *Co-Location* mit den Flink JVM-Prozessen. In einem solchen Setup spricht jeder Flink TaskManager mit einem Funktionsprozess, der *"neben ihnen"* sitzt. Ein gängiger Weg, dies zu tun, ist ein System wie Kubernetes und Deploy Pods bestehend aus einem Flink-Container und dem Funktions-Seitenwagencontainer; die beiden kommunizieren über das pod-lokale Netzwerk.

Dieser Modus unterstützt verschiedene Sprachen und vermeidet das Routen von Anrufen durch einen Service/LoadBalancer, aber er kann den Zustand nicht skalieren und Teile unabhängig berechnen.

![koloziert](/fig/concepts/arch_funs_colocated.svg)

Dieser Bereitstellungsstil ähnelt dem von Flink's Table API und API-Beams Portability-Layer bereitzustellen und nicht-JVM-Funktionen auszuführen.

#### Embedded Functions

*Eingebettete Funktionen* ähneln dem Ausführungsmodus von Stateful Functions 1.0 und Flinks Java/Scala Stream Processing APIs. Die Funktionen werden im JVM ausgeführt und werden direkt mit den Nachrichten und dem Zustandszugriff aufgerufen. Dies ist der performanteste Weg, allerdings nur um den Preis der Unterstützung von JVM-Sprachen. Updates für Funktionen bedeuten die Aktualisierung des Flink-Clusters.

![embedded](/fig/concepts/arch_funs_embedded.svg)

Nach der Analogie der Datenbank sind eingebettete Funktionen ein bisschen wie *Gespeicherte Prozeduren*, aber prinzipieller: Die Funktionen hier sind normale Java/Scala/Kotlin Funktionen, die Standardschnittstellen implementieren und können in jeder IDE entwickelt bzw. getestet werden.