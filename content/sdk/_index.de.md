---
title: "SDKs"
draft: falsch
weight: 3
---

Stateful Function Applikationen bestehen aus einem oder mehreren Modulen. Ein Modul ist ein Paket von Funktionen, die von der Laufzeit geladen werden und zur Nachricht verfügbar sind. Funktionen aus allen geladenen Modulen sind multiplexiert und können sich beliebig mitteilen. Jedes Modul kann Funktionen enthalten, die mit einer anderen Sprache SDK geschrieben werden, was eine nahtlose sprachübergreifende Kommunikation ermöglicht.

Stateful Functions unterstützt zwei Arten von Modulen: Embedded und Remote.

## Remote-Modul

Remote-Module werden als externe Prozesse vom Apache Flink® runtime; im selben Container, als Sidecar oder anderen externen Standorten. Dieser Modultyp kann eine beliebige Anzahl von SDK unterstützen. Remote-Module werden über `YAML` Konfigurationsdateien mit dem System registriert.

### Spezifikation

Eine Remote-Modulkonfiguration besteht aus einem `Meta-` Abschnitt und einer `Spezifikation` Abschnitt. `meta` enthält zusätzliche Informationen über das Modul. Die `-Spezifikation` beschreibt die Funktionen des Moduls und definiert deren Dauerwerte.

### Funktionen definieren

`module.spec.functions` deklariert eine Liste der `Funktion` Objekte, die vom entfernten Modul implementiert werden. Eine `-Funktion` wird über eine Reihe von Eigenschaften beschrieben.

* `function.meta.kind`
    * Das Protokoll, das zur Kommunikation mit der Remote-Funktion verwendet wird.
    * Unterstützte Werte - `http`
* `function.meta.type`
    * Der Funktionstyp, definiert als `<namespace>/<name>`.
* `function.spec.endpoint`
    * Der Endpunkt, an dem die Funktion erreichbar ist.
    * Unterstützte Schemata sind: `http`, `https`.
    * Transport über UNIX-Domain-Sockets wird durch die Schemata `http+unix` oder `https+unix` unterstützt.
    * Bei Verwendung von UNIX-Domain-Sockets lautet das Endpunkt-Format: `http+unix://<socket-file-path>/<serve-url-path>`. Zum Beispiel `http+unix:///uds.sock/path/of/url`.
* `function.spec.states`
    * Eine Liste der in der Remote-Funktion deklarierten persistierten Werte.
    * Jeder Eintrag besteht aus einem `Namen` und einem optionalen `Ablaufdatum Nach` Eigenschaften.
    * Standard für `läuft ab nach` - 0, was bedeutet, dass das Ablaufen des Status deaktiviert ist.
* `function.spec.maxNumBatchAnfragen`
    * The maximum number of records that can be processed by a function for a particular `address` before invoking backpressure on the system.
    * Standard - 1000
* `function.spec.timeout`
    * Gibt an, wie lange die Laufzeit dauert, bis die entfernte Funktion zurückgegeben wird, bevor sie scheitert.
    * Standard - 1 min

```yaml
Version: "2. "

Modul:
  meta:
    type: remote
  spec:
    Funktionen:
      - Funktion:
        meta:
          kind: http
          type: example/greeter
        spec:
          endpoint: http://<host-name>/statefun
          states:
            - name: seen_count
              expireAfter: 5min
          maxNumBatchRequests: 500
          timeout: 2min
```

## Embedded Module

