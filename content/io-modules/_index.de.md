---
title: "I/O-Module"
draft: falsch
weight: 4
---

Die I/O-Module von Stateful Functions ermöglichen Funktionen zum Empfangen und Senden von Nachrichten an externe Systeme. Basierend auf dem Konzept von Ingress (Input) und Egress (Output) Punkten und gebaut auf dem Apache Flink® Connector ecosystem, I/O-Module ermöglichen die Interaktion mit der Außenwelt durch den Stil des Nachrichtenübergangs.

## Einzug

Ein Ingress ist ein Eingangspunkt, bei dem Daten von einem externen System verbraucht und auf Null oder mehr Funktionen weitergeleitet werden. Es wird über einen __IngressIdentifier__ und einen __IngressSpec__ definiert.

Ein ingress-Identifizierer, ähnlich wie ein Funktionstyp, identifiziert eindeutig eine ingress durch Angabe seines Eingabe-Typs, eines Namensraumes und eines Namens.

Die Spezifikation legt fest, wie eine Verbindung zum externen System hergestellt werden soll, was für jedes einzelne I/O-Modul spezifisch ist. Jedes Paar ist an das System innerhalb eines zugewiesenen Funktionsmoduls gebunden.

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
          id: beispiel/user-ingress
          type: # ingress type
        spec: # ingress specific configurations
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
public class ModuleWithIngress implementiert StatefulFunctionModule {

    public static final IngressIdentifier<User> INGRESS =
        new IngressIdentifier<>(User. lasst, "Beispiel", "user-ingress");

    @Override
    public void configure(Map<String, String> globalConfiguration, Binder binder) {
        IngressSpec<User> spec = ...
        binder.bindIngress(spec);
    }
}
```
{{% /tab %}}
{{< /tabs >}}

## Router

Ein Router ist ein zustandsloser Operator, der jeden Datensatz aus einer ingress nimmt und ihn auf Null oder mehr Funktionen umleitet. Router sind über ein zustandsfähiges Funktionsmodul an das System gebunden, und im Gegensatz zu anderen Komponenten kann ein ingress eine beliebige Anzahl von Routern haben.

{{< tabs >}}
{{% tab name="remote module" %}}
Wenn in yaml definiert wird, werden Router durch eine Liste von Funktionstypen definiert. Die Id-Komponente der Adresse wird aus dem Schlüssel gezogen, der mit jedem Datensatz in seiner zugrunde liegenden Quellimplementierung verknüpft ist.
```yaml
Ziele:
    - beispiel-namespace/my-function-1
    - beispiel-namespace/my-function-2
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
public class UserRouter implementiert Router<User> {

    @Override
    public void route(User Nachricht, Downstream<User> Downstream) {
        Downstream. orward(FnUser.TYPE, message.getUserId(), message);
    }
}
```
{{% /tab %}}
{{< /tabs >}}

## Egress

Egress ist das Gegenteil von ingress; es ist ein Punkt, der Botschaften nimmt und sie in externe Systeme schreibt. Jede Eizelle ist mit zwei Komponenten definiert, einem __EgressIdentifier__ und einem __EgressSpec__.

Eine Eierkennung identifiziert eine Eier, basierend auf einem Namensraum, Namen und Herstellungsart. Eine Eierspektrum definiert die Details der Verbindung zum externen System, die Details sind für jedes einzelne I/O-Modul spezifisch. Jedes Paar der Identifikatoren ist an das System innerhalb eines zugewiesenen Funktionsmoduls gebunden.

{{< tabs >}}
{{% tab name="remote module" %}}
Wenn in yaml definiert wird, werden Router durch eine Liste von Funktionstypen definiert. Die Id-Komponente der Adresse wird aus dem Schlüssel gezogen, der mit jedem Datensatz in seiner zugrunde liegenden Quellimplementierung verknüpft ist.
```yaml
Version: "1. "

Modul:
  meta:
    typ: remote
  spec:
    egresses:
      - egress:
          meta:
            id: beispiel/user-egress
            type: # egress type
          spec: # egress specific configurations  
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
public class ModuleWithEgress implementiert StatefulFunctionModule {

    public static final EgressIdentifier<User> EGRESS =
            new EgressIdentifier<>("example", "egress", Benutzer. Glas);

    @Override
    public void configure(Map<String, String> globalConfiguration, Binder binder) {
        EgressSpec<User> spec = ...
        binder.bindEgress(spec);
    }
}
```
{{% /tab %}}
{{< /tabs >}}