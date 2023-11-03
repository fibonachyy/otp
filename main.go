package main

import (
	"fmt"
	"otp/services/otp"
	"otp/services/otp/providers/kavenegar"
	"time"
)

func Sum(a int64, b int64) int64 {
	sum := a + b
	return sum
}
func Mul(a int64, b int64) int64 {
	mul := a * b
	return mul
}
func Min(a, b int64) int64 {
	min := a - b
	return min
}
func Div(a int64, b int64) int64 {
	div := a / b
	return div
}

//  min , div , mul

func AddTax(price, percentTax int64) int64 {
	// tax := price / 100 * percentTax
	d := Div(price, 100)
	tax := Mul(d, percentTax)
	return Sum(price, tax)

}
func AddDeliveryCost(price, cost int64) int64 {
	return Sum(price, cost)
}

type Test struct {
	OTP otp.IOtpService
}

func newTest(o otp.IOtpService) *Test {
	return &Test{
		OTP: o,
	}
}

func main() {
	key := fmt.Sprintf("6677586D34555975446F6D5A6432646D45755A6C4F727066663252396752776E6B4151744F4455425A54733D")
	url := fmt.Sprintf("https://api.kavenegar.com/v1")
	sender := fmt.Sprintf("10008663")
	provider := kavenegar.New(url, key, sender, time.Minute)
	o := otp.New(provider)
	otpService := newTest(o)
	newOTP := otpService.OTP.NewOTP(6)
	fmt.Println(newOTP)
	err := otpService.OTP.SendSMS(newOTP, fmt.Sprintf("09215369016"))
	if err != nil {
		panic(err)
	}
}
