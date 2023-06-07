## force-md completion

Generate completion script

### Synopsis

To load completions:

Bash:

$ source <(force-md completion bash)

# To load completions for each session, execute once:
Linux:
  $ force-md completion bash > /etc/bash_completion.d/force-md
MacOS:
  $ force-md completion bash > /usr/local/etc/bash_completion.d/force-md

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ force-md completion zsh > "${fpath[1]}/_force-md"

# You will need to start a new shell for this setup to take effect.


```
force-md completion [bash|zsh|fish|powershell]
```

### Options

```
  -h, --help   help for completion
```

### Options inherited from parent commands

```
      --convert-xml-entities   convert numeric xml entities to character entities (default true)
      --silent                 show errors only
```

### SEE ALSO

* [force-md](force-md.md)	 - force-md manipulate Salesforce metadata

