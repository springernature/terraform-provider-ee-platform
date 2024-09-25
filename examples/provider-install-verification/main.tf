terraform {
  required_providers {
    ee-platform = {
      source = "hashicorp.com/edu/ee-platform"
    }
  }
}

provider "ee-platform" {}

data "ee-platform_teams" "teams" {}

output "all_teams" {
  value = data.ee-platform_teams.teams
}