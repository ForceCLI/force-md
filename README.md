# force-md

Manipulate Salesforce metadata.

## Profiles

### Tidy

Clean up metadata by sorting groups of elements in natural order.

```
force-md profile tidy src/profiles/*
```

### Clone Field Permissions

Add field permissions for a new field to Profiles by copying the permissions
from another field.

```
force-md profile field-permissions clone -s Account.My_Field__c -f Account.New_Field__c src/profiles/*
```

## Developing

To add support for a new metadata type, [zek](https://github.com/miku/zek) can
be useful for getting started by generating a `struct` that matches the XML
structure, e.g.

```
$ zek -C -m src/queues/*
```
