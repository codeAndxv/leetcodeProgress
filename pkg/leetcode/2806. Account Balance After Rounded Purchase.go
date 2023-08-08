package leetcode

func AccountBalanceAfterPurchase(purchaseAmount int) int {
	amount1 := (purchaseAmount / 10) * 10
	amount2 := (purchaseAmount/10 + 1) * 10
	var amount int
	if amount2-purchaseAmount <= purchaseAmount-amount1 {
		amount = amount2
	} else {
		amount = amount1
	}
	return 100 - amount
}
