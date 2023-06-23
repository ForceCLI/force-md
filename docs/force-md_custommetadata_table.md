## force-md custommetadata table

List custom metadata in a table

```
force-md custommetadata table [flags] [filename]...
```

### Examples

```

$ force-md custommetadata table -f 'dlrs__CalculationMode__c != "Realtime"' src/customMetadata/dlrs__LookupRollupSummary2.*

```

### Options

```
  -f, --filter string   expr boolean expression to filter records (default "true")
  -h, --help            help for table
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
```

### SEE ALSO

* [force-md custommetadata](force-md_custommetadata.md)	 - Manage Custom Metadata

