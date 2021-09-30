package stats

import (
	"fmt"
	"log"
	"testing"

	"github.com/NeverlandMJ/bank/pkg/types"
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


func Test(t *testing.T) {
	category := types.Category{}
	log.Println("category => ", category)
	log.Println("category => ", types.Category)
	log.Println("category => ", types.Category)
}