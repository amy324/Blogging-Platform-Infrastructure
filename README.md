# Blogging Platform Infrastructure

An infrastructure-focused blogging platform built with **Go, Docker, Terraform, PostgreSQL/CockroachDB and Prometheus**.

The project uses a simple REST API to demonstrate how an application can be **containerised, provisioned through Infrastructure as Code, connected to persistent storage and instrumented for observability**. The application functionality is intentionally lightweight so that the infrastructure can be the primary focus.

## Project Overview

The platform combines several core DevOps and SRE practices:

* **Infrastructure as Code:** Terraform defines the Docker network and application infrastructure, making the environment reproducible and version controlled.
* **Containerisation:** the Go application is packaged as a Docker image, providing a consistent runtime environment.
* **Service networking:** the application runs within a dedicated Docker network, providing an isolated environment for communicating services.
* **Persistent storage:** the API uses PostgreSQL-compatible CockroachDB for storing blog posts.
* **Observability:** Prometheus instrumentation exposes application and Go runtime metrics through `/metrics`.
* **Configuration management:** database connection details are supplied through environment variables rather than being embedded in the application.

The result is a small but complete infrastructure stack that demonstrates the workflow of taking an application from source code to a reproducible, containerised deployment with basic observability.

## Architecture

```text
                    ┌─────────────────────┐
                    │      Client         │
                    │  curl / Postman     │
                    └──────────┬──────────┘
                               │
                               ▼
                    ┌─────────────────────┐
                    │    Go REST API      │
                    │       :8080         │
                    └──────┬─────────┬────┘
                           │         │
                 ┌─────────▼───┐   ┌─▼──────────────┐
                 │ CockroachDB │   │  Prometheus    │
                 │ / PostgreSQL│   │   /metrics     │
                 └─────────────┘   └────────────────┘
                           
                  ┌─────────────────────────────┐
                  │          Terraform          │
                  │ Docker network + container  │
                  └─────────────────────────────┘
```

## Key Features

* REST API supporting full CRUD operations for blog posts.
* Go application built with the Gin web framework.
* GORM-based database access and automatic schema migration.
* Docker containerisation with a dedicated application network.
* Terraform-managed Docker infrastructure.
* Environment-based configuration for database credentials and connection settings.
* Prometheus metrics endpoint for application and runtime monitoring.
* Standard Go runtime and process metrics provided through the Prometheus client library.

## Technologies

| Technology                   | Purpose                               |
| ---------------------------- | ------------------------------------- |
| **Go**                       | REST API and application logic        |
| **Gin**                      | HTTP routing and request handling     |
| **GORM**                     | Database access and schema management |
| **CockroachDB / PostgreSQL** | Persistent data storage               |
| **Docker**                   | Application containerisation          |
| **Terraform**                | Infrastructure as Code                |
| **Prometheus**               | Metrics and observability             |

## Prerequisites

* [Terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)
* [Docker](https://docs.docker.com/engine/install/)
* PostgreSQL-compatible database
* Linux, macOS or Windows with [WSL](https://learn.microsoft.com/en-us/windows/wsl/install)

The project was developed and tested on Ubuntu Linux.

## Usage

Clone the repository:

```bash
git clone https://github.com/amy324/Blogging-Platform-Infrastructure.git
cd Blogging-Platform-Infrastructure
```

Create a `.env` file containing the database connection details:

```env
DB_HOST=<database_host>
DB_PORT=<database_port>
DB_USER=<database_user>
DB_PASSWORD=<database_password>
DB_NAME=<database_name>
DB_SSL_MODE=<ssl_mode>
```

Initialise Terraform and review the infrastructure plan:

```bash
terraform init
terraform plan
```

Deploy the infrastructure:

```bash
terraform apply
```

Terraform creates the Docker network and application container defined in `main.tf`.

## API

| Method   | Endpoint     | Description               |
| -------- | ------------ | ------------------------- |
| `GET`    | `/posts`     | Retrieve all posts        |
| `POST`   | `/posts`     | Create a new post         |
| `PUT`    | `/posts/:id` | Update an existing post   |
| `DELETE` | `/posts/:id` | Delete a post             |
| `GET`    | `/metrics`   | Expose Prometheus metrics |

### Example: Create a post

```json
{
  "title": "New Post Title",
  "content": "Lorem ipsum dolor sit amet.",
  "author": "Author Name"
}
```

## Monitoring

The application exposes a Prometheus-compatible `/metrics` endpoint.

A custom `requests_total` counter tracks the total number of requests received by the application, while the Prometheus Go client also exposes runtime and process metrics such as:

* Goroutines
* Memory allocation and heap usage
* CPU usage
* Garbage collection
* Process memory usage
* HTTP metric collection activity

These metrics provide visibility into application behaviour and resource usage and provide a foundation for further monitoring and alerting.

## Infrastructure

Terraform is used to manage the Docker infrastructure rather than manually creating containers.

The current configuration provisions:

* A dedicated Docker bridge network.
* The Go application container.
* Port mapping from host `8080` to the application's `8080`.
* The application container attached to the Terraform-managed network.

Using Terraform allows these resources to be recreated consistently and keeps the infrastructure configuration alongside the application code.

## Testing

The API can be tested using `curl`, Postman or Go's testing framework.

The running container can be verified with:

```bash
docker ps
```

The API endpoints can then be used to test CRUD operations, while `/metrics` can be queried to verify that Prometheus metrics are being exposed correctly.

## License

This project is licensed under the MIT License.
