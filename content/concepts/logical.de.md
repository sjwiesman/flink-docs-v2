---
title: "Logische Funktionen"
draft: falsch
weight: 1
---

Stateful Function's werden logisch zugewiesen, was bedeutet, dass das System eine unbegrenzte Anzahl von Instanzen mit einer begrenzten Menge an Ressourcen unterstützen kann. Logische Instanzen verwenden keine CPU, Speicher oder Threads wenn nicht aktiv aufgerufen wird, so gibt es keine theoretische Obergrenze für die Anzahl der Instanzen, die geschaffen werden können. Benutzer werden ermutigt, ihre Anwendungen so granulär wie möglich zu modellieren basierend auf dem, was für ihre Anwendung am sinnvollsten ist, anstatt Anwendungen um Ressourcenbeschränkungen herum zu gestalten.

## Funktionsadresse

In einer lokalen Umgebung ist die Adresse eines Objekts die gleiche wie eine Referenz darauf. Aber in einer Stateful Function's Applikation sind Funktionsinstanzen virtuell und deren Laufzeitort ist dem Benutzer nicht ausgesetzt. Stattdessen wird eine `-Adresse` verwendet, um eine bestimmte Zustandsfunktion im System zu referenzieren.

![adresse](/fig/concepts/address.svg)

Eine Adresse besteht aus zwei Komponenten, einem `Funktionstyp` und `ID`. Ein Funktionstyp ähnelt einer Klasse in einer objektorientierten Sprache; er legt fest, welche Funktion die Adressenreferenzen besitzen. Die ID ist ein Primärschlüssel, der den Funktionsaufruf auf eine bestimmte Instanz des Funktionstyps ausdehnt.

Wenn eine Funktion aufgerufen wird, werden alle Aktionen - einschließlich des Lesens und Schreibens von andauernden Zuständen - auf die aktuelle Adresse übertragen.

Stellen Sie sich zum Beispiel vor, es gäbe eine Stateful Function Anwendung, um das Inventar eines Lagers zu verfolgen. Eine mögliche Implementierung könnte eine `Inventory` Funktion enthalten, die die Nummerneinheiten auf Lager für einen bestimmten Artikel verfolgt; dies wäre der Funktionstyp. Es gäbe dann eine logische Instanz dieses Typs für jede SKU, die das Lager verwaltet. Wäre es Kleidung, könnte es eine Instanz für Shirts und ein anderes für Hose geben; "shirt" und "pant" wären zwei ids. Jede Instanz kann unabhängig voneinander interagiert und kommuniziert werden. Die Anwendung ist frei, so viele Instanzen wie es gibt Arten von Gegenständen im Inventar.

## Funktionslebenszyklus

Logische Funktionen werden weder erstellt noch zerstört, sondern existieren immer während der gesamten Lebensdauer einer Anwendung. Wenn eine Anwendung startet, erstellt jeder parallele Arbeiter des Frameworks ein physikalisches Objekt pro Funktionstyp. Dieses Objekt wird verwendet, um alle logischen Instanzen dieses Typs auszuführen, die von diesem bestimmten Arbeiter ausgeführt werden. Das erste Mal eine Nachricht an eine Adresse gesendet wird wird es sein, als ob diese Instanz immer existierte, wenn ihre fortgesetzten Staaten leer waren.

Die Beseitigung aller fortgesetzten Zustände einer Art ist das Gleiche wie die Vernichtung. Wenn eine Instanz keinen Status hat und nicht aktiv läuft, dann belegt sie keine CPU, keine Threads und keinen Speicher.

Eine Instanz mit Daten in einem oder mehreren ihrer anhaltenden Werte belegt nur die Ressourcen, die für die Speicherung dieser Daten notwendig sind. Zustandsspeicher wird vom Apache Flink Runtime verwaltet und im konfigurierten Zustands-Backend gespeichert.
