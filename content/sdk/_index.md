---
title: "SDKs"
draft: false
weight: 3
---

Stateful Function applications are composed of one or more Modules.
A module is a bundle of functions that are loaded by the runtime and available to be messaged.
Functions from all loaded modules are multiplexed and free to message each other arbitrarily.
Each module may contain functions written using a different language SDK, allowing for seamless cross language comunication.

Stateful Functions supports two types of modules: Embedded and Remote.

## Remote Module

Remote modules are run as external processes from the Apache Flink® runtime; in the same container, as a sidecar, or other external location.
This module type can support any number of language SDK's.
Remote modules are registered with the system via ``YAML`` configuration files.

### Specification

A remote module configuration consists of a ``meta`` section and a ``spec`` section.
``meta`` contains auxillary information about the module.
The ``spec`` describes the functions contained within the module and defines their persisted values.

### Defining Functions

``module.spec.functions`` declares a list of ``function`` objects that are implemented by the remote module.
A ``function`` is described via a number of properties.

* ``function.meta.kind``
    * The protocol used to communicate with the remote function.
    * Supported Values - ``http``
* ``function.meta.type``
    * The function type, defined as ``<namespace>/<name>``.
* ``function.spec.endpoint``
    * The endpoint at which the function is reachable.
    * Supported schemes are: ``http``, ``https``.
    * Transport via UNIX domain sockets is supported by using the schemes ``http+unix`` or ``https+unix``.
    * When using UNIX domain sockets, the endpoint format is: ``http+unix://<socket-file-path>/<serve-url-path>``. For example, ``http+unix:///uds.sock/path/of/url``.
* ``function.spec.states``
    * A list of the persisted values declared within the remote function.
    * Each entry consists of a `name` property and an optional `expireAfter` property.
    * Default for `expireAfter` - 0, meaning that state expiration is disabled.
* ``function.spec.maxNumBatchRequests``
    * The maximum number of records that can be processed by a function for a particular ``address`` before invoking backpressure on the system.
    * Default - 1000
* ``function.spec.timeout``
    * The maximum amount of time for the runtime to wait for the remote function to return before failing.
    * Default - 1 min

```yaml
version: "2.0"

module:
  meta:
    type: remote
  spec:
    functions:
      - function:
        meta:
          kind: http
          type: example/greeter
        spec:
          endpoint: http://<host-name>/statefun
          states:
            - name: seen_count
              expireAfter: 5min
          maxNumBatchRequests: 500
          timeout: 2min
```

## Embedded Module

