## force-md permissionset tidy

Tidy Permission Set metadata

### Synopsis


Tidy permission set metadata.

	The --wide and --ignore-errors flags can be used to help manage
	Permission Set metadata stored in a git repository.

	Configure clean and smudge git filters to use force-md:
	$ git config --local filter.salesforce-permissionset.clean 'force-md permissionset tidy --wide --ignore-errors -'
	$ git config --local filter.salesforce-permissionset.smudge 'force-md permissionset tidy --ignore-errors -'

	Update .gitattributes to use the salesforce-permissionset filter:
	*.permissionset-meta.xml filter=salesforce-permissionset

	The --wide flag will cause the Permission Set metadata to be stored in a
	flattened format that makes it easier to resolve merge conflicts.  If a child
	of a fieldPermissions element changes, for example, the entire
	fieldPermissions element will show up as changed because it's stored on a single line.

	The smudge filter will cause the metadata to be unflattened so it's available
	in the normal "long" format in the working copy.



```
force-md permissionset tidy [flags] [filename]...
```

### Options

```
  -h, --help            help for tidy
  -i, --ignore-errors   ignore errors
  -l, --list            list files that need tidying
  -w, --wide            flatten into wide format
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
      --verbose                show debugging output
```

### SEE ALSO

* [force-md permissionset](force-md_permissionset.md)	 - Manage Permission Sets

