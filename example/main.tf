terraform {
  required_providers {
    bumpy = {
      source  = "bumpycorp/bumpy" # short-hand for registry.terraform.io/bumpycorp/bumpy
      version = "0.0.1"
    }
  }
}

variable "input_version" {
  type    = string
  default = "0.1.0"
}

data "bumpy" "major" {
  major_version = var.input_version
}

data "bumpy" "minor" {
  minor_version = var.input_version
}

data "bumpy" "patch" {
  patch_version = var.input_version
}

output "major" {
  value = data.bumpy.major.version
}

output "minor" {
  value = data.bumpy.minor.version
}

output "patch" {
  value = data.bumpy.patch.version
}
