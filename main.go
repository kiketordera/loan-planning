package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/", func(c *gin.Context) {
		var json Request
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		loan, err := strconv.ParseFloat(json.LoanAmount, 64)
		nominalRate, err := strconv.ParseFloat(json.NominalRate, 64)
		duration, err := json.Duration.Float64()
		if err != nil {
			fmt.Print("It has beeen an error")
		}
		plan := createPlan(loan, nominalRate, duration, json.StartDate)

		c.JSON(http.StatusOK, gin.H{"borrowerPayments": plan})

	})
	router.Run(":8080") // listen an
}

func createPlan(loanAmount, nominalRate, duration float64, startDate string) []Mensuality {

	t, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		fmt.Println(err)
	}

	// We create the variables outside the loop since is a costly operation, so we can reuse them
	mensualities := []Mensuality{}
	m := Mensuality{}
	initialOutstandingPrincipal := loanAmount
	interest := 0.0
	principal := 0.0
	monthlyRate := nominalRate / 12
	fromPorcentualToDecimal := monthlyRate / 100
	annuity := (loanAmount * fromPorcentualToDecimal) / (1 - math.Pow((1+fromPorcentualToDecimal), -duration))
	annuity = round(annuity)
	m.BorrowerPaymentAmount = annuity

	for i := 0.0; i < duration; i++ {
		interest = ((nominalRate * 30 * initialOutstandingPrincipal) / 360) / 100
		principal = annuity - interest
		m.Date = t
		m.Principal = round(principal)
		m.Interest = round(interest)
		m.InitialOutstandingPrincipal = round(initialOutstandingPrincipal)
		m.RemainingOutstandingPrincipal = round(initialOutstandingPrincipal - principal)
		// To ajust the last payment
		if m.RemainingOutstandingPrincipal < 0 {
			m.BorrowerPaymentAmount = round(m.BorrowerPaymentAmount + m.RemainingOutstandingPrincipal)
			m.Principal = round(m.Principal + m.RemainingOutstandingPrincipal)
			m.RemainingOutstandingPrincipal = 0
		}
		mensualities = append(mensualities, m)
		initialOutstandingPrincipal -= principal
		// We add 1 month for the next payment
		t = t.AddDate(0, 1, 0)
	}

	JSON, err := json.Marshal(mensualities)
	if err != nil {
		fmt.Print("Fail to convert to JSON")
		return nil
	}
	fmt.Print("These are the mensualities: ", string(JSON))
	return mensualities
}

// Mensuality is every payment the customer need to pay back
type Mensuality struct {
	BorrowerPaymentAmount         float64   `json:"borrowerPaymentAmount"`
	Date                          time.Time `json:"date"`
	InitialOutstandingPrincipal   float64   `json:"initialOutstandingPrincipal"`
	Interest                      float64   `json:"interest"`
	Principal                     float64   `json:"principal"`
	RemainingOutstandingPrincipal float64   `json:"remainingOutstandingPrincipal"`
}

// Request is the query made to the server in the POST
type Request struct {
	LoanAmount  string      `json:"loanAmount"`
	NominalRate string      `json:"nominalRate"`
	Duration    json.Number `json:"duration"`
	StartDate   string      `json:"startDate"`
}

// Rounds the float number in 2 decimals
func round(n float64) float64 {
	return math.Round(n*100) / 100
}
