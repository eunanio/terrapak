# Terrapak
Terrapak is a private terraform registry. It integrates with your Github pull requests to automatically manage your terraform modules and make version control a much easier task.

## Feature Overview
- Automatic versioning of Terraform modules
- Monorepo friendly CI/CD workflow
- Supports S3 as storage backend
- Support GitHub Organisations for Authorization with `terraform login`

## Requirements
- Postgres 16
- Redis 7
- Github OAuth App
- S3 Bucket

## Getting Started

Terrapak uses a configuration file to define the modules you want to publish. Create a file named `terrapak.hcl` in the root of your repository. The file should contain a list of modules you want to publish. Each module should have a name and a path to the module directory. The path is relative to the root of the repository.
Example `terrapak.hcl` file:

```hcl
# terrapak.hcl

terrapak {
    organization = "myorg"
}

module "aws-bucket" {
    provider = "aws"
    path = "modules/aws/bucket"
    version = "1.0.0"
    # Example url: terrapak.dev/myorg/aws-bucket/aws
}

```

Example usage as module source:
```hcl
# main.tf

module "bucket" {
    source = "terrapak.dev/myorg/aws-bucket/aws"
    version = "1.0.0"
    bucket_name = "my-bucket"
}
```

### Installation
Terrapak is available as a docker image on [Docker Hub](https://hub.docker.com/r/monoci/terrapak). You can run the server with the following command:

```bash
docker run -p 5551:80 --mount type=bind,source="./config.yml",target=/tmp/config.yml -e TP_CONFIG_FILE=/tmp/config.yml monoci/terrapak
```

Docker Compose
```yaml
version: "1"
# This compose file is targeting local development.For deployments please use dedicated services like RDS for Postgres
services:
  terrapak:
    image: monoci/terrapak:v2
    ports:
      - "5551:5551"
    depends_on:
      - redis
      - postgres
    volumes:
      - ./config.yml:/tmp/config.yml
    environment:
     - TP_CONFIG_FILE=/tmp/config.yml
  redis:
    image: redis
    command: redis-server --requirepass ${REDIS_PASSWORD}
    ports:
      - "6379:6379"
    volumes:
      - /redis:/data
    env_file:
      - .env
  postgres:
    image: postgres:16
    restart: unless-stopped
    ports:
      - "5432:5432"
    volumes:
      - /postgres:/var/lib/postgresql/data
    env_file:
      - .env

```

*HTTPS is required to use this application, I recommend using a reverse proxy such as ngrok for local development.*

