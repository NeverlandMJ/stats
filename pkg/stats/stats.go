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


func FilterByCategory2(payments []types.Payment, category types.Category) types.Money{
	occurance := types.Money(0) 
	
	for _, payment := range payments{
		if payment.Category == category{
			occurance++
		}
	}
	return occurance
}

func TotalInCategory2(payments []types.Payment, category types.Category) types.Money {
	total := types.Money(0)
	for _, amounts := range payments {
		
		if amounts.Category == category {
			total = total + amounts.Amount
		}
	}
	
	return total
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money{

	categories := map[types.Category]types.Money{}

	for _, payment := range payments{
		
		categories[payment.Category] = TotalInCategory(payments, payment.Category)/FilterByCategory2(payments, payment.Category)
	}

		
	return categories
}



func PeriodsDynamic(first map[types.Category]types.Money, 
	second map[types.Category]types.Money, ) map[types.Category]types.Money{ 

		result := map[types.Category]types.Money{}
	if len(first) >= len(second){
		for category1, amount1 := range first{

			for category2, amount2 := range second{

		     
				if category1 == category2{
					result[category1] = amount2 - amount1
					break
				}else if category1 != category2{
					result[category1] = -amount1
				}							
	
	}
}
	}else if len(first) < len(second){
			for category2, amount2 := range second{

				for category1, amount1 := range first{

		     
				if category1 == category2{
					result[category1] = amount2 - amount1
					break
				}else if category1 != category2{
					result[category2] = amount2
				}							
	
			}
		}
	}
		return result
}

