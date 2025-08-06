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
  default = "4.1.0"
}

data "bumpy_major_version" "bump" {
  version     = var.input_version
  build       = substr(md5("hello"), 0, 6)
  pre_release = "backend-feature-x"
}

data "bumpy_minor_version" "bump" {
  version     = var.input_version
  build       = substr(md5("hello"), 0, 6)
  pre_release = "backend-feature-x"
}

data "bumpy_patch_version" "bump" {
  version     = var.input_version
  build       = substr(md5("hello"), 0, 6)
  pre_release = "backend-feature-x"
}

output "bumpy_major" {
  value = data.bumpy_major_version.bump.result
}

output "bumpy_minor" {
  value = data.bumpy_minor_version.bump.result
}

output "bumpy_patch" {
  value = data.bumpy_patch_version.bump.result
}
