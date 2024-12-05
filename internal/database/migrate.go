package database

import (
	"log"

	aboutus "github.com/sawalreverr/recything/internal/about-us"
	achievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/admin/entity"
	"github.com/sawalreverr/recything/internal/article"
	customdata "github.com/sawalreverr/recything/internal/custom-data"
	"github.com/sawalreverr/recything/internal/faq"
	"github.com/sawalreverr/recything/internal/report"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	user "github.com/sawalreverr/recything/internal/user"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(
		&user.User{},
		&entity.Admin{},

		&report.Report{},
		&report.WasteMaterial{},
		&report.ReportWasteMaterial{},
		&report.ReportImage{},

		&article.WasteCategory{},
		&article.ContentCategory{},

		&faq.FAQ{},
		&achievement.Achievement{},
		&customdata.CustomData{},
		&aboutus.AboutUs{},
		&aboutus.AboutUsImage{},

		&task.TaskChallenge{},
		&task.TaskStep{},

		&user_task.UserTaskChallenge{},
		&user_task.UserTaskImage{},
		&user_task.UserTaskStep{},

		&video.Video{},
		&video.VideoCategory{},
		&video.Comment{},

		&article.Article{},
		&article.ArticleSection{},
		&article.ArticleCategories{},
		&article.ArticleComment{},
	); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
