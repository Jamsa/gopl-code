package bank

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

func Balance() int {
	sema <- struct{}{} //获取锁
	b := balance
	<-sema //释放锁
	return b
}
