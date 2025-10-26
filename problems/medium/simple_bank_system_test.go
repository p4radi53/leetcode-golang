package medium

import "testing"

func TestConstructor(t *testing.T) {
	tests := []struct {
		name    string
		balance []int64
	}{
		{
			name:    "create bank with single account",
			balance: []int64{100},
		},
		{
			name:    "create bank with multiple accounts",
			balance: []int64{10, 100, 20, 50, 30},
		},
		{
			name:    "create bank with empty balance",
			balance: []int64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := Constructor(tt.balance)
			if len(bank.balance) != len(tt.balance) {
				t.Errorf("Constructor() balance length = %v, want %v", len(bank.balance), len(tt.balance))
			}
			for i := range tt.balance {
				if bank.balance[i] != tt.balance[i] {
					t.Errorf("Constructor() balance[%d] = %v, want %v", i, bank.balance[i], tt.balance[i])
				}
			}
		})
	}
}

func TestIsValidAccount(t *testing.T) {
	bank := Constructor([]int64{10, 100, 20, 50, 30})

	tests := []struct {
		name    string
		account int
		want    bool
	}{
		{
			name:    "invalid account 0",
			account: 0,
			want:    false,
		},
		{
			name:    "valid account first",
			account: 1,
			want:    true,
		},
		{
			name:    "valid account in middle",
			account: 3,
			want:    true,
		},
		{
			name:    "valid account last",
			account: 5,
			want:    true,
		},
		{
			name:    "invalid account negative",
			account: -1,
			want:    false,
		},
		{
			name:    "invalid account out of bounds",
			account: 6,
			want:    false,
		},
		{
			name:    "invalid account way out of bounds",
			account: 100,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bank.isValidAccount(tt.account); got != tt.want {
				t.Errorf("isValidAccount(%d) = %v, want %v", tt.account, got, tt.want)
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	tests := []struct {
		name           string
		initialBalance []int64
		account        int
		money          int64
		want           bool
		expectedBalance int64
	}{
		{
			name:           "deposit to valid account",
			initialBalance: []int64{10, 100, 20},
			account:        2,
			money:          50,
			want:           true,
			expectedBalance: 150,
		},
		{
			name:           "deposit zero",
			initialBalance: []int64{10, 100, 20},
			account:        1,
			money:          0,
			want:           true,
			expectedBalance: 10,
		},
		{
			name:           "deposit large amount",
			initialBalance: []int64{10, 100, 20},
			account:        3,
			money:          1000000,
			want:           true,
			expectedBalance: 1000020,
		},
		{
			name:           "deposit to invalid account negative",
			initialBalance: []int64{10, 100, 20},
			account:        -1,
			money:          50,
			want:           false,
			expectedBalance: 0,
		},
		{
			name:           "deposit to invalid account zero",
			initialBalance: []int64{10, 100, 20},
			account:        0,
			money:          50,
			want:           false,
			expectedBalance: 0,
		},
		{
			name:           "deposit to invalid account out of bounds",
			initialBalance: []int64{10, 100, 20},
			account:        4,
			money:          50,
			want:           false,
			expectedBalance: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := Constructor(tt.initialBalance)
			got := bank.Deposit(tt.account, tt.money)

			if got != tt.want {
				t.Errorf("Deposit(%d, %d) = %v, want %v", tt.account, tt.money, got, tt.want)
			}

			if tt.want && bank.balance[tt.account-1] != tt.expectedBalance {
				t.Errorf("After Deposit(%d, %d), balance[%d] = %v, want %v",
					tt.account, tt.money, tt.account, bank.balance[tt.account-1], tt.expectedBalance)
			}
		})
	}
}

func TestWithdraw(t *testing.T) {
	tests := []struct {
		name           string
		initialBalance []int64
		account        int
		money          int64
		want           bool
		expectedBalance int64
	}{
		{
			name:           "withdraw with sufficient funds",
			initialBalance: []int64{10, 100, 20},
			account:        2,
			money:          50,
			want:           true,
			expectedBalance: 50,
		},
		{
			name:           "withdraw exact balance",
			initialBalance: []int64{10, 100, 20},
			account:        1,
			money:          10,
			want:           true,
			expectedBalance: 0,
		},
		{
			name:           "withdraw with insufficient funds",
			initialBalance: []int64{10, 100, 20},
			account:        3,
			money:          30,
			want:           false,
			expectedBalance: 20,
		},
		{
			name:           "withdraw from invalid account negative",
			initialBalance: []int64{10, 100, 20},
			account:        -1,
			money:          5,
			want:           false,
			expectedBalance: 0,
		},
		{
			name:           "withdraw from invalid account zero",
			initialBalance: []int64{10, 100, 20},
			account:        0,
			money:          5,
			want:           false,
			expectedBalance: 0,
		},
		{
			name:           "withdraw from invalid account out of bounds",
			initialBalance: []int64{10, 100, 20},
			account:        4,
			money:          5,
			want:           false,
			expectedBalance: 0,
		},
		{
			name:           "withdraw zero",
			initialBalance: []int64{10, 100, 20},
			account:        2,
			money:          0,
			want:           true,
			expectedBalance: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := Constructor(tt.initialBalance)
			got := bank.Withdraw(tt.account, tt.money)

			if got != tt.want {
				t.Errorf("Withdraw(%d, %d) = %v, want %v", tt.account, tt.money, got, tt.want)
			}

			if tt.want && bank.balance[tt.account-1] != tt.expectedBalance {
				t.Errorf("After Withdraw(%d, %d), balance[%d] = %v, want %v",
					tt.account, tt.money, tt.account, bank.balance[tt.account-1], tt.expectedBalance)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	tests := []struct {
		name              string
		initialBalance    []int64
		account1          int
		account2          int
		money             int64
		want              bool
		expectedBalance1  int64
		expectedBalance2  int64
	}{
		{
			name:             "successful transfer",
			initialBalance:   []int64{10, 100, 20},
			account1:         2,
			account2:         3,
			money:            50,
			want:             true,
			expectedBalance1: 50,
			expectedBalance2: 70,
		},
		{
			name:             "transfer all funds",
			initialBalance:   []int64{10, 100, 20},
			account1:         1,
			account2:         2,
			money:            10,
			want:             true,
			expectedBalance1: 0,
			expectedBalance2: 110,
		},
		{
			name:             "transfer with insufficient funds",
			initialBalance:   []int64{10, 100, 20},
			account1:         3,
			account2:         1,
			money:            30,
			want:             false,
			expectedBalance1: 20,
			expectedBalance2: 10,
		},
		{
			name:             "transfer from invalid account negative",
			initialBalance:   []int64{10, 100, 20},
			account1:         -1,
			account2:         2,
			money:            10,
			want:             false,
			expectedBalance1: 0,
			expectedBalance2: 100,
		},
		{
			name:             "transfer from invalid account zero",
			initialBalance:   []int64{10, 100, 20},
			account1:         0,
			account2:         2,
			money:            10,
			want:             false,
			expectedBalance1: 0,
			expectedBalance2: 100,
		},
		{
			name:             "transfer to invalid account",
			initialBalance:   []int64{10, 100, 20},
			account1:         2,
			account2:         6,
			money:            10,
			want:             false,
			expectedBalance1: 100,
			expectedBalance2: 0,
		},
		{
			name:             "transfer to same account",
			initialBalance:   []int64{10, 100, 20},
			account1:         2,
			account2:         2,
			money:            50,
			want:             true,
			expectedBalance1: 100,
			expectedBalance2: 100,
		},
		{
			name:             "transfer zero",
			initialBalance:   []int64{10, 100, 20},
			account1:         1,
			account2:         3,
			money:            0,
			want:             true,
			expectedBalance1: 10,
			expectedBalance2: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := Constructor(tt.initialBalance)
			got := bank.Transfer(tt.account1, tt.account2, tt.money)

			if got != tt.want {
				t.Errorf("Transfer(%d, %d, %d) = %v, want %v",
					tt.account1, tt.account2, tt.money, got, tt.want)
			}

			if tt.want {
				if bank.balance[tt.account1-1] != tt.expectedBalance1 {
					t.Errorf("After Transfer(%d, %d, %d), balance[%d] = %v, want %v",
						tt.account1, tt.account2, tt.money, tt.account1,
						bank.balance[tt.account1-1], tt.expectedBalance1)
				}
				if bank.balance[tt.account2-1] != tt.expectedBalance2 {
					t.Errorf("After Transfer(%d, %d, %d), balance[%d] = %v, want %v",
						tt.account1, tt.account2, tt.money, tt.account2,
						bank.balance[tt.account2-1], tt.expectedBalance2)
				}
			}
		})
	}
}

func TestBankIntegrationScenario(t *testing.T) {
	// Test a complete scenario with multiple operations
	bank := Constructor([]int64{10, 100, 20, 50, 30})

	// Withdraw from account 3 (balance[2] = 20)
	if !bank.Withdraw(3, 10) {
		t.Error("Expected withdraw to succeed")
	}
	if bank.balance[2] != 10 {
		t.Errorf("balance[2] = %v, want 10", bank.balance[2])
	}

	// Transfer from account 5 to account 2 (balance[4] = 30, balance[1] = 100)
	if !bank.Transfer(5, 2, 20) {
		t.Error("Expected transfer to succeed")
	}
	if bank.balance[4] != 10 {
		t.Errorf("balance[4] = %v, want 10", bank.balance[4])
	}
	if bank.balance[1] != 120 {
		t.Errorf("balance[1] = %v, want 120", bank.balance[1])
	}

	// Deposit to account 1 (balance[0] = 10)
	if !bank.Deposit(1, 20) {
		t.Error("Expected deposit to succeed")
	}
	if bank.balance[0] != 30 {
		t.Errorf("balance[0] = %v, want 30", bank.balance[0])
	}

	// Try to withdraw more than available from account 4 (balance[3] = 50) (should fail)
	if bank.Withdraw(4, 60) {
		t.Error("Expected withdraw to fail due to insufficient funds")
	}
	if bank.balance[3] != 50 {
		t.Errorf("balance[3] = %v, want 50 (unchanged)", bank.balance[3])
	}

	// Transfer from invalid account (should fail)
	if bank.Transfer(10, 2, 5) {
		t.Error("Expected transfer to fail due to invalid account")
	}
}
