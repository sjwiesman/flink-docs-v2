---
title: "Zustand Bootstrapping"
draft: falsch
weight: 4
---

Oftmals benötigen Anwendungen einen intialen Zustand, der von historischen Daten in einer Datei, Datenbank oder einem anderen System bereitgestellt wird. Weil der Status durch den Snapshoting-Mechanismus von Apache Flink verwaltet wird, für Stateful Function Anwendungen, das bedeutet, dass den intialen Zustand in einen [Savepoint](https://ci.apache.org/projects/flink/flink-docs-stable/ops/state/savepoints.html) schreibt, der verwendet werden kann, um den Job zu starten. Benutzer können den Anfangsstatus für Stateful Functions Anwendungen mittels Flink's [State Processor API](https://ci.apache.org/projects/flink/flink-docs-release-1.10/dev/libs/state_processor_api.html) und einem `StatefulFunctionSavepointCreator` bootstrapieren.

Um loszulegen, fügen Sie folgende Bibliotheken in Ihre Anwendung ein:

```xml
<dependency>
  <groupId>org.apache. link</groupId>
  <artifactId>statefun-flink-state-processor</artifactId>
  <version>{{< version >}}</version>
</dependency>
<dependency>
  <groupId>org. pache. link</groupId>
  <artifactId>flink-state-processor-api_{{< scala_version >}}</artifactId>
  <version>{{< flink_version >}}</version>
</dependency>
```

{{% notice note %}}
Der Savepoint-Ersteller unterstützt derzeit nur die Initialisierung des Zustands für Java-Module.
{{% /Nachricht %}}

## Status Bootstrap-Funktion

Eine `StateBootstrapFunction` legt fest, wie der Bootstrap-Status für eine `StatefulFunction` Instanz mit einer gegebenen Eingabe verwendet werden kann.

Jede Instanz der Bootstrap-Funktionen entspricht direkt einem `StatefulFunction` Typ. Ebenso wird jede Instanz eindeutig durch eine Adresse identifiziert, die durch den Typ und die ID der Funktion dargestellt wird, die bootstrapst wird. Jeder Zustand, der von einer Instanz der Bootstrap-Funktionen beibehalten wird, wird der entsprechenden Live-Instanz `StatefulFunction` mit der gleichen Adresse zur Verfügung stehen.

Denken Sie zum Beispiel an die folgende State bootstrap Funktion:

```java
public class MyStateBootstrapFunction implementiert StateBootstrapFunction {

    @Persisted
    private PersistedValue<MyState> state = PersistedValue.of("my-state", MyState. Glas);

    @Override
    public void bootstrap(Context context Objekteingabe) {
        Status. et(extractStateFromInput(input));
    }
}
```

Angenommen, diese Bootstrap-Funktion wurde für die Funktion Typ `MyFunctionType`bereitgestellt, und die ID der Bootstrap-Funktionsinstanz war `id-13`. Die Funktion schreibt den Status des Namens `my-state` unter Verwendung der angegebenen Bootstrap-Daten. Nach dem Wiederherstellen einer Stateful Functions Applikation aus dem Speicherpunkt, der mit dieser Bootstrap-Funktion generiert wurde, der zugewiesenen Funktionsinstanz mit Adresse `(MyFunctionType, id-13)` wird bereits unter dem Status `my-state` verfügbar sein.

## Einen Speicherpunkt erstellen

Speichernpunkte werden durch die Definition bestimmter Metadaten erzeugt, wie zum Beispiel der maximalen Parallelität und des Status-Backends. Das Standardstatus-Backend ist [RocksDB](https://ci.apache.org/projects/flink/flink-docs-stable/ops/state/state_backends.html#the-rocksdbstatebackend).

```java
int maxParallelismus = 128;
StatefulFunctionsSavepointCreator newSavepoint = neuer StatefulFunctionsSavepointCreator(maxParallelismus);
```

Jeder Eingabedatensatz wird im Savepoint-Ersteller mit einem [-Router](/io-module/#router) registriert, der jeden Datensatz auf Null oder mehr Funktionsinstanzen umleitet. Sie können dann eine beliebige Anzahl von Funktionstypen in den Savepoint-Ersteller registrieren, ähnlich wie bei der Registrierung von Funktionen in einem zustandsfähigen Funktionsmodul. Geben Sie schließlich einen Ausgangsort für den resultierenden Speicherpunkt an.

```java
// Daten aus einer Datei, einer Datenbank oder einem anderen Ort lesen
Final ExecutionEnvironment env = ExecutionEnvironment. etExecutionEnvironment();

final DataSet<Tuple2<String, Integer>> userSeenCounts = env. romElements(
    Tuple2.of("foo", 4), Tuple2.of("bar", 3), Tuple2.of("joe", 2));

// Datensatz mit einem Router registrieren
newSavepoint. ithBootstrapData(userSeenCounts, MyStateBootstrapFunctionRouter::new);

// Erstelle eine Bootstrap-Funktion zur Verarbeitung der Datensätze
newSavepoint. ithStateBootstrapFunctionProvider(
        new FunctionType("apache", "my-function"),
        ignoriert -> new MyStateBootstrapFunction());

newSavepoint. rite("file:///savepoint/path/");

env.execute();
```

Weitere Informationen zur Verwendung von Flink's `DataSet` API finden Sie in der offiziellen [Dokumentation](https://ci.apache.org/projects/flink/flink-docs-stable/dev/batch/).

## Einsatz

Nach dem Erstellen eines neuen Savpepoints kann der Startzustand für eine Stateful Functions Anwendung verwendet werden.

{{< tabs >}}
{{% tab name="image deployment" %}}
Beim Deployment basierend auf einem Bild übergeben Sie den `-s` Befehl an das Flink [JobManager](https://ci.apache.org/projects/flink/flink-docs-stable/concepts/glossary.html#flink-master) Bild.
```yaml
version: "2.1"
Dienste:
  master:
    image: my-statefun-application-image
    Befehl: -s file:///savepoint/path
```
{{% /tab %}}
{{% tab name="session cluster" %}}
Wenn Sie in einem Flink Session-Cluster verteilen, geben Sie das Savepoint Argument im Flink CLI an.
```bash
$ ./bin/flink run -s file:///savepoint/path stateful-functions-job.jar
```
{{% /tab %}}
{{< /tabs >}}
