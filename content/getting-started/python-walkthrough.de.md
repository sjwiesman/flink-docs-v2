---
title: "Python-Durchgang"
draft: falsch
weight: 1
---

Stateful Functions bietet eine Plattform für robuste und zeitgemäße Event-basierte Anwendungen. Es bietet eine feinkörnige Kontrolle über Staat und Zeit, was die Implementierung fortgeschrittener Systeme ermöglicht. In dieser Schritt-für-Schritt Anleitung erfahren Sie, wie Sie mit der Stateful Functions API eine Stateful Functions Applikation erstellen.

## Was baust du?

Wie alle großen Einführungen in der Software wird auch dieser Weg am Anfang beginnen: Hallo sagen. Die Anwendung wird eine einfache Funktion ausführen, die eine Anfrage akzeptiert und mit einem Gruß reagiert. Sie wird nicht versuchen, alle Komplexitäten der Anwendungsentwicklung abzudecken. aber konzentrieren Sie sich stattdessen auf den Aufbau einer zustandsfähigen Funktion — wo Sie Ihre Geschäftslogik umsetzen.

## Voraussetzungen

Dieser Durchgang setzt voraus, dass du Python kennst aber du solltest auch dann weiterverfolgen können, wenn du aus einer anderen Programmiersprache kommst.

## Hilfe, ich stehe fest!

