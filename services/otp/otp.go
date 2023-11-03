package otp

import (
	"crypto/rand"
	"fmt"
	"io"
)

type OTPService struct {
	Provider IProvider
	Redis    string
}

func New(p IProvider) *OTPService {
	return &OTPService{
		Provider: p,
		Redis:    fmt.Sprintf("redis"),
	}
}
func (o OTPService) NewOTP(digit int) string {
	return encodeToString(digit)
}
func (o OTPService) SendSMS(otp string, receptor string) error {

	err := o.Provider.SendSMS(otp, receptor)
	if err != nil {
		return fmt.Errorf("failed to send SMS: %v", err)
	}
	return nil
}
func encodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
