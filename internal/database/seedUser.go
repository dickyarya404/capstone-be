package database

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/sawalreverr/recything/internal/helper"
	userEntity "github.com/sawalreverr/recything/internal/user"
)

var addresses = []string{
	"Jalan Jendral Sudirman, Jakarta, Indonesia",
	"Jalan MH Thamrin, Jakarta, Indonesia",
	"Jalan Gatot Subroto, Jakarta, Indonesia",
	"Jalan Medan Merdeka Selatan, Jakarta, Indonesia",
	"Jalan Merdeka Utara, Jakarta, Indonesia",
	"Jalan Kuningan, Jakarta, Indonesia",
	"Jalan Mangga Dua, Jakarta, Indonesia",
	"Jalan Kebon Jeruk, Jakarta, Indonesia",
	"Jalan Panglima Polim, Jakarta, Indonesia",
	"Jalan Senopati, Jakarta, Indonesia",
	"Jalan Braga, Bandung, Indonesia",
	"Jalan Asia Afrika, Bandung, Indonesia",
	"Jalan Dago, Bandung, Indonesia",
	"Jalan Riau, Bandung, Indonesia",
	"Jalan Malioboro, Yogyakarta, Indonesia",
	"Jalan Solo, Yogyakarta, Indonesia",
	"Jalan Cendrawasih, Surabaya, Indonesia",
	"Jalan Tunjungan, Surabaya, Indonesia",
	"Jalan Diponegoro, Surabaya, Indonesia",
	"Jalan Darmo, Surabaya, Indonesia",
}

func randomBadge(points int) string {
	if points <= 50000 {
		return "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189121/user_badge/htaemsjtlhfof7ww01ss.png"
	} else if points <= 150000 {
		return "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189221/user_badge/oespnjdgoynkairlutbk.png"
	} else if points <= 300000 {
		return "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189184/user_badge/jshs1s2fwevahgtvjkgj.png"
	} else {
		return "https://res.cloudinary.com/dymhvau8n/image/upload/v1718188250/user_badge/icureiapdvtzyu5b99zu.png"
	}
}

func randomDate(start, end time.Time) time.Time {
	delta := end.Sub(start)
	seconds := rand.Int63n(int64(delta.Seconds()))
	return start.Add(time.Duration(seconds) * time.Second)
}

func randomGender() string {
	genders := []string{"laki-laki", "perempuan"}
	return genders[rand.Intn(len(genders))]
}

func generateUser(password string) []userEntity.User {
	gofakeit.Seed(0)
	users := make([]userEntity.User, 50)
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 6, 30, 23, 59, 59, 999, time.UTC)

	for i := 0; i < 50; i++ {
		points := rand.Intn(400001)
		users[i] = userEntity.User{
			ID:         fmt.Sprintf("USR%04d", i+1),
			Name:       gofakeit.Name(),
			Email:      gofakeit.Email(),
			Password:   password,
			Point:      uint(points),
			Gender:     randomGender(),
			BirthDate:  gofakeit.DateRange(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2006, 12, 31, 23, 59, 59, 999, time.UTC)),
			Address:    addresses[rand.Intn(len(addresses))],
			PictureURL: gofakeit.ImageURL(200, 200),
			OTP:        uint(gofakeit.Number(100000, 999999)),
			IsVerified: true,
			Badge:      randomBadge(points),
			CreatedAt:  randomDate(startDate, endDate),
		}
	}

	return users
}

func (m *mysqlDatabase) InitUser() {
	if err := m.GetDB().Migrator().DropTable(&userEntity.User{}); err != nil {
		return
	}

	if err := m.GetDB().AutoMigrate(&userEntity.User{}); err != nil {
		return
	}

	hashed, _ := helper.GenerateHash("password@123")
	users := generateUser(hashed)

	for _, user := range users {
		m.GetDB().FirstOrCreate(&user, user)
	}

	log.Println("User data added!")
}
