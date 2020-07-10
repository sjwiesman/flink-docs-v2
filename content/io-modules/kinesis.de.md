---
title: "AWS Kinesis"
draft: falsch
weight: 2
---



Stateful Functions bietet ein AWS Kinesis I/O Modul zum Lesen und Schreiben von Kinesis Streams. Es basiert auf dem [Kinesis Connector](https://ci.apache.org/projects/flink/flink-docs-release-1.10/dev/connectors/kinesis.html) des Apache Flinks. Das I/O-Modul Kinesis ist in Yaml oder Java konfigurierbar.
## Abhängigkeit

Um das I/O-Modul von Kinesis in Java zu verwenden, fügen Sie bitte die folgende Abhängigkeit in Ihren Pom ein.

```xml
<dependency>
    <groupId>org.apache.flink</groupId>
    <artifactId>statefun-kinesis-io</artifactId>
    <version>{{< version >}}</version>
    <scope>bereitgestellt</scope>
</dependency>
```

## Kinesis Ingress Spec

Ein `KinesisIngressSpec` erklärt eine ingress-Spec für den Verbrauch aus Kinesis-Stream.

Es akzeptiert die folgenden Argumente:

1. Die AWS-Region
2. Ein AWS-Anmeldedatenanbieter
3. Ein `KinesisIngressDeserializer` für die Deserialisierung von Daten aus Kinesis (nur Java)
4. Die Startposition des Streams
5. Eigenschaften für den Kinesis-Client
6. Der Name des zu verbrauchenden Streams

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
          type: statefun. inesis. o/routable-protobuf-ingress
          id: example-namespace/messages
        spec:
          awsRegion:
            type: specific
            id: us-west-1
          awsCredentials:
            type: basic
            accessKeyId: my_access_key_id
            secretAccessKey: my_secret_access_key
          startupPosition:
            type: earliest
          streams:
            - stream-1
              typeUrl: com. oogleapis/org.apache.flink.statefun.docs.models. ser
              Ziele:
                - beispiel-namespace/my-function-1
                - beispiel-namespace/my-function-2
            - stream: stream-2
              typeUrl: com. oogleapis/org.apache.flink.statefun.docs.models. ser
              Ziele:
                - beispiel-namespace/my-function-1
          clientConfigProperties:
            - SocketTimeout: 9999
            - MaxConnections: 15
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
  IngressIdentifier<User> ID =
      new IngressIdentifier<>(User. lass, "Beispiel", "input-ingress");

  IngressSpec<User> kinesisIngress =
      KinesisIngressBuilder. orIdentifier(ID)
          .withAwsRegion("us-west-1")
          .withAwsCredentials(AwsCredentials. romDefaultProviderChain())
          .withDeserializer(UserDeserializer. Glas)
          .withStream("Stream-Name")
          . ithStartupPosition(KinesisIngressStartupPosition.fromEarliest())
          . ithClientConfigurationProperty("key", "value")
          .build();
```
{{% /tab %}}
{{< /tabs >}}

Die ingress akzeptiert auch Eigenschaften, um den Kinesis-Client direkt zu konfigurieren, mit `KinesisIngressBuilder#withClientConfigurationProperty()`. Please refer to the Kinesis [client configuration](https://docs.aws.amazon.com/AWSJavaSDK/latest/javadoc/com/amazonaws/ClientConfiguration.html) documentation for the full list of available properties. Beachten Sie, dass die mit benannten Methoden übergebene Konfiguration eine höhere Priorität hat und ihre jeweiligen Einstellungen in den angegebenen Eigenschaften überschreibt.

### Startup Position

Die ingress erlaubt es, die Startposition als eine der folgenden zu konfigurieren:

#### Neueste (Standard)

Beginnen Sie mit dem Verbrauch von der neuesten Position, also dem Kopf der Stream-Scherben.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
startupPosition:
    Typ: neueste
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
KinesisIngressStartupPosition#fromNeuest();
```
{{% /tab %}}
{{< /tabs >}}

#### Früher

Beginnen Sie mit dem Verbrauch von der frühestmöglichen Position.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
startupPosition:
    Typ: früheste
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
KinesisIngressStartupPosition#fromEarliest();
```
{{% /tab %}}
{{< /tabs >}}

#### Datum

Beginnt von Ausgleichsständen, die eine Einnahmezeit größer oder gleich einem bestimmten Datum haben.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
startupPosition:
    Typ: Datum
    Datum: 2020-02-01 04:15:00.00 Z
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
KinesisIngressStartupPosition#fromDate(ZonedDateTime.now());
```
{{% /tab %}}
{{< /tabs >}}

### Kinesis Deserializer (nur Java)

Die Kinesis ingress muss wissen, wie man die Binärdaten in Kinesis in Java-Objekte verwandeln kann. Der `KinesisIngressDeserializer` erlaubt Benutzern ein solches Schema anzugeben. Die `T deserialize(IngressRecord ingressRecord)` Methode wird für jeden Kinesis-Eintrag aufgerufen und die Binärdaten und Metadaten von Kinesis übergeben.

```java
Paket org.apache.flink.statefun.docs.io.kinesis;

import com.fasterxml.jackson.databasind.ObjectMapper;
import java.io.IOException;
import org.apache.flink.statefun.docs.models.User;
import org.apache.flink.statefun.sdk.kinesis.ingress.IngressRecord;
import org. pache.flink.statefun.sdk.kinesis.ingress.KinesisIngressDeserializer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class UserDeserializer implementiert KinesisIngressDeserializer<User> {

  private static Logger LOG = LoggerFactory. etLogger(UserDeserializer. Glas);

  private letzte ObjectMapper Zuordnung = new ObjectMapper();

  @Override
  public User deserialize(IngressRecord ingressRecord) {
    try {
      return mapper. eadValue(ingressRecord.getData(), User.class);
    } catch (IOException e) {
      LOG. ebug("Fehler beim Deserialisieren des Datensatzs", e);
      return null;
    }
  }
}
```

## Kinesis Egress Spec

Ein `KinesisEgressBuilder` erklärt eine Eierspeiche, um Daten in einen Kinesis-Stream zu schreiben.

Es akzeptiert die folgenden Argumente:

1. Die Eierkennung, die dieser Eizelle zugeordnet ist
2. Der AWS-Anmeldedatenanbieter
3. Ein `KinesisEgressSerializer` für die Serialisierung von Daten in Kinesis (nur Java)
4. Die AWS-Region
5. Eigenschaften für den Kinesis-Client
6. Die Anzahl der maximal ausstehenden Datensätze bevor Rückdruck angewendet wird

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
Version: "1. "

Modul:
  meta:
    type: remote
  spec:
    egresses:
      - egress:
        meta:
          type: statefun. inesis. o/generic-egress
          id: example/output-messages
        spec:
          awsRegion:
            type: custom-endpoint
            endpoint: https://localhost:4567
            id: us-west-1
          awsCredentials:
            type: profile
            profileName: john-doe
            profilePath: /path/to/profile/config
          maxOutstandingRecords: 9999
          clientConfigProperties:
            - ThreadingModel: POOLED
            - ThreadPoolSize: 10 10
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
    EgressIdentifier<User> ID =
      new EgressIdentifier<>("example", "output-egress", User. lass);

  EgressSpec<User> kinesisEgress =
      KinesisEgressBuilder. orIdentifier(ID)
          .withAwsCredentials(AwsCredentials. romDefaultProviderChain())
          .withAwsRegion("us-west-1")
          . ithMaxOutstandingRecords(100)
          . ithClientConfigurationProperty("key", "value")
          . ithSerializer(UserSerializer.class)
          .build();
```
{{% /tab %}}
{{< /tabs >}}

Please refer to the Kinesis [producer default configuration properties](https://github.com/awslabs/amazon-kinesis-producer/blob/master/java/amazon-kinesis-producer-sample/default_config.properties) documentation for the full list of available properties.

### Kinesis Serializer (nur Java)

Das Kinesis-Ei muss wissen, wie Java-Objekte in binäre Daten umgewandelt werden. Der `KinesisEgressSerializer` erlaubt Benutzern ein solches Schema anzugeben. Die `EgressRecord serialize(T-Wert)` Methode wird für jede Nachricht aufgerufen, so dass Benutzer einen Wert und andere Metadaten festlegen können.

```java
Paket org.apache.flink.statefun.docs.io.kinesis;

import com.fasterxml.jackson.databasind.ObjectMapper;
import java.io.IOException;
import org.apache.flink.statefun.docs.models.User;
import org.apache.flink.statefun.sdk.kinesis.egress.EgressRecord;
import org. pache.flink.statefun.sdk.kinesis.egress.KinesisEgressSerializer;
org.slf4j.Logger;
import org.slf4j. oggerFactory;

public class UserSerializer implementiert KinesisEgressSerializer<User> {

  private statische endgültige Logger LOG = LoggerFactory.getLogger(UserSerializer. lass);

  private statische String STREAM = "user-stream";

  private ObjectMapper Mapper = new ObjectMapper();

  @Override
  public EgressRecord serialize(User value) {
    try {
      return EgressRecord. ewBuilder()
          .withPartitionKey(Wert. etUserId())
          .withData(mapper. riteValueAsBytes(value))
          .withStream(STREAM)
          . uild();
    } catch (IOException e) {
      LOG. nfo("Fehler beim serializer user", e);
      return null;
    }
  }
}
```

## AWS Region

Sowohl das Kinesis-Einband als auch die Eier können auf eine bestimmte AWS-Region eingestellt werden.

#### Standard-Anbieterkette (Standard)

Konsultiert die Standard-Anbieterkette von AWS, um die AWS-Region zu ermitteln.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
awsCredentials:
    Typ: Standard
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
AwsRegion.fromDefaultProviderChain();
```
{{% /tab %}}
{{< /tabs >}}

#### Spezifisch

Bestimmt eine AWS-Region mit der einzigartigen ID der Region.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
awsCredentials:
    Typ: spezifisch
    id: us-west-1
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
AwsRegion.of("us-west-1");
```
{{% /tab %}}
{{< /tabs >}}

#### Eigener Endpunkt

Verbindet sich mit einer AWS Region über einen nicht standardmäßigen AWS Service Endpunkt. Dies wird in der Regel nur für Entwicklungs- und Testzwecke verwendet.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
awsCredentials:
    Typ: custom-endpoint
    Endpunkt: https://localhost:4567
    id: us-west-1
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
awsCredentials:
    Typ: custom-endpoint
    Endpunkt: https://localhost:4567
    id: us-west-1
```
{{% /tab %}}
{{< /tabs >}}

## AWS Zugangsdaten

Sowohl die Kinesis ingress als auch die Eier können über Standard-AWS Zugangsdaten konfiguriert werden.

#### Standard-Anbieterkette (Standard)

Konsultiert die Standard-Anbieterkette von AWS, um die Zugangsdaten von AWS zu ermitteln.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
awsCredentials:
    Typ: Standard
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
AwsCredentials.fromDefaultProviderChain();
```
{{% /tab %}}
{{< /tabs >}}

#### Einfache

Legt die AWS-Zugangsdaten direkt mit der angegebenen Zugangsschlüssel-ID und geheimen Zugangsschlüssel fest.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
awsCredentials:
    Typ: Basic
    accessKeyId: access-key-id
    secretAccessKey: Secret-Access-Key
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
AwsCredentials.basic("accessKeyId", "secretAccessKey");
```
{{% /tab %}}
{{< /tabs >}}

#### Profil

Bestimmt die AWS-Zugangsdaten mit Hilfe eines AWS-Konfigurationsprofils und des Konfigurationspfades des Profils.

{{< tabs >}}
{{% tab name="remote module" %}}
```yaml
awsCredentials:
    Typ: Basic
    profileName: profile-name
    profilePath: /path/to/profile/config
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
AwsCredentials.profile("profile-name", "/path/to/profile/config");
```
{{% /tab %}}
{{< /tabs >}}
