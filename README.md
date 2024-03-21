

# Blogging Platform Infrastructure

This project implements the infrastructure for a simple blogging platform using Docker containers orchestrated with Terraform. The platform includes a RESTful API for managing posts and utilizes Prometheus for monitoring purposes, highlighting key aspects of Site Reliability Engineering (SRE) practices. As the primary focus of this project is on infrastructure rather than application functionality, the application provides basic CRUD operations. These operations serve to illustrate the fundamental principles of Infrastructure as Code (IaC) and its workings However, the infrastructure is designed to seamlessly accommodate the development of more sophisticated applications with ease.

## Table of Contents

- [Project Overview](#project-overview)
- [Key Features](#key-features)
- [Prerequisites](#prerequisites)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Monitoring](#monitoring)
- [Testing](#testing)
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
- **Database Connectivity**: Connects to a Postgres CockroachDB database backend for persistent storage of blog post data, ensuring data integrity and reliability.
- **Prometheus Metrics**: Exposes application metrics via the `/metrics` endpoint for monitoring performance, availability, and resource utilization.
- **Terraform Configuration**: Defines infrastructure components and dependencies using Terraform configuration files (`main.tf`), enabling infrastructure as code (IaC) practices for version-controlled, repeatable deployments.


## Prerequisites

Before running the infrastructure setup, ensure you have the following installed:

- **Terraform**: Version 0.12 or higher.
- **Docker**: For building and running the application containers.
- **Prometheus**: Required for monitoring the application.

  For best results, I recommend you run this on a UNIX-like environement. For reference, this project was built using a [Linux Ubuntu](https://ubuntu.com/) environment. Alternatively, you can run on Windows using [Terraform Cloud](https://www.hashicorp.com/products/terraform?utm_source=google&utm_channel_bucket=paid&utm_medium=sem&utm_campaign=CLOUD_EMEA_UKI_ENG_BOFU_PRACTITIONER_SEM_A_ALL_TERRAFORM_CLD_GG_BRAND_-_Obility&utm_content=terraform%20cloud-144080651246-645564370290&utm_offer=signup&gad_source=1&gclid=CjwKCAjwte-vBhBFEiwAQSv_xXBDmsqlQnMqEhksoTd7_uncrOtBNF7AntDMdJlMJexWBITA5bh34xoCkaEQAvD_BwE) and [Docker Desktop](https://www.docker.com/products/docker-desktop/), however everything will run much more smoothly as Docker and terraform are designed to work best in UNIX-like environments.

If you have a PC, I recommend installing [WSL](https://learn.microsoft.com/en-us/windows/wsl/install) which allows you to run programs on a Linux distro of your choice and seemlessly switch between UNIX-like and Windows envrionments without having to dual boot

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
3. Create a new file named `.env` in the root directory of the project.
   
4. Add the following environment variables to the `.env` file:

    ```
    DB_HOST=<database_host>
    DB_PORT=<database_port>
    DB_USER=<database_user>
    DB_PASSWORD=<database_password>
    DB_NAME=<database_name>
    DB_SSL_MODE=<ssl_mode>
    ```

    Replace `<database_host>`, `<database_port>`, `<database_user>`, `<database_password>`, `<database_name>`, and `<ssl_mode>` with your actual database connection details. Here's an example:

    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=myuser
    DB_PASSWORD=mypassword
    DB_NAME=mydatabase
    DB_SSL_MODE=disable
    ```

5. Modify the Terraform configuration file (`main.tf`) if necessary to customize the infrastructure settings.

6. Initialize Terraform:

    ```bash
    terraform init
    ```

7. Review the Terraform execution plan:

    ```bash
    terraform plan
    ```

8. Apply the Terraform configuration to create the infrastructure:

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

The monitoring setup includes Prometheus for collecting and storing application metrics. Prometheus scrapes metrics from the `/metrics` endpoint exposed by the application, providing insights into various aspects of its behavior and performance. Below are some of the key metrics collected by Prometheus:

- **go_gc_duration_seconds**: This metric provides information about the duration of garbage collection cycles in Go. It includes quantiles such as minimum, 25th percentile, median, 75th percentile, and maximum durations, helping to identify potential performance bottlenecks caused by garbage collection.

- **go_goroutines**: Indicates the current number of Goroutines (lightweight threads) that are actively running, offering insights into the concurrency level of the application.

- **go_memstats_alloc_bytes**: Represents the number of bytes allocated and still in use by the Go runtime, reflecting the current memory usage of the application.

- **go_memstats_heap_inuse_bytes**: Shows the number of heap bytes that are actively in use by the Go runtime, providing visibility into the heap memory usage.

- **go_threads**: Reflects the number of operating system threads created by the Go runtime, which can impact the application's parallelism and resource utilization.

- **process_cpu_seconds_total**: Represents the total user and system CPU time spent by the process, indicating the overall CPU usage of the application.

- **process_resident_memory_bytes**: Indicates the size of resident memory (physical memory) used by the process, helping to assess its memory footprint.

- **promhttp_metric_handler_requests_in_flight**: Shows the current number of scrapes (data collection requests) being served by the Prometheus HTTP metric handler, reflecting the workload on the monitoring system.

- **promhttp_metric_handler_requests_total**: Displays the total number of scrapes categorized by HTTP status code, providing insights into the frequency and success rate of data collection requests.

- **requests_total**: Represents the total number of requests received by the application, serving as a fundamental metric for assessing its workload and usage patterns.

These metrics offer insights into various aspects of the application's performance, resource utilization, and runtime behavior, facilitating effective monitoring and troubleshooting.


## Testing

If you decide to use this infastructure for your own projects, you can ensure the reliability and functionality of the infrastructure and API endpoints, you can perform the following tests:

- **Infrastructure Testing**: After deploying the infrastructure using Terraform, verify that the Docker containers are running as expected and accessible. You can use tools like `docker ps` to list running containers and `curl` or a web browser to access the API endpoints.

- **API Endpoint Testing**: Use tools like `curl`, Postman, or automated testing frameworks to send requests to the API endpoints (`GET`, `POST`, `PUT`, `DELETE`) and validate the responses. Ensure that the CRUD operations behave as expected and handle error cases gracefully.

- **Integration Testing**: Conduct integration tests to validate the interactions between different components of the infrastructure, such as the API service, database backend, and monitoring system. You can use testing frameworks like Go's built-in testing package or tools like Selenium for web UI testing if you decide to create a front-end.

By conducting thorough testing, you can ensure the reliability, scalability, and security of the blogging platform infrastructure and API services, providing a better user experience and minimizing the risk of issues in production environments.


## Contributing

Contributions to this project are welcome! If you encounter any issues or have suggestions for improvements, please feel free to open an issue or create a pull request.

## License

This project is licensed under the MIT License.


