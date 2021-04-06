package stats

import (
	"github.com/fm2901/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var allSum types.Money
	var allCount types.Money

	if len(payments) < 1 {
		return 0
	}

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		allSum += payment.Amount
		allCount += 1
	}

	return allSum / allCount
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sumInCategory types.Money

	if len(payments) < 1 {
		return 0
	}

	for _, payment := range payments {
		if payment.Category != category || payment.Status == types.StatusFail {
			continue
		}
		sumInCategory += payment.Amount
	}

	return sumInCategory
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	counter := map[types.Category]int{}
	
	for _, payment := range payments {
		if payment.Amount > 0 {
			categories[payment.Category] += payment.Amount
			counter[payment.Category] += 1
		}
	}

	for cat, value := range categories {
		categories[cat] = int(categories[cat]) / counter[cat]
	}

	return categories
}