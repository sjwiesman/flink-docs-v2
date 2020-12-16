---
title: The Broadcast State Pattern
type: docs
weight: 2
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

# The Broadcast State Pattern

## Provided APIs

### BroacastProcessFunction & KeyedBroadcastProcessFunction

## Important Considerations

After describing the offered APIs, this section focuses on the important things to keep in mind when using broadcast state. These are:

* **There is no cross-task communication:** As stated earlier, this is the reason why only the broadcast side of a `(Keyed)-BroadcastProcessFunction` can modify the contents of the broadcast state. In addition, the user has to make sure that all tasks modify the contents of the broadcast state in the same way for each incoming element. Otherwise, different tasks might have different contents, leading to inconsistent results.

* **Order of events in Broadcast State may differ across tasks:** Although broadcasting the elements of a stream guarantees that all elements will (eventually) go to all downstream tasks, elements may arrive in a different order to each task. So the state updates for each incoming element **MUST NOT** depend on the ordering of the incoming events.

* **All tasks checkpoint their broadcast state:** Although all tasks have the same elements in their broadcast state when a checkpoint takes place (checkpoint barriers do not overpass elements), all tasks checkpoint their broadcast state, and not just one of them. This is a design decision to avoid having all tasks read from the same file during a restore (thus avoiding hotspots), although it comes at the expense of increasing the size of the checkpointed state by a factor of p (= parallelism). Flink guarantees that upon restoring/rescaling there will be **no duplicates** and **no missing data**. In case of recovery with the same or smaller parallelism, each task reads its checkpointed state. Upon scaling up, each task reads its own state, and the remaining tasks (p_new-p_old) read checkpoints of previous tasks in a round-robin manner.

* **No RocksDB state backend:** Broadcast state is kept in-memory at runtime and memory provisioning should be done accordingly. This holds for all operator states.
