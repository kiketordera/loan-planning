package main

import (
	"math/rand"
	"testing"
)

func TestProfit(t *testing.T) {
	m := generatorPayments()
	loan := 0.0
	paid := 0.0
	for _, plan := range m {
		for i, montlyPayment := range plan {
			if i == 0 {
				loan = montlyPayment.InitialOutstandingPrincipal
			}
			paid += montlyPayment.BorrowerPaymentAmount
		}
		if loan > paid {
			t.Errorf("ERROR: The Bank is loosing money with this credit")
		}
	}
}

func TestNegatives(t *testing.T) {
	m := generatorPayments()
	for _, plan := range m {
		for _, payment := range plan {
			if payment.RemainingOutstandingPrincipal < 0 {
				t.Errorf("ERROR: RemainingOutstandingPrincipal is negative")
			}
			if payment.Principal < 0 {
				t.Errorf("ERROR: Principal is negative")
			}
			if payment.InitialOutstandingPrincipal < 0 {
				t.Errorf("ERROR: InitialOutstandingPrincipal is negative")
			}
			if payment.BorrowerPaymentAmount < 0 {
				t.Errorf("ERROR: BorrowerPaymentAmount is negative")
			}
		}
	}
}

func generatorPayments() [][]Mensuality {
	payments := [][]Mensuality{}
	for i := 0; i < 99999; i++ {
		payments = append(payments, createPlan(float64(random(50, 99999)), float64(random(0, 7))+rand.Float64(), random(2, 36), "2018-01-01T00:00:01Z"))
	}
	return payments
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
