provider "docker" {
  source = "kreuzwerker/docker"
  # Specify Docker API version (optional)
  # api_version = "1.40"
}

resource "docker_network" "app_network" {
  name   = "app-network"
}
