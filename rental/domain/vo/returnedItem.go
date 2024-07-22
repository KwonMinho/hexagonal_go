package vo

import (
	"time"
)

type ReturnItem struct {
	RentalItem RentalItem
	ReturnDate time.Time
}

func CreateReturnItem(item RentalItem) ReturnItem {
	return ReturnItem {
		RentalItem: item,
		ReturnDate: time.Now(),
	}
}