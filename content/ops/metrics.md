---
title: "Metrics"
draft: false
weight: 3
---

Stateful Functions includes a number of SDK specific metrics.
Along with the [standard metric scopes](https://ci.apache.org/projects/flink/flink-docs-stable/monitoring/metrics.html#system-scope), Stateful Functions supports ``Function Scope`` which one level below operator scope.

#### metrics.scope.function

* Default: &lt;host&gt;.taskmanager.&lt;tm_id&gt;.&lt;job_name&gt;.&lt;operator_name&gt;.&lt;subtask_index&gt;.&lt;function_namespace&gt;.&lt;function_name&gt;
* Applied to all metrics that were scoped to a function.

Metrics | Scope | Description | Type 
--------|-------|-------------|-----
in | Function | The number of incoming messages. | Counter
inRate | Function | The average number of incoming messages per second. | Meter
out-local | Function | The number of messages sent to a function on the same task slot. | Counter 
out-localRate | Function | The average number of messages sent to a function on the same task slot. | Meter
out-remote | Function | The number of messages sent to a function on a different task slot. | Counter 
out-remoteRate | Function | The average number of messages sent to a function on a different task slot. | Meter
out-egress | Function | The number of messages sent to an egress. | Counter
feedback.produced | Operator | The number of messages read from the feedback channel. | Meter
feedback.producedRate | Operator | The average number of messages read from the feedback channel per second. | Meter
