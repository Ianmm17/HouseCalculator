package service_test

import (
	"HouseCalculator/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalDebt(t *testing.T) {
	testTotalDebt := service.TotalDebt(100.00, 100.00, 100.00, 100.00, 100.01)
	assert.Equal(t, 500.01, testTotalDebt)
}
