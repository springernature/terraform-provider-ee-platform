# EE Platform Terraform Provider

A Terraform Provider for working with EE platform resources.

For user documentation see the terraform registry:
https://registry.terraform.io/providers/springernature/ee-platform


## Development

See the makefile commands

```shell
# run the tests
$ make test

# build the binary in your $GOHOME/bin 
$ make install

# tell terraform to use this binary instead of from the actual registry
$ cat ~/.terraformrc
provider_installation {

  dev_overrides {
    "registry.terraform.io/springernature/ee-platform" = "/Users/<USERNAME>/go/bin/"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}

# run the example 
$ cd examples/provider-install-verification
$ terraform init && terraform plan
```

## Releasing

[![Release](https://github.com/springernature/terraform-provider-ee-platform/actions/workflows/release.yml/badge.svg)](https://github.com/springernature/terraform-provider-ee-platform/actions/workflows/release.yml)

To publish a new release, create a git tag in the semver format `vx.y.z`. The workflow will run which creates a github release and updates the terraform registry.

```shell
$ git tag -a v0.0.2 -m "message"
$ git push --tags
```
