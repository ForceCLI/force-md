## force-md objects tidy

Tidy object metadata

### Synopsis


Tidy object metadata.

	The --fix-missing flag can be used to add missing object metadata.  This includes:
	* picklist fields missing from Record Types


```
force-md objects tidy [flags] [filename]...
```

### Options

```
      --fix-missing   fix missing configuration (record type picklist options)
  -h, --help          help for tidy
  -l, --list          list files that need tidying
      --warn          warn about possibly bad metadata (unassiged record type picklist options)
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
```

### SEE ALSO

* [force-md objects](force-md_objects.md)	 - Manage Custom and Standard Objects

