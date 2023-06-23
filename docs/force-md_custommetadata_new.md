## force-md custommetadata new

Create new custom metadata record

```
force-md custommetadata new [filename]... [flags]
```

### Examples

```

$ force-md custommetadata new src/customMetadata/My_Metadata.Example.md -l 'My Example' -v '{My_Field__c: "My Value", Default__c: true}'

```

### Options

```
  -h, --help            help for new
  -l, --label string    label
  -v, --values string   object describing values
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
```

### SEE ALSO

* [force-md custommetadata](force-md_custommetadata.md)	 - Manage Custom Metadata

