---
title: "Java-Durchgang"
draft: falsch
weight: 2
---

Wie alle großen Einführungen in der Software wird auch dieser Weg am Anfang beginnen: Hallo sagen. Die Anwendung wird eine einfache Funktion ausführen, die eine Anfrage akzeptiert und mit einem Gruß reagiert. Sie wird nicht versuchen, alle Komplexitäten der Anwendungsentwicklung abzudecken. aber konzentrieren Sie sich stattdessen auf den Aufbau einer zustandsfähigen Funktion — wo Sie Ihre Geschäftslogik umsetzen.
## Ein einfaches Hallo

Begrüßungsaktionen werden durch Verbrauchs-, Routing- und Übergabe-Nachrichten ausgelöst, die mittels ProtoBuf definiert werden.

```protobuf
syntax = "proto3";

Nachricht GreetRequest {
    string who = 1;
}

Nachricht GreetResponse {
    string who = 1;
    String Gruß = 2;
}
```

Unter der Haube werden Nachrichten mit [zugewiesenen Funktionen](/sdk/java/)verarbeitet, per Definition jede Klasse, die die `StatefulFunction` Schnittstelle implementiert.

```java
Paket org.apache.flink.statefun.examples.greeter;

import org.apache.flink.statefun.sdk.Context;
import org.apache.flink.statefun.sdk. tatefulFunction;

public final class GreetFunction implementiert StatefulFunction {

    @Override
    public void invoke(Context context Objekteingabe) {
        GreetRequest greetMessage = (GreetRequest) eingehen;

        GreetResponse Antwort = GreetResponse. ewBuilder()
            .setWho(greetMessage. etWho())
            .setGreeting("Hallo " + greetMessage. etWho())
            .build();

        Kontext. end(greetingConstants.GREETING_EGRESS_ID, antworten);
    }
}
```

