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

# DataSet API 

DataSet programs in Flink are regular programs that implement transformations on data sets (e.g., filtering, mapping, joining, grouping). The data sets are initially created from certain sources (e.g., by reading files, or from local collections). Results are returned via sinks, which may for example write the data to (distributed) files, or to standard output (for example the command line terminal). Flink programs run in a variety of contexts, standalone, or embedded in other programs. The execution can happen in a local JVM, or on clusters of many machines.

Please refer to the DataStream API overview for an introduction to the basic concepts of the Flink API. That overview is for the DataStream API but the basic concepts of the two APIs are the same.

In order to create your own Flink DataSet program, we encourage you to start with the anatomy of a Flink Program and gradually add your own transformations. The remaining sections act as references for additional operations and advanced features.

{{< hint info >}}
Starting with Flink 1.12 the DataSet has been soft deprecated. We recommend that you use the DataStream API with `BATCH` execution mode. The linked section also outlines cases where it makes sense to use the DataSet API but those cases will become rarer as development progresses and the DataSet API will eventually be removed. Please also see FLIP-131 for background information on this decision. 
{{< /hint >}}

## Example Program

The following program is a complete, working example of WordCount.
You can copy & paste the code to run it locally.
You only have to include the correct Flink’s library into your project and specify the imports.
Then you are ready to go!

{{< tabs "basic-example" >}}
{{< tab "Java" >}}
```java
public class WordCountExample {
    public static void main(String[] args) throws Exception {
        final ExecutionEnvironment env = ExecutionEnvironment.getExecutionEnvironment();

        DataSet<String> text = env.fromElements(
            "Who's there?",
            "I think I hear them. Stand, ho! Who's there?");

        DataSet<Tuple2<String, Integer>> wordCounts = text
            .flatMap(new LineSplitter())
            .groupBy(0)
            .sum(1);

        wordCounts.print();
    }

    public static class LineSplitter implements FlatMapFunction<String, Tuple2<String, Integer>> {
        @Override
        public void flatMap(String line, Collector<Tuple2<String, Integer>> out) {
            for (String word : line.split(" ")) {
                out.collect(new Tuple2<String, Integer>(word, 1));
            }
        }
    }
}
```
{{< /tab >}}
{{< tab "Scala" >}}
```scala
import org.apache.flink.api.scala._

object WordCount {
  def main(args: Array[String]) {

    val env = ExecutionEnvironment.getExecutionEnvironment
    val text = env.fromElements(
      "Who's there?",
      "I think I hear them. Stand, ho! Who's there?")

    val counts = text
      .flatMap { _.toLowerCase.split("\\W+") filter { _.nonEmpty } }
      .map { (_, 1) }
      .groupBy(0)
      .sum(1)

    counts.print()
  }
}
```
{{< /tab >}}
{{< /tabs >}}

## DataSet Transformations

Data transformations transform one or more DataSets into a new DataSet.
Programs can combine multiple transformations into sophisticated assemblies.

#### Map

Takes one element and produces one element.

{{< tabs "mapfun" >}}
{{< tab "Java" >}}
```java
data.map(new MapFunction<String, Integer>() {
  public Integer map(String value) { return Integer.parseInt(value); }
});
```
{{< /tab >}}
{{< tab "Scala" >}}
```scala
data.map { x => x.toInt }
```
{{< /tab >}}
{{< /tabs >}}

#### FlatMap

Takes one element and produces zero, one, or more elements. 

{{< tabs "flatmapfunc" >}}
{{< tab "Java" >}}
```java
data.flatMap(new FlatMapFunction<String, String>() {
  public void flatMap(String value, Collector<String> out) {
    for (String s : value.split(" ")) {
      out.collect(s);
    }
  }
});
```
{{< /tab >}}
{{< tab "Scala" >}}
```scala
data.flatMap { str => str.split(" ") }
```
{{< /tab >}}
{{< /tabs >}}

## Specifying Keys

## Data Sources

### Read Compressed Files

## Data Sinks

## Iteration Operators

## Operating on Data Objects in Functions

### Object-Reuse Disabled (DEFAULT)

### Object-Reuse Enabled

## Debugging

### Local Execution Envronment

### Collection Data Sources and Sinks

## Semantic Annotations

## Broadcast Variables

## Distributed Cache

## Passing Parameters to Functions
