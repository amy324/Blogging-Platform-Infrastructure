

# Blogging Platform Infrastructure

This project implements the infrastructure for a simple blogging platform using Docker containers orchestrated with Terraform. The platform includes a RESTful API for managing posts and utilizes Prometheus for monitoring purposes, highlighting key aspects of Site Reliability Engineering (SRE) practices. As the primary focus of this project is on infrastructure rather than application functionality, the application provides basic CRUD operations. These operations serve to illustrate the fundamental principles of Infrastructure as Code (IaC) and its workings However, the infrastructure is designed to seamlessly accommodate the development of more sophisticated applications with ease.

## Table of Contents

- [Project Overview](#project-overview)
- [Key Features](#key-features)
- [Prerequisites](#prerequisites)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Monitoring](#monitoring)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

The blogging platform infrastructure is designed with a focus on reliability, scalability, and observability, following best practices of Site Reliability Engineering (SRE). Key aspects of the infrastructure include:

- **Containerization**: The application components are containerized using Docker, providing consistency and isolation across environments.
- **Orchestration**: Terraform is used for infrastructure as code (IaC) to define and manage the deployment of Docker containers, ensuring consistency and reproducibility of the infrastructure setup.
- **Monitoring**: Prometheus is integrated into the infrastructure to collect and store application metrics, allowing for real-time monitoring and alerting on key performance indicators (KPIs) and service-level objectives (SLOs).
- **Automation**: The infrastructure setup is automated using Terraform, enabling seamless provisioning, scaling, and teardown of resources, reducing manual intervention and human error.
- **Resilience**: The platform is designed with fault tolerance and redundancy in mind, utilizing features such as container health checks, automatic container restarts, and database replication for data durability.
- **Scalability**: The infrastructure architecture supports horizontal scalability, allowing for dynamic scaling of application instances based on demand using Docker Swarm or Kubernetes in production environments.

## Key Features

- **RESTful API**: The platform provides a RESTful API for CRUD operations on blog posts, adhering to REST principles for resource management and HTTP methods for actions.
- **Database Connectivity**: Connects to a CockroachDB database backend for persistent storage of blog post data, ensuring data integrity and reliability.
- **Prometheus Metrics**: Exposes application metrics via the `/metrics` endpoint for monitoring performance, availability, and resource utilization.
- **Terraform Configuration**: Defines infrastructure components and dependencies using Terraform configuration files (`main.tf`), enabling infrastructure as code (IaC) practices for version-controlled, repeatable deployments.
- **Continuous Integration/Continuous Deployment (CI/CD)**: Integrates with CI/CD pipelines for automated testing, building, and deployment of application changes, facilitating rapid and reliable software delivery.

## Prerequisites

Before running the infrastructure setup, ensure you have the following installed:

- **Terraform**: Version 0.12 or higher.
- **Docker**: For building and running the application containers.
- **Prometheus**: Required for monitoring the application.

## Usage

To deploy the blogging platform infrastructure:

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/amy324/Blogging-Platform-Infrastructure.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Blogging-Platform-Infrastructure
    ```

3. Modify the Terraform configuration file (`main.tf`) if necessary to customize the infrastructure settings.

4. Initialize Terraform:

    ```bash
    terraform init
    ```

5. Review the Terraform execution plan:

    ```bash
    terraform plan
    ```

6. Apply the Terraform configuration to create the infrastructure:

    ```bash
    terraform apply
    ```

## API Endpoints

The API service provides the following endpoints for managing blog posts:

- `GET /posts`: Fetches all blog posts.
    - Retrieves all posts from the database, eg:
    ```json
    {
        "id": 950923362343550977,
        "title": "New Post Title",
        "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
        "author": "Author Name",
        "created_at": "2024-03-12T18:46:20.011684Z",
        "updated_at": "0001-01-01T00:00:00Z"
    },
    ```
- `POST /posts`: Creates a new blog post.
   - Use the body structure 
      ```json
      {
        "title": "New Post Title",
        "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
        "author": "Author Name",
      }
        ```
- `PUT /posts/:id`: Updates an existing blog post.
- `DELETE /posts/:id`: Deletes a blog post by ID.


## Monitoring

The monitoring setup includes Prometheus for collecting and storing application metrics. Metrics are exposed via the `/metrics` endpoint and can be queried for analysis and visualization.



## Contributing

Contributions to this project are welcome! If you encounter any issues or have suggestions for improvements, please feel free to open an issue or create a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
