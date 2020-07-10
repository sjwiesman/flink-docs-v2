---
title: "Flink Connectors"
draft: false
weight: 3
---

The source-sink I/O module allows you to plug in existing, or custom, Flink connectors that are not already integrated into a dedicated I/O module. For details of how to build a custom connector see the official [Apache Flink documentation](https://ci.apache.org/projects/flink/flink-docs-stable).

## Dependency

To use a custom Flink connector, please include the following dependency in your pom.

```xml
<dependency>
    <groupId>org.apache.flink</groupId>
    <artifactId>statefun-flink-io</artifactId>
    <version>{{< version >}}</version>
    <scope>provided</scope>
</dependency>
```

## Source Spec

A source function specs creates an ingress from a Flink source function.

```java
package org.apache.flink.statefun.docs.io.flink;

import java.util.Map;
import org.apache.flink.statefun.docs.models.User;
import org.apache.flink.statefun.flink.io.datastream.SourceFunctionSpec;
import org.apache.flink.statefun.sdk.io.IngressIdentifier;
import org.apache.flink.statefun.sdk.io.IngressSpec;
import org.apache.flink.statefun.sdk.spi.StatefulFunctionModule;

public class ModuleWithSourceSpec implements StatefulFunctionModule {

    @Override
    public void configure(Map<String, String> globalConfiguration, Binder binder) {
        IngressIdentifier<User> id = new IngressIdentifier<>(User.class, "example", "users");
        IngressSpec<User> spec = new SourceFunctionSpec<>(id, new FlinkSource<>());
        binder.bindIngress(spec);
    }
}
```

## Sink Spec

A sink function spec creates an egress from a Flink sink function.

```java
package org.apache.flink.statefun.docs.io.flink;

import java.util.Map;
import org.apache.flink.statefun.docs.models.User;
import org.apache.flink.statefun.flink.io.datastream.SinkFunctionSpec;
import org.apache.flink.statefun.sdk.io.EgressIdentifier;
import org.apache.flink.statefun.sdk.io.EgressSpec;
import org.apache.flink.statefun.sdk.spi.StatefulFunctionModule;

public class ModuleWithSinkSpec implements StatefulFunctionModule {

    @Override
    public void configure(Map<String, String> globalConfiguration, Binder binder) {
        EgressIdentifier<User> id = new EgressIdentifier<>("example", "user", User.class);
        EgressSpec<User> spec = new SinkFunctionSpec<>(id, new FlinkSink<>());
        binder.bindEgress(spec);
    }
}
```
