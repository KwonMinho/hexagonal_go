package vo

import (
	"time"
)	
type RentalItem struct {
	item Item
	rentalDate time.Time
	returnDate time.Time
	overdued bool
}

func CreateRentalItem(item Item) RentalItem {
	rentalDays:= 14

	return RentalItem{
		item: item,
		rentalDate: time.Now(),
		returnDate: time.Time{}.AddDate(0, 0, rentalDays),
		overdued: false,
	}
}