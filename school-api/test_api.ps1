# Test script for School API
$env:DB_HOST="localhost"
$env:DB_PORT="5432"
$env:DB_USER="postgres"
$env:DB_PASSWORD="pass"
$env:DB_NAME="SportRental"
$env:DB_SSLMODE="disable"
$env:PORT="8000"

Write-Host "Starting API server..."
Start-Process -FilePath "go" -ArgumentList "run", "./cmd" -WindowStyle Hidden

Start-Sleep -Seconds 5

Write-Host "Testing all endpoints..."

# Test Classes
Write-Host "`n1. Testing Classes endpoints:"
Write-Host "GET /api/v1/classes"
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/v1/classes" -Method GET
Write-Host "Response: $($response | ConvertTo-Json -Depth 2)"

Write-Host "`nPOST /api/v1/classes"
$classData = @{
    grade = 10
    letter = "A"
} | ConvertTo-Json
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/v1/classes" -Method POST -Body $classData -ContentType "application/json"
Write-Host "Response: $($response | ConvertTo-Json -Depth 2)"

# Test Students
Write-Host "`n2. Testing Students endpoints:"
Write-Host "GET /api/v1/students"
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/v1/students" -Method GET
Write-Host "Response: $($response | ConvertTo-Json -Depth 2)"

# Test Teachers
Write-Host "`n3. Testing Teachers endpoints:"
Write-Host "GET /api/v1/teachers"
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/v1/teachers" -Method GET
Write-Host "Response: $($response | ConvertTo-Json -Depth 2)"

# Test Subjects
Write-Host "`n4. Testing Subjects endpoints:"
Write-Host "GET /api/v1/subjects"
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/v1/subjects" -Method GET
Write-Host "Response: $($response | ConvertTo-Json -Depth 2)"

# Test Teacher Assignments
Write-Host "`n5. Testing Teacher Assignments endpoints:"
Write-Host "GET /api/v1/teacher-assignments"
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/v1/teacher-assignments" -Method GET
Write-Host "Response: $($response | ConvertTo-Json -Depth 2)"

# Test Lesson Schedules
Write-Host "`n6. Testing Lesson Schedules endpoints:"
Write-Host "GET /api/v1/lesson-schedules"
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/v1/lesson-schedules" -Method GET
Write-Host "Response: $($response | ConvertTo-Json -Depth 2)"

# Test Lesson Logs
Write-Host "`n7. Testing Lesson Logs endpoints:"
Write-Host "GET /api/v1/lesson-logs"
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/v1/lesson-logs" -Method GET
Write-Host "Response: $($response | ConvertTo-Json -Depth 2)"

# Test Student Lessons
Write-Host "`n8. Testing Student Lessons endpoints:"
Write-Host "GET /api/v1/student-lessons"
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/v1/student-lessons" -Method GET
Write-Host "Response: $($response | ConvertTo-Json -Depth 2)"

# Test Attendance Statuses
Write-Host "`n9. Testing Attendance Statuses endpoints:"
Write-Host "GET /api/v1/attendance-statuses"
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/v1/attendance-statuses" -Method GET
Write-Host "Response: $($response | ConvertTo-Json -Depth 2)"

Write-Host "`nAPI testing completed!"
