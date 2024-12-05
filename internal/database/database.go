package database

import "gorm.io/gorm"

type Database interface {
	GetDB() *gorm.DB
	InitSuperAdmin()
	InitUser()

	InitWasteCategories()
	InitContentCategories()
	InitWasteMaterials()

	InitFaqs()
	InitCustomDatas()
	InitAboutUs()
	InitAchievements()

	InitTasks()
	InitVideos()
	InitArticle()
	InitReport()

	InitComment()
}
