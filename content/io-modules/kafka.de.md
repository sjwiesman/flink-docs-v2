---
title: "Apache Kafka"
draft: falsch
weight: 1
---

Stateful Functions bietet ein Apache Kafka I/O Modul zum Lesen und Schreiben von Kafka Themen. Er basiert auf dem universellen [Kafka-Anschluss von Apache Flink](https://ci.apache.org/projects/flink/flink-docs-stable/dev/connectors/kafka.html) und stellt genau einmal Verarbeitungssemantik zur Verfügung. Das Kafka I/O Modul ist in Yaml oder Java konfigurierbar.

## Abhängigkeit
{{< tabs >}}
{{% tab name="remote module" %}}
Wenn Sie als Remote-Modul konfiguriert sind, gibt es keine zusätzlichen Abhängigkeiten zum Projekt hinzuzufügen.
{{% /tab %}}
{{% tab name="embedded module" %}}
Um das I/O-Modul von Kafka in Java zu verwenden, fügen Sie bitte die folgende Abhängigkeit in Ihren Pom ein.

```bash
<dependency>
    <groupId>org.apache.flink</groupId>
    <artifactId>statefun-kafka-io</artifactId>
    <version>{{ site.version }}</version>
    <scope>bereitgestellt</scope>
</dependency>
```
{{% /tab %}}
{{< /tabs >}}

## Kafka Ingres Spec

Ein `KafkaIngressSpec` erklärt eine ingress-Spec zum Verzehr aus Kafka-Cluster.

Es akzeptiert die folgenden Argumente:

1. Die ingress-Kennung, die dieser ingress zugeordnet ist
2. Der Themenname / Liste der Themennamen
3. Die Adresse der Bootstrap-Server
4. Die zu verwendende Verbrauchergruppe Id
5. Ein `KafkaIngressDeserializer` für die Deserialisierung von Daten aus Kafka (nur Java)
6. Die Position, von der verbraucht werden soll

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
Version: "1. "

Modul:
  meta:
  type: remote
spec:
  ingress:
    - ingress:
      meta:
        type: statefun. afka. o/routable-protobuf-ingress
        id: beispiel/user-ingress
      spec:
        address: kafka-broker:9092
        consumerGroupId: routable-kafka-e2e
        startupPosition:
          type: früheste
          topics:
            - topic: messages-1
              typeUrl: org. pache.flink.statefun.docs.models. ser
              Ziele:
                - beispiel-namespace/my-function-1
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
  IngressIdentifier<User> ID =
      new IngressIdentifier<>(User. lass, "Beispiel", "input-ingress");

  IngressSpec<User> kafkaIngress =
      KafkaIngressBuilder. orIdentifier(ID)
          . ithKafkaAddress("localhost:9092")
          . ithConsumerGroupId("Grüße")
          .withTopic("my-topic")
          . ithDeserializer(UserDeserializer.class)
          .withStartupPosition(KafkaIngressStartupPosition.fromLatest())
          .build();
```
{{% /tab %}}
{{< /tabs >}}

Die ingress akzeptiert auch Eigenschaften, um den Kafka-Client direkt zu konfigurieren, mit `KafkaIngressBuilder#withProperties(Properties)`. Die vollständige Liste der verfügbaren Eigenschaften finden Sie in der Kafka [Verbraucherkonfiguration](https://docs.confluent.io/current/installation/configuration/consumer-configs.html) Dokumentation. Beachten Sie, dass die Konfiguration mit benannten Methoden wie `KafkaIngressBuilder#withConsumerGroupId(String)`übergeben wurde, hat eine höhere Priorität und überschreibt ihre jeweiligen Einstellungen in den angegebenen Eigenschaften.

### Startup Position

Die ingress erlaubt es, die Startposition als eine der folgenden zu konfigurieren:

#### Von Gruppen-Versatz (Standard)

Beginnt von Ausgleichsaktionen, die Kafka für die bestimmte Verbrauchergruppe verpflichtet waren.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
startupPosition:
    Typ: Gruppen-Offsets
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
KafkaIngressStartupPosition#fromGroupOffsets();
```
{{% /tab %}}
{{< /tabs >}}

#### Ohrenliste

Beginnt ab dem frühesten Offset.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
startupPosition:
    Typ: früheste
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
KafkaIngressStartupPosition#fromEarliest();
```
{{% /tab %}}
{{< /tabs >}}

#### Neueste

Beginnt mit dem neuesten Offset.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
startupPosition:
    Typ: neueste
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
KafkaIngressStartupPosition#fromNeuest();
```
{{% /tab %}}
{{< /tabs >}}
