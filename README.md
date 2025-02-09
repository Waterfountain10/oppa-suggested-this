# oppa-suggested-this
Oppa Suggested is a microservices-based recommendation platform tailored for fans of K-dramas.


here is the structure (helped by Claude)

```
oppa-suggested-this/
├── cmd/                     # Main application entry points for each service
├── internal/               # Private code only used within this project
├── pkg/                    # Public code that could be imported by other projects
├── api/                   # API documentation and schemas
├── deployments/           # Docker, Kubernetes, and infrastructure configs
├── configs/               # Application configuration files
├── test/                  # Test files and test data
├── web/                   # Frontend assets and templates
└── scripts/               # Development and deployment scripts
├── go.mod                     # Go modules file
├── go.sum                     # Go modules checksum
├── Makefile                   # Build automation
└── README.md                  # Project documentation

```

where

Key Code Files:
pkg/models/content.go:

- Core data structures for content (dramas/games/music)
- Rating and recommendation request/response types

internal/recommendation/engine.go:

- Main recommendation logic
- Handles content storage and genre-based recommendations
- Thread-safe data management

internal/recommendation/handlers.go:

- HTTP endpoints for adding content/ratings
- Getting recommendations
- Request/response handling

cmd/recommendation-engine/main.go:

- Service entry point
- Server setup and routing
- Middleware configuration

configs/recommendation.yaml:

- Service configuration
- Rating thresholds
- Content type settings

configs/aggregator.yaml:

- Data collection settings
- External API configs
- Database settings
