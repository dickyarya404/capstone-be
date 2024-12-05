package helper

func BonusTask(badegUser string, userPoint int) int {
	switch badegUser {
	case "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189121/user_badge/htaemsjtlhfof7ww01ss.png":
		return userPoint + userPoint*10/100
	case "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189221/user_badge/oespnjdgoynkairlutbk.png":
		return userPoint + userPoint*15/100
	case "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189184/user_badge/jshs1s2fwevahgtvjkgj.png":
		return userPoint + userPoint*20/100
	case "https://res.cloudinary.com/dymhvau8n/image/upload/v1718188250/user_badge/icureiapdvtzyu5b99zu.png":
		return userPoint + userPoint*25/100
	default:
		return userPoint
	}
}
