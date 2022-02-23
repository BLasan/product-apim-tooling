## apictl mi delete role

Delete a role from the Micro Integrator

### Synopsis

Delete a role with the name specified by the command line argument [role-name] from a Micro Integrator in the environment specified by the flag --environment, -e

```
apictl mi delete role [role-name] [flags]
```

### Examples

```
To delete a role
  apictl mi delete role [role-name] -e dev
To delete a role in a secondary user store
  apictl mi delete role [role-name] -d [domain] -e dev
NOTE: The flag (--environment (-e)) is mandatory
```

### Options

```
  -d, --domain string        Select the domain of the role
  -e, --environment string   Environment of the Micro Integrator from which a role should be deleted
  -h, --help                 help for role
```

### Options inherited from parent commands

```
  -k, --insecure   Allow connections to SSL endpoints without certs
      --verbose    Enable verbose mode
```

### SEE ALSO

* [apictl mi delete](apictl_mi_delete.md)	 - Delete users from a Micro Integrator instance

