package helper

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/sawalreverr/recything/config"
	"gopkg.in/gomail.v2"
)

func SendMail(receiverEmail string, otp uint) error {
	conf := config.GetConfig()

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "Recything <service@recything.com>")
	mailer.SetHeader("To", receiverEmail)
	mailer.SetHeader("Subject", "Recything - OTP Verififcation")

	msg := fmt.Sprintf("Hello, This is your OTP <b>%v</b>", otp)
	mailer.SetBody("text/html", msg)

	dialer := gomail.NewDialer(
		conf.SMTP.Host,
		conf.SMTP.Port,
		conf.SMTP.AuthEmail,
		conf.SMTP.AuthPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

func GenerateOTP() uint {
	const digits = "0123456789"
	const length = 6

	var otp string
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		otp += string(digits[num.Int64()])
	}

	var otpInt uint
	fmt.Sscanf(otp, "%d", &otpInt)

	return otpInt
}
