terraform {
  required_providers {
    hashicups = {
      source = "hashicorp.com/edu/ee-platform"
    }
  }
}

provider "ee-platform" {}

data "ee_platform_teams" "teams" {}

