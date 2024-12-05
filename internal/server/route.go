package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	aboutusHandler "github.com/sawalreverr/recything/internal/about-us/handler"
	aboutusRepo "github.com/sawalreverr/recything/internal/about-us/repository"
	aboutusUsecase "github.com/sawalreverr/recything/internal/about-us/usecase"
	achievementHandler "github.com/sawalreverr/recything/internal/achievements/manage_achievements/handler"
	achievementRepo "github.com/sawalreverr/recything/internal/achievements/manage_achievements/repository"
	achievementUsecase "github.com/sawalreverr/recything/internal/achievements/manage_achievements/usecase"
	userAchievementHandler "github.com/sawalreverr/recything/internal/achievements/user_achievements/handler"
	userAchievementRepo "github.com/sawalreverr/recything/internal/achievements/user_achievements/repository"
	userAchievementUsecase "github.com/sawalreverr/recything/internal/achievements/user_achievements/usecase"
	"github.com/sawalreverr/recything/internal/admin/handler"
	"github.com/sawalreverr/recything/internal/admin/repository"
	"github.com/sawalreverr/recything/internal/admin/usecase"
	articleHandler "github.com/sawalreverr/recything/internal/article/handler"
	articleRepository "github.com/sawalreverr/recything/internal/article/repository"
	articleUsecase "github.com/sawalreverr/recything/internal/article/usecase"
	authHandler "github.com/sawalreverr/recything/internal/auth/handler"
	authUsecase "github.com/sawalreverr/recything/internal/auth/usecase"
	customDataHandler "github.com/sawalreverr/recything/internal/custom-data/handler"
	customDataRepository "github.com/sawalreverr/recything/internal/custom-data/repository"
	customDataUsecase "github.com/sawalreverr/recything/internal/custom-data/usecase"
	dashboardHandler "github.com/sawalreverr/recything/internal/dashboard/handler"
	dashboardRepo "github.com/sawalreverr/recything/internal/dashboard/repository"
	dashboardUsecase "github.com/sawalreverr/recything/internal/dashboard/usecase"
	faqHandler "github.com/sawalreverr/recything/internal/faq/handler"
	faqRepo "github.com/sawalreverr/recything/internal/faq/repository"
	faqUsecase "github.com/sawalreverr/recything/internal/faq/usecase"
	homepageHandler "github.com/sawalreverr/recything/internal/homepage/handler"
	homepageRepo "github.com/sawalreverr/recything/internal/homepage/repository"
	homepageUsecase "github.com/sawalreverr/recything/internal/homepage/usecase"
	leaderboardHandler "github.com/sawalreverr/recything/internal/leaderboard/handler"
	leaderboardRepo "github.com/sawalreverr/recything/internal/leaderboard/repository"
	leaderboardUsecase "github.com/sawalreverr/recything/internal/leaderboard/usecase"
	"github.com/sawalreverr/recything/internal/middleware"
	reminaiHandler "github.com/sawalreverr/recything/internal/remin-ai/handler"
	reminaiUsecase "github.com/sawalreverr/recything/internal/remin-ai/usecase"
	reportHandler "github.com/sawalreverr/recything/internal/report/handler"
	reportRepo "github.com/sawalreverr/recything/internal/report/repository"
	reportUsecase "github.com/sawalreverr/recything/internal/report/usecase"
	approvalTaskHandler "github.com/sawalreverr/recything/internal/task/approval_task/handler"
	approvalTaskRepo "github.com/sawalreverr/recything/internal/task/approval_task/repository"
	approvalTaskUsecase "github.com/sawalreverr/recything/internal/task/approval_task/usecase"
	taskHandler "github.com/sawalreverr/recything/internal/task/manage_task/handler"
	taskRepo "github.com/sawalreverr/recything/internal/task/manage_task/repository"
	taskUsecase "github.com/sawalreverr/recything/internal/task/manage_task/usecase"
	userTaskHandler "github.com/sawalreverr/recything/internal/task/user_task/handler"
	userTaskRepo "github.com/sawalreverr/recything/internal/task/user_task/repository"
	userTaskUsecase "github.com/sawalreverr/recything/internal/task/user_task/usecase"
	userHandler "github.com/sawalreverr/recything/internal/user/handler"
	userRepo "github.com/sawalreverr/recything/internal/user/repository"
	userUsecase "github.com/sawalreverr/recything/internal/user/usecase"
	videoHandler "github.com/sawalreverr/recything/internal/video/manage_video/handler"
	videoRepo "github.com/sawalreverr/recything/internal/video/manage_video/repository"
	videoUsecase "github.com/sawalreverr/recything/internal/video/manage_video/usecase"
	userVideoHandler "github.com/sawalreverr/recything/internal/video/user_video/handler"
	userVideoRepo "github.com/sawalreverr/recything/internal/video/user_video/repository"
	userVideoUsecase "github.com/sawalreverr/recything/internal/video/user_video/usecase"
)

