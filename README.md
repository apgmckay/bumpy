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

This project uses [task](https://taskfile.dev/) to manage it's builds.

### Terraform Provider

Setup a `$HOME/.terraformrc` file something like the below:

```
provider_installation {
  dev_overrides {
    "registry.terraform.io/bumpycorp/bumpy" = "/absolute/path/to/bumpy/terraform_provider"
  }
  direct {}
}
```

This allows for overriding of the registry config for local development, which means that you don't need to install terraform providers via `terraform init`.

Compile the terraform provider using the [task](https://taskfile.dev) file.

Start the grpc server for the compiled provder, this will return you a `TF_REATTACH_PROVIDERS` environment variable that you will need to set in the shell where you are going to run terraform, in this doc we will use the example/ dir.

Change directories in to example and run terraform.

```
# Set TF_REATTACH_PROVIDERS returned from setting up the grpc server for the terraform provider 
cd example
terraform plan
terraform apply
```
