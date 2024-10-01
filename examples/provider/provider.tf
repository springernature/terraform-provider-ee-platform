terraform {
  required_providers {
    ee-platform = {
      source  = "springernature/ee-platform"
      version = "0.0.2"
    }
  }
}

provider "ee-platform" {
  platform_api = "https://ee-platform-dev.springernature.app"
}
