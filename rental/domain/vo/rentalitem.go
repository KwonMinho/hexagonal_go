package vo

import (
	"time"

	"hexagonal_go/rental/domain/consts"
)

type RentalItem struct {
	Item Item
	RentalDate time.Time
	ReturnDate time.Time
	Overdue bool
}

func CreateRentalItem(item Item) RentalItem {
	rentalDays:= consts.MaxRentalPeriod

	return RentalItem{
		Item: item,
		RentalDate: time.Now(),
		ReturnDate: time.Time{}.AddDate(0, 0, rentalDays),
		Overdue: false,
	}
}

func (r *RentalItem) GetReturnDate() time.Time {
	return r.ReturnDate
}