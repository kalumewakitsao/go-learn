package pointers

import "errors"
import "fmt"

type Bitcoin int

type Stringer interface {
    String() string
}

type Wallet struct {
    balance Bitcoin
}

var ErrorInsufficientFunds = errors.New("Sorry,  have insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
    if amount > w.balance {
        return ErrorInsufficientFunds
    }

    w.balance -= amount
    return nil
}

func (w *Wallet) Balance() Bitcoin{
    return w.balance
}

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}
