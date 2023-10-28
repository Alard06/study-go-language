package customer

type Customer struct {
	Name         string
	Age          int
	balance      int
	Debt         int
	Discount     int
	CalcDiscount func() (int, error)
}

func New(name string, age, balance, Debt, Discount int) {

}
