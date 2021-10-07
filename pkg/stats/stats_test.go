package stats

import (
	"reflect"
	"testing"

	"github.com/NeverlandMJ/bank/v2/pkg/types"
	
)

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment

	result := FilterByCategory(payments, "mobile")

	if len(result) != 0{
		t.Error("resul len != 0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}

	result := FilterByCategory(payments, "mobile")

	if len(result) !=0 {
		t.Error("result len != 0")
	}

}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}

	result := FilterByCategory(payments, "mobile")

	if len(result) != 0{
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 2, Category: "food"},
	}

	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestFilterByCategory_moreThanOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}

	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}

	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}



func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "fun", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 6_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
		{ID: 6, Category: "food", Amount: 6_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 5_000_00,
		"food": 4_000_00,
		"fun": 3_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic1(t *testing.T) {
	firstPeriod := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}

	secondPeriod := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
		
	}

	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
		
	}

	result := PeriodsDynamic(firstPeriod, secondPeriod)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}

}

func TestPeriodsDynamic2(t *testing.T) {
	firstPeriod := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}

	secondPeriod := map[types.Category]types.Money{
		"auto": 20,
		"food": 20,
	}

	expected := map[types.Category]types.Money{
		"auto": 10,
		"food": 0,
	}

	result := PeriodsDynamic(firstPeriod, secondPeriod)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}

}

func TestPeriodsDynamic3(t *testing.T) {
	firstPeriod := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}

	secondPeriod := map[types.Category]types.Money{
		"food": 20,
		
	}

	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
		
	}

	result := PeriodsDynamic(firstPeriod, secondPeriod)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}

}
func TestPeriodsDynamic4(t *testing.T) {
	firstPeriod := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}

	secondPeriod := map[types.Category]types.Money{
		"auto": 10,
		"food": 25,
		"mobile": 5,
	}

	expected := map[types.Category]types.Money{
		"auto": 0,
		"food": 5,
		"mobile": 5,
	}

	result := PeriodsDynamic(firstPeriod, secondPeriod)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}

}