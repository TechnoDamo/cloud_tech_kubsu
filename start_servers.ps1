# Start both API and Frontend servers
Write-Host "Starting School Management System..." -ForegroundColor Green

# Start API server
Write-Host "Starting API server on port 8000..." -ForegroundColor Yellow
$env:DB_HOST="localhost"
$env:DB_PORT="5432"
$env:DB_USER="postgres"
$env:DB_PASSWORD="pass"
$env:DB_NAME="SportRental"
$env:DB_SSLMODE="disable"
$env:PORT="8000"

Start-Process -FilePath "powershell" -ArgumentList "-Command", "cd school-api; ./school-api" -WindowStyle Minimized

# Wait a moment for API to start
Start-Sleep -Seconds 3

# Start Frontend server
Write-Host "Starting Frontend server on port 3000..." -ForegroundColor Yellow
Start-Process -FilePath "python3" -ArgumentList "frontend/server.py" -WindowStyle Minimized

# Wait for servers to start
Start-Sleep -Seconds 2

Write-Host "`nServers started successfully!" -ForegroundColor Green
Write-Host "Frontend: http://localhost:3000" -ForegroundColor Cyan
Write-Host "API: http://localhost:8000" -ForegroundColor Cyan
Write-Host "`nPress any key to stop servers..." -ForegroundColor Yellow
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")

# Stop servers
Write-Host "`nStopping servers..." -ForegroundColor Red
Get-Process | Where-Object {$_.ProcessName -eq "school-api" -or $_.ProcessName -eq "python3"} | Stop-Process -Force
Write-Host "Servers stopped." -ForegroundColor Green
