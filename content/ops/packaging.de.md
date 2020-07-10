---
title: "Verpackung"
draft: falsch
weight: 1
---

Stateful Functions Anwendungen können entweder als eigenständige Anwendungen oder als Flink Jobs verpackt werden, die an einen Cluster gesendet werden können.

## Bilder

Der empfohlene Bereitstellungsmodus für Stateful Functions Anwendungen ist das Erstellen eines Docker-Bildes. Auf diese Weise muss Benutzercode keine Apache Flink Komponenten paketieren. Das angegebene Basis-Image erlaubt es Teams, ihre Anwendungen mit allen notwendigen Laufzeit-Abhängigkeiten schnell zu paketieren.

Unten ist ein Beispiel Dockerfile für das Erstellen eines Stateful Functions Bildes mit einem [Embedded-Modul](/sdk/#embedded-module) und einem [Remote-Modul](/sdk/#remote-module) für eine Anwendung namens `State fun-example`.

```dockerfile
FROM flink-statefun:{{< version >}}

RUN mkdir -p /opt/statefun/modules/statefun-example
RUN mkdir -p /opt/statefun/modules/remote

COPY target/statefun-example*jar /opt/statefun/modules/statefun-example/
COPY module.yaml /opt/statefun/modules/remote/module.yaml
```

{{< stabil >}}
{{% notice note %}}
Die Flink-Community wartet derzeit darauf, dass die offiziellen Docker-Bilder im Docker Hub veröffentlicht werden. In der Zwischenzeit Ververica hat freiwillig die Bilder von Stateful Function's über ihre öffentliche Registrierung zur Verfügung gestellt: `FROM ververica/flink-statefun:{{< version >}}` Sie können dem Status des Docker Hub-Beitrags folgen [hier](https://github.com/docker-library/official-images/pull/7749).
{{% /Nachricht %}}
{{</>}}

{{< unstable >}}
{{% notice warning %}}
Die Flink-Community veröffentlicht keine Bilder für Snapshot-Versionen. Sie können diese Version lokal erstellen, indem Sie das [Repo](https://github.com/apache/flink-statefun) klonen und den Anweisungen unter `tools/docker/README.md` folgen.
{{% /Nachricht %}}
{{</>}}


## Flink Jar

Wenn Sie Ihren Job lieber an einen existierenden Flink-Cluster senden möchten, fügen Sie einfach `statefun-flink-distribution` als Abhängigkeit zu Ihrer Anwendung hinzu.

```xml
<dependency>
    <groupId>org.apache.flink</groupId>
    <artifactId>statefun-flink-distribution</artifactId>
    <version>{{< version >}}</version>
</dependency>
```

Es enthält alle Abhängigkeiten zur Laufzeit von Stateful Functions und konfiguriert den Haupteinstieg der Anwendung.

{{% notice note %}}
Die Distribution muss in Ihrer Anwendung Fett JAR gebündelt werden, so dass sie auf Flinks [User Code Klassenlader](https://ci.apache.org/projects/flink/flink-docs-stable/monitoring/debugging_classloading.html#inverted-class-loading-and-classloader-resolution-order) ist.
{{% /Nachricht %}}

```bash
$ ./bin/flink run -c org.apache.flink.statefun.flink.core.StatefulFunctionsJob ./statefun-example.jar
```
Die folgenden Konfigurationen sind für die Ausführung der StateFun Applikation unbedingt erforderlich.

```yaml
classloader.parent-first-patterns.additional: org.apache.flink.statefun;org.apache.kafka;com.google.protobuf
```
