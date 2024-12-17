package exmath

import "fmt"

// 个人存款计算器，参考：https://fin.paas.cmbchina.com/fininfo/calsaving
func Income(startAmount int, monthlyAmount int, year int, yearlyPer float64) {
	amount := float64(startAmount)
	interest := float64(0)
	fmt.Printf("启动资金=%0.2f,每月投入=%d,年利率=%0.2f%%,年限=%d\n", amount, monthlyAmount, yearlyPer, year)
	for i := 1; i <= year; i++ {
		yearInterest := amount * yearlyPer / 100
		for j := 1; j <= 12; j++ {
			amount += float64(monthlyAmount)
			yearInterest += float64(monthlyAmount) * yearlyPer / 100 * float64(12-j+1) / 12
		}
		interest = interest + yearInterest
		fmt.Printf("第%d年,累计本金=%0.2f,本年利息=%0.2f,累计利息=%0.2f,合计=%0.2f\n", i, amount, yearInterest, interest, interest+amount)
	}
}
