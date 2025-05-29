## force-md tidy

Tidy Metadata

```
force-md tidy [flags]
```

### Examples

```

$ force-md tidy sfdx/main/default/objects/*/{fields,validationRules}/* sfdx/main/default/flows/*

$ force-md tidy src/objects/*

```

### Options

```
  -h, --help   help for tidy
  -l, --list   list files that need tidying
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
      --verbose                show debugging output
```

### SEE ALSO

* [force-md](force-md.md)	 - force-md manipulate Salesforce metadata

