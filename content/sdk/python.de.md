---
title: "Python SDK"
draft: falsch
weight: 1
---


Stateful Funktionen sind die Bausteine der Anwendungen; sie sind atomare Einheiten der Isolation, Verteilung und Beharrlichkeit. Als Objekte kapseln sie den Zustand einer einzelnen Entität (z. B. ein bestimmter Benutzer, ein Gerät oder eine Sitzung) und kodieren ihr Verhalten. Stateful Funktionen können miteinander interagieren, und externe Systeme durch Nachrichtenübermittlung. Das Python SDK wird als [Remote-Modul](/sdk/#remote-module) unterstützt.

Um loszulegen, fügen Sie das Python-SDK als Abhängigkeit zu Ihrer Anwendung hinzu.

```bash
apache-flink-statefun=={{< version >}}
```


## Definition einer Stateful Funktion

Eine zugewiesene Funktion ist jede Funktion, die zwei Parameter benötigt, eine `Kontext` und `Nachricht`. Die Funktion ist an die Laufzeit durch den zustandsfähigen Funktionsdekorator gebunden. Das Folgende ist ein Beispiel für eine einfache Hallo Welt-Funktion.

```python
von StatefulFunctions

Funktionen = StatefulFunctions()

@functions. ind("beispiel/hallo")
def hello_function(context message):
    """A simple hallo world function"""
    user = User()
    message npack(user)

    print("Hallo " + user.name)
```

This code declares a function with in the namespace `example` and of type `hello` and binds it to the `hello_function` Python instance. Die Nachrichten sind untypisiert und durchlaufen das System als Protobuf {{< protoany >}} so kann eine Funktion möglicherweise mehrere Arten von Nachrichten verarbeiten.

Der `Kontext` bietet Metadaten über die aktuelle Nachricht und Funktion, und wie Sie andere Funktionen oder externe Systeme aufrufen können. Eine vollständige Referenz aller vom Kontextobjekt unterstützten Methoden wird am [Ende dieser Seite aufgelistet](/sdk/python.html#context-reference).

## Tippe Hinweise

Wenn die Funktion einen statischen Satz bekannter unterstützter Typen hat, können diese als [Tipphinweise](https://docs.python.org/3/library/typing.html) angegeben werden. Dies beinhaltet [Gewerkschaftstypen](https://docs.python.org/3/library/typing.html#typing.Union) für Funktionen, die mehrere Eingabetypen unterstützen.

```python
importiere
von StatefulFunctions Functions

Funktionen = StatefulFunctions()

@functions. ind("beispiel/hallo")
def hello_function(context meldung: Benutzer):
    """A simple hallo world function with typing"""

    print("Hallo " + Nachricht. ame)

@funktion. ind("example/goodbye")
def goodbye_function(context, message: typing. nion[User, Admin]):
    """Eine Funktion, die Typen""" verschickt

    wenn isinstance(message, Benutzer):
        Drucken ("Goodbye Benutzer")
    elif isinstance(Nachricht, Admin):
        Drucken ("Goodbye Admin")
```

## Funktionstypen und Nachrichten

Der Dekorator `bindet` registriert jede Funktion mit der Laufzeit unter einem Funktionstyp. Der Funktionstyp muss das Formular `<namespace>/<name>` benutzen. Funktionstypen können dann von anderen Funktionen referenziert werden, um eine Adresse und Nachricht einer bestimmten Instanz zu erstellen.

```python
von google.protobuf.any_pb2 importieren Sie beliebige
vom Zustandsimport StatefulFunctions

Funktionen = StatefulFunctions()

@functions. ind("Beispiel/Anrufer")
def caller_function(context message):
    """Eine einfache Stateful Funktion, die eine Nachricht an den Benutzer mit der ID `user1`""" schickt

    user = User()
    Benutzer. ser_id = "user1"
    user.name = "Seth"

    envelope = Any()
    envelope. ack(user)

    context.send("example/hello", user.user_id, envelope)
```

Alternativ können Funktionen manuell an die Laufzeit gebunden werden.

```python
functions.register("Beispiel/Anrufer", Anrufer_funktion)
```
## Verspätete Nachrichten senden

Funktionen sind in der Lage, Nachrichten auf eine Verzögerung zu senden, so dass sie nach einer gewissen Zeit ankommen. Funktionen können sich sogar selbst verspätete Nachrichten schicken, die als Rückruf dienen können. Die verzögerte Nachricht ist nicht-blockierend, so dass die Funktionen die Datensätze zwischen dem Sende- und Empfangszeitpunkt einer verzögerten Nachricht weiterverarbeiten. Die Verzögerung wird über eine [Python Timedelta](https://docs.python.org/3/library/datetime.html#datetime.timedelta) angegeben.

```python
von google.protobuf.any_pb2 importieren Sie beliebige
vom Zustandsimport StatefulFunctions

Funktionen = StatefulFunctions()

@functions. ind("Beispiel/Anrufer")
def caller_function(context message):
    """Eine einfache Stateful Funktion, die eine Nachricht an den Benutzer mit der ID `user1`""" schickt

    user = User()
    Benutzer. ser_id = "user1"
    user.name = "Seth"

    envelope = Any()
    envelope. ack(user)

    context.send("example/hello", user.user_id, envelope)
```

## Dauerhaftigkeit

Stateful Functions behandelt den Zustand als Bürger erster Klasse, so dass alle zugehörigen Funktionen den Zustand leicht definieren können, der durch die Laufzeit automatisch toleriert wird. Alle zugewiesenen Funktionen können Zustände enthalten, indem sie lediglich Werte im `Kontext` Objekt speichern. Die Daten werden immer auf einen bestimmten Funktionstyp und Bezeichner übertragen. Zustandswerte konnten fehlen, `Keine`, oder ein {{< protoany >}}.

**Achtung:** [Remote-Module](/sdk/#remote-module) erfordert, dass alle Statuswerte bei module.yaml eifrig registriert sind. Es wird auch das Konfigurieren anderer Status-Eigenschaften, wie zum Beispiel das Auslaufen von Zuständen, erlauben. Weitere Details finden Sie auf dieser Seite.

Unten ist eine zustandsfähige Funktion, die Benutzer grüßt, basierend auf der Anzahl der Male, die sie gesehen haben.

```python
von google.protobuf.any_pb2 importieren Sie beliebige
vom Zustandsimport StatefulFunctions

Funktionen = StatefulFunctions()

@functions. ind("beispiel/count")
def count_greeter(kontext, message):
    """Funktion, die einen Benutzer basierend auf
    grüßt, wie oft er aufgerufen wurde"""
    user = User()
    Nachricht. npack(user)


    state = context["count"]
    if state is None:
        state = Any()
        state. ack(Count(1))
        Ausgabe = generate_message(1, user)
    else:
        counter = Count()
        state. npack(counter)
        Zähler. alue += 1
        Ausgabe = generate_message(counter. alue, Benutzer)
        Status. ack(counter)

    context["count"] = state
    print(output)

def generate_message(count, user):
    if count == 1:
        return "Hallo " + user. ame
    elif count == 2:
        return "Hallo erneut!"
    elif count == 3:
        gibt "Dritter Mal den Charme"
    weiter:
        return "Hallo für die " + Anzahl + "te Mal" zurück
```

Darüber hinaus können anhaltende Werte durch Löschen des Wertes gelöscht werden.

```python
del Kontext["Zähler"]
```

## Exposing Funktionen

Das Python-SDK wird mit einem `RequestReplyHandler` ausgeliefert, das automatisch Funktionsaufrufe basierend auf RESTful HTTP `POSTS` verschickt. Der `RequestReplyHandler` kann mit jedem HTTP-Framework freigelegt werden.

```python
von Statefun Import RequestReplyHandler

Handler RequestReplyHandler(Funktionen)
```

#### Funktionen mit Flask bedienen

Ein beliebtes Python Web Framework ist [Flask](https://palletsprojects.com/p/flask/). Es kann verwendet werden, um einen `RequestResponseHandler` schnell und einfach zu belichten.

```python
@app.route('/statefun', methods=['POST'])
def handle():
    response_data = handler(request. ata)
    response = make_response(response_data)
    Antwort. eaders.set('Content-Type', 'application/octet-stream')
    gibt die Antwort zurück


wenn __name__ == "__main__":
    app.run()
```

## Kontext-Referenz

Das `Kontext` Objekt das an jede Funktion übergeben wurde, hat folgende Attribute / Methoden.

* send(Selbst, Typname: str, id: str, Nachricht: Any)
    * Send a message to any function with the function type of the form `<namesapce>/<type>` and message of type {{< protoany >}}
* pack_and_send(Selbst, Typname: str, id: str, message)
    * Das Gleiche wie oben, aber es wird die Protobuf-Nachricht in einer {{< protoany >}} packen
* antwort(selber, nachricht: any)
    * Sendet eine Nachricht an die aufrufende Funktion
* pack_and_reply(selber, Nachricht)
    * Das Gleiche wie oben, aber es wird die Protobuf-Nachricht in einer {{< protoany >}} packen
* send_after(self, delay: timedelta, typname: str, id: str, message: Any)
    * Sendet eine Nachricht nach einer Verzögerung
* pack_and_send_after(self, delay: timedelta, typename: str, id: str, message)
    * Das Gleiche wie oben, aber es wird die Protobuf-Nachricht in einer {{< protoany >}} packen
* send_egress(Selbst, Typenname, Nachricht: Any)
    * Sendet eine Nachricht an eine Eizelle mit einem Typennamen des Formulars `<namespace>/<name>`
* pack_and_send_egress(Selbst, Typenname, Nachricht)
    * Das Gleiche wie oben, aber es wird die Protobuf-Nachricht in einer {{< protoany >}} packen
* \_\_getitem\_\_(Selbst, Name)
    * Ruft das unter dem Namen {{< protoany >}} oder `Keine` registrierte Bundesland ab, wenn kein Wert gesetzt ist
* \_\_delitem\_\_(Selbst, Name)
    * Löscht das unter dem Namen registrierte Bundesland
* \_\_setitem\_\_(Selbst, Name, Wert: Any)
    * Speichert den Wert unter dem angegebenen Namen in Status.