Diese Funktion nimmt eine Anfrage auf und sendet eine Antwort an ein externes System (oder [Ei](/io-module/#egress)). Das ist zwar schön, zeigt aber nicht die wirkliche Macht der staatlichen Funktionen: die Handhabung Zustand.

## Ein Stateful Hallo

Angenommen, Sie wollen für jeden Benutzer eine personalisierte Antwort generieren, je nachdem, wie oft er eine Anfrage gesendet hat.

```java
privater statischer String greetText(String-Name, int seen) {
    switch (seen) {
        case 0:
            return String. ormat("Hallo %s ! , Name);
        Fall 1:
            gibt String zurück. ormat("Hallo wieder %s ! , Name);
        Fall 2:
            gibt String zurück. ormat("drittes Mal der Charme! %s! , Name);
        Fall 3:
            gibt String zurück. ormat("Viel Glück, Sie wieder zu sehen %s ! , Name);
        Standard:
            gibt die Zeichenkette zurück. ormat("Hallo zum %d-ten Mal %s", gesehen + 1, name);
}
```

## Routing-Nachrichten

Um einem Benutzer eine personalisierte Grüße zu senden, muss das System den Überblick behalten, wie oft es bisher jeden Benutzer gesehen hat. Allgemein gesprochen die einfachste Lösung wäre, für jeden Benutzer eine Funktion zu erstellen und die Anzahl der bereits gesehenen Zeiten unabhängig zu verfolgen. Mit den meisten Frameworks wäre dies unerschwinglich teuer. Zugewiesene Funktionen sind jedoch virtuell und verbrauchen keine CPU oder Speicher, wenn sie nicht aktiv aufgerufen werden. Das bedeutet, dass Ihre Anwendung so viele Funktionen wie nötig erstellen kann – in diesem Fall Nutzer –, ohne sich um den Ressourcenverbrauch zu kümmern.

Wann immer Daten von einem externen System verbraucht werden (oder [ingress](/io-module/#ingress)), wird auf der Basis eines gegebenen Funktionstyps und eines Identifikators zu einer bestimmten Funktion weitergeleitet. Der Funktionstyp repräsentiert die Klasse der Funktion, die aufgerufen werden soll, wie die Greeter-Funktion, während der Identifikator (`GreetRequest#getWho`) den Aufruf auf eine bestimmte virtuelle Instanz basierend auf einem Schlüssel umfasst.

```java
Paket org.apache.flink.statefun.examples.greeter;

import org.apache.flink.statefun.examples.kafka.generated.GreetRequest;
import org.apache.flink.statefun.sdk.io. Externe;

letzte Klasse GreetRouter implementiert Router<GreetRequest> {

    @Override
    public void route(GreetRequest message, Downstream<GreetRequest> Downstream) {
        Downstream. orward(greetingConstants.GREETER_FUNCTION_TYPE, message.getWho(), message);
    }
}
```

Wenn also eine Nachricht für einen Benutzer namens John kommt, wird sie an Johns spezielle Greeter-Funktion verschickt. Falls es eine folgende Nachricht für einen Benutzer namens Jane gibt, wird eine neue Instanz der Greeter Funktion erzeugt.

## Dauerhaftigkeit

[Dauerhafter Wert](/sdk/#persistence) ist ein spezieller Datentyp, der es zustandsfähigen Funktionen ermöglicht, den fehlertoleranten Zustand ihrer Bezeichner zu erhalten, so dass jede Instanz einer Funktion den Status unabhängig verfolgen kann. Um Informationen über mehrere Grußnachrichten zu „speichern“ müssen Sie dann ein Feld mit dem Wert "Weiteres Wert" (`Anzahl`) der Greet Funktion zuordnen. Für jeden Benutzer können Funktionen nun verfolgen, wie oft sie gesehen wurden.

```java
Paket org.apache.flink.statefun.examples.greeter;

import org.apache.flink.statefun.sdk.Context;
import org.apache.flink.statefun.sdk.StatefulFunction;
import org.apache.flink.statefun.sdk.annotations. ersisted;
import org.apache.flink.statefun.sdk.state. erfasster Wert;

public final class GreetFunction implementiert StatefulFunction {

    @Persisted
    private final PersistedValue<Integer> count = PersistedValue. f("count", Integer. Glas);

    @Override
    Öffentliche Leerzeichen aufrufen (Kontext-Kontext, Objekteingabe) {
        GreetRequest greetMessage = (GreetRequest) eingehen;

        GreetResponse Antwort = computePersonalizedGreeting(greetMessage);

        Kontext. end(Grüße Konstanten. REETING_EGRESS_ID, Antwort);
    }

    private GreetResponse computePersonalizedGreeting(GreetRequest greetMessage) {
        final String name = greetMessage. etWho();
        final int seen = count.getOrDefault(0);
        count. et(seen + 1);

        String greeting = greetText(name, seen);

        greetResponse zurückgeben. ewBuilder()
            . etWho(name)
            . eting(grüßen)
            . uild();
    }
}
```

Jedes Mal, wenn eine Nachricht verarbeitet wird, berechnet die Funktion eine personalisierte Nachricht für diesen Benutzer. Es liest und aktualisiert die Anzahl der Male, die der Benutzer gesehen hat, und sendet eine Begrüßung an den egress.

Sie können den vollständigen Code für die Anwendung in diesem Durchgang [hier einsehen]({{ site.github_url }}/tree/{{ site.github_branch }}/statefun-examples/statefun-greeter-example). Werfen Sie einen Blick auf das Modul `GreetingModule`, das ist der Haupteinstiegspunkt für die vollständige Anwendung, um zu sehen, wie alles zusammengesetzt wird. Sie können dieses Beispiel lokal mit dem angegebenen Docker Setup ausführen.

```bash
$ Docker-Compound Build 
$ Docker-Compound hoch
```

Dann senden Sie einige Nachrichten zum Thema "Namen" und beobachten, was aus "Grüße".

```bash
$ docker-compose exec kafka-broker kafka-console-producer. h \
    --broker-list localhost:9092 \
    --topic names

docker-compose exec kafka-broker kafka-console-consumer. h \
     --bootstrap-server localhost:9092 \
     --isolation-level read_committed \
     --from-beginning \
     --topic Grüße
```

![grüner](/fig/greeter-function.gif)

## Willst du gehen <unk> ?

Dieser Greeter vergisst niemals einen Benutzer. Versuchen Sie die Funktion so zu ändern, dass sie die `-Anzahl` für jeden Benutzer zurücksetzt, der mehr als 60 Sekunden verbringt, ohne mit dem System zu interagieren.