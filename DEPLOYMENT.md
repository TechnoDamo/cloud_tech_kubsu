# Deployment Guide

This guide covers different deployment strategies for the School Management System.

## 🐳 Docker Deployment

### Prerequisites
- Docker and Docker Compose installed
- Git (to clone repository)

### 1. Single Container Deployment

#### Create docker-compose.yml
```yaml
version: '3.8'

services:
  database:
    image: postgres:15
    environment:
      POSTGRES_DB: SportRental
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: pass
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./db/init.sql:/docker-entrypoint-initdb.d/init.sql
    networks:
      - school-network

  api:
    build: ./school-api
    ports:
      - "8000:8000"
    environment:
      DB_HOST: database
      DB_PORT: 5432
      DB_USER: postgres
      DB_PASSWORD: pass
      DB_NAME: SportRental
      DB_SSLMODE: disable
      PORT: 8000
      GIN_MODE: release
    depends_on:
      - database
    networks:
      - school-network

volumes:
  postgres_data:

networks:
  school-network:
    driver: bridge
```

#### Deploy
```bash
# Clone repository
git clone <repository-url>
cd school-management-system

# Start services
docker-compose up -d

# Check status
docker-compose ps

# View logs
docker-compose logs -f api
```

### 2. Multi-Stage Build

#### Create Dockerfile in school-api/
```dockerfile
# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 8000

# Run application
CMD ["./main"]
```

## ☁️ Cloud Deployment

### AWS Deployment

#### 1. ECS with Fargate

##### Create ECS Task Definition
```json
{
  "family": "school-api",
  "networkMode": "awsvpc",
  "requiresCompatibilities": ["FARGATE"],
  "cpu": "256",
  "memory": "512",
  "executionRoleArn": "arn:aws:iam::account:role/ecsTaskExecutionRole",
  "containerDefinitions": [
    {
      "name": "school-api",
      "image": "your-account.dkr.ecr.region.amazonaws.com/school-api:latest",
      "portMappings": [
        {
          "containerPort": 8000,
          "protocol": "tcp"
        }
      ],
      "environment": [
        {
          "name": "DB_HOST",
          "value": "your-rds-endpoint.amazonaws.com"
        },
        {
          "name": "DB_PORT",
          "value": "5432"
        },
        {
          "name": "DB_USER",
          "value": "postgres"
        },
        {
          "name": "DB_NAME",
          "value": "SportRental"
        },
        {
          "name": "DB_SSLMODE",
          "value": "require"
        },
        {
          "name": "GIN_MODE",
          "value": "release"
        }
      ],
      "secrets": [
        {
          "name": "DB_PASSWORD",
          "valueFrom": "arn:aws:secretsmanager:region:account:secret:db-password"
        }
      ],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/school-api",
          "awslogs-region": "us-east-1",
          "awslogs-stream-prefix": "ecs"
        }
      }
    }
  ]
}
```

##### Deploy Steps
```bash
# 1. Build and push to ECR
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin your-account.dkr.ecr.us-east-1.amazonaws.com

docker build -t school-api ./school-api
docker tag school-api:latest your-account.dkr.ecr.us-east-1.amazonaws.com/school-api:latest
docker push your-account.dkr.ecr.us-east-1.amazonaws.com/school-api:latest

# 2. Create ECS cluster
aws ecs create-cluster --cluster-name school-cluster

# 3. Create task definition
aws ecs register-task-definition --cli-input-json file://task-definition.json

# 4. Create service
aws ecs create-service \
  --cluster school-cluster \
  --service-name school-api-service \
  --task-definition school-api:1 \
  --desired-count 1 \
  --launch-type FARGATE \
  --network-configuration "awsvpcConfiguration={subnets=[subnet-12345],securityGroups=[sg-12345],assignPublicIp=ENABLED}"
```

#### 2. RDS PostgreSQL Setup
```bash
# Create RDS instance
aws rds create-db-instance \
  --db-instance-identifier school-db \
  --db-instance-class db.t3.micro \
  --engine postgres \
  --master-username postgres \
  --master-user-password your-secure-password \
  --allocated-storage 20 \
  --vpc-security-group-ids sg-12345 \
  --db-subnet-group-name your-subnet-group
```

### Google Cloud Platform

#### 1. Cloud Run Deployment

##### Create cloudbuild.yaml
```yaml
steps:
  # Build container
  - name: 'gcr.io/cloud-builders/docker'
    args: ['build', '-t', 'gcr.io/$PROJECT_ID/school-api', './school-api']
  
  # Push container
  - name: 'gcr.io/cloud-builders/docker'
    args: ['push', 'gcr.io/$PROJECT_ID/school-api']
  
  # Deploy to Cloud Run
  - name: 'gcr.io/cloud-builders/gcloud'
    args: [
      'run', 'deploy', 'school-api',
      '--image', 'gcr.io/$PROJECT_ID/school-api',
      '--region', 'us-central1',
      '--platform', 'managed',
      '--allow-unauthenticated'
    ]
```

##### Deploy
```bash
# Build and deploy
gcloud builds submit --config cloudbuild.yaml

# Set environment variables
gcloud run services update school-api \
  --region us-central1 \
  --set-env-vars DB_HOST=your-cloud-sql-ip,DB_PORT=5432,DB_USER=postgres,DB_NAME=SportRental,DB_SSLMODE=require
```

#### 2. Cloud SQL Setup
```bash
# Create Cloud SQL instance
gcloud sql instances create school-db \
  --database-version=POSTGRES_15 \
  --tier=db-f1-micro \
  --region=us-central1

# Create database
gcloud sql databases create SportRental --instance=school-db

# Create user
gcloud sql users create postgres --instance=school-db --password=your-secure-password
```

### Azure Deployment

#### 1. Container Instances

##### Create deployment template
```json
{
  "apiVersion": "2019-12-01",
  "type": "Microsoft.ContainerInstance/containerGroups",
  "name": "school-api",
  "location": "East US",
  "properties": {
    "containers": [
      {
        "name": "school-api",
        "properties": {
          "image": "your-registry.azurecr.io/school-api:latest",
          "ports": [
            {
              "port": 8000
            }
          ],
          "environmentVariables": [
            {
              "name": "DB_HOST",
              "value": "your-postgres-server.postgres.database.azure.com"
            },
            {
              "name": "DB_PORT",
              "value": "5432"
            },
            {
              "name": "DB_USER",
              "value": "postgres@your-postgres-server"
            },
            {
              "name": "DB_NAME",
              "value": "SportRental"
            },
            {
              "name": "DB_SSLMODE",
              "value": "require"
            }
          ],
          "resources": {
            "requests": {
              "cpu": 1,
              "memoryInGb": 1
            }
          }
        }
      }
    ],
    "osType": "Linux",
    "ipAddress": {
      "type": "Public",
      "ports": [
        {
          "port": 8000,
          "protocol": "TCP"
        }
      ]
    }
  }
}
```

##### Deploy
```bash
# Create resource group
az group create --name school-rg --location eastus

# Deploy container
az container create --resource-group school-rg --file deployment-template.json
```

## 🔧 Production Configuration

### Environment Variables
```bash
# Production environment
export DB_HOST=your-production-db-host
export DB_PORT=5432
export DB_USER=your-production-user
export DB_PASSWORD=your-secure-production-password
export DB_NAME=SportRental
export DB_SSLMODE=require
export PORT=8000
export GIN_MODE=release
```

### Security Considerations
- Use strong, unique passwords
- Enable SSL/TLS for database connections
- Use secrets management (AWS Secrets Manager, Azure Key Vault, etc.)
- Implement proper network security groups
- Regular security updates
- Monitor access logs

### Monitoring and Logging
- Set up CloudWatch, Stackdriver, or Azure Monitor
- Configure log aggregation
- Set up alerts for errors
- Monitor database performance
- Track API response times

### Scaling
- Horizontal scaling with load balancers
- Database read replicas for read-heavy workloads
- Caching layer (Redis) for frequently accessed data
- CDN for static assets

## 🚀 CI/CD Pipeline

### GitHub Actions Example
```yaml
name: Deploy to Production

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v2
    
    - name: Build Docker image
      run: docker build -t school-api ./school-api
    
    - name: Push to ECR
      env:
        AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
        AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
      run: |
        aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin $ECR_REGISTRY
        docker tag school-api:latest $ECR_REGISTRY/school-api:latest
        docker push $ECR_REGISTRY/school-api:latest
    
    - name: Deploy to ECS
      run: |
        aws ecs update-service --cluster school-cluster --service school-api-service --force-new-deployment
```

## 📊 Health Checks

### Application Health Check
```bash
# Basic health check
curl -f http://your-api-url/api/v1/classes || exit 1

# Database connectivity check
curl -f http://your-api-url/api/v1/attendance-statuses || exit 1
```

### Load Balancer Health Check
```bash
# Configure health check endpoint
# GET /health
# Expected: 200 OK
# Timeout: 5s
# Interval: 30s
# Unhealthy threshold: 3
```

## 🔄 Backup and Recovery

### Database Backup
```bash
# Automated backup script
#!/bin/bash
DATE=$(date +%Y%m%d_%H%M%S)
pg_dump -h $DB_HOST -U $DB_USER -d $DB_NAME > backup_$DATE.sql
aws s3 cp backup_$DATE.sql s3://your-backup-bucket/
```

### Disaster Recovery
1. Regular automated backups
2. Cross-region replication
3. Point-in-time recovery
4. Database failover procedures
5. Application rollback procedures

## 📈 Performance Optimization

### Database Optimization
- Proper indexing on frequently queried columns
- Query optimization and monitoring
- Connection pooling
- Read replicas for read-heavy workloads

### Application Optimization
- Response compression
- Caching strategies
- CDN for static content
- Load balancing
- Horizontal scaling

## 🛠️ Troubleshooting

### Common Deployment Issues

#### Container Won't Start
```bash
# Check container logs
docker logs container-name

# Check environment variables
docker exec container-name env

# Verify database connectivity
docker exec container-name nc -zv database-host 5432
```

#### Database Connection Issues
```bash
# Test database connection
psql -h $DB_HOST -U $DB_USER -d $DB_NAME

# Check network connectivity
telnet $DB_HOST 5432

# Verify credentials
echo $DB_USER $DB_PASSWORD
```

#### Performance Issues
```bash
# Check resource usage
docker stats container-name

# Monitor database performance
# Check slow query logs
# Monitor connection pool usage
```

This deployment guide provides comprehensive instructions for deploying the School Management System across different platforms and environments.
