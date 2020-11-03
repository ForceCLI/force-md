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
