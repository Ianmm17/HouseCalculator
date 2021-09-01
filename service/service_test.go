package service_test

import (
	"HouseCalculator/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalDebt(t *testing.T) {
	testTotalDebt := service.TotalDebt(100.00, 100.00, 100.00, 100.00, 100.00, 100.00, 100.00, 100.00, 100.00, 100.00)
	assert.Equal(t, 1000.00, testTotalDebt)
}

func TestCalculatesDTIPercent(t *testing.T) {
	expectedDTI := "64"
	testCalcDebt := service.CalculatesDTI("4500", "85000")
	assert.Equal(t, expectedDTI, testCalcDebt)
}
