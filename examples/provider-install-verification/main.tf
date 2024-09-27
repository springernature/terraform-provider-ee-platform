terraform {
  required_providers {
    ee-platform = {
      source = "springernature/ee-platform"
    }
  }
}

provider "ee-platform" {
    platform_api = "https://ee-platform.springernature.app:8080"
}

data "ee-platform_teams" "all_teams" {}

output "teams" {
  value = data.ee-platform_teams.all_teams.teams
}

# example looping over teams
# key === team id
resource "local_file" "foo" {
  for_each = data.ee-platform_teams.all_teams.teams
  filename = "/tmp/test/${each.key}"
  content  = "name = ${each.value.name}, id = ${each.value.id}"
}