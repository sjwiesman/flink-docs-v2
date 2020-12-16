---
title: State Schema Evolution
type: docs
weight: 6
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

# State Schema Evolution

pache Flink streaming applications are typically designed to run indefinitely or for long periods of time.
As with all long-running services, the applications need to be updated to adapt to changing requirements.
This goes the same for data schemas that the applications work against; they evolve along with the application.

This page provides an overview of how you can evolve your state type’s data schema.
The current restrictions varies across different types and state structures (`ValueState`, `ListState`, etc.).

## Evolving State Schema

## Supported Data Types for Schema Evolution

### POJO Types

### Avro Types

Flink fully supports evolving schema of Avro type state, as long as the schema change is considered compatible by [Avro’s rules for schema resolution](http://avro.apache.org/docs/current/spec.html#Schema+Resolution).

One limitation is that Avro generated classes used as the state type cannot be relocated or have different namespaces when the job is restored.

## Limitations

* Schema evolution of keys is not supported. RocksDB state backend relies on binary objects ientity, rather than `hashCode` method implementation. Any changes to the key object structure could lead to non-deterministic behavior. 

* **Kryo** cannot be used for schema evolution. When Kryo is used, there is no possibility for the framework to verify if an incompatible change has been made.