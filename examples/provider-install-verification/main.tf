terraform {
  required_providers {
    ee-platform = {
      source = "springernature/ee-platform"
    }
  }
}

provider "ee-platform" {}

data "ee-platform_teams" "teams" {}

output "all_teams" {
  value = data.ee-platform_teams.teams.teams
}

resource "local_file" "foo" {
  for_each = data.ee-platform_teams.teams.teams
  filename = "/tmp/test/${each.key}"
  content  = each.value.name
}