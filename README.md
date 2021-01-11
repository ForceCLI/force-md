# force-md

Manipulate Salesforce metadata.

## Permission Sets

### Tidy

Clean up metadata by sorting groups of elements in natural order.

```
force-md permissionset tidy src/permissionsets/*
```

### Clone Field Permissions

Add field permissions for a new field to Permission Sets by copying the
permissions from another field.

```
force-md permissionset field-permissions clone -s Account.My_Field__c -f Account.New_Field__c src/permissionsets/*
```

### Add Class

Enable access to an apex class

```
force-md permissionset add-class -c MyClass src/permissionsets/My_Permission_Set.permissionset
```


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

### Add Object Permissions

Add object permissions to Profiles.  All permissions will default to false; use `profile object-permissions` to update.

```
force-md profile add-object-permissions -o Account src/profiles/*
```


### Update Object Permissions

Update the Read, Create, Edit, Delete, View All, and Modify All permissions on
Profiles.  Any permissions not specified on the command line will be left
unchanged.

```
force-md profile object-permissions -o Account -e -D src/profiles/*
```

### Delete Visibility/Access Permissions

```
force profile flow delete -f MyFlow src/profiles/*
force profile apex delete -c MyClass src/profiles/*
force profile application delete -a MyApplication src/profiles/*
force profile visualforce delete -p MyPage src/profiles/*
```

## Developing

To add support for a new metadata type, [zek](https://github.com/miku/zek) can
be useful for getting started by generating a `struct` that matches the XML
structure, e.g.

```
$ zek -C -m src/queues/*
```
