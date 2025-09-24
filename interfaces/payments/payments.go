package payments

type PaymentMetod interface {
	Pay(usd int) int
	Cansel(id int)
}

type PaymentModul struct {
	paymentMetod PaymentMetod
	// как сохранять информацию об оплатах
	paymentsInfo map[int]PaymentInfo
}

func NewPaymentModul(paymentMetod PaymentMetod) *PaymentModul {
	return &PaymentModul{
		paymentsInfo: map[int]PaymentInfo{},
		paymentMetod: paymentMetod,
	}
}

// принимает описание операции, сумму оплаты - возвращает ID проведенной операции
func (p PaymentModul) Pay(description string, usd int) int {
	//проводить оплpayments
	// получать ID операции
	id := p.paymentMetod.Pay(usd)

	// сохранять информацию о проведенной опеарции
	// -описание операции
	// - сколько потрачено
	// - была ли отменена данная операция
	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}

	p.paymentsInfo[id] = info

	// операции возвращять id

	return id

}

// принимает ID операции которую необходимо отменить
func (p PaymentModul) Cansel(id int) {
	p.paymentMetod.Cansel(id)

	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}
	//отменили
	info.Cancelled = true

	// и запомнили в map
	p.paymentsInfo[id] = info
}

// принимает ID операции которую необходимо отменить, возвращает инфо об этой операции
func (p PaymentModul) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

// возвращает инфо о всех проведенных операциях
func (p PaymentModul) AllInfo() map[int]PaymentInfo {
	//создаем копию map и передаем скопированные данные,
	// что бы исходные данные остались без изменений

	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))

	// пробегаем циклом по оригинальной мапе и все ключи и данные записываем в новыю tempMap
	for key, val := range p.paymentsInfo {
		tempMap[key] = val
	}
	return tempMap
}
