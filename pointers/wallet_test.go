package main

import (
	"testing"
)

func TestWall(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)

		// 理解 fmt %s 与 Stringer interface
		// fmt.Printf("Want Balance: %d", want)
		// fmt.Printf("Want Balance: %s", want)
		// fmt.Println("Want Balance: ", want)

	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcom(20)}
		err := wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcom(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcom(20)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(100)

		assertBalance(t, wallet, startBalance)
		assertError(t, err, InsufficientFundError)
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcom) {
	got := wallet.balance
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("get an error but didn't want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		// t.Fatal 被调用将停止测试，不需要返回更多的错误
		// 否则，got.Error 将会因为空指针而引起Panic
		t.Fatal("didn't get an error")
	}
	if got != want {
		t.Errorf("got '%s', wnat '%s'", got, want)
	}
}
