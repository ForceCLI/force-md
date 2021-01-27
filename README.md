# force-md

Manipulate Salesforce metadata.

## Commands

See [docs/force-md.md](docs/force-md.md) for all supported commands.

## Developing

To add support for a new metadata type, [zek](https://github.com/miku/zek) can
be useful for getting started by generating a `struct` that matches the XML
structure, e.g.

```
$ zek -C -m src/queues/*
```
