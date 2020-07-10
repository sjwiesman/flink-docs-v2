---
title: "Connectoren flinken"
draft: falsch
weight: 3
---

Das I/O-Modul ermöglicht die Anbindung bestehender oder benutzerdefinierter Flink-Anschlüsse, die noch nicht in ein dediziertes I/O-Modul integriert sind. Details zum Erstellen eines benutzerdefinierten Konnektors finden Sie in der offiziellen [Apache Flink Dokumentation](https://ci.apache.org/projects/flink/flink-docs-stable).

## Abhängigkeit

Um einen benutzerdefinierten Flink-Konnektor zu verwenden, fügen Sie bitte die folgende Abhängigkeit in Ihren Pom.

```xml
<dependency>
    <groupId>org.apache.flink</groupId>
    <artifactId>statefun-flink-io</artifactId>
    <version>{{< version >}}</version>
    <scope>bereitgestellt</scope>
</dependency>
```

## Quell-Spec

Eine Quellfunktion erzeugt eine ingress aus einer Flink-Quellfunktion.

```java
Paket org.apache.flink.statefun.docs.io.flink;

import java.util.Map;
import org.apache.flink.statefun.docs.models.User;
import org.apache.flink.statefun.io.datastream.SourceFunctionSpec;
import org.apache.flink.statefun.sdk.io.IngressIdentifier;
import org.apache.flink.statefun.sdk.io.IngressSpec;
import org.apache.flink.statefun.sdk.spi. tatefulFunctionModule

Die öffentliche Klasse ModuleWithSourceSpec implementiert StatefulFunctionModule {

    @Override
    public void configure(Map<String, String> globalConfiguration, Binder binder) {
        IngressIdentifier<User> id = new IngressIdentifier<>(User. lasst, "Beispiel", "Benutzer");
        IngressSpec<User> spec = new SourceFunctionSpec<>(id, neue FlinkSource<>());
        Binder. indIngress(spec);
    }
}
```

## Sink Spec

Eine Spülfunktion erzeugt ein Ei aus einer Flink Spülfunktion.

```java
Paket org.apache.flink.statefun.docs.io.flink;

import java.util.Map;
import org.apache.flink.statefun.docs.models.User;
import org.apache.flink.statefun.io.datastream.SinkFunctionSpec;
import org.apache.flink.statefun.sdk.io.EgressIdentifier;
import org.ache.flink.statefun.sdk.io.EgressSpec;
import org.apache.flink.statefun.sdk.spi. tatefulFunctionModule

public class ModuleWithSinkSpec implementiert StatefulFunctionModule {

    @Override
    public void configure(Map<String, String> globalConfiguration, Binder binder) {
        EgressIdentifier<User> id = new EgressIdentifier<>("example", "user", User. Glas);
        EgressSpec<User> spec = new SinkFunctionSpec<>(id, neuer FlinkSink<>());
        Binder. indEgress(spec);
    }
}
```
