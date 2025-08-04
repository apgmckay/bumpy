terraform {
  required_providers {
    bumpy = {
      source  = "registry.terraform.io/bumpycorp/bumpy" # short-hand for registry.terraform.io/bumpycorp/bumpy
      version = "0.0.1"
    }
  }
}

data "bumpy_major" "version" {
  version = "1.0.0"
}

output "version" {
  value = data.bumpy_major.version.version
}
