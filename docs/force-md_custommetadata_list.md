## force-md custommetadata list

List custom metadata

```
force-md custommetadata list [flags] [filename]...
```

### Examples

```

$ force-md custommetadata list -f 'dlrs__CalculationMode__c != "Realtime"' src/customMetadata/dlrs__LookupRollupSummary2.*

```

### Options

```
  -f, --filter string   expr boolean expression to filter records (default "true")
  -h, --help            help for list
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
```

### SEE ALSO

* [force-md custommetadata](force-md_custommetadata.md)	 - Manage Custom Metadata

