package stats

import (
	"github.com/NeverlandMJ/bank/v2/pkg/types"
	"fmt"
	
)

func ExampleAvg(){
	payment := []types.Payment{
		{
			Amount: 10_000_00,
		},
		{
			Amount: 20_000_00,
		},
		{
			Amount: 30_000_00,
		},

	}
	averageAmount := Avg(payment)
	fmt.Println(averageAmount)

	// Output: 2000000

}

func ExampleTotalInCategory() {
	cards := []types.Payment{
		{
			Amount:   1_000_000,
			Category: "TJS",
		},
		{
			Amount:   2_000_000,
			Category: "TJS",
		},
		{
			Amount:   3_000_000,
			Category: "USD",
		},
	}
	total := TotalInCategory(cards, "TJS")
	fmt.Println(total)

	//Output: 3000000
}


