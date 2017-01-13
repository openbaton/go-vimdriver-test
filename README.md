# Go Test VimDriver plugin
`go-vimdriver-test` is a drop in replacement for [the Java VimDriver plugin for OpenBaton][java-test-vimdriver], written using Go and [go-openbaton].

This plugin emulates and behaves like a real vim-driver plugin, for tasks such testing the NFVO without having access to an actual VIM (ie, Openstack). 

It works with the `dummy` VNFM implementations, like [go-dummy-vnfm], [java-dummy-vnfm] and [python-dummy-vnfm]. 

## How to install `go-vimdriver-test`
`go-vimdriver-test` is meant as a starting point for new plugins, or as a tool to debug and test the [plugin package of go-openbaton][go-openbaton-plugin].

If you wish to use it instead of [java-test-vimdriver], you need to launch it separately after launching the NFVO, or to wrap it into a Jar file and place it into 
the `plugins` directory of the main OpenBaton distribution (see [How to use `thunks` to make a Jar](#thunks) section below).

The plugin can be built using `go`, with 

```shell
go get -u github.com/mcilloni/go-vimdriver-test
```

This will fetch and build the source and its dependencies, creating a `go-vimdriver-test` in the `bin` directory of your _GOPATH_.

## How to launch the Test plugin

```shell
go-vimdriver-test --log "logfile" "name" "rabbit host" "port" "# of workers to be spawned" "username" "password"
```

If `--log` is omitted, the plugin will log on `/var/log/openbaton/type-plugin.log` on *NIX and in the Event Logger on Windows.
Use `--log -` to log only on `stdout`.

An empty command line is equivalent with the invocation below:

```shell
go-vimdriver-test --log "-" openbaton localhost 5672 10 admin openbaton
```

## How to use the Test plugin

After launching the plugin (see above), follow the [Dummy NSR Tutorial][tutorial].

## How to use `thunks` to make a Jar

OpenBaton expects its plugins to be contained in `.jar` files. If you wish to autolaunch this plugin together with the NFVO, install [thunks] using go, and then invoke

```shell
# set the environment variable GOOS = linux
go build github.com/mcilloni/go-vimdriver-test
thunks go-vimdriver-test
``` 

This will create a self extracting Jar named _go-vimdriver-test.jar_ that will extract and launch the VimDriver, executing it with the arguments provided by the NFVO.

## Issue tracker

Issues and bug reports should be posted to the GitHub Issue Tracker of this project.

## What is Open Baton?

Open Baton is an open source project providing a comprehensive implementation of the ETSI Management and Orchestration (MANO) specification and the TOSCA Standard.

Open Baton provides multiple mechanisms for interoperating with different VNFM vendor solutions. It has a modular architecture which can be easily extended for supporting additional use cases. 

It integrates with OpenStack as standard de-facto VIM implementation, and provides a driver mechanism for supporting additional VIM types. It supports Network Service management either using the provided Generic VNFM and Juju VNFM, or integrating additional specific VNFMs. It provides several mechanisms (REST or PUB/SUB) for interoperating with external VNFMs. 

It can be combined with additional components (Monitoring, Fault Management, Autoscaling, and Network Slicing Engine) for building a unique MANO comprehensive solution.

## Source Code and documentation

The Source Code of the other Open Baton projects can be found [on their GitHub page][openbaton-github], and the documentation can be found [on the official website][openbaton-doc].

## News and Website

Check the [Open Baton Website][openbaton]!

## Licensing and distribution
Licensed under the Apache License. See LICENSE for further details.

[openbaton]: http://openbaton.org
[openbaton-doc]: http://openbaton.org/documentation
[openbaton-github]: http://github.org/openbaton
[java-test-vimdriver]: https://github.com/openbaton/test-plugin
[java-dummy-vnfm]: https://github.com/openbaton/dummy-vnfm-amqp
[python-dummy-vnfm]: https://github.com/openbaton/python-vnfm-dummy
[go-dummy-vnfm]: https://github.com/mcilloni/go-dummy-vnfm
[go-openbaton]: https://github.com/mcilloni/go-openbaton
[go-openbaton-plugin]: https://github.com/mcilloni/go-openbaton/tree/master/plugin
[thunks]: https://github.com/mcilloni/thunks
[tutorial]: https://openbaton.github.io/documentation/dummy-NSR

