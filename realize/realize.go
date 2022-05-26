package realize

import (
	"Max/types"
)

func Max(payments []types.Payment) types.Payment {
	var max types.Payment

	for i := 0; i < len(payments)-1; i++ {
		if payments[i].Amount < payments[i+1].Amount {
			max = payments[i+1]
		}
	}

	return max

}
