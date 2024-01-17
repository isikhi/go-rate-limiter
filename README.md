# Go Rate Limiter

This project, implementing rate limit functionality using Golang, PostgreSQL, Redis, Prometheus Docker, Kubernetes, and Chi framework with gRPC and HTTP endpoints.

Also, as a go learner, project architecture was initiated with inspiration from and based on the https://github.com/gmhafiz/go8 project. Thanks to this good starter pack.

I kept the Authentication layer so that I could save some configuration files in the future. But it wasn't necessary here. An extra module.

## Getting Started

Follow these steps to get the project up and running in your local environment.
### TODOs
- [ ] Overall tests
- [ ] Improve logging mechanisms
- [ ] Improve the rate limit duration value type(minutes, hour etc.)
- [ ] Add rate limit unit/functional tests
- [ ] Docker Hub push
- [ ] K8s migration and seeder jobs
- [ ] K8s application bootstrap
- [ ] K8s readme documentations


### Prerequisites

- Golang
- PostgreSQL
- Redis
- (Optional) Docker
- (Optional) Kubernetes/Minikube

### Installation


1. Navigate to the project directory:

   ```bash
    git clone https://github.com/isikhi/go-rate-limiter.git
    cd ./go-rate-limiter
    ```

2. Configuration(Environment)
   ```bash
    cp ./.env.example .env
    ```

3. Install dependencies:

    ```bash
   go mod download
    ```

4. Run Migrations and Seed DB

    ```bash
    go run cmd/migrate/main.go
    go run cmd/seed/main.go
   ```

5. Run App

    ```bash
    go run cmd/app/main.go
   ```

### Install & Run with docker(compose)
1. Clone the repository:

    ```bash
    git clone https://github.com/isikhi/go-rate-limiter.git
    cd ./go-rate-limiter
    ```
2. Configuration(Environment)
   ```bash
    cp ./.env.example .env
    ```

3. Docker Compose up!

    ```bash
    docker compose -f docker-compose.yml up --build
    ```


### Install & Run with kubernetes(minikube) [DRAFT - WIP]
1. Push to repository to docker hub or use mine.
   - App: isikhi/go-rate-limiter:app-1.0.0
   - Migrator: isikhi/go-rate-limiter:migrate-1.0.0
   - Seeder: isikhi/go-rate-limiter:seed-1.0.0

2. Run
   1. Postgres
      ```bash
       ./k8s/postgres/apply.sh
       ```
   2. Redis
      ```bash
       ./k8s/postgres/apply.sh
       ```
   3. App
      ```bash
       ./k8s/postgres/apply.sh
       ```


# How it works ? Test and Example Usage

1. Define Rate Limit Options
   Define rate limit options by referring to the example in [Rate Limit Example Http](./examples/rate-limit.http). You can find an HTTP request example that creates a client, token, and duration (in minutes).

2. Define Rate Limit Options
   Use the RPC client found in [Rate Limit Example RPC Client](./client/main.go) to make a request and determine whether rate limiting is applied or not.

3. Monitor Prometheus Metrics
   Retrieve Prometheus metrics from the endpoint `http://{{BASE_URL}}/metric`.



### Project Statement and Basic Architectural drawings
![Project Statements](./assets/project-statements.png)

