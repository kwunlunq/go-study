package main

func main() {

	payAmount := 1000

	cathayPay := NewPayment("cathay pay", "4284300128421824", payAmount, &CreditCard{})
	cathayPay.Pay()

	ctbcPay := NewPayment("ctbc pay", "4284300128421824", payAmount, &CreditCard{})
	ctbcPay.Pay()

	linePay := NewPayment("line pay", "4284300128421824", payAmount, &QRCode{})
	linePay.Pay()

	jkoPay := NewPayment("jko pay", "4284300128421824", payAmount, &QRCode{})
	jkoPay.Pay()

	println("===")

	linePay.SetStrategy(&CreditCard{})
	linePay.Pay()
}
