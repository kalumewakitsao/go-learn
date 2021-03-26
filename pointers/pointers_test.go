package pointers

import "testing"

func TestWallet(t *testing.T) {

    assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
        t.Helper()
        got := wallet.Balance()

        if got != want {
            t.Errorf("got %s but wanted %s", got, want)
        }
    }

    assertError := func(t testing.TB, got error, want error) {
        t.Helper()
        if got != want {
            t.Errorf("got %q, wanted %q", got, want)
        }

        if got == nil {
            t.Fatal("wanted error but didn't get one")
        }
    }

    t.Run("Test Depositing", func(t *testing.T){
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))

        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Test Withdraw", func(t *testing.T){
        wallet := Wallet{balance: Bitcoin(100)}
        wallet.Withdraw(Bitcoin(20))

        assertBalance(t, wallet, Bitcoin(80))
    })

    t.Run("Test Withdraw: Insufficient Funds", func(t *testing.T){
        startingBalance := Bitcoin(100)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Bitcoin(120))

        assertBalance(t, wallet, startingBalance)
        assertError(t, err, ErrorInsufficientFunds)
    })
}
