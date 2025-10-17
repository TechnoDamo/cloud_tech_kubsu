# School API

A Go REST API implementing CRUD for the school database.

## Run locally

1. Set environment variables (Postgres):

```bash
export DB_HOST=89.169.166.74
export DB_PORT=5432
export DB_USER=school-admin
export DB_PASSWORD=password123!
export DB_NAME=school-management
export DB_SSLMODE=disable
export PORT=8000
```

2. Start API:

```bash
cd school-api
go mod tidy
go run ./cmd
```

The server runs at http://localhost:$PORT

## Endpoints (per /api/v1)
- Classes: `GET/POST /classes`, `GET/PUT/DELETE /classes/{id}`
- Students: `GET/POST /students`, `GET/PUT/DELETE /students/{id}`
- Teachers: `GET/POST /teachers`, `GET/PUT/DELETE /teachers/{id}`
- Subjects: `GET/POST /subjects`, `GET/PUT/DELETE /subjects/{id}`
- TeacherAssignments: `GET/POST /teacher-assignments`, `GET/PUT/DELETE /teacher-assignments/{id}`
- LessonSchedules: `GET/POST /lesson-schedules`, `GET/PUT/DELETE /lesson-schedules/{id}`
- LessonLogs: `GET/POST /lesson-logs`, `GET/PUT/DELETE /lesson-logs/{id}`
- StudentLessons: `GET/POST /student-lessons`, `GET/PUT/DELETE /student-lessons/{id}`
- AttendanceStatuses: `GET/POST /attendance-statuses`, `GET/PUT/DELETE /attendance-statuses/{code}`

Refer to `api-docs/swagger/openapi.yaml` for detailed schemas.


- url: https://d5ds9jru50sk4al4aiaa.svoluuab.apigw.yandexcloud.net
    description: Local development server