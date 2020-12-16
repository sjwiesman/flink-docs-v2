---
title: Checkpointing
type: docs
weight: 3
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

# Checkpointing

## Prerequisites

## Enabling and Configuring Checkpointing

### Related Config Options

## Selecting a State Backend

## State Checkpoints in Iterative Jobs

Flink currently only provides processing guarantees for jobs without iterations. Enabling checkpointing on an iterative job causes an exception.
In order to force checkpointing on an iterative program the user needs to set a special flag when enabling checkpointing: 

```java
env.enableCheckpointing(interval, CheckpointingMode.EXACTLY_ONCE, force = true)
```

Please note that records in flight in the loop edges (and the state changes associated with them) will be lost during failure.

## Restart Strategies

