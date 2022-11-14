ci-lint:
    docker-compose -f docker-compose-ci.yml run build golangci-lint run