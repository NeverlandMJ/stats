package stats

import (
	"fmt"
	"github.com/NeverlandMJ/bank/pkg/bank/types"
)

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
