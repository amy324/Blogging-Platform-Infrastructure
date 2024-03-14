terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
}

resource "docker_network" "app_network" {
  name   = "app-network"
}

resource "docker_container" "myapp_container" {
  name  = "myapp"
  image = "app:latest"
  ports {
    internal = 8080
    external = 8080
  }
  network_mode = docker_network.app_network.name
  # Add more configurations as needed
}
