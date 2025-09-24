package payments

// структура описания информации об оплате 
type PaymentInfo struct{
	Description string
	Usd int
	Cancelled bool
}