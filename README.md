  <img src="https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/openBaton.png" width="250"/>

  Copyright © 2015-2016 [Open Baton](http://openbaton.org).
  Licensed under [Apache v2 License](http://www.apache.org/licenses/LICENSE-2.0).
  
# Test VIM Driver in Go 
`go-vimdriver-test` is a drop in replacement for [the Java VimDriver plugin for Open Baton][java-test-vimdriver], written using Go and [go-openbaton].

This plugin emulates and behaves like a real vim-driver plugin, for tasks such testing the NFVO without having access to an actual VIM (ie, OpenStack). 

It works with the `dummy` VNFM implementations, like [go-dummy-vnfm], [java-dummy-vnfm] and [python-dummy-vnfm]. 

## How to install the Test VIM Driver
`go-vimdriver-test` is meant as a starting point for new plugins, or as a tool to debug and test the [plugin package of go-Open Baton][go-openbaton-plugin].

If you wish to use it instead of [java-test-vimdriver], you need to launch it separately after launching the NFVO, or to wrap it into a Jar file and place it into 
the `plugins` directory of the main OpenBaton distribution (see [How to use `thunks` to make a Jar](#thunks) section below).

The plugin can be built using `go`, with 

```shell
go get -u github.com/openbaton/go-vimdriver-test
```

This will fetch and build the source and its dependencies, creating a `go-vimdriver-test` in the `bin` directory of your _GOPATH_.

## How to launch the Test VIM Driver

```shell
go-vimdriver-test --log "logfile" "name" "rabbit host" "port" "# of workers to be spawned" "username" "password"
```

If `--log` is omitted, the plugin will log on `/var/log/openbaton/type-plugin.log` on *NIX and in the Event Logger on Windows.
Use `--log -` to log only on `stdout`.

An empty command line is equivalent with the invocation below:

```shell
go-vimdriver-test --log "-" openbaton localhost 5672 10 admin openbaton
```

## How to use the Test VIM Driver

After launching the plugin (see above), follow the [Dummy NSR Tutorial][tutorial].

## How to use `thunks` to make a Jar

OpenBaton expects its plugins to be contained in `.jar` files. If you wish to autolaunch this plugin together with the NFVO, install [thunks] using go, and then invoke

```shell
# set the environment variable GOOS = linux
go build github.com/openbaton/go-vimdriver-test
thunks go-vimdriver-test
``` 

This will create a self extracting Jar named _go-vimdriver-test.jar_ that will extract and launch the VimDriver, executing it with the arguments provided by the NFVO.

## Issue tracker

Issues and bug reports should be posted to the GitHub Issue Tracker of this project.
# What is Open Baton?

Open Baton is an open source project providing a comprehensive implementation of the ETSI Management and Orchestration (MANO) specification and the TOSCA Standard.

Open Baton provides multiple mechanisms for interoperating with different VNFM vendor solutions. It has a modular architecture which can be easily extended for supporting additional use cases. 

It integrates with OpenStack as standard de-facto VIM implementation, and provides a driver mechanism for supporting additional VIM types. It supports Network Service management either using the provided Generic VNFM and Juju VNFM, or integrating additional specific VNFMs. It provides several mechanisms (REST or PUB/SUB) for interoperating with external VNFMs. 

It can be combined with additional components (Monitoring, Fault Management, Autoscaling, and Network Slicing Engine) for building a unique MANO comprehensive solution.

## Source Code and documentation

The Source Code of the other Open Baton projects can be found [here][openbaton-github] and the documentation can be found [here][openbaton-doc] .

## News and Website

Check the [Open Baton Website][openbaton]
Follow us on Twitter @[openbaton][openbaton-twitter].

## Licensing and distribution
Copyright © [2015-2016] Open Baton project

Licensed under the Apache License, Version 2.0 (the "License");

you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

## Support
The Open Baton project provides community support through the Open Baton Public Mailing List and through StackOverflow using the tags openbaton.

## Supported by
  <img src="https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/fokus.png" width="250"/><img src="https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/tu.png" width="150"/>

[fokus-logo]: https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/fokus.png
[go-dummy-vnfm]: https://github.com/openbaton/go-dummy-vnfm
[go-openbaton]: https://github.com/openbaton/go-openbaton
[go-openbaton-plugin]: https://github.com/openbaton/go-openbaton/tree/master/plugin
[openbaton]: http://openbaton.org
[openbaton-doc]: http://openbaton.org/documentation
[openbaton-github]: http://github.org/openbaton
[openbaton-logo]: https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/openBaton.png
[openbaton-mail]: mailto:users@openbaton.org
[openbaton-twitter]: https://twitter.com/openbaton
[python-dummy-vnfm]: https://github.com/openbaton/python-vnfm-dummy
[tub-logo]: https://raw.githubusercontent.com/openbaton/openbaton.github.io/master/images/tu.png
[thunks]: https://github.com/mcilloni/thunks
[tutorial]: https://openbaton.github.io/documentation/dummy-NSR
[java-test-vimdriver]: https://github.com/openbaton/test-plugin
[java-dummy-vnfm]: https://github.com/openbaton/dummy-vnfm-amqp
