package stats

import "github.com/NeverlandMJ/bank/v2/pkg/types"


func Avg(payments []types.Payment) types.Money {

	total := types.Money(0)
   
	for _, amounts := range payments {
		if amounts.Status != types.StatusFail{
		total = total + amounts.Amount
		}
	}
	average := types.Money(total) / types.Money(len(payments))

	return types.Money(average)

}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	total := types.Money(0)
	for _, amounts := range payments {
		
		if amounts.Status != types.StatusFail && amounts.Category == category {
			total = total + amounts.Amount
		}
	}
	

	return total
}
