## force-md copy

Copy metadata between source and metadata formats

### Synopsis

Copy and convert metadata between source format (SFDX) and metadata format (MDAPI).

Examples:
  force-md copy src -t sfdx -f source      # Convert from metadata to source format
  force-md copy sfdx -t src -f metadata    # Convert from source to metadata format
  force-md copy src -t sfdx -f source -x package.xml  # Convert with package.xml filter

```
force-md copy [source directory] [flags]
```

### Options

```
  -f, --format string    target format (source or metadata)
  -h, --help             help for copy
  -x, --package string   package.xml file to filter metadata
  -t, --target string    target directory
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
      --verbose                show debugging output
```

### SEE ALSO

* [force-md](force-md.md)	 - force-md manipulate Salesforce metadata

