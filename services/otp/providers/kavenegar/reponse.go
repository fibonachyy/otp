package kavenegar

type SendOtpResponse struct {
	Return  ReturnStatus `json:"return"`
	Entries []Entries    `json:"entries"`
}
type ReturnStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type Entries struct {
	MessageID  int    `json:"messageid"`
	Message    string `json:"message"`
	Status     int    `json:"status"`
	StatusText string `json:"statustext"`
	Sender     string `json:"sender"`
	Receptor   string `json:"receptor"`
	Date       int    `json:"date"`
	Cost       int    `json:"cost"`
}