var (
	SuperAdminMiddleware        = middleware.RoleBasedMiddleware("super admin")
	SuperAdminOrAdminMiddleware = middleware.RoleBasedMiddleware("super admin", "admin")
	UserMiddleware              = middleware.RoleBasedMiddleware("user")
	AllRoleMiddleware           = middleware.RoleBasedMiddleware("super admin", "admin", "user")
)

func (s *echoServer) publicHttpHandler() {
	// Healthy Check
	s.app.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Swagger
	s.app.Static("/assets", "web/assets")
	s.app.Static("/docs", "docs")

	s.app.GET("/", func(c echo.Context) error {
		return c.File("web/index.html")
	})

	s.app.GET("/terms-and-conditions", func(c echo.Context) error {
		return c.File("web/syarat-dan-ketentuan.html")
	})

	// Example need user auth
	s.app.GET("/needUser", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}, UserMiddleware)
}

func (s *echoServer) authHttpHandler() {
	userRepository := userRepo.NewUserRepository(s.db)
	adminRepository := repository.NewAdminRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(userRepository, adminRepository)
	handler := authHandler.NewAuthHandler(usecase)

	// Register User
	s.gr.POST("/register", handler.Register)

	// Verify OTP after Register
	s.gr.POST("/verify-otp", handler.VerifyOTP)

	// Resend OTP
	s.gr.POST("/resend-otp", handler.ResendOTP)

	// Login User
	s.gr.POST("/login", handler.LoginUser)

	// Login Admin
	s.gr.POST("/admin/login", handler.LoginAdmin)
}

