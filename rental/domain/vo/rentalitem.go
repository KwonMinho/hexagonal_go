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