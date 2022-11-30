package main

import (
    "log"

    "github.com/Drack112/Crud-Golang-API/config"
    "github.com/Drack112/Crud-Golang-API/controllers"
    "github.com/Drack112/Crud-Golang-API/middlewares"
    "github.com/Drack112/Crud-Golang-API/repository"
    "github.com/Drack112/Crud-Golang-API/service"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gorm.io/gorm"
)

var (
    db             *gorm.DB                   = config.SetupDatabasePostgres()
    UserRepository repository.UserRepository  = repository.NewUserRepository(db)
    BookRepository repository.BookRepository  = repository.NewBookRepository(db)
    AuthService    service.AuthService        = service.NewAuthService(UserRepository)
    UserService    service.UserService        = service.NewUserService(UserRepository)
    JWTService     service.JWTService         = service.NewJWTService()
    BookService    service.BookService        = service.NewBookService(BookRepository)
    AuthController controllers.AuthController = controllers.NewAuthController(AuthService, JWTService)
    UserController controllers.UserController = controllers.NewUserController(UserService, JWTService)
    BookController controllers.BookController = controllers.NewBookController(BookService, JWTService)
)

func init() {
    err := godotenv.Load()

    if err != nil {
        log.Panicf("Error in loading the .env file\n %s", err.Error())
    }
}

func main() {
    defer config.CloseDatabasePostgres(db)

    r := gin.Default()

    authRoutes := r.Group("api/auth")
    {
        authRoutes.POST("/login", AuthController.Login)
        authRoutes.POST("/register", AuthController.Register)
    }

    userRoutes := r.Group("api/users")
    {
        userRoutes.GET("/profile", UserController.Profile)
        userRoutes.PATCH("/profile", UserController.Update)
    }

    bookRoutes := r.Group("api/books", middlewares.AuthorizeJWT(JWTService))
    {
        bookRoutes.GET("/", BookController.All)
        bookRoutes.POST("/", BookController.Insert)
        bookRoutes.GET("/:id", BookController.FindByID)
        bookRoutes.PUT("/:id", BookController.Update)
        bookRoutes.DELETE("/:id", BookController.Delete)
    }
    r.Run()
}
