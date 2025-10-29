# Bumpy

A package that is designed to allow for simple [semantic version](https://semver.org/) bumping.

cli input can be provided using the provided `--version` flag, this can also be provided via stdin like so:

```
$ echo "1.0.0" | bumpy major
2.0.0
```

This allows you to use bumpy to create automation from things like on disk `Version` files as part of CI/CD pipelines.

Pre-Release and Build versions are also supported using the `--build=<name>` and `--pre-release=<name>` flags, further specification can be found in the [BNF section of the semver doc](https://semver.org/#backusnaur-form-grammar-for-valid-semver-versions).

## Development

At the moment you will need to set `GOPRIVATE=github.com/apgmckay` in your environment to pull the client.

This project uses [task](https://taskfile.dev/) to manage it's builds.

This project uses [Kafka](https://kafka.apache.org/) to drive events, you will need to run the terraform code under `_terraform` to create the required topics.

There is a script under `scripts` that can be used to consume events.

### Bumpy Client

Check [here](https://github.com/apgmckay/bumpy-client) for Bumpy client.

### Terraform Provider

Check [here](https://github.com/apgmckay/terraform-provider-bumpy) for terraform provider. 
