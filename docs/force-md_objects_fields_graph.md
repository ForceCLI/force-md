## force-md objects fields graph

List relationship between fields and other objects

### Synopsis

List relationship between fields and objects for graph analysis using
digraph (https://github.com/golang/tools/blob/gopls/v0.4.4/cmd/digraph/digraph.go)

```
force-md objects fields graph [flags] [filename]...
```

### Examples

```

 $ force-md objects fields graph src/objects/* | digraph transpose

 $ force-md objects fields graph --object-only src/objects/* | digraph degree

```

### Options

```
  -f, --filtered-lookup       filtered lookup fields only
  -m, --formula               formula fields only
  -h, --help                  help for graph
  -k, --history-tracking      with history tracking
  -l, --label string          label
  -K, --no-history-tracking   without history tracking
  -R, --no-required           not required fields
  -D, --no-trending           without trending tracking
  -U, --no-unique             non-unique fields only
  -o, --object-only           show relationships between objects (default fields)
  -L, --references string     references object
  -r, --required              required fields
  -d, --trending              with trending tracking
  -t, --type strings          field type
  -u, --unique                unique fields only
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
      --verbose                show debugging output
```

### SEE ALSO

* [force-md objects fields](force-md_objects_fields.md)	 - Manage object field metadata

