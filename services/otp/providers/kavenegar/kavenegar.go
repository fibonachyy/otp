package kavenegar

import (
	"fmt"
	"time"
)

type Kavenegar struct {
	Http httpClient
}

func New(url, apiKey, sender string, timeout time.Duration) *Kavenegar {
	client := newHttpClient(url, apiKey, sender, timeout)
	return &Kavenegar{
		Http: *client,
	}
}

func (k Kavenegar) SendSMS(otp string, receptor string) error {
	r := []string{receptor}
	response, err := k.Http.SendOtp(r, otp)
	fmt.Println(response)
	if err != nil {
		return err
	}
	return nil
}
