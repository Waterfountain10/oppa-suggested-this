# oppa-suggested-this
Oppa Suggested is a microservices-based recommendation platform tailored for fans of K-dramas.


here is the structure (helped by Claude)

```
k-rec-hub/
├── .github/                      # GitHub Actions workflows
│   └── workflows/
│       └── ci.yml
├── cmd/                          # Main applications
│   ├── recommendation-engine/    # Recommendation service
│   │   └── main.go
│   ├── data-aggregator/         # Data aggregation service
│   │   └── main.go
│   └── api-gateway/             # API Gateway service
│       └── main.go
├── internal/                     # Private application and library code
│   ├── recommendation/          # Recommendation engine logic
│   │   ├── engine.go
│   │   ├── models.go
│   │   └── algorithms.go
│   ├── aggregator/              # Data aggregation logic
│   │   ├── collector.go
│   │   └── processor.go
│   └── common/                  # Shared internal code
│       ├── config/
│       └── middleware/
├── pkg/                         # Public library code
│   └── models/                  # Shared data models
│       └── content.go
├── api/                         # API documentation and schemas
│   └── openapi/
│       └── api.yaml
├── deployments/                 # Deployment configurations
│   ├── docker/                  # Dockerfiles
│   │   ├── recommendation/
│   │   │   └── Dockerfile
│   │   ├── aggregator/
│   │   │   └── Dockerfile
│   │   └── gateway/
│   │       └── Dockerfile
│   ├── kubernetes/             # Kubernetes manifests
│   │   ├── recommendation.yaml
│   │   ├── aggregator.yaml
│   │   └── gateway.yaml
│   └── terraform/              # Infrastructure as Code
│       ├── main.tf
│       └── variables.tf
├── configs/                    # Configuration files
│   ├── recommendation.yaml
│   ├── aggregator.yaml
│   └── gateway.yaml
├── test/                      # Additional external test apps and test data
│   ├── integration/
│   └── testdata/
├── web/                       # Web assets
│   ├── static/
│   └── templates/
├── scripts/                   # Scripts for development
│   ├── setup.sh
│   └── test.sh
├── go.mod                     # Go modules file
├── go.sum                     # Go modules checksum
├── Makefile                   # Build automation
└── README.md                  # Project documentation

```
