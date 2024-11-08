package main

/* feature that allows you to define a contract for behavior */
// type Payment interface {
// 	Pay(int)
// }

// no need to use implement keyword
// type CreditCard struct {
// }

// func (c CreditCard) Pay(v int) {
// 	fmt.Printf("pay with credit card: %v\n", v)
// }

// type QRCode struct {
// }

// func (q QRCode) Pay(v int) {
// 	fmt.Printf("pay with qr code: %v\n", v)
// }

/* ----- implement interface for external library ----*/
// external lib can't modify source code
// type MobileBanking struct {
// }

// func (m *MobileBanking) Transfer(v int) {
// 	fmt.Printf("Transfer: %v", v)
// }

// type MyTransfer interface {
// 	Transfer(int)
// }

func main() {

	// c := CreditCard{}

	// payToCompany(c)

	// q := QRCode{}

	// payToCompany(q)

}

// func payToCompany(p Payment) {
// 	price := 100
// 	p.Pay(price)
// }

// func myPayment(m MyTransfer) {
// 	m.Transfer(100)
// }
