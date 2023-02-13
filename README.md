## DSKit Examples
This repository contains examples from [dskit](https://github.com/grafana/dskit/) repository.

DSKit is a collection of Go packages that are useful for building distributed systems.

It's mainly maintained by the Grafana Labs team.

## Caution
This is a unofficial repository. The examples are not guaranteed to be correct. 
For more information about dskit, please refer to the [official repository](https://github.com/grafana/dskit/).

## Examples
- [Module manager](./module_manager): ModuleManager and Services are the main components of System. It manages entire app's dependencies and lifecycle.
- [Memberlist](./memberlist): Actually, The memberlist is not a core part of dskit, but it depends on [hashicorp/memberlist](https://github.com/hashicorp/memberlist)
- [Ring](./ring_default): A simple example of using dskit/ring package.