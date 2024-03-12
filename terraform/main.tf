provider "docker" {
  version = ">= 2.0.0"
  # Specify Docker API version (optional)
  # api_version = "1.40"
}

resource "docker_network" "app_network" {
  name   = "app-network"
}
