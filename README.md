### Project Structure for [https://btabaku.nl/](https://btabaku.nl/)

```
personal-website/
├── cmd/
│   └── server/
│       └── main.go          # Entry point for the application
├── internal/
│   ├── config/              # Configuration handling
│   ├── handlers/            # HTTP handlers
│   ├── models/              # Data models
│   ├── services/            # Business logic
│   └── storage/             # Database interactions
├── web/
│   ├── static/              # Static assets (CSS, JS, images)
│   └── templates/           # HTML templates
├── scripts/                 # Scripts for setup/deployment
├── Dockerfile               # Dockerfile for containerization
├── docker-compose.yml       # Docker Compose configuration
├── go.mod                   # Go module file
└── README.md                # Project documentation
```

- **cmd/server/main.go**: The main entry point for the application.
- **internal/config**: Handles configuration loading and management.
- **internal/handlers**: Contains HTTP handlers for different routes.
- **internal/models**: Defines data models used in the application.
- **internal/services**: Contains business logic and service implementations.
- **internal/storage**: Manages database interactions and queries.
- **web/static**: Contains static assets like CSS, JavaScript, and images.
- **web/templates**: Contains HTML templates for rendering views.
- **scripts**: Contains scripts for setup and deployment.
- **Dockerfile**: Defines the Docker image for containerization.
- **docker-compose.yml**: Configuration for Docker Compose to manage multi-container applications.
- **go.mod**: Go module file that defines dependencies.
- **README.md**: Project documentation.

### How to Run

To run the application using Docker Compose, follow these steps:

1. **Clean up previous images and containers**:
    ```sh
    docker-compose down --rmi all
    ```

2. **Build the new multi-stage image**:
    ```sh
    docker-compose build
    ```

3. **Start the container**:
    ```sh
    docker-compose up
    ```

This will clean up any previous Docker images and containers, build the new Docker image, and start the application in a container.