package vo

import (
	"time"
)

type ReturnedItem struct {
	RentalItem RentalItem
	ReturnDate time.Time
}