# School Management System

A complete school management system with PostgreSQL database, Go REST API, and OpenAPI documentation.

## 📁 Project Structure

```
Cloud/
├── db/                          # Database schema and scripts
│   ├── init.sql                 # Database schema (tables, constraints)
│   ├── population.sql           # Sample data insertion
│   ├── deletion.sql             # Data cleanup scripts
│   └── *.png, *.puml           # Database diagrams
├── school-api/                  # Go REST API server
│   ├── cmd/
│   │   └── main.go             # Application entry point
│   ├── internal/
│   │   ├── db/
│   │   │   └── db.go           # Database connection
│   │   ├── models/             # GORM data models
│   │   │   ├── class.go
│   │   │   ├── student.go
│   │   │   ├── teacher.go
│   │   │   ├── subject.go
│   │   │   ├── teacher_assignment.go
│   │   │   ├── lesson_schedule.go
│   │   │   ├── lesson_log.go
│   │   │   ├── student_lesson.go
│   │   │   └── attendance_status.go
│   │   ├── handlers/           # HTTP request handlers
│   │   │   ├── class_handler.go
│   │   │   ├── student_handler.go
│   │   │   ├── teacher_handler.go
│   │   │   ├── subject_handler.go
│   │   │   ├── teacher_assignment_handler.go
│   │   │   ├── lesson_schedule_handler.go
│   │   │   ├── lesson_log_handler.go
│   │   │   ├── student_lesson_handler.go
│   │   │   └── attendance_status_handler.go
│   │   └── router/
│   │       └── router.go        # HTTP routing setup
│   ├── go.mod                  # Go module dependencies
│   ├── go.sum                  # Go module checksums
│   └── README.md               # API-specific documentation
└── api-docs/                   # API documentation
    └── swagger/
        ├── openapi.yaml        # OpenAPI 3.0 specification
        ├── index.html          # Swagger UI
        └── *.js, *.css         # Swagger UI assets
```

## 🗄️ Database Schema

### Core Entities
- **Classes**: School classes (grade + letter)
- **Students**: Student information linked to classes
- **Teachers**: Teacher profiles
- **Subjects**: School subjects
- **TeacherAssignments**: Links teachers to subjects
- **LessonSchedules**: Weekly lesson timetables
- **LessonLogs**: Actual lessons that occurred
- **StudentLessons**: Individual student participation, grades, attendance
- **AttendanceStatuses**: Attendance codes (P/A/L/E/S)

### Key Relationships
- Students belong to Classes
- Teachers are assigned to Subjects via TeacherAssignments
- LessonSchedules define when subjects are taught to classes
- LessonLogs record actual lessons that happened
- StudentLessons track individual student participation

## 🚀 Quick Start

### Prerequisites
- Go 1.23+
- PostgreSQL 12+
- Docker (optional, for containerized database)

### 1. Database Setup

#### Option A: Using Docker (Recommended)
```bash
# Start PostgreSQL container
docker run --name school-db \
  -e POSTGRES_DB=SportRental \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=pass \
  -p 5432:5432 \
  -d postgres:15

# Wait for container to start
sleep 10

# Run schema
psql -h localhost -U postgres -d SportRental -f db/init.sql

# Optional: Add sample data
psql -h localhost -U postgres -d SportRental -f db/population.sql
```

#### Option B: Local PostgreSQL
```bash
# Create database
createdb -U postgres SportRental

# Run schema
psql -U postgres -d SportRental -f db/init.sql

# Optional: Add sample data
psql -U postgres -d SportRental -f db/population.sql
```

### 2. API Server Setup

```bash
# Navigate to API directory
cd school-api

# Install dependencies
go mod tidy

# Set environment variables
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=pass
export DB_NAME=SportRental
export DB_SSLMODE=disable
export PORT=8000

# Run server
go run ./cmd
```

### 3. Verify Installation

```bash
# Test API health
curl http://localhost:8000/api/v1/classes

# Expected response: {"data":[]}
```

## 📚 API Documentation

### Interactive Documentation
- **Swagger UI**: Open `api-docs/swagger/index.html` in browser
- **OpenAPI Spec**: `api-docs/swagger/openapi.yaml`

### Base URL
```
http://localhost:8000/api/v1
```

### Available Endpoints

| Entity | GET | POST | GET by ID | PUT | DELETE |
|--------|-----|------|-----------|-----|--------|
| Classes | ✅ | ✅ | ✅ | ✅ | ✅ |
| Students | ✅ | ✅ | ✅ | ✅ | ✅ |
| Teachers | ✅ | ✅ | ✅ | ✅ | ✅ |
| Subjects | ✅ | ✅ | ✅ | ✅ | ✅ |
| Teacher Assignments | ✅ | ✅ | ✅ | ✅ | ✅ |
| Lesson Schedules | ✅ | ✅ | ✅ | ✅ | ✅ |
| Lesson Logs | ✅ | ✅ | ✅ | ✅ | ✅ |
| Student Lessons | ✅ | ✅ | ✅ | ✅ | ✅ |
| Attendance Statuses | ✅ | ✅ | ✅ | ✅ | ✅ |

### Example API Calls

#### Create a Class
```bash
curl -X POST http://localhost:8000/api/v1/classes \
  -H "Content-Type: application/json" \
  -d '{"grade": 10, "letter": "A"}'
```

#### Create a Student
```bash
curl -X POST http://localhost:8000/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{"class_id": 1, "first_name": "John", "last_name": "Doe", "patronymic": "Smith"}'
```

#### Get All Teachers
```bash
curl http://localhost:8000/api/v1/teachers
```

## 🛠️ Development

### Project Architecture

#### Models Layer (`internal/models/`)
- GORM models representing database entities
- JSON tags for API serialization
- Database constraints and validations

#### Handlers Layer (`internal/handlers/`)
- HTTP request/response handling
- Input validation and error handling
- Business logic implementation

#### Database Layer (`internal/db/`)
- PostgreSQL connection management
- Environment-based configuration

#### Router Layer (`internal/router/`)
- HTTP route registration
- Middleware setup
- API versioning

### Adding New Features

1. **Add Model**: Create new file in `internal/models/`
2. **Add Handler**: Create new file in `internal/handlers/`
3. **Register Routes**: Add to `internal/router/router.go`
4. **Update OpenAPI**: Add endpoints to `api-docs/swagger/openapi.yaml`

### Code Style
- Follow Go conventions
- Use meaningful variable names
- Add comments for complex logic
- Handle errors explicitly

## 🚀 Deployment

### Production Environment Variables
```bash
export DB_HOST=your-db-host
export DB_PORT=5432
export DB_USER=your-db-user
export DB_PASSWORD=your-secure-password
export DB_NAME=your-db-name
export DB_SSLMODE=require
export PORT=8000
export GIN_MODE=release
```

### Docker Deployment

#### 1. Create Dockerfile
```dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]
```

#### 2. Build and Run
```bash
# Build image
docker build -t school-api .

# Run container
docker run -p 8000:8000 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=5432 \
  -e DB_USER=postgres \
  -e DB_PASSWORD=pass \
  -e DB_NAME=SportRental \
  -e DB_SSLMODE=disable \
  school-api
```

### Cloud Deployment

#### AWS ECS/Fargate
1. Push Docker image to ECR
2. Create ECS task definition
3. Configure environment variables
4. Set up RDS PostgreSQL instance
5. Deploy ECS service

#### Google Cloud Run
1. Build and push to Google Container Registry
2. Deploy with Cloud SQL PostgreSQL
3. Configure environment variables
4. Set up load balancing

#### Azure Container Instances
1. Build and push to Azure Container Registry
2. Create container group with Azure Database for PostgreSQL
3. Configure environment variables
4. Deploy container instance

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `DB_HOST` | Database host | localhost | ✅ |
| `DB_PORT` | Database port | 5432 | ✅ |
| `DB_USER` | Database username | - | ✅ |
| `DB_PASSWORD` | Database password | - | ✅ |
| `DB_NAME` | Database name | - | ✅ |
| `DB_SSLMODE` | SSL mode | disable | ❌ |
| `PORT` | API server port | 8000 | ❌ |
| `GIN_MODE` | Gin framework mode | debug | ❌ |

### Database Configuration
- **Connection Pool**: Configured via GORM
- **Migrations**: Auto-migrate on startup
- **SSL**: Configurable via `DB_SSLMODE`

## 🧪 Testing

### Manual Testing
```bash
# Test all endpoints
curl http://localhost:8000/api/v1/classes
curl http://localhost:8000/api/v1/students
curl http://localhost:8000/api/v1/teachers
# ... etc
```

### Automated Testing
```bash
# Run tests (when implemented)
go test ./...

# Run with coverage
go test -cover ./...
```

## 📊 Monitoring

### Health Checks
```bash
# Basic health check
curl http://localhost:8000/api/v1/classes

# Database connectivity
# Check server logs for connection status
```

### Logging
- Structured logging via Gin framework
- Database connection logs
- Request/response logging
- Error logging with stack traces

## 🔒 Security

### Database Security
- Use strong passwords
- Enable SSL in production
- Restrict database access
- Regular security updates

### API Security
- Input validation
- SQL injection prevention (GORM)
- CORS configuration
- Rate limiting (recommended)

## 📈 Performance

### Database Optimization
- Proper indexing
- Connection pooling
- Query optimization
- Regular maintenance

### API Optimization
- Response compression
- Caching strategies
- Load balancing
- Horizontal scaling

## 🐛 Troubleshooting

### Common Issues

#### Database Connection Failed
```bash
# Check database is running
docker ps | grep postgres

# Test connection
psql -h localhost -U postgres -d SportRental

# Check environment variables
echo $DB_HOST $DB_PORT $DB_USER $DB_NAME
```

#### API Not Responding
```bash
# Check if server is running
netstat -an | grep :8000

# Check server logs
# Look for error messages in console output
```

#### Port Already in Use
```bash
# Find process using port 8000
lsof -i :8000

# Kill process
kill -9 <PID>

# Or use different port
export PORT=8001
```

## 📝 License

This project is part of a university course assignment.

## 👥 Contributing

1. Fork the repository
2. Create feature branch
3. Make changes
4. Test thoroughly
5. Submit pull request

## 📞 Support

For issues and questions:
1. Check troubleshooting section
2. Review logs for error messages
3. Verify environment configuration
4. Test database connectivity
