
# TC-Chat Application

This is a Go-based chat application that uses Docker and Docker Compose for containerization. The application is built with a microservices architecture and leverages PostgreSQL as its database.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Makefile Commands](#makefile-commands)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

Before you begin, ensure you have the following software installed on your machine:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/your-username/tc-chat.git
    cd tc-chat
    ```

2. **Build the Docker images**:

   Use the `Makefile` to build the Docker images.

    ```bash
    make build
    ```

## Usage

### Running the Application

To start the application along with the PostgreSQL database, run:

```bash
make up
 
This will start the services defined in the docker-compose.yml file.

Stopping the Application
To stop the running services, execute:
make down
```
Makefile Commands
make env: Exports the CONFIG_PATH environment variable.
make build: Builds the Docker images using Docker Compose.
make up: Starts the Docker containers in detached mode.
make down: Stops and removes the Docker containers.
make logs: Displays the logs from the Docker containers.
make restart: Restarts the Docker containers.
Configuration
The application configuration is managed through the local.yaml file located in the config/ directory. This file includes database connection details, server configurations, and other application settings.