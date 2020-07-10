---
title: "Java SDK"
draft: falsch
weight: 2
---

Stateful Funktionen sind die Bausteine der Anwendungen; sie sind atomare Einheiten der Isolation, Verteilung und Beharrlichkeit. Als Objekte kapseln sie den Zustand einer einzelnen Entität (z. B. ein bestimmter Benutzer, ein Gerät oder eine Sitzung) und kodieren ihr Verhalten. Stateful Funktionen können miteinander interagieren, und externe Systeme durch Nachrichtenübermittlung. Das Java SDK wird als [Embedded-Modul](/sdk/#embedded-module) unterstützt.

Um loszulegen, fügen Sie das Java SDK als Abhängigkeit zu Ihrer Anwendung hinzu.

```xml
<dependency>
    <groupId>org.apache.flink</groupId>
    <artifactId>statefun-sdk</artifactId>
    <version>{{< version >}}</version>
</dependency>
```

## Definition einer Stateful Funktion

Eine zugewiesene Funktion ist jede Klasse, die die `StatefulFunction` Schnittstelle implementiert. Das Folgende ist ein Beispiel für eine einfache Hallo Welt-Funktion.

```java
Paket org.apache.flink.statefun.docs;

import org.apache.flink.statefun.sdk.Context;
import org.apache.flink.statefun.sdk. tatefulFunction;

public class FnHelloWorld implementiert StatefulFunction {

    @Override
    public void invoke(Context-Kontext, Objekteingabe) {
        System. ut.println("Hallo " + input.toString());
    }
}
```

Funktionen verarbeiten jede eingehende Nachricht über ihre `Methode und rufen` auf. Eingaben werden als `java.lang.Object` nicht eingegeben und durch das System übergeben, so dass eine Funktion möglicherweise mehrere Arten von Nachrichten verarbeiten kann.

Der `Kontext` liefert Metadaten über die aktuelle Nachricht und Funktion, und wie Sie andere Funktionen oder externe Systeme aufrufen können. Funktionen werden basierend auf einem Funktionstyp und einem eindeutigen Identifikator aufgerufen.

### Zugelassene Match-Funktion

Stateful Funktionen bieten eine mächtige Abstraktion für die Arbeit mit Events und Zuständen, die es Entwicklern ermöglichen, Komponenten zu erstellen, die auf jede Art von Nachricht reagieren können. Normalerweise müssen Funktionen nur mit bekannten Nachrichtentypen umgehen, und die `StatefulMatchFunction` Schnittstelle bietet eine empfohlene Lösung für dieses Problem.

#### Einfache Match-Funktion

Stateful Match-Funktionen sind eine angesehene Variante Stateful Funktionen für genau dieses Muster. Entwickler skizzieren erwartete Typen, optionale Prädikate und gut sortierte Geschäftslogik und lassen das System jede Eingabe in die richtige Aktion versenden. Varianten werden innerhalb einer `configure` Methode gebunden, die beim ersten Laden einer Instanz ausgeführt wird.

```java
Paket org.apache.flink.statefun.docs.match;

import org.apache.flink.statefun.sdk.Context;
import org.apache.flink.statefun.sdk.match.MatchBinder;
import org.apache.flink.statefun.sdk.match. tatefulMatchFunction;

public class FnMatchGreeter erweitert StatefulMatchFunction {

    @Override
    public void configure(MatchBinder binder) {
        binder
            . redicate(Customer.class, this::greetCustomer)
            .predicate(Employee. lass, Employee::isManager, this::greetManager)
            . redicate(Employee. lass, das::greetEmployee);
    }

    Privat ungültiger Gruß Manager(Kontext-Kontext, Mitarbeiternachricht) {
        System. ut. rintln("Hallo Manager " + Nachricht. etEmployeeId());
    }

    Private ungültige greetEmployee(Kontextkontext, Mitarbeiternachricht) {
        System. ut.println("Hallo Mitarbeiter " + Nachricht. etEmployeeId());
    }

    private ungültige grüße Kunden (Kontext-Kontext, Kundennachricht) {
        System. ut.println("Hallo Kunde " + message.getName());
    }
}
```

#### Ihre Funktion abschließen

Ähnlich wie beim ersten Beispiel übereinstimmende Funktionen sind standardmäßig teilweise und werfen `IllegalStateException` für alle Eingabe, die keinem Branch entsprechen. Sie können vollständig gemacht werden, indem eine `-Klausel, die andernfalls` als "Catch-All" für unübertroffene Eingaben dient halten Sie es als eine Standardklausel in einer Java-Schalter-Anweisung vor. Die `Aktion andernfalls` nimmt ihre Nachricht als ein untypisiertes `java.lang.Object`auf, das es dir erlaubt, unerwartete Nachrichten zu bearbeiten.

```java
Paket org.apache.flink.statefun.docs.match;

import org.apache.flink.statefun.sdk.Context;
import org.apache.flink.statefun.sdk.match.MatchBinder;
import org.apache.flink.statefun.sdk.match. tatefulMatchFunction;

public class FnMatchGreeterWithCatchAll extends StatefulMatchFunction {

    @Override
    public void configure(MatchBinder binder) {
        binder
            . redicate(Customer.class, this::greetCustomer)
            .predicate(Employee. lass, Employee::isManager, this::greetManager)
            . redicate(Employee.class, this::greetEmployee)
            . therwise(this::catchAll);
    }

    Private Leerzeichen catchAll(Kontext-Kontext, Objektnachrichtung) {
        System. ut. rintln("Hallo unerwartete Nachricht");
    }

    Private ungültige grüße Manager(Kontext-Kontext, Mitarbeiternachricht) {
        System. ut. rintln("Hallo Manager");
    }

    private void greetEmployee(Kontext-Kontext, Mitarbeiternachricht) {
        System. ut. rintln("Hallo Mitarbeiter");
    }

    Privat nichtig grüßCustomer(Kontext-Kontext, Kundennachricht) {
        System. ut.println("Hallo Kunde");
    }
}
```

#### Aktionsauflösungsbefehl

Passende Funktionen stimmen immer mit den folgenden Auflösungsregeln von den meisten bis zu den wenigsten spezifischen Aktionen überein.

Zuerst finden Sie eine Aktion, die dem Typ entspricht und dem Vorhersagen entspricht. Wenn zwei Prädikate für einen bestimmten Beitrag wahr sind, gewinnt der im Bindemittel registrierte Eintrag. Als nächstes suchen Sie nach einer Aktion, die mit dem Typ übereinstimmt, aber kein zugehöriges Prädikat hat. Schließlich, wenn ein Catch-All existiert, wird es ausgeführt oder eine `IllegalState-Ausnahme` geworfen werden.

## Funktionstypen und Nachrichten

In Java sind Funktionstypen als _strings_ definiert, die einen Namensraum und Namen enthalten. Der Typ ist an die Implementierungsklasse in der [-Modul-](/sdk/#embedded-module) Definition gebunden. Unten ist ein Beispiel-Funktionstyp für die Hallo Welt-Funktion.

```java
Paket org.apache.flink.statefun.docs;

import org.apache.flink.statefun.sdk.FunctionType;

/** Ein Funktionstyp, der an {@link FnHelloWorld} gebunden wird. */
public class Identifiers {

  public static final FunctionType HELLO_TYPE = new FunctionType("apache/flink", "hello");
}
```

Dieser Typ kann dann von anderen Funktionen referenziert werden, um eine Adresse und Nachricht einer bestimmten Instanz zu erstellen.

```java
Paket org.apache.flink.statefun.docs;

import org.apache.flink.statefun.sdk.Context;
import org.apache.flink.statefun.sdk. tatefulFunction;

/** Eine einfache Zustandsfunktion, die eine Nachricht an den Benutzer mit der ID "user1" */
öffentliche Klasse FnCaller implementiert StatefulFunction {

  @Override
  public void invoke(Context-Kontext, Objekteingabe) {
    Kontext. end(Identifiers.HELLO_TYPE, "user1", new MyUserMessage());
  }
}
```

## Verspätete Nachrichten senden

Funktionen sind in der Lage, Nachrichten auf eine Verzögerung zu senden, so dass sie nach einer gewissen Zeit ankommen. Funktionen können sich sogar selbst verspätete Nachrichten schicken, die als Rückruf dienen können. Die verzögerte Nachricht ist nicht-blockierend, so dass die Funktionen die Datensätze zwischen dem Sende- und Empfangszeitpunkt einer verzögerten Nachricht weiterverarbeiten.

```java
Paket org.apache.flink.statefun.docs.delay;

import java.time.Dauern;
import org.apache.flink.statefun.sdk.Context;
import org.apache.flink.statefun.sdk. tatefulFunction;

public class FnDelayedMessage implementiert StatefulFunction {

    @Override
    public void invoke(Context context Objekteingabe) {
        if (input instanceof Message) {
            System. ut.println("Hallo");
            context.sendAfter(Duration.ofMinutes(1), context. elf(), neue DelayedMessage());
        }

        if (input instanceof DelayedMessage) {
            System. ut.println("Willkommen in der Zukunft!");
        }
    }
}
```

## Abgeschlossene Async-Anfragen

Bei der Interaktion mit externen Systemen, wie einer Datenbank oder API, Man muss darauf achten, dass die Kommunikationsverzögerung mit dem externen System nicht die gesamte Arbeit der Anwendung beherrscht. Stateful Functions erlaubt es, eine Java `CompletableFuture` zu registrieren, die zu einem späteren Zeitpunkt zu einem Wert aufgelöst wird. Future's werden zusammen mit einem Metadaten-Objekt registriert, das einen zusätzlichen Kontext über den Anrufer bietet.

Wenn die Zukunft erfolgreich oder ausnahmsweise abgeschlossen ist, wird der Anrufer-Funktionstyp und die ID mit einem `AsyncOperationResult` aufgerufen. Ein asynchrones Ergebnis kann in einem von drei Staaten abgeschlossen werden:

### Erfolg

Die asynchrone Operation war erfolgreich, und das erzeugte Ergebnis kann über `AsyncOperationResult#value` abgerufen werden.

### Fehler

Die asynchrone Operation ist fehlgeschlagen und die Ursache kann über `AsyncOperationResult#throwable` erreicht werden.

### Unbekannt

Die zugewiesene Funktion wurde neu gestartet, möglicherweise auf einem anderen Computer, bevor die `CompletableFuture` abgeschlossen wurde Daher ist es unbekannt, was der Status der asynchronen Operation.

```java
Paket org.apache.flink.statefun.docs.async;

import java.util.concurrent.CompletableFuture;
import org.apache.flink.statefun.sdk.AsyncOperationResult;
import org.apache.flink.statefun.sdk.Context;
import org.apache.flink.statefun. dk.StatefulFunction;

@SuppressWarnings("unmarkiert")
public class EnrichmentFunction implementiert StatefulFunction {

    private Final QueryService client;

    public EnrichmentFunction(QueryService client) {
        dies. lient = Client;
    }

    @Override
    public void invoke(Context context Objekteingabe) {
        if (input instanceof User) {
            onUser(context (Benutzer) Eingabe);
        } Sonst wenn (Eingangsinstanz von AsyncOperationResult) {
            onAsyncResult((AsyncOperationResult) eingeben);
        }
    }

    private void onUser(Kontext-Kontext, Benutzer) {
        CompletableFuture<UserEnrichment> future = client. etDataAsync(user.getUserId());
        Kontext. egisterAsyncOperation(Benutzer, Zukunft);
    }

    private void onAsyncResult(AsyncOperationResult<User, UserEnrichment> result) {
        if (result). uccessful()) {
            Benutzer Metadaten = Ergebnis. etadata();
            UserEnrichment value = Resultat. alue();
            System.out. rintln(
                String. ormat("Erfolgreich die Zukunft erfolgreich abgeschlossen: %s %s", Metadaten, Wert));
        } anders wenn (Ergebnis. ailure()) {
            System.out. rintln(
                String. ormat("Es ist etwas furchtbar schief gelaufen %s", Ergebnis. hrowable());
        } else {
            System. ut.println("Nicht sicher was passiert ist, vielleicht wiederholen");
        }
    }
}
```

## Dauerhaftigkeit

Stateful Functions behandelt den Zustand als Bürger erster Klasse, so dass alle zugehörigen Funktionen den Zustand leicht definieren können, der durch die Laufzeit automatisch toleriert wird. Alle zugewiesenen Funktionen können Zustände enthalten, indem lediglich ein oder mehrere vorhandene Felder definiert werden.

Der einfachste Weg, um loszulegen ist mit einem `PersistedValue`, , der durch seinen Namen und die Klasse des Typs definiert ist, den er speichert. Die Daten werden immer auf einen bestimmten Funktionstyp und Bezeichner übertragen. Unten ist eine zustandsfähige Funktion, die Benutzer grüßt, basierend auf der Anzahl der Male, die sie gesehen haben.

<div class="alert alert-info">
  <strong>Achtung:</strong> Alle <b>Persistenten Wert</b>, <b>Persistente Tabelle</b>, und <b>PersistedAppendingBuffer</b> Felder müssen mit einer <b>@Persisted</b> Anmerkung markiert werden oder sie werden von der Laufzeit nicht tolerant gemacht.
</div>

```java
Paket org.apache.flink.statefun.docs;

import org.apache.flink.statefun.sdk.Context;
import org.apache.flink.statefun.sdk.FunctionType;
import org.apache.flink.statefun.sdk.StatefulFunction;
import org.apache.flink.statefun.sdk.annotations.Persisted;
import org.apache.flink.statefun.sdk.state. ersistedValue;

public class FnUserGreeter implementiert StatefulFunction {

    public static FunctionType TYPE = new FunctionType("example", "greeter");

    @Persistiert
    privat PersistedValue<Integer> count = PersistedValue. f("count", Integer. Glas);

    public void invoke(Kontextkontext, Object Input) {
        String userId = context. elf(). d();
        int gesehen = zählen. etOrDefault(0);

        Schalter (gesehen) {
            case 0:
                System. ut. rintln(String. ormat("Hallo %s! , userId));
                break;
            Fall 1:
                System. ut. rintln("Hallo nochmal! );
                Pausen;
            Fall 2:
                System. ut. rintln("drittes Mal ist der Charme :)");
                Pause;
            Standard:
                System. ut.println(String. ormat("Hallo zum %d-ten Mal", gesehen + 1));
        }

        zählen. et(gesehen + 1);
    }
}
```

Dauerhafter Wert kommt mit den richtigen primitiven Methoden zur Erstellung leistungsstarker Stateful Anwendungen. Wenn Sie `PersistedValue#get` aufrufen, wird der aktuelle Wert eines in Status gespeicherten Objekts zurückgegeben, oder `null` , wenn nichts gesetzt ist. Umgekehrt wird `PersistedValue#set` den Wert im Status aktualisieren und `PersistedValue#clear` wird den Wert vom Status löschen.

### Sammlungstypen

Zusammen mit `PersistedValue`unterstützt das Java SDK zwei anhaltende Sammlungstypen. `PersistedTable` ist eine Sammlung von Schlüsseln und Werten und `PersistedAppendingBuffer` ist ein Nur-Anhängselpuffer.

Diese Typen entsprechen funktionell `PersistedValue<Map>` und `PersistedValue<Collection>` , können aber in manchen Situationen zu einer besseren Performance führen.

```java
@Persisted
PersistedTable<String, Integer> Tabelle = PersistedTable.of("my-table", String.class, Integer.class);

@Persisted
PersistedAppendingBuffer<Integer> buffer = PersistedAppendingBuffer.of("my-buffer", Integer.class);
```

### Status Ablauf

Dauerhafte Zustände können so konfiguriert werden, dass sie ablaufen und nach einer bestimmten Dauer gelöscht werden. Dies wird von allen Arten von Status unterstützt:

```java
@Persistent
PersistedValue<Integer> Tabelle = PersistedWert. f(
    "my-value",
    Integer.class,
    Ablauf. xpireAfterWriting(Duration.ofHours(1)));

@PersistedTable
PersistedTable<String, Integer> table = PersistedTable.of(
    "my-table",
    String. lass,
    Integer.class,
    Expiration.expireAfterWriting(Duration.ofMinutes(5));

@Persisted
PersistedAppendingBuffer<Integer> buffer = PersistedAppendingBuffer. f(
    "my-buffer",
    Integer.class,
    Expiration.expireAfterWriting(Duration.ofSeconds(30)));
```

Es werden zwei Ablaufmodi unterstützt:

```java
Ablaufdatum (...)

Expiration.expireAfterReadingOrWriting(...)
```

State TTL Konfigurationen werden durch die Laufzeit fehlertolerant gemacht. Bei Ausfallzeiten werden bei einem Neustart die während der Ausfallzeit zu löschenden Zustandseinträge sofort gelöscht.

## Funktionsanbieter und Abhängigkeitsinjektion

Zugewiesene Funktionen werden über einen verteilten Knotencluster erzeugt. `StatefulFunctionProvider` ist eine Werksklasse um beim ersten Aktivieren eine neue Instanz einer zugewiesenen Funktion zu erstellen.

```java
Paket org.apache.flink.statefun.docs;

import org.apache.flink.statefun.docs.dependency.ProductionDependency;
import org.apache.flink.docs.dependency.RuntimeDependency;
import org.apache.flink.statefun.sdk.FunctionType;
import org.apache.flink.statefun.sdk.StatefulFunction;
import org.apache.flink. tatefun.sdk. tatefulFunctionProvider

CustomProvider der öffentlichen Klasse implementiert StatefulFunctionProvider {

    public StatefulFunction functionOfType(FunctionType type) {
        RuntimeDependency dependency = new ProductionDependency();
        gibt neue FnWithDependency(Abhängigkeit);
    }
}
```

Anbieter werden auf jedem parallelen Arbeiter einmal pro Typ aufgerufen, nicht für jede ID. Wenn eine zustandsfähige Funktion benutzerdefinierte Konfigurationen erfordert, können diese innerhalb eines Providers definiert und an den Konstruktor der Funktionen übergeben werden. Hier können auch gemeinsam genutzte physikalische Ressourcen wie eine Datenbankverbindung erstellt werden, die von einer beliebigen Anzahl virtueller Funktionen genutzt wird. Jetzt können Tests schnell Mock oder Test-Abhängigkeiten liefern, ohne dass komplexe Abhängigkeits-Injection-Frameworks benötigt werden.

```java
Paket org.apache.flink.statefun.docs;

import org.apache.flink.statefun.docs.dependency.RuntimeDependency;
import org.apache.flink.docs.dependency.TestDependency;
import org.junit.Assert;
import org.junit. est;

public class FunctionTest {

    @Test
    public void testFunctionWithCustomDependency() {
        RuntimeDependency dependency = new TestDependency();
        FnWithDependency Funktion = neue FnWithDependency(Abhängigkeit);

        Assert. ssertEquals("Es scheint, dass Mathematik kaputt ist", 1 + 1, 2);
    }
}
```
