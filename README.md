 Nodefy

 A cloud-based platform to explore Linux environments, build development instances, and deploy applications using GitHub repositories.

---

📌 Overview

Nodefy is a lightweight cloud platform that simplifies learning, development, and deployment by providing on-demand Linux environments and container-based infrastructure.

It enables users to:
- 🐧 Explore Linux operating systems in real-time  
- 🏗 Create and manage development instances  
- 🚀 Deploy applications using GitHub repositories  

Nodefy bridges the gap between **theoretical learning** and **real-world cloud experience**.

---

Key Features

- 🐧 **Linux OS Exploration**
  - Run Linux environments like Ubuntu
  - Hands-on terminal access

- 🏗 **Instance Management**
  - Create, start, stop, and delete containers
  - Manage multiple instances

- 🌐 **Multi-Node Development**
  - Build distributed systems (backend, database, cache)
  - Docker network integration

- 🚀 **Deployment using GitHub Repositories**
  - Deploy apps directly from repository links
  - Automatic build and execution

- 🔐 **Secure Access**
  - SSH access to containers
  - Public access via tunneling

- 📊 **Dashboard**
  - Monitor and control instances easily

---

## 🏗 Architecture

![Architecture Diagram](./assets/architecture.png)

Frontend → Backend → Docker Engine → Network → Database → User Access

---

## 🔁 Workflow

![Workflow Diagram](./assets/workflow.png)

1. User selects OS / creates instance / provides GitHub repository  
2. Backend validates and processes request  
3. Docker container is created  
4. Multi-node setup (if needed)  
5. Deployment using GitHub repository  
6. Access provided via SSH or URL  
7. Data stored in database  
8. User manages instances via dashboard  

---

## 🖼 Screenshots

### 🔹 Dashboard
![Dashboard](./assets/dashboard.png)

### 🔹 OS Selection
![OS Selection](./assets/os-selection.png)

### 🔹 Instance Management
![Instances](./assets/instances.png)

---

## ⚙️ Tech Stack

- **Frontend:** HTML, CSS, JavaScript  
- **Backend:** Go (Golang)
- **Database:** PostgreSQL  
- **Containerization:** Docker  
- **Caching:** Redis  
- **Networking:** SSH, Tunneling  

---

## 🚀 Getting Started

Follow these steps to run Nodefy locally.

---

## 📌 Prerequisites

Make sure you have installed:

- Docker  
- Go (Golang)  
- PostgreSQL  
- Git  

---

## 🔧 Installation & Setup

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/johngithiyon/Nodefy.git
cd nodefy

