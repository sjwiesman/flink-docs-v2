---
title: "I/O Modules"
draft: false
weight: 4
---

Stateful Functions’ I/O modules allow functions to receive and send messages to external systems.
Based on the concept of Ingress (input) and Egress (output) points, and built on top of the Apache Flink® connector ecosystem, I/O modules enable functions to interact with the outside world through the style of message passing.

## Ingress

An Ingress is an input point where data is consumed from an external system and forwarded to zero or more functions.
It is defined via an __IngressIdentifier__ and an __IngressSpec__.

An ingress identifier, similar to a function type, uniquely identifies an ingress by specifying its input type, a namespace, and a name.

The spec defines the details of how to connect to the external system, which is specific to each individual I/O module.
Each identifier-spec pair is bound to the system inside an stateful function module.

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
          id: example/user-ingress
          type: # ingress type
        spec: # ingress specific configurations
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
public class ModuleWithIngress implements StatefulFunctionModule {

    public static final IngressIdentifier<User> INGRESS =
        new IngressIdentifier<>(User.class, "example", "user-ingress");

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

A router is a stateless operator that takes each record from an ingress and routes it to zero or more functions. Routers are bound to the system via a stateful function module, and unlike other components, an ingress may have any number of routers.

{{< tabs >}}
{{% tab name="remote module" %}}
When defined in yaml, routers are defined by a list of function types.
The id component of the address is pulled from the key associated with each record in its underlying source implementation.
```yaml
targets:
    - example-namespace/my-function-1
    - example-namespace/my-function-2
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
public class UserRouter implements Router<User> {

    @Override
    public void route(User message, Downstream<User> downstream) {
        downstream.forward(FnUser.TYPE, message.getUserId(), message);
    }
}
```
{{% /tab %}}
{{< /tabs >}}

## Egress

Egress is the opposite of ingress; it is a point that takes messages and writes them to external systems.
Each egress is defined using two components, an __EgressIdentifier__ and an __EgressSpec__.

An egress identifier uniquely identifies an egress based on a namespace, name, and producing type.
An egress spec defines the details of how to connect to the external system, the details are specific to each individual I/O module.
Each identifier-spec pair are bound to the system inside a stateful function module.

{{< tabs >}}
{{% tab name="remote module" %}}
When defined in yaml, routers are defined by a list of function types.
The id component of the address is pulled from the key associated with each record in its underlying source implementation.
```yaml
version: "1.0"

module:
  meta:
    type: remote
  spec:
    egresses:
      - egress:
          meta:
            id: example/user-egress
            type: # egress type
          spec: # egress specific configurations  
```
{{% /tab %}}
{{% tab name="embedded module" %}}
```java
public class ModuleWithEgress implements StatefulFunctionModule {
    
    public static final EgressIdentifier<User> EGRESS =
            new EgressIdentifier<>("example", "egress", User.class);

    @Override
    public void configure(Map<String, String> globalConfiguration, Binder binder) {
        EgressSpec<User> spec = ...
        binder.bindEgress(spec);
    }
}
```
{{% /tab %}}
{{< /tabs >}}