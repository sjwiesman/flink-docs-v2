---
title: "Overview"
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

# Concepts

The [Hands-on Training]({{< ref "docs/learn-flink/overview" >}}) explains the basic concepts of stateful and timely stream processing that underlie Flink’s APIs, and provides examples of how these mechanisms are used in applications.
Stateful stream processing is introduced in the context of [Data Pipelines & ETL]({{< ref "docs/learn-flink/etl" >}}) and is further developed in the section on [Fault Tolerance]({{< ref "docs/learn-flink/fault_tolerance" >}}).
Timely stream processing is introduced in the section on Streaming Analytics.

This **Concepts in Depth** section provides a deeper understanding of how Flink’s architecture and runtime implement these concepts.

# Flink’s APIs

Flink offers different levels of abstraction for developing streaming/batch applications.

![Levels of Abstraction](/fig/concepts/levels_of_abstraction.svg)


