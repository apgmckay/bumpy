# Bumpy

A package that is designed to allow for simple [semantic version](https://semver.org/) bumping.

cli input can be provided using the provided `--version` flag, this can also be provided via stdin like so:

```
$ echo "1.0.0" | bumpy major
2.0.0
```

This allows you to use bumpy to create automation from things like on disk `Version` files as part of CI/CD pipelines.
