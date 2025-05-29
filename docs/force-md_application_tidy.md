## force-md application tidy

Tidy custom application

### Synopsis


Tidy custom application metadata.

	The --wide and --ignore-errors flags can be used to help manage
	CustomApplication metadata stored in a git repository.

	Configure clean and smudge git filters to use force-md:
	$ git config --local filter.salesforce-application.clean 'force-md application tidy --wide --ignore-errors -'
	$ git config --local filter.salesforce-application.smudge 'force-md application tidy --ignore-errors -'

	Update .gitattributes to use the salesforce-application filter:
	*.app-meta.xml filter=salesforce-application

	The --wide flag will cause the CustomApplication metadata to be stored in a
	flattened format that makes it easier to resolve merge conflicts.  If a child
	of a profileActionOverrides element changes, for example, the entire
	profileActionOverrides element will show up as changed because it's stored on a single line.

	The smudge filter will cause the metadata to be unflattened so it's available
	in the normal "long" format in the working copy.



```
force-md application tidy [filename]...
```

### Options

```
  -h, --help            help for tidy
  -i, --ignore-errors   ignore errors
  -w, --wide            flatten into wide format
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
      --verbose                show debugging output
```

### SEE ALSO

* [force-md application](force-md_application.md)	 - Manage Applications

