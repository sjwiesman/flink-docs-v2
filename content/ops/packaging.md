---
title: "Packaging"
draft: false
weight: 1
---

Stateful Functions applications can be packaged as either standalone applications or Flink jobs that can be submitted to a cluster.

## Images

The recommended deployment mode for Stateful Functions applications is to build a Docker image.
This way, user code does not need to package any Apache Flink components.
The provided base image allows teams to package their applications with all the necessary runtime dependencies quickly.

Below is an example Dockerfile for building a Stateful Functions image with both an [embedded module](/sdk/#embedded-module) and a [remote module](/sdk/#remote-module) for an application called ``statefun-example``.

```dockerfile
FROM flink-statefun:{{< version >}}

RUN mkdir -p /opt/statefun/modules/statefun-example
RUN mkdir -p /opt/statefun/modules/remote

COPY target/statefun-example*jar /opt/statefun/modules/statefun-example/
COPY module.yaml /opt/statefun/modules/remote/module.yaml
```

{{< stable >}}
{{% notice note %}}
The Flink community is currently waiting for the official Docker images to be published to Docker Hub.
In the meantime, Ververica has volunteered to make Stateful Function's images available via their public registry: 
`FROM ververica/flink-statefun:{{< version >}}`
You can follow the status of Docker Hub contribution [here](https://github.com/docker-library/official-images/pull/7749).
{{% /notice %}}
{{</ stable >}}

{{< unstable >}}
{{% notice warning %}}
The Flink community does not publish images for snapshot versions.
You can build this version locally by cloning the [repo](https://github.com/apache/flink-statefun) and following the instructions in 
`tools/docker/README.md`.
{{% /notice %}}
{{</ unstable >}}


## Flink Jar

If you prefer to package your job to submit to an existing Flink cluster, simply include ``statefun-flink-distribution`` as a dependency to your application.

```xml
<dependency>
	<groupId>org.apache.flink</groupId>
	<artifactId>statefun-flink-distribution</artifactId>
	<version>{{< version >}}</version>
</dependency>
```

It includes all of Stateful Functions' runtime dependencies and configures the application's main entry-point.

{{% notice note %}}
The distribution must be bundled in your application fat JAR so that it is on Flink's [user code class loader](https://ci.apache.org/projects/flink/flink-docs-stable/monitoring/debugging_classloading.html#inverted-class-loading-and-classloader-resolution-order).
{{% /notice %}}

```bash
$ ./bin/flink run -c org.apache.flink.statefun.flink.core.StatefulFunctionsJob ./statefun-example.jar
```
The following configurations are strictly required for running StateFun application.

```yaml
classloader.parent-first-patterns.additional: org.apache.flink.statefun;org.apache.kafka;com.google.protobuf
```
