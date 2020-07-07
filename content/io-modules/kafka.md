---
title: "Apache Kafka"
draft: false
weight: 1
---

Stateful Functions offers an Apache Kafka I/O Module for reading from and writing to Kafka topics.
It is based on Apache Flink's universal [Kafka connector](https://ci.apache.org/projects/flink/flink-docs-stable/dev/connectors/kafka.html) and provides exactly-once processing semantics.
The Kafka I/O Module is configurable in Yaml or Java.

## Dependency
{{< tabs >}}
{{% tab name="remote module" %}}
When configured as a remote module, there are no additional dependencies to add to your project.
{{% /tab %}}
{{% tab name="embedded module" %}}
To use the Kafka I/O Module in Java, please include the following dependency in your pom.

```bash
<dependency>
	<groupId>org.apache.flink</groupId>
	<artifactId>statefun-kafka-io</artifactId>
	<version>{{ site.version }}</version>
	<scope>provided</scope>
</dependency>
```
{{% /tab %}}
{{< /tabs >}}

## Kafka Ingress Spec

A ``KafkaIngressSpec`` declares an ingress spec for consuming from Kafka cluster.

It accepts the following arguments:

1. The ingress identifier associated with this ingress
2. The topic name / list of topic names
3. The address of the bootstrap servers
4. The consumer group id to use
5. A ``KafkaIngressDeserializer`` for deserializing data from Kafka (Java only)
6. The position to start consuming from

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
version: "1.0"

module:
  meta:
  type: remote
spec:
  ingresses:
    - ingress:
      meta:
        type: statefun.kafka.io/routable-protobuf-ingress
        id: example/user-ingress
      spec:
        address: kafka-broker:9092
        consumerGroupId: routable-kafka-e2e
        startupPosition:
          type: earliest
          topics:
            - topic: messages-1
              typeUrl: org.apache.flink.statefun.docs.models.User
              targets:
                - example-namespace/my-function-1
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
  IngressIdentifier<User> ID =
      new IngressIdentifier<>(User.class, "example", "input-ingress");

  IngressSpec<User> kafkaIngress =
      KafkaIngressBuilder.forIdentifier(ID)
          .withKafkaAddress("localhost:9092")
          .withConsumerGroupId("greetings")
          .withTopic("my-topic")
          .withDeserializer(UserDeserializer.class)
          .withStartupPosition(KafkaIngressStartupPosition.fromLatest())
          .build();
```
{{% /tab %}}
{{< /tabs >}}

The ingress also accepts properties to directly configure the Kafka client, using ``KafkaIngressBuilder#withProperties(Properties)``.
Please refer to the Kafka [consumer configuration](https://docs.confluent.io/current/installation/configuration/consumer-configs.html) documentation for the full list of available properties.
Note that configuration passed using named methods, such as ``KafkaIngressBuilder#withConsumerGroupId(String)``, will have higher precedence and overwrite their respective settings in the provided properties.

### Startup Position

The ingress allows configuring the startup position to be one of the following:

#### From Group Offset (default)

Starts from offsets that were committed to Kafka for the specified consumer group.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
startupPosition:
    type: group-offsets
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
KafkaIngressStartupPosition#fromGroupOffsets();
```
{{% /tab %}}
{{< /tabs >}}

#### Earlist

Starts from the earliest offset.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
startupPosition:
    type: earliest
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
KafkaIngressStartupPosition#fromEarliest();
```
{{% /tab %}}
{{< /tabs >}}

#### Latest

Starts from the latest offset.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
startupPosition:
    type: latest
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
KafkaIngressStartupPosition#fromLatest();
```
{{% /tab %}}
{{< /tabs >}}
