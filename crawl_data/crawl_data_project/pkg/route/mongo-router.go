package route

import (
	"crawl_data/pkg/handlers"
	"crawl_data/pkg/middleware"
	"crawl_data/pkg/repo"
	serv "crawl_data/pkg/service"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
func MongoRouter(incomingRoutes *gin.Engine, db *gorm.DB) {
	var (
		userRepositoryPostgres repo.UserRepositoryPostgres = repo.NewUserRepositoryPostgres(db)
		userRepository         repo.UserRepository         = repo.NewUserRepositoryMon(userRepositoryPostgres)
		jwtService             serv.JWTService             = serv.NewJWTService()
		userService            serv.UserService            = serv.NewUserService(userRepository)
		authService            serv.AuthService            = serv.NewAuthService(userRepository)
		authController         handlers.AuthController     = handlers.NewAuthController(authService, jwtService)
		userController         handlers.UserController     = handlers.NewUserController(userService, jwtService, authService)

		vascaraPostgresRepository repo.VascaraPostgresRepository = repo.NewVascaraPostgresRepository(db)
		historyRepository         repo.HistoryRepository         = repo.NewHistoryRepository()
		vascaraRepository         repo.VascaraMongoRepository    = repo.NewVascaraMongoRepository(vascaraPostgresRepository)
		vascaraService            serv.VascaraService            = serv.NewVascaraService(vascaraRepository)
		vascaraController         handlers.VascaraController     = handlers.NewVascaraController(vascaraService, jwtService, historyRepository)
	)

	authRoutes := incomingRoutes.Group("api/auth/")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := incomingRoutes.Group("api/user/", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	vascavaRoutes := incomingRoutes.Group("api/vascara/", middleware.AuthorizeJWT(jwtService))
	{
		vascavaRoutes.GET("/sendmessage", vascaraController.InsertProductSendMessage)
		vascavaRoutes.GET("/receivemessage", vascaraController.InsertProductReceiveMessage)
	}
	vascavaTestRoutes := incomingRoutes.Group("api/vascara/test/")
	{
		vascavaTestRoutes.GET("/insertproduct", vascaraController.InsertProductTest)
	}

}