func (s *echoServer) userHttpHandler() {
	repository := userRepo.NewUserRepository(s.db)
	usecase := userUsecase.NewUserUsecase(repository)
	handler := userHandler.NewUserHandler(usecase)

	// Profile user based on JWT user token
	s.gr.GET("/user/profile", handler.Profile, UserMiddleware)

	// Edit detail user based on JWT user token
	s.gr.PUT("/user/profile", handler.UpdateDetail, UserMiddleware)

	// Upload avatar user based on JWT user token
	s.gr.POST("/user/uploadAvatar", handler.UploadAvatar, UserMiddleware)

	// Find user data using param userId, doesnt need jwt
	s.gr.GET("/user/:userId", handler.FindUser, AllRoleMiddleware)

	// Find all user data with pagination, need JWT admin or superadmin token
	s.gr.GET("/users", handler.FindAllUser, SuperAdminOrAdminMiddleware)

	// Delete user data using param userId
	s.gr.DELETE("/user/:userId", handler.DeleteUser, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) supAdminHttpHandler() {
	repository := repository.NewAdminRepository(s.db)
	usecase := usecase.NewAdminUsecase(repository)
	handler := handler.NewAdminHandler(usecase)

	// register admin by super admin
	s.gr.POST("/admin", handler.AddAdminHandler, SuperAdminMiddleware)

	// get all admin by super admin
	s.gr.GET("/admins", handler.GetDataAllAdminHandler, SuperAdminMiddleware)

	// get data admin by id by super admin
	s.gr.GET("/admin/:adminId", handler.GetDataAdminByIdHandler, SuperAdminMiddleware)

	// update admin by super admin
	s.gr.PATCH("/admin/:adminId", handler.UpdateAdminHandler, SuperAdminMiddleware)

	// delete admin by super admin
	s.gr.DELETE("/admin/:adminId", handler.DeleteAdminHandler, SuperAdminMiddleware)

	// get profile admin or super admin
	s.gr.GET("/admin/profile", handler.GetProfileAdminHandler, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) reportHttpHandler() {
	reportRepository := reportRepo.NewReportRepository(s.db)
	userRepository := userRepo.NewUserRepository(s.db)
	usecase := reportUsecase.NewReportUsecase(reportRepository, userRepository)
	handler := reportHandler.NewReportHandler(usecase)

	// User create new report
	s.gr.POST("/report", handler.NewReport, UserMiddleware)

	// User get all history reports
	s.gr.GET("/report", handler.GetHistoryUserReports, UserMiddleware)

	// Admin update status approved or reject
	s.gr.PUT("/report/:reportId", handler.UpdateStatus, SuperAdminOrAdminMiddleware)

	// Admin get all with pagination and filter
	s.gr.GET("/reports", handler.GetAllReports, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) faqHttpHandler() {
	repository := faqRepo.NewFaqRepository(s.db)
	usecase := faqUsecase.NewFaqUsecase(repository)
	handler := faqHandler.NewFaqHandler(usecase)

	// User get all faqs
	s.gr.GET("/faqs", handler.GetAllFaqs, UserMiddleware)

	// User get all faqs by category
	s.gr.GET("/faqs/category", handler.GetFaqsByCategory, UserMiddleware)

	// User get all faqs by keyword
	s.gr.GET("/faqs/search", handler.GetFaqsByKeyword, UserMiddleware)
}

func (s *echoServer) manageTask() {
	repository := taskRepo.NewManageTaskRepository(s.db)
	usecase := taskUsecase.NewManageTaskUsecase(repository)
	handler := taskHandler.NewManageTaskHandler(usecase)

	// create task by admin or super admin
	s.gr.POST("/tasks", handler.CreateTaskHandler, SuperAdminOrAdminMiddleware)

	// get task challenge by pagination
	s.gr.GET("/tasks", handler.GetTaskChallengePaginationHandler, SuperAdminOrAdminMiddleware)

	// get task challenge by id
	s.gr.GET("/tasks/:taskId", handler.GetTaskByIdHandler, SuperAdminOrAdminMiddleware)

	// update task challenge
	s.gr.PATCH("/tasks/:taskId", handler.UpdateTaskHandler, SuperAdminOrAdminMiddleware)

	// delete task challenge
	s.gr.DELETE("/tasks/:taskId", handler.DeleteTaskHandler, SuperAdminOrAdminMiddleware)

}

func (s *echoServer) userTask() {
	repository := userTaskRepo.NewUserTaskRepository(s.db)
	usecase := userTaskUsecase.NewUserTaskUsecase(repository)
	handler := userTaskHandler.NewUserTaskHandler(usecase)

	// get all tasks
	s.gr.GET("/user/tasks", handler.GetAllTasksHandler, UserMiddleware)

	// get task by id
	s.gr.GET("/user/tasks/:taskId", handler.GetTaskByIdHandler, UserMiddleware)

	// create task by user or start task
	s.gr.POST("/user/tasks/:taskChallengeId", handler.CreateUserTaskHandler, UserMiddleware)

	// get task in progress by user current
	s.gr.GET("/user-current/tasks/in-progress", handler.GetUserTaskByUserIdHandler, UserMiddleware)

	// send task done by user current
	s.gr.POST("/user-current/tasks/:userTaskId", handler.UploadImageTaskHandler, UserMiddleware)

	// get task done by user current
	s.gr.GET("/user-current/tasks/done", handler.GetUserTaskDoneByUserIdHandler, UserMiddleware)

	// update user task if reject
	s.gr.PUT("/user-current/tasks/:userTaskId", handler.UpdateUserTaskHandler, UserMiddleware)

	// get user task details if repair
	s.gr.GET("/user-current/task/:userTaskId", handler.GetUserTaskDetailsHandler, UserMiddleware)

	// get history point by user current
	s.gr.GET("/user-current/tasks/history", handler.GetHistoryPointByUserIdHandler, UserMiddleware)

	// update user task step
	s.gr.PUT("/user-current/steps", handler.UpdateTaskStepHandler, UserMiddleware)

	// get user task by user task id
	s.gr.GET("/user/task/:userTaskId", handler.GetUserTaskByUserTaskIdHandler, UserMiddleware)

	// get user task rejected by user current
	s.gr.GET("/user/tasks/rejected/:userTaskId", handler.GetUserTaskRejectedByUserIdHandler, UserMiddleware)
}

func (s *echoServer) approvalTask() {
	repository := approvalTaskRepo.NewApprovalTaskRepositoryImpl(s.db)
	usecase := approvalTaskUsecase.NewApprovalTaskUsecase(repository)
	handler := approvalTaskHandler.NewApprovalTaskHandler(usecase)

	// get all pagination user task
	s.gr.GET("/approval-tasks", handler.GetAllApprovalTaskPaginationHandler, SuperAdminOrAdminMiddleware)

	// approve user task
	s.gr.PUT("/approve-tasks/:userTaskId", handler.ApproveUserTaskHandler, SuperAdminOrAdminMiddleware)

	// reject user task
	s.gr.PUT("/reject-tasks/:userTaskId", handler.RejectUserTaskHandler, SuperAdminOrAdminMiddleware)

	// get user task details
	s.gr.GET("/user-task/:userTaskId", handler.GetUserTaskDetailsHandler, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) manageAchievement() {
	repository := achievementRepo.NewManageAchievementRepository(s.db)
	usecase := achievementUsecase.NewManageAchievementUsecase(repository)
	handler := achievementHandler.NewManageAchievementHandler(usecase)

	// get all achievement
	s.gr.GET("/achievements", handler.GetAllAchievementHandler, SuperAdminOrAdminMiddleware)

	// get achievement by id
	s.gr.GET("/achievements/:achievementId", handler.GetAchievementByIdHandler, SuperAdminOrAdminMiddleware)

	// update achievement
	s.gr.PATCH("/achievements/:achievementId", handler.UpdateAchievementHandler, SuperAdminOrAdminMiddleware)

	// delete achievement
	s.gr.DELETE("/achievements/:achievementId", handler.DeleteAchievementHandler, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) customDataHandler() {
	repository := customDataRepository.NewCustomDataRepository(s.db)
	usecase := customDataUsecase.NewCustomDataUsecase(repository)
	handler := customDataHandler.NewCustomDataHandler(usecase)

	// Create new custom data for admin
	s.gr.POST("/custom-data", handler.NewCustomData, SuperAdminOrAdminMiddleware)

	// Update custom data for admin
	s.gr.PUT("/custom-data/:dataId", handler.UpdateData, SuperAdminOrAdminMiddleware)

	// Delete custom data for admin
	s.gr.DELETE("/custom-data/:dataId", handler.DeleteData, SuperAdminOrAdminMiddleware)

	// Get custom data by id for admin
	s.gr.GET("/custom-data/:dataId", handler.GetDataByID, SuperAdminOrAdminMiddleware)

	// Get all custom data for admin
	s.gr.GET("/custom-datas", handler.GetAllData, SuperAdminMiddleware)
}

func (s *echoServer) reminAIHandler() {
	repository := customDataRepository.NewCustomDataRepository(s.db)
	usecase := reminaiUsecase.NewReminAIUsecase(repository)
	handler := reminaiHandler.NewReminAIHandler(usecase)

	// ReMin AI Chatbot with user access
	s.gr.POST("/remin-ai", handler.AskGPT, UserMiddleware)
}

func (s *echoServer) userAchievement() {
	repository := userAchievementRepo.NewUserAchievementRepository(s.db)
	usecase := userAchievementUsecase.NewUserAchievementUsecase(repository)
	handler := userAchievementHandler.NewUserAchievementHandler(usecase)

	// get achievement by user
	s.gr.GET("/user/achievements", handler.GetAvhievementsByUserhandler, UserMiddleware)
}

func (s *echoServer) manageVideo() {
	repository := videoRepo.NewManageVideoRepository(s.db)
	usecase := videoUsecase.NewManageVideoUsecaseImpl(repository)
	handler := videoHandler.NewManageVideoHandlerImpl(usecase)

	// create data video
	s.gr.POST("/videos/data", handler.CreateDataVideoHandler, SuperAdminOrAdminMiddleware)

	// get all category video
	s.gr.GET("/videos/categories", handler.GetAllCategoryVideoHandler, AllRoleMiddleware)

	// get all data video pagination
	s.gr.GET("/videos/data", handler.GetAllDataVideoPaginationHandler, SuperAdminOrAdminMiddleware)

	// get details data video by id
	s.gr.GET("/videos/data/:videoId", handler.GetDetailsDataVideoByIdHandler, SuperAdminOrAdminMiddleware)

	// update data video
	s.gr.PATCH("/videos/data/:videoId", handler.UpdateDataVideoHandler, SuperAdminOrAdminMiddleware)

	// delete data video
	s.gr.DELETE("/videos/data/:videoId", handler.DeleteDataVideoHandler, SuperAdminOrAdminMiddleware)
}

func (s *echoServer) userVideo() {
	repository := userVideoRepo.NewUserVideoRepository(s.db)
	usecase := userVideoUsecase.NewUserVideoUsecase(repository)
	handler := userVideoHandler.NewUserVideoHandler(usecase)

	// get all video
	s.gr.GET("/videos", handler.GetAllVideoHandler, UserMiddleware)

	// search video by title
	s.gr.GET("/videos/search", handler.SearchVideoByKeywordHandler, UserMiddleware)

	// search video by category
	s.gr.GET("/videos/category", handler.SearchVideoByCategoryHandler, UserMiddleware)

	// get video detail
	s.gr.GET("/video/:videoId", handler.GetVideoDetailHandler, UserMiddleware)

	// add comment
	s.gr.POST("/videos/comment", handler.AddCommentHandler, UserMiddleware)
}

func (s *echoServer) aboutUsHandler() {
	repository := aboutusRepo.NewAboutUsRepository(s.db)
	usecase := aboutusUsecase.NewAboutUsUsecase(repository)
	handler := aboutusHandler.NewAboutUsHandler(usecase)

	// Get about us by category
	s.gr.GET("/about-us/category", handler.GetAboutUsByCategory, UserMiddleware)
}

func (s *echoServer) leaderboardHandler() {
	repository := leaderboardRepo.NewLeaderboardRepository(s.db)
	usecase := leaderboardUsecase.NewLeaderboardUsecase(repository)
	handler := leaderboardHandler.NewLeaderboardHandler(usecase)

	// Get leaderboard
	s.gr.GET("/leaderboard", handler.GetLeaderboardHandler, AllRoleMiddleware)
}

func (s *echoServer) articleHandler() {
	repositoryArticle := articleRepository.NewArticleRepository(s.db)
	repositoryAdmin := repository.NewAdminRepository(s.db)
	repositoryUser := userRepo.NewUserRepository(s.db)
	usecase := articleUsecase.NewArticleUsecase(repositoryArticle, repositoryAdmin, repositoryUser)
	handler := articleHandler.NewArticleHandler(usecase)

	// Get all article
	s.gr.GET("/articles", handler.GetAllArticle, AllRoleMiddleware)

	// Get by keyword
	s.gr.GET("/article/search", handler.GetArticleByKeyword, AllRoleMiddleware)

	// Get by category
	s.gr.GET("/article/category", handler.GetArticleByCategory, AllRoleMiddleware)

	// Get article by id
	s.gr.GET("/article/:articleId", handler.GetArticleByID, AllRoleMiddleware)

	// Create new article by admin
	s.gr.POST("/article", handler.NewArticle, SuperAdminOrAdminMiddleware)

	// Update article by admin
	s.gr.PUT("/article/:articleId", handler.UpdateArticle, SuperAdminOrAdminMiddleware)

	// Delete article by admin
	s.gr.DELETE("/article/:articleId", handler.DeleteArticle, SuperAdminOrAdminMiddleware)

	// Add new comment by user
	s.gr.POST("/article/:articleId/comment", handler.NewArticleComment, UserMiddleware)

	// Upload image
	s.gr.POST("/article/upload", handler.ArticleUploadImage, SuperAdminOrAdminMiddleware)

	// Get all categories
	s.gr.GET("/categories", handler.GetAllCategories)
}

func (s *echoServer) homepageHandler() {
	repository := homepageRepo.NewHomepageRepository(s.db)
	usecase := homepageUsecase.NewHomepageUsecase(repository)
	handler := homepageHandler.NewHomePageHandler(usecase)

	// Get homepage
	s.gr.GET("/homepage", handler.GetHomepageHandler, UserMiddleware)
}

func (s *echoServer) dashboardHandler() {
	repository := dashboardRepo.NewDashboardRepository(s.db)
	usecase := dashboardUsecase.NewDashboardUsecase(repository)
	handler := dashboardHandler.NewDashboardHandler(usecase)

	// Get dashboard
	s.gr.GET("/dashboards", handler.GetDashboardHandler, SuperAdminOrAdminMiddleware)
}
