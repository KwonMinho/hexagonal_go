package vo

import (
	"time"
)

type RentalItem struct {
	Item Item
	RentalDate time.Time
	ReturnDate time.Time
	Overdued bool
}

func CreateRentalItem(item Item) RentalItem {
	rentalDays:= 14

	return RentalItem{
		Item: item,
		RentalDate: time.Now(),
		ReturnDate: time.Time{}.AddDate(0, 0, rentalDays),
		Overdued: false,
	}
}

func (r *RentalItem) GetReturnDate() time.Time {
	return r.ReturnDate
}