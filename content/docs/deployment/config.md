---
title: "Configuration"
type: docs
weight: 1
---
<!--
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
-->

All configuration is done in `conf/flink-conf.yaml`, which is expected to be a flat collection of [YAML key value pairs](http://www.yaml.org/spec/1.2/spec.html) with format `key: value`.

The configuration is parsed and evaluated when the Flink processes are started.
Changes to the configuration file require restarting the relevant processes.

The out of the box configuration will use your default Java installation.
You can manually set the environment variable `JAVA_HOME` or the configuration key `env.java.home` in `conf/flink-conf.yaml` if you want to manually override the Java runtime to use.

{{< toc >}}

# Common Setup Options

*Common options to configure your Flink application or cluster.*

### Hosts and Ports

Options to configure hostnames and ports for the different Flink components.

The JobManager hostname and port are only relevant for standalone setups without high-availability.
In that setup, the config values are used by the TaskManagers to find (and connect to) the JobManager.
In all highly-available setups, the TaskManagers discover the JobManager via the High-Availability-Service (for example ZooKeeper).

Setups using resource orchestration frameworks (K8s, Yarn, Mesos) typically use the framework's service discovery facilities.

You do not need to configure any TaskManager hosts and ports, unless the setup requires the use of specific port ranges or specific network interfaces to bind to.

{{< generated/common_host_port_section >}}

### Fault Tolerance

These configuration options control Flink's restart behaviour in case of failures during the execution. 
By configuring these options in your `flink-conf.yaml`, you define the cluster's default restart strategy. 

The default restart strategy will only take effect if no job specific restart strategy has been configured via the `ExecutionConfig`.

{% include generated/restart_strategy_configuration.html %}

**Fixed Delay Restart Strategy**

{% include generated/fixed_delay_restart_strategy_configuration.html %}

**Failure Rate Restart Strategy**

{% include generated/failure_rate_restart_strategy_configuration.html %}
