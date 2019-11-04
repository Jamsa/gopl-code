package bank

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func teller() {
	var balance int // 限定在teller goroutine中
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
