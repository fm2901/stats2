package stats

import (
	"github.com/fm2901/bank/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	var payments []types.Payment
	payments = append(payments,
		types.Payment {
			ID: 1,
			Amount: 100,
			Category: "mobile",
		},
		types.Payment {
			ID: 1,
			Amount: 1000,
			Category: "komunnal",
		},
	)

	result := Avg(payments)
	fmt.Println(result)
	// Output: 550
}

func ExampleTotalInCategory() {
	var payments []types.Payment
	payments = append(payments,
		types.Payment {
			ID: 1,
			Amount: 100,
			Category: "mobile",
		},
		types.Payment {
			ID: 1,
			Amount: 1000,
			Category: "komunnal",
		},
	)

	result := TotalInCategory(payments, "mobile")
	fmt.Println(result)
	// Output: 100
}