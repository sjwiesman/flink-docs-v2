---
title: "Flink Architecture"
type: docs
weight: 4
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

# Flink Architecture

Flink is a distributed system and requires effective allocation and management
of compute resources in order to execute streaming applications. It integrates
with all common cluster resource managers such as [Hadoop
YARN](https://hadoop.apache.org/docs/stable/hadoop-yarn/hadoop-yarn-site/YARN.html),
[Apache Mesos](https://mesos.apache.org/) and
[Kubernetes](https://kubernetes.io/), but can also be set up to run as a
standalone cluster or even as a library.

This section contains an overview of Flink’s architecture and describes how its
main components interact to execute applications and recover from failures.

## Anatomy of a Flink Cluster

The Flink runtime consists of two types of processes: a _JobManager_ and one or more _TaskManagers_.

<img src="{% link /fig/processes.svg %}" alt="The processes involved in executing a Flink dataflow" class="offset" width="70%" />

The *Client* is not part of the runtime and program execution, but is used to
prepare and send a dataflow to the JobManager.  After that, the client can
disconnect (_detached mode_), or stay connected to receive progress reports
(_attached mode_). The client runs either as part of the Java/Scala program
that triggers the execution, or in the command line process `./bin/flink run
...`.

### JobManager

he _JobManager_ has a number of responsibilities related to coordinating the distributed execution of Flink Applications:
it decides when to schedule the next task (or set of tasks), reacts to finished
tasks or execution failures, coordinates checkpoints, and coordinates recovery on
failures, among others. This process consists of three different components:

  * **ResourceManager** 

    The _ResourceManager_ is responsible for resource de-/allocation and
    provisioning in a Flink cluster — it manages **task slots**, which are the
    unit of resource scheduling in a Flink cluster (see [TaskManagers](#taskmanagers)).
    Flink implements multiple ResourceManagers for different environments and
    resource providers such as YARN, Mesos, Kubernetes and standalone
    deployments. In a standalone setup, the ResourceManager can only distribute
    the slots of available TaskManagers and cannot start new TaskManagers on
    its own.  

  * **Dispatcher** 

    The _Dispatcher_ provides a REST interface to submit Flink applications for
    execution and starts a new JobMaster for each submitted job. It
    also runs the Flink WebUI to provide information about job executions.

  * **JobMaster** 

    A _JobMaster_ is responsible for managing the execution of a single
    [JobGraph]({{< ref "/docs/concepts/glossary#logical-graph" >}}).
    Multiple jobs can run simultaneously in a Flink cluster, each having its
    own JobMaster.

### TaskManagers


The *TaskManagers* (also called *workers*) execute the tasks of a dataflow, and buffer and exchange the data streams.

There must always be at least one TaskManager. The smallest unit of resource scheduling in a TaskManager is a task _slot_. The number of task slots in a
TaskManager indicates the number of concurrent processing tasks. Note that
multiple operators may execute in a task slot (see [Tasks and OperatorChains](#tasks-and-operator-chains)).

## Tasks and Operator Chains

## Task Slots and Resources

## Flink Application Execution

A _Flink Application_ is any user program that spawns one or multiple Flink
jobs from its ``main()`` method. The execution of these jobs can happen in a
local JVM (``LocalEnvironment``) or on a remote setup of clusters with multiple
machines (``RemoteEnvironment``). For each program, the
[``ExecutionEnvironment``]({{ site.javadocs_baseurl }}/api/java/) provides methods to
control the job execution (e.g. setting the parallelism) and to interact with
the outside world (see [Anatomy of a Flink Program]({%
link dev/datastream_api.md %}#anatomy-of-a-flink-program)).

The jobs of a Flink Application can either be submitted to a long-running
[Flink Session Cluster]({{< ref "/docs/concepts/glossary#flink-session-cluster" >}}), 
a dedicated [Flink Job Cluster]({{< ref "/docs/concepts/glossary#flink-job-cluster" >}}), or a
[Flink Application Cluster]({{< ref "/docs/concepts/glossary#flink-application-cluster" >}}).
The difference between these options is mainly related to the cluster’s lifecycle and to resource
isolation guarantees.

### Flink Session Cluster

* **Cluster Lifecycle**: in a Flink Session Cluster, the client connects to a
  pre-existing, long-running cluster that can accept multiple job submissions.
  Even after all jobs are finished, the cluster (and the JobManager) will
  keep running until the session is manually stopped. The lifetime of a Flink
  Session Cluster is therefore not bound to the lifetime of any Flink Job.

* **Resource Isolation**: TaskManager slots are allocated by the
  ResourceManager on job submission and released once the job is finished.
  Because all jobs are sharing the same cluster, there is some competition for
  cluster resources — like network bandwidth in the submit-job phase. One
  limitation of this shared setup is that if one TaskManager crashes, then all
  jobs that have tasks running on this TaskManager will fail; in a similar way, if
  some fatal error occurs on the JobManager, it will affect all jobs running
  in the cluster.

* **Other considerations**: having a pre-existing cluster saves a considerable
  amount of time applying for resources and starting TaskManagers. This is
  important in scenarios where the execution time of jobs is very short and a
  high startup time would negatively impact the end-to-end user experience — as
  is the case with interactive analysis of short queries, where it is desirable
  that jobs can quickly perform computations using existing resources.

{{< hint info >}}
**Note:** Formerly, a Flink Session Cluster was also known as a Flink Cluster in _session mode_.
{{< /hint >}}

### Flink Job Cluster

* **Cluster Lifecycle**: in a Flink Job Cluster, the available cluster manager
  (like YARN or Kubernetes) is used to spin up a cluster for each submitted job
  and this cluster is available to that job only. Here, the client first
  requests resources from the cluster manager to start the JobManager and
  submits the job to the Dispatcher running inside this process. TaskManagers
  are then lazily allocated based on the resource requirements of the job. Once
  the job is finished, the Flink Job Cluster is torn down.

* **Resource Isolation**: a fatal error in the JobManager only affects the one job running in that Flink Job Cluster.

* **Other considerations**: because the ResourceManager has to apply and wait
  for external resource management components to start the TaskManager
  processes and allocate resources, Flink Job Clusters are more suited to large
  jobs that are long-running, have high-stability requirements and are not
  sensitive to longer startup times.

{{< hint info >}}
**Note:** Formerly, a Flink Job Cluster was also known as a Flink Cluster in _job (or per-job) mode_.
{{< /hint >}}

### Flink Application Cluster

* **Cluster Lifecycle**: a Flink Application Cluster is a dedicated Flink
  cluster that only executes jobs from one Flink Application and where the
  ``main()`` method runs on the cluster rather than the client. The job
  submission is a one-step process: you don’t need to start a Flink cluster
  first and then submit a job to the existing cluster session; instead, you
  package your application logic and dependencies into a executable job JAR and
  the cluster entrypoint (``ApplicationClusterEntryPoint``)
  is responsible for calling the ``main()`` method to extract the JobGraph.
  This allows you to deploy a Flink Application like any other application on
  Kubernetes, for example. The lifetime of a Flink Application Cluster is
  therefore bound to the lifetime of the Flink Application.

* **Resource Isolation**: in a Flink Application Cluster, the ResourceManager
  and Dispatcher are scoped to a single Flink Application, which provides a
  better separation of concerns than the Flink Session Cluster.

{{< hint info >}}
**Note:** A Flink Job Cluster can be seen as a _run-on-client_ alternative to Flink Application Clusters. 
{{< /hint >}}
