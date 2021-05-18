package account

import (
	"sync"
)

type Account struct {
	sync.RWMutex
	closed bool
	balance int64
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{closed:false, balance:initialDeposit}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.RLock()
	defer a.RUnlock()

	if a.closed {
		return 0, false
	}
	if a.balance < 0 {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed || a.balance + amount < 0 {
		return 0, false
	}

	a.balance += amount

	return a.balance, true
}
