package server

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sawalreverr/recything/config"
	"github.com/sawalreverr/recything/internal/database"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
	gr   *echo.Group
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	app := echo.New()
	app.Validator = &CustomValidator{validator: validator.New()}

	group := app.Group("/api/v1")

	return &echoServer{
		app:  app,
		db:   db,
		conf: conf,
		gr:   group,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method} | ${status}] ~ ${uri} ~ ${remote_ip} ~ ${latency_human}\n",
	}))
	s.app.Use(middleware.CORS())

	// Public Handler
	s.publicHttpHandler()

	// Authentication Handler
	s.authHttpHandler()

	// Users Handler
	s.userHttpHandler()

	// super admin handler
	s.supAdminHttpHandler()

	// report handler
	s.reportHttpHandler()

	// FAQs Handler
	s.faqHttpHandler()

	// manage task handler
	s.manageTask()

	// user task handler
	s.userTask()

	// approval task handler
	s.approvalTask()

	// manage achievement handler
	s.manageAchievement()

	// reminai handler
	s.reminAIHandler()

	// custom data handler
	s.customDataHandler()

	// user achievement handler
	s.userAchievement()

	// manage video handler
	s.manageVideo()

	// user video handler
	s.userVideo()

	// About-us handler
	s.aboutUsHandler()

	// leaderboard handler
	s.leaderboardHandler()

	// Article Handler
	s.articleHandler()

	// homepage handler
	s.homepageHandler()

	// dashboard handler
	s.dashboardHandler()

	serverPORT := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverPORT))
}
