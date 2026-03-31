<div align="center">

# Nodefy

**A cloud-based platform to explore Linux environments, build development instances, and deploy applications using GitHub repositories.**

[![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-DC382D?logo=redis&logoColor=white)](https://redis.io/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

</div>

---

## Overview

Nodefy is a lightweight cloud platform that simplifies learning, development, and deployment by providing on-demand Linux environments and container-based infrastructure.

It enables users to:

- **Explore** Linux operating systems in real-time
- **Create** and manage development instances
- **Deploy** applications using GitHub repositories

Nodefy bridges the gap between theoretical learning and real-world cloud experience.

---

## Features

| Feature | Description |
|---|---|
|  **Linux OS Exploration** | Run Linux environments like Ubuntu with hands-on terminal access |
|  **Instance Management** | Create, start, stop, and delete containers; manage multiple instances |
|  **Multi-Node Development** | Build distributed systems (backend, database, cache) with Docker network integration |
|  **GitHub Deployment** | Deploy apps directly from repository links with automatic build and execution |
|  **Secure Access** | SSH access to containers with public access via tunneling |
|  **Dashboard** | Monitor and control instances from a unified interface |

---


## Workflow

### 1️⃣ Choose Stack
Select the required technology stack and environment for development.

![Choose Stack](/assests/choose.png)

---

### 2️⃣ Explore OS
Interact with Linux environments and explore system-level operations.

![Explore OS](/assests/exploreos.png)

---

### 3️⃣ Build Instances
Create and configure containerized instances for development.

![Build Instances](/assests/build.png)

---

### 4️⃣ Deploy Using GitHub Repositories
Provide your GitHub repository link to automatically build and deploy your application.

![Deploy](/assests/deploy.png)

---

### 5️⃣ Instance Management
Start, stop, and manage multiple running containers through the dashboard.

![Instance Management](/assests/instancesmanage.png)

1. User selects OS / creates instance / provides GitHub repository
2. Backend validates and processes the request
3. Docker container is created
4. Multi-node setup is configured (if needed)
5. Deployment is executed using the GitHub repository
6. Access is provided via SSH or public URL
7. Data is stored in the database
8. User manages instances via the dashboard

---


## Tech Stack

| Layer | Technology |
|---|---|
| **Frontend** | HTML, CSS, JavaScript |
| **Backend** | Go (Golang) |
| **Database** | PostgreSQL |
| **Containerization** | Docker |
| **Caching** | Redis |
| **Networking** | SSH, Tunneling |

---

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Go](https://go.dev/doc/install) (v1.21+)
- [Git](https://git-scm.com/downloads)
- [Postgres Docker Container](https://hub.docker.com/_/postgres)
- [Redis Docker Container](https://hub.docker.com/_/redis)
- [Make](https://www.gnu.org/software/make/)

### Installation

**1. Clone the repository**

```bash
git clone https://github.com/johngithiyon/Nodefy.git
cd nodefy
```

**2. Configure environment variables**

Create a `.env` file in the project root:

```env
CONN_STR=postgres://user:password@localhost/nodefy?sslmode=disable
REDIS_ADDRESS=localhost:6379
REDIS_PASSWORD=yourpassword
MAIL_FROM=your@email.com
MAIL_PASS=yourpassword
```

**3. Install dependencies**

```bash
go mod tidy
```

## Start the Required Containers

Before running Nodefy, you need to start the required services: **PostgreSQL** and **Redis**. You can start them using Docker with the following commands:

```bash
docker start postgres_containername  redis_containername
```

**4. Run the application**

```bash
go run cmd/server/main.go
```
      Or

**Using Make Command**
```bash
make run
```

**5. Open in browser**

```
http://localhost:8080
```

---

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

<div align="center">
</div>
