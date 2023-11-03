package otp

type IProvider interface {
	SendSMS(otp string, receptor string) error
}

type IOtpService interface {
	NewOTP(digit int) string
	SendSMS(otp string, receptor string) error
	// VerifyOtp(otp string, receptor string) (bool, error)
}
