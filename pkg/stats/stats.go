package stats

import "github.com/NeverlandMJ/bank/v2/pkg/types"


func Avg(payments []types.Payment) types.Money {

	total := types.Money(0)
	soni := 0
	for _, amounts := range payments {
		if amounts.Status != types.StatusFail{
		total = total + amounts.Amount
		soni++
		}
	}
	average := types.Money(total) / types.Money(soni)

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

func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment{
	var filtered []types.Payment
	for _, payment := range payments{
		if payment.Category == category{
			filtered = append(filtered, payment)
		}
	}
	return filtered
}
func CategoriesTotal(payments []types.Payment)map[types.Category]types.Money{
	categories := map[types.Category]types.Money{}

	for _, payment := range payments{
		categories[payment.Category] += payment.Amount
	}
	return categories
}