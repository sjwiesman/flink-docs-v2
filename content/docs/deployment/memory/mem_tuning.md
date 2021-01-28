---
title: "Memory Tuning Guide"
weight: 5
type: docs
aliases:
  - /deployment/memory/mem_tuning.html
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

# Memory tuning guide

In addition to the [main memory setup guide]({{< ref "/deployment/memory/mem_setup" >}}), this section explains how to set up memory
depending on the use case and which options are important for each case.

## Configure memory for standalone deployment

It is recommended to configure [total Flink memory]({{< ref "/deployment/memory/mem_setup" >}}#configure-total-memory)
([`taskmanager.memory.flink.size`]({{< ref "/deployment/config" >}}#taskmanager-memory-flink-size) or [`jobmanager.memory.flink.size`]({{< ref "/deployment/config" >}}#jobmanager-memory-flink-size))
or its components for [standalone deployment]({{< ref "/deployment/resource-providers/standalone/index" >}}) where you want to declare how much memory
is given to Flink itself. Additionally, you can adjust *JVM metaspace* if it causes [problems]({{< ref "/deployment/memory/mem_trouble" >}}#outofmemoryerror-metaspace).

The *total Process memory* is not relevant because *JVM overhead* is not controlled by Flink or the deployment environment,
only physical resources of the executing machine matter in this case.

## Configure memory for containers

It is recommended to configure [total process memory]({{< ref "/deployment/memory/mem_setup" >}}#configure-total-memory)
([`taskmanager.memory.process.size`]({{< ref "/deployment/config" >}}#taskmanager-memory-process-size) or [`jobmanager.memory.process.size`]({{< ref "/deployment/config" >}}#jobmanager-memory-process-size))
for the containerized deployments ([Kubernetes]({{< ref "/deployment/resource-providers/standalone/kubernetes" >}}), [Yarn]({{< ref "/deployment/resource-providers/yarn" >}}) or [Mesos]({{< ref "/deployment/resource-providers/mesos" >}})).
It declares how much memory in total should be assigned to the Flink *JVM process* and corresponds to the size of the requested container.

<span class="label label-info">Note</span> If you configure the *total Flink memory* Flink will implicitly add JVM memory components
to derive the *total process memory* and request a container with the memory of that derived size.

<div class="alert alert-warning">
  <strong>Warning:</strong> If Flink or user code allocates unmanaged off-heap (native) memory beyond the container size
  the job can fail because the deployment environment can kill the offending containers.
</div>
See also description of [container memory exceeded]({{< ref "/deployment/memory/mem_trouble" >}}#container-memory-exceeded) failure.

## Configure memory for state backends

This is only relevant for TaskManagers.

When deploying a Flink streaming application, the type of [state backend]({{< ref "/ops/state/state_backends" >}}) used
will dictate the optimal memory configurations of your cluster.

### Heap state backend

When running a stateless job or using a heap state backend ([MemoryStateBackend]({{< ref "/ops/state/state_backends" >}}#the-memorystatebackend)
or [FsStateBackend]({{< ref "/ops/state/state_backends" >}}#the-fsstatebackend)), set [managed memory]({{< ref "/deployment/memory/mem_setup_tm" >}}#managed-memory) to zero.
This will ensure that the maximum amount of heap memory is allocated for user code on the JVM.

### RocksDB state backend

The [RocksDBStateBackend]({{< ref "/ops/state/state_backends" >}}#the-rocksdbstatebackend) uses native memory. By default,
RocksDB is set up to limit native memory allocation to the size of the [managed memory]({{< ref "/deployment/memory/mem_setup_tm" >}}#managed-memory).
Therefore, it is important to reserve enough *managed memory* for your state. If you disable the default RocksDB memory control,
TaskManagers can be killed in containerized deployments if RocksDB allocates memory above the limit of the requested container size
(the [total process memory]({{< ref "/deployment/memory/mem_setup" >}}#configure-total-memory)).
See also [how to tune RocksDB memory]({{< ref "/ops/state/large_state_tuning" >}}#tuning-rocksdb-memory)
and [state.backend.rocksdb.memory.managed]({{< ref "/deployment/config" >}}#state-backend-rocksdb-memory-managed).

## Configure memory for batch jobs

This is only relevant for TaskManagers.

Flink's batch operators leverage [managed memory]({{< ref "/deployment/memory/mem_setup_tm" >}}#managed-memory) to run more efficiently.
In doing so, some operations can be performed directly on raw data without having to be deserialized into Java objects.
This means that [managed memory]({{< ref "/deployment/memory/mem_setup_tm" >}}#managed-memory) configurations have practical effects
on the performance of your applications. Flink will attempt to allocate and use as much [managed memory]({{< ref "/deployment/memory/mem_setup_tm" >}}#managed-memory)
as configured for batch jobs but not go beyond its limits. This prevents `OutOfMemoryError`'s because Flink knows precisely
how much memory it has to leverage. If the [managed memory]({{< ref "/deployment/memory/mem_setup_tm" >}}#managed-memory) is not sufficient,
Flink will gracefully spill to disk.

## Configure memory for sort-merge blocking shuffle

The number of required network buffers per sort-merge blocking result partition is controlled by 
[taskmanager.network.sort-shuffle.min-buffers]({{< ref "/deployment/config" >}}#taskmanager-network-sort-shuffle-min-buffers)
and the default value is 64 which is quite small. Though it can work for arbitrary parallelism, the 
performance may not be the best. For large scale jobs, it is suggested to increase this config value 
to improve compression ratio and reduce small network packets which is good for performance. To increase 
this value, you may also need to increase the size of total network memory by adjusting the config 
values of [taskmanager.memory.network.fraction]({{< ref "/deployment/config" >}}#taskmanager-memory-network-fraction),
[taskmanager.memory.network.min]({{< ref "/deployment/config" >}}#taskmanager-memory-network-min) and [taskmanager.
memory.network.max]({{< ref "/deployment/config" >}}#taskmanager-memory-network-max) to avoid `insufficient number of 
network buffers` error.

Except for network memory, the sort-merge blocking shuffle implementation also uses some unmanaged 
direct memory for shuffle data writing and reading. So to use sort-merge blocking shuffle, you may 
need to reserve some direct memory for it by increasing the config value of [taskmanager.memory.task
.off-heap.size]({{< ref "/deployment/config" >}}#taskmanager-memory-task-off-heap-size). If direct memory OOM error 
occurs after you enable the sort-merge blocking shuffle, you can just give more direct memory until 
the OOM error disappears.