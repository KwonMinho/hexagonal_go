package domain

import "hexagonal_go/rental/domain/vo"

type RentalCard struct {
	no 	vo.RentalCardNum
	member vo.IDName
	status vo.RentalStatus
	lateFee vo.LateFee
	items []vo.RentalItem
	returnedItems []vo.ReturnedItem
}

func NewRentalCard(no vo.RentalCardNum, member vo.IDName, status vo.RentalStatus, lateFee vo.LateFee, items []vo.RentalItem, returnedItems []vo.ReturnedItem) *RentalCard {
	return &RentalCard{
		no:  no,
		member: member,
		status: status,
		lateFee: lateFee,
		items: items,
		returnedItems: returnedItems,
	}
}