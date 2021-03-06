## apictl k8s change registry

Change the registry

### Synopsis

Change the registry to be pushed the built micro-gateway image

```
apictl k8s change registry [flags]
```

### Examples

```
apictl k8s change registry
```

### Options

```
  -h, --help                   help for registry
  -c, --key-file string        Credentials file
  -p, --password string        Password of the given user
      --password-stdin         Prompt for password of the given user in the stdin
  -R, --registry-type string   Registry type: DOCKER_HUB | AMAZON_ECR |GCR | HTTP
  -r, --repository string      Repository name or URI
  -u, --username string        Username of the repository
```

### Options inherited from parent commands

```
  -k, --insecure   Allow connections to SSL endpoints without certs
      --verbose    Enable verbose mode
```

### SEE ALSO

* [apictl k8s change](apictl_k8s_change.md)	 - Change a configuration in K8s cluster resource

