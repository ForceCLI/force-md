## force-md profile tidy

Tidy profile metadata

### Synopsis


Tidy profile metadata.

	The --wide and --ignore-errors flags can be used to help manage
	Profile metadata stored in a git repository.

	Configure clean and smudge git filters to use force-md:
	$ git config --local filter.salesforce-profile.clean 'force-md profile tidy --wide --ignore-errors -'
	$ git config --local filter.salesforce-profile.smudge 'force-md profile tidy --ignore-errors -'

	Update .gitattributes to use the salesforce-profile filter:
	*.profile-meta.xml filter=salesforce-profile

	The --wide flag will cause the Profile metadata to be stored in a
	flattened format that makes it easier to resolve merge conflicts.  If a child
	of a fieldPermissions element changes, for example, the entire
	fieldPermissions element will show up as changed because it's stored on a single line.

	The smudge filter will cause the metadata to be unflattened so it's available
	in the normal "long" format in the working copy.



```
force-md profile tidy [flags] [filename]...
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
```

### SEE ALSO

* [force-md profile](force-md_profile.md)	 - Manage Profiles

