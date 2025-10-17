package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	dbpkg "school-api/internal/db"
	"school-api/internal/models"
	"school-api/internal/router"
)

// CORS middleware для всех запросов
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")

		// Preflight запросы (OPTIONS) просто возвращаем 200
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Загружаем .env (локально)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, continuing...")
	}

	// Подключение к базе данных
	db, err := dbpkg.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Миграция моделей
	if err := db.AutoMigrate(
		&models.Class{},
		&models.Student{},
		&models.Teacher{},
		&models.Subject{},
		&models.TeacherAssignment{},
		&models.LessonSchedule{},
		&models.LessonLog{},
		&models.StudentLesson{},
		&models.AttendanceStatus{},
	); err != nil {
		log.Fatalf("Auto migrate failed: %v", err)
	}

	// Настройка маршрутов
	r := router.Setup(db)

	// Оборачиваем маршрутизатор в CORS middleware
	handler := corsMiddleware(r)

	// Получаем порт из окружения (Serverless Containers)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // fallback для локальной разработки
	}
	addr := ":" + port

	log.Printf("Server listening on port %s", port)

	// Запуск HTTP сервера
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
