# force-md

Manipulate Salesforce metadata.

## Usage

See [docs/force-md.md](docs/force-md.md) for all supported commands.

The commands are long, but tab completion makes them relatively painless.
Enable bash completion or see `force-md completion --help` for other options.

```
$ source <(force-md completion bash)
```

## Examples

Below are some basic examples.  See the [wiki](https://github.com/ForceCLI/force-md/wiki/Recipes) for higher level
examples.

### Tidy Permission Sets

Clean up metadata by sorting groups of elements in natural order.

```
$ force-md permissionset tidy src/permissionsets/*
```

### Clone Field Permissions

Add field permissions for a new field to Permission Sets by copying the
permissions from another field.

```
$ force-md permissionset field-permissions clone -s Account.My_Field__c -f Account.New_Field__c src/permissionsets/*
```

### Merge Permission Sets

Grant all permissions from a source permission set to another permission set.

```
$ force-md permissionset merge -s src/permissionsets/Subset.permissionset src/permissionsets/Superset.permissionset
```

### Add Class

Enable access to an apex class

```
$ force-md permissionset apex add -c MyClass src/permissionsets/My_Permission_Set.permissionset
```

### Add Tab

Enable tab visibility

```
$ force-md permissionset tab add -t My_Tab src/permissionsets/My_Permission_Set.permissionset
```

### Add Object Permissions to Profiles

Add object permissions to Profiles.  All permissions will default to false; use `profile object-permissions edit` to update.

```
$ force-md profile object-permissions add -o Account src/profiles/*
```

### Update Object Permissions

Update the Read, Create, Edit, Delete, View All, and Modify All permissions on
Profiles.  Any permissions not specified on the command line will be left
unchanged.

```
$ force-md profile object-permissions edit -o Account -e -D src/profiles/*
```

### Copy and Convert Metadata

Copy metadata between source format (SFDX) and metadata format (MDAPI).
Automatically handles merging/splitting of CustomObjects and CustomObjectTranslations.

```
$ force-md copy src -t sfdx -f source      # Convert from metadata to source format
$ force-md copy sfdx -t src -f metadata    # Convert from source to metadata format
```

## Developing

To add support for a new metadata type, [zek](https://github.com/miku/zek) can
be useful for getting started by generating a `struct` that matches the XML
structure, e.g.

```
$ zek -C -m src/queues/*
```