Wenn du nicht weiterkommst, schaue dir die [Community-Support-Ressourcen](https://flink.apache.org/gettinghelp.html) an. Insbesondere Die [Benutzer-Mailingliste von Apache Flink](https://flink.apache.org/community.html#mailing-lists) ist durchweg als eines der aktivsten aller Apache-Projekte eingestuft und eine großartige Möglichkeit, schnell Hilfe zu erhalten.

## Wie man mitverfolgt

Wenn du weiterverfolgen möchtest, benötigst du einen Computer mit [Python 3](https://www.python.org/) zusammen mit [Docker](https://www.docker.com/).

{{% notice note %}}
Jeder Code-Block in diesem Gehweg darf nicht die volle umgebende Klasse enthalten, um kürzer zu werden. Der vollständige Code ist am [unten auf dieser Seite](#full-application) verfügbar.
{{% /notice %}}

Sie können eine Zip-Datei mit einem Skelett-Projekt herunterladen, indem Sie hier [klicken](/downloads/walkthrough.zip).

{{< unstable >}}
Das Projekt Stateful Functions veröffentlicht keine Snapshot-Versionen des Python-SDK auf PyPI. Bitte erwägen Sie die Verwendung einer stabilen Version dieser Anleitung.
{{< /unstable >}}

Nach dem Entpacken des Pakets finden Sie eine Reihe von Dateien. Dazu gehören Dockerfiles und Datengeneratoren, um diesen Durchgang in einer lokal integrierten Umgebung durchzuführen.

```bash
$tree Statefun-walkthrough
State fun-walkthrough
<unk> 文<unk> Dockerfile
<unk> 文<unk> Docker-compose. ml
<unk> 文<unk> Generator
<unk> <unk> <unk> <unk> Dockerfile
<unk> <unk> <unk> <unk> event-generator.py
<unk> <unk> <unk> <unk> <unk> messages_pb2. y
<unk> 文<unk> greeter
<unk> <unk> <unk> <unk> Dockerfile
<unk> <unk> <unk> <unk> greeter.py
<unk> <unk> <unk> 日<unk> Nachrichten. roto
<unk> <unk> 本<unk> messages_pb2.py
<unk> <unk> <unk> <unk> <unk> requirements.txt
<unk> 本<unk> module.yaml
```

## Mit Events beginnen

Stateful Functions ist ein ereignisgetriebenes System, daher beginnt die Entwicklung mit der Definition unserer Ereignisse. Die grünere Anwendung wird ihre Ereignisse mithilfe von [Protokollpuffern](https://developers.google.com/protocol-buffers) definieren. Wenn eine Begrüßungsanfrage für einen bestimmten Benutzer aufgenommen wird, wird sie zur entsprechenden Funktion weitergeleitet. Die Antwort wird mit einem entsprechenden Gruß zurückgegeben. Der dritte Typ, `SeenCount`, ist eine Utility-Klasse, die letztere benutzt wird, um die Anzahl der bisher gesehenen Benutzer zu verwalten.

```protobuf
syntax = "Proto3";

Paketbeispiel;

// Externe Anfrage von einem Benutzer, der begrüßt werden möchte
Nachricht GreetRequest {
    // Der Name des zu grüßenden Benutzers
    String Name = 1;
}
// Eine benutzerdefinierte Antwort an den Benutzer gesendet
Nachricht GreetResponse {
    // Der Name des Benutzers wird begrüßt
    String Name = 1;
    // Die Benutzer haben den Gruß angepasst
    String Gruß = 2;
}
// Eine interne Nachricht zum Speichern des Status
Nachricht SeenCount {
    // Die Anzahl der Male, die ein Benutzer bisher gesehen hat,
    int64 seen = 1;
}
```

## Unsere erste Funktion

Unter der Haube werden Nachrichten mit [zugewiesenen Funktionen](/sdk/python/)verarbeitet, die jede zwei Argumentfunktion ist, die an die `StatefulFunction` Laufzeit gebunden ist. Funktionen sind mit dem `@function.bind` Dekorator an die Laufzeit gebunden. Beim Binden einer Funktion wird sie mit einem Funktionstyp kommentiert. Dies ist der Name, mit dem diese Funktion beim Senden von Nachrichten referenziert wird.

Wenn Sie die Datei `greeter/greeter.py` öffnen, sollten Sie den folgenden Code sehen.

```python
von StatefulFunctions

Funktionen = StatefulFunctions()

@functions.bind("Beispiel/Greeter")
def greet(context greet_request):
    Pass
```

Eine zugewiesene Funktion benötigt zwei Argumente, einen Kontext und eine Nachricht. Der [-Kontext](/sdk/python/#context-reference) bietet Zugriff auf zustandsfähige Funktionen zur Laufzeitumgebung, wie zum Beispiel Statusverwaltung und Nachrichtenübergabe. Sie werden einige dieser Funktionen während dieses Durchlaufs erforschen.

Der andere Parameter ist die Eingabemeldung, die an diese Funktion übergeben wurde. Standardmäßig werden Nachrichten als Protobuf weitergegeben [Alle](https://developers.google.com/protocol-buffers/docs/reference/python-generated#wkt). Wenn eine Funktion nur einen bekannten Typ akzeptiert, können Sie den Nachrichtentyp mittels Python 3 Syntax überschreiben. Auf diese Weise müssen Sie die Nachricht oder Überprüfungstypen nicht entpacken.

```python
von messages_pb2 import GreetRequest
von statefun import StatefulFunctions

functions = StatefulFunctions()

@functions.bind("Beispiel/greeter")
def greet(context greet_request: GreetRequest):
    Pass
```

## Sende eine Antwort

Stateful Functions akzeptieren Nachrichten und können diese auch verschicken. Nachrichten können sowohl an andere Funktionen als auch an externe Systeme gesendet werden (oder [Ei]({/io-modules/#egress)).

Ein beliebtes externes System ist [Apache Kafka](http://kafka.apache.org/). Als ersten Schritt lassen Sie uns unsere Funktion in `grüner/greeter.py` aktualisieren, um auf jede Eingabe mit einem Gruß an ein Kafka-Thema zu reagieren.

```python
von messages_pb2 Import GreetRequest, GreetResponse
von StatefulFunctions

functions = StatefulFunctions()

@functions. ind("example/greeter")
def greet(context message: GreetRequest):
    response = GreetResponse()
    response ame = message.name
    response.greeting = "Hallo {}".format(message.name)

    egress_message = kafka_egress_record(topic="greetings", key=message ame, value=Antwort)
    context.pack_and_send_egress("Beispiel/greets", egress_message)
```
Für jede Nachricht wurde eine Antwort erstellt und an einen Kafka-Themenanruf `Grüße` mit `Namen` partitioniert. Die `egress_message` wird an ein `Ei gesendet` namens `Beispiel/Grüße`. Dieser Bezeichner verweist auf einen bestimmten Kafka-Cluster und ist bei der Bereitstellung unten konfiguriert.

## Ein Stateful Hallo

Das ist ein guter Anfang, zeigt aber nicht die wirkliche Macht staatlicher Funktionen - mit Staat zu arbeiten. Angenommen, Sie wollen für jeden Benutzer eine personalisierte Antwort generieren, je nachdem, wie oft er eine Anfrage gesendet hat.

```python
def compute_greeting(name, seen):
    """
    Berechnen Sie eine personalisierte Begrüßung, basierend auf der Anzahl der Male, die dieser @name zuvor gesehen hatte.
    """
    Vorlagen = ["", "Willkommen %s", "Schön, Sie wieder zu sehen %s", "Dritte Zeit ist ein Charme %s"]
    wenn < len(templates):
        Gruß = Vorlagen[seen] % Name
    else:
        Gruß = "Schön, Sie zum %d-nten Mal %s zu sehen" % (Gesehen, Name)

    Antwort = GreetResponse()
    response.name = name
    response.greeting = Gruß

    Rückantwort
```

Um Informationen über mehrere Grußnachrichten zu „speichern“ müssen Sie dann ein Feld mit dem Wert "persistiert" (`seen_count`) der Greet Funktion zuordnen. Für jeden Benutzer können Funktionen nun verfolgen, wie oft sie gesehen wurden.

```python
@functions.bind("example/greeter")
def greet(context, greet_message: GreetRequest):
    state = context.state('seen_count').unpack(SeenCount)
    if not state:
        state = SeenCount()
        state.seen = 1
    else:
        state.seen += 1
    context.state('seen_count').pack(state)

    response = compute_greeting(greet_request.name, state.seen)

    egress_message = kafka_egress_record(topic="greetings", key=greet_request.name, value=response)
    context.pack_and_send_egress("example/greets", egress_message)
```

Der Status `seen_count` wird immer auf den aktuellen Namen gesetzt, so dass er jeden Benutzer unabhängig verfolgen kann.

## Alles zusammen verkoppeln

Stateful Function Applikationen kommunizieren mit der Apache Flink Laufzeit mit `http`. Das Python-SDK wird mit einem `RequestReplyHandler` ausgeliefert, das automatisch Funktionsaufrufe basierend auf RESTful HTTP `POSTS` verschickt. Der `RequestReplyHandler` kann mit jedem HTTP-Framework freigelegt werden.

Ein beliebtes Python Web Framework ist [Flask](https://palletsprojects.com/p/flask/). Es kann verwendet werden, um schnell und einfach eine Anwendung der Apache Flink Laufzeit zu entlarven.

```python
aus statefun Import StatefulFunctions
von statefun import RequestReplyHandler

functions = StatefulFunctions()

@functions. ind("walkthrough/greeter")
def greeter(kontext, Meldung: GreetRequest):
    pass

handler = RequestReplyHandler(functions)

# Endpunkt bedienen

von flask import request
from flask import make_response
from flask import Flask Flask

app = Flask(__name__)

@app. oute('/statefun', methods=['POST'])
def handle():
    response_data = handler(request. ata)
    response = make_response(response_data)
    response eaders.set('Content-Type', 'application/octet-stream')
    gibt die Antwort zurück


wenn __name__ == "__main__":
    app.run()
```

## Konfiguration der Laufzeit

Die Stateful Function runtime stellt Anfragen an die greeter Funktion, indem `http` auf dem `Flask` Server aufgerufen wird. Um dies zu tun, muss er wissen, welchen Endpunkt er nutzen kann, um den Server zu erreichen. Dies ist auch ein guter Zeitpunkt, um unsere Verbindung zu den Ein- und Ausgabe Kafka Themen zu konfigurieren. Die Konfiguration befindet sich in einer Datei namens `module.yaml`.

```yaml
Version: "1. "
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
            endpoint: http://python-worker:8000/statefun
            states:
              - seen_count
            maxNumBatchRequests: 500
            timeout: 2min
    ingress:
      - ingress:
          meta:
            type: statefun. afka. o/routable-protobuf-ingress
            id: Beispiel/Namen
          spec:
            Adresse: kafka-broker:9092
            consumerGroupId: my-group-id
            topics:
              - Thema: Namen
                typeUrl: com. oogleapis/Beispiel. reetRequest
                Ziele:
                  - beispiel/greeter
    egresses:
      - egress:
          meta:
            type: statefun. afka. o/generic-egress
            id: example/greets
          spec:
            address: kafka-broker:9092
            deliverySemantic:
              type: exactly-once
              transactionTimeoutMillis: 100000
```

Diese Konfiguration macht ein paar interessante Dinge.

Die erste ist, unsere Funktion zu erklären, `Beispiel/greeter`. Es enthält den Endpunkt, zu dem es zusammen mit den Zuständen erreichbar ist, zu denen die Funktion Zugriff hat.

Die ingress ist das eingegebene Kafka-Thema, das `GreetRequest` Nachrichten an die Funktion weiterleitet. Neben grundlegenden Eigenschaften wie Brokeradresse und Verbrauchergruppe enthält sie eine Liste von Zielen. Dies sind die Funktionen, an die jede Nachricht gesendet wird.

Die Eier ist die Ausgabe Kafka Cluster. Es enthält Brokerspezifische Konfigurationen, erlaubt aber jede Nachricht zu einem beliebigen Thema weiterzuleiten.

## Einsatz

Jetzt, da die grünere Anwendung erstellt wurde, ist es Zeit zu implementieren. Der einfachste Weg, eine Stateful Function Applikation bereitzustellen, ist die Verwendung des von der Community bereitgestellten Basisbildes und das Laden des Moduls. Das Basisbild stellt die Stateful Function Laufzeit zur Verfügung, es wird das `module.yaml` verwenden, um für diesen speziellen Job zu konfigurieren. Dies kann im `Dockerfile` im Root-Verzeichnis gefunden werden.

```dockerfile
FROM flink-statefun:{{ site.version }}

RUN mkdir -p /opt/statefun/modules/greeter
ADD module.yaml /opt/statefun/modules/greeter
```

Sie können diese Anwendung nun lokal mit dem angegebenen Docker Setup ausführen.

```bash
$ Docker-Compose up -d
```

Um das Beispiel in Aktionen zu sehen, sehen Sie, was aus dem Thema `grüße`:

```bash
$ Docker-Compose Logs -f Eventgenerator
```

## Willst du gehen <unk> ?

Dieser Greeter vergisst niemals einen Benutzer. Versuchen Sie die Funktion so zu ändern, dass sie die `seen_count` für jeden Benutzer zurücksetzt, der mehr als 60 Sekunden mit dem System verbringt, ohne mit dem System zu interagieren.

## Vollständige Anwendung

```python
von messages_pb2 Import SeenCount, GreetRequest GreetResponse

von statefun Import StatefulFunctions
von statefun import RequestReplyHandler
von statefun import kafka_egress_record

functions = StatefulFunctions()

@functions. ind("beispiel/greeter")
def greet(context greet_request: greetRequest):
    state = context. tate('seen_count'). npack(SeenCount)
    wenn nicht status:
        state = SeenCount()
        status. een = 1
    else:
        status. een += 1
    Kontext. tate('seen_count').pack(state)

    response = compute_greeting(greet_request. ame, state.seen)

    egress_message = kafka_egress_record(topic="greetings", key=greet_request. ame, value=Antwort)
    Kontext. ack_and_send_egress("Beispiel/greets", egress_message)


def compute_greeting(name, gesehen):
    """
    Personalisierte Grüße berechnen basierend auf der Anzahl der Male, die dieser @name zuvor gesehen wurde.
    """
    Vorlagen = ["", "Willkommen %s", "Schön, Sie wieder zu sehen %s", "Dritte Zeit ist ein Charme %s"]
    wenn < len(templates):
        Gruß = Vorlagen[seen] % Name
    else:
        Gruß = "Schön, Sie zum %d-nten Mal %s zu sehen" % (Gesehen, Name)

    Antwort = GreetResponse()
    response.name = name
    Antwort. reeting = greeting = greeting

    return response


handler = RequestReplyHandler(functions)

#
# Serve the endpoint
#

from flask import import request
from flask import make_response
from flask import Flask

app = Flask(__name__)


@app. oute('/statefun', methods=['POST'])
def handle():
    response_data = handler(request. ata)
    response = make_response(response_data)
    response eaders.set('Content-Type', 'application/octet-stream')
    gibt Antwort


zurück, wenn __name__ == "__main__":
    app.run()
```
