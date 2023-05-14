package tax_test

import (
	"testing"

	tax "github.com/IcaroSilvaFK/testify_go_lang/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {

	tx, err := tax.CalculateTax(1_000)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, tx)

	tx, err = tax.CalculateTax(0)

	assert.Error(t, err, "invalid amount")
	assert.Equal(t, 0.0, tx)
}
