package tax

import "errors"

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("invalid amount")
	}
	if amount >= 1_000 && amount < 20_000 {
		return 10.0, nil
	}
	if amount <= 20_000 {
		return 20.0, nil
	}
	return 5.0, nil
}
