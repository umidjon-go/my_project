package main

import (
	"edu-crm/config"
	"edu-crm/handler"
	"edu-crm/handler/middleware"
	"edu-crm/pkg/db"
	"edu-crm/storage"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
)

func main() {
	cfg := config.Load()

	db, err := db.ConnectDb(*cfg)
	if err != nil {
		log.Fatalf("failed to connect db: %v\n", err)
	}
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("failed to goose set: %v\n", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	storage := storage.NewStorage(db)

	handler := handler.NewHandler(storage)

    route := gin.Default()
	route.Use(middleware.CORSMiddleware())

	branch := route.Group("/branches")
	{
		branch.POST("", handler.CreateBranch)
		branch.GET("", handler.GetBranches)
		branch.GET("/:branch_id", handler.GetBranch)
		branch.PUT("/:branch_id", handler.PutBranch)
		branch.DELETE("/:branch_id", handler.DeleteBranch)
	}


	user := route.Group("/users")
	{
		user.POST("", handler.CreateUser)
		user.GET("", handler.GetUsers)
		user.GET("/:user_id", handler.GetUser)
		user.PUT("/:user_id", handler.PutUser)
		user.DELETE("/:user_id", handler.DeleteUser)
	}
	course := route.Group("/courses")
	{
		course.POST("", handler.CreateCourse)
		course.GET("", handler.GetCourses)
		course.GET("/:course_id", handler.GetCourse)
		course.PUT("/:course_id", handler.PutCourse)
		course.DELETE("/:course_id", handler.DeleteCourse)
	}
	student := route.Group("/students")
	{
		student.POST("", handler.CreateStudent)
		student.GET("", handler.GetStudents)
		student.GET("/:student_id", handler.GetStudent)
		student.PUT("/:student_id", handler.PutStudent)
		student.DELETE("/:student_id", handler.DeleteStudent)
	}
	lesson := route.Group("/lessons")
	{
		lesson.POST("", handler.CreateLesson)
		lesson.GET("", handler.GetLessons)
		lesson.GET("/:lesson_id", handler.GetLesson)
		lesson.PUT("/:lesson_id", handler.PutLesson)
		lesson.PATCH("/:lesson_id/complete", handler.CompleteLesson)
		lesson.DELETE("/:lesson_id", handler.DeleteLesson)
	}
	lessonStudent := route.Group("/lesson_students")
	{
		lessonStudent.POST("", handler.CreateLessonStudent)
		lessonStudent.GET("/:lesson_student_id", handler.GetLessonStudent)
	}
	group := route.Group("/groups")
	{
		group.POST("", handler.CreateGroup)
		group.GET("", handler.GetGroups)
		group.GET("/:group_id", handler.GetGroup)
		group.PUT("/:group_id", handler.PutGroup)
		group.DELETE("/:group_id", handler.DeleteGroup)
	}
	payment := route.Group("/payments")
	{
		payment.POST("", handler.CreatePayment)
		payment.GET("", handler.GetPayments)
		payment.GET("/:payment_id", handler.GetPayment)
		payment.PUT("/:payment_id", handler.PutPayment)
		payment.DELETE("/:payment_id", handler.DeletePayment)
	}
	

	if err := route.Run(":8080"); err != nil {
		log.Fatalf("failed to run port 8080: %v\n", err)
	}
}
