---
title: Windows
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

# Windows

Windows are at the heart of processing infinite streams.
Windows split the stream into “buckets” of finite size, over which we can apply computations.
This document focuses on how windowing is performed in Flink and how the programmer can benefit to the maximum from its offered functionality.

The general structure of a windowed Flink program is presented below.
The first snippet refers to keyed streams, while the second to non-keyed ones.
As one can see, the only difference is the `keyBy(...)` call for the keyed streams and the `window(...)` which becomes `windowAll(...)` for non-keyed streams.
This is also going to serve as a roadmap for the rest of the page.

**Keyed Windows**

```
stream
       .keyBy(...)               <-  keyed versus non-keyed windows
       .window(...)              <-  required: "assigner"
      [.trigger(...)]            <-  optional: "trigger" (else default trigger)
      [.evictor(...)]            <-  optional: "evictor" (else no evictor)
      [.allowedLateness(...)]    <-  optional: "lateness" (else zero)
      [.sideOutputLateData(...)] <-  optional: "output tag" (else no side output for late data)
       .reduce/aggregate/fold/apply()      <-  required: "function"
      [.getSideOutput(...)]      <-  optional: "output tag"
```

**Non-Keyed Windows**

```
stream
       .windowAll(...)           <-  required: "assigner"
      [.trigger(...)]            <-  optional: "trigger" (else default trigger)
      [.evictor(...)]            <-  optional: "evictor" (else no evictor)
      [.allowedLateness(...)]    <-  optional: "lateness" (else zero)
      [.sideOutputLateData(...)] <-  optional: "output tag" (else no side output for late data)
       .reduce/aggregate/fold/apply()      <-  required: "function"
      [.getSideOutput(...)]      <-  optional: "output tag"
```

In the above, the commands in square brackets (`[…]`) are optional.
This reveals that Flink allows you to customize your windowing logic in many different ways so that it best fits your needs.

## Window Lifecycle

## Keyed vs Non-Keyed Windows

## Window Assigners

### Tumbling Windows

### Sliding Windows

### Session Windows

### Global Windows

## Window Functions

### Reduce Function

### AggregateFunction

### ProcessWindowFunction

### ProcessWindowFunction with Incremental Aggregation

### WindowFunction (Legacy)

## Triggers

### Fire and Purge

### Default Triggers of Window Assigners

### Built-In and Custom Triggers

## Evictors

## Allowed Lateness

### Getting Late Data as a Side Output

### Late Element Consierations

## Working with Window Results

### Interaction of Watermarks and Windows

### Consecutive Windowed Operations

## userful State Size Considerations
