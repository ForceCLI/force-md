## force-md sharingrules delete

Delete both criteria-based and owner-based sharing rules

### Synopsis

Delete both criteria-based and owner-based sharing rules. Use --type to specify the rule type, or omit to try deleting from both types.

```
force-md sharingrules delete -r RuleName [flags] [filename]...
```

### Options

```
  -h, --help          help for delete
  -r, --rule string   rule name to delete
  -t, --type string   rule type (criteria, owner) - if not specified, will try to delete from both
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
      --verbose                show debugging output
```

### SEE ALSO

* [force-md sharingrules](force-md_sharingrules.md)	 - Manage Sharing Rules

