terraform {
  required_providers {
    ee-platform = {
      source = "springernature/ee-platform"
    }
  }
}

provider "ee-platform" {
  platform_api = "https://ee-platform-dev.springernature.app"
}
