---
title: "Stateful Stream Processing"
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

# Stateful Stream Processing

## What is State?

## Keyed State

## State Persistence

Flink implements fault tolerance using a combination of **stream replay** and **checkpointing**.
A checkpoint marks a specific point in each of the input streams along with the corresponding state for each of the operators.
A streaming dataflow can be resumed from a checkpoint while maintaining consistency (exactly-once processing semantics) by restoring the state of the operators and replaying the records from the point of the checkpoint.

The checkpoint interval is a means of trading off the overhead of fault tolerance during execution with the recovery time (the number of records that need to be replayed).

The fault tolerance mechanism continuously draws snapshots of the distributed streaming data flow.
For streaming applications with small state, these snapshots are very light-weight and can be drawn frequently without much impact on performance.
The state of the streaming applications is stored at a configurable place, usually in a distributed file system.

In case of a program failure (due to machine-, network-, or software failure), Flink stops the distributed streaming dataflow.
The system then restarts the operators and resets them to the latest successful checkpoint.
The input streams are reset to the point of the state snapshot.
Any records that are processed as part of the restarted parallel dataflow are guaranteed to not have affected the previously checkpointed state.

{{< hint warning >}}
By default, checkpointing is disable. See Checkpointing for details on how to configure checkpointing.
{{< /hint >}}

{{< hint info >}}
Because Flink’s checkpoints are realized through distributed snapshots, we use the words snapshot and checkpoint interchangeably. Often we also use the term snapshot to mean either checkpoint or savepoint.
{{< /hint >}}

### Checkpointing

The central part of Flink’s fault tolerance mechanism is drawing consistent snapshots of the distributed data stream and operator state.
These snapshots act as consistent checkpoints to which the system can fall back in case of a failure.
Flink’s mechanism for drawing these snapshots is described in “[Lightweight Asynchronous Snapshots for Distributed Dataflows](http://arxiv.org/abs/1506.08603)”.
It is inspired by the standard [Chandy-Lamport algorithm](http://research.microsoft.com/en-us/um/people/lamport/pubs/chandy.pdf) for distributed snapshots and is specifically tailored to Flink’s execution model.

Keep in mind that everything to do with checkpointing can be done asynchronously.
The checkpoint barriers don’t travel in lock step and operations can asynchronously snapshot their state.

Since Flink 1.11, checkpoints can be taken with or without alignment.
In this section, we describe aligned checkpoints first.

#### Barriers

A core element in Flink’s distributed snapshotting are the stream barriers.
These barriers are injected into the data stream and flow with the records as part of the data stream. Barriers never overtake records, they flow strictly in line.
A barrier separates the records in the data stream into the set of records that goes into the current snapshot, and the records that go into the next snapshot.
Each barrier carries the ID of the snapshot whose records it pushed in front of it.
Barriers do not interrupt the flow of the stream and are hence very lightweight.
Multiple barriers from different snapshots can be in the stream at the same time, which means that various snapshots may happen concurrently.

{{< img src="/fig/concepts/stream_barriers.svg" alt="Stream Barriers">}}

Stream barriers are injected into the parallel data flow at the stream sources.
The point where the barriers for snapshot n are injected (let’s call it `Sn`) is the position in the source stream up to which the snapshot covers the data. 
For example, in Apache Kafka, this position would be the last record’s offset in the partition.
This position `Sn` is reported to the checkpoint coordinator (Flink’s JobManager).

The barriers then flow downstream. 
When an intermediate operator has received a barrier for snapshot `n` from all of its input streams, it emits a barrier for snapshot `n` into all of its outgoing streams.
Once a sink operator (the end of a streaming DAG) has received the barrier `n` from all of its input streams, it acknowledges that snapshot `n` to the checkpoint coordinator.
After all sinks have acknowledged a snapshot, it is considered completed.

Once snapshot `n` has been completed, the job will never again ask the source for records from before `Sn`, since at that point these records (and their descendant records) will have passed through the entire data flow topology.

{{< img src="/fig/concepts/stream_aligning.svg" alt="Stream Aligning" width="95%">}}

Operators that receive more than one input stream need to align the input streams on the snapshot barriers. The figure above illustrates this:

* As soon as the operator receives snapshot barrier n from an incoming stream, it cannot process any further records from that stream until it has received the barrier n from the other inputs as well. Otherwise, it would mix records that belong to snapshot `n` and with records that belong to snapshot `n+1`.
* Once the last stream has received barrier `n`, the operator emits all pending outgoing records, and then emits snapshot `n` barriers itself.
* It snapshots the state and resumes processing records from all input streams, processing records from the input buffers before processing the records from the streams.
* Finally, the operator writes the state asynchronously to the state backend.

Note that the alignment is needed for all operators with multiple inputs and for operators after a shuffle when they consume output streams of multiple upstream subtasks.

### Unaligned Checkpointing

### State Backends

### Savepoints

### Exactly Once vs. At Least Once

The alignment step may add latency to the streaming program.
Usually, this extra latency is on the order of a few milliseconds, but we have seen cases where the latency of some outliers increased noticeably.
For applications that require consistently super low latencies (few milliseconds) for all records, Flink has a switch to skip the stream alignment during a checkpoint.
Checkpoint snapshots are still drawn as soon as an operator has seen the checkpoint barrier from each input.

When the alignment is skipped, an operator keeps processing all inputs, even after some checkpoint barriers for checkpoint n arrived.
That way, the operator also processes elements that belong to checkpoint `n+1` before the state snapshot for checkpoint `n` was taken.
On a restore, these records will occur as duplicates, because they are both included in the state snapshot of checkpoint `n`, and will be replayed as part of the data after checkpoint `n`.

{{< hint info >}}
**Note:** Alignment happens only for operators with multiple predecessors (`joins`) as well as operators with multiple senders (after a stream repartitioning/shuffle). Because of that, dataflows with only embarrassingly parallel streaming operations (`map()`, `flatMap()`,` filter()`, …) actually give exactly once guarantees even in at least once mode.
{{< /hint >}}

## State and Fault Tolerance in Batch Programs