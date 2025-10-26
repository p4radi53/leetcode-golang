package medium

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{
		balance: balance,
	}
}

func (this *Bank) isValidAccount(account int) bool {
	if account > len(this.balance) || account <= 0 {
		return false
	} else {
		return true
	}
}

func (this *Bank) hasEnoughMoney(account int, money int64) bool {
	if this.balance[account-1] < money {
		return false
	} else {
		return true
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.isValidAccount(account1) || !this.isValidAccount(account2) {
		return false
	}

	if !this.hasEnoughMoney(account1, money) {
		return false
	}

	this.balance[account1-1] = this.balance[account1-1] - money
	this.balance[account2-1] = this.balance[account2-1] + money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.isValidAccount(account) {
		return false
	}
	this.balance[account-1] = this.balance[account-1] + money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.isValidAccount(account) {
		return false
	}

	if !this.hasEnoughMoney(account, money) {
		return false
	}

	this.balance[account-1] = this.balance[account-1] - money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
