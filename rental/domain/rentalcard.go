package domain

import (
	"errors"
	"time"

	"hexagonal_go/rental/domain/consts"
	"hexagonal_go/rental/domain/vo"
)

type RentalCard struct {
	no 	vo.RentalCardNum
	member vo.IDName
	status vo.RentalStatus
	lateFee vo.LateFee
	items []vo.RentalItem
	returnedItems []vo.ReturnItem
}

func NewRentalCard(owner vo.IDName) *RentalCard {
	return &RentalCard{
		no:  vo.CreateRentalCardNum(),
		member: owner,
		status: vo.RentalAvailable,
		lateFee: vo.CreateLateFee(),
		items: []vo.RentalItem{},
		returnedItems: []vo.ReturnItem{},
	}
}

func (r *RentalCard) addRentalItem(item vo.RentalItem) {
	r.items = append(r.items, item)
}

func (r *RentalCard) removeRentalItem(rentalItem vo.RentalItem) error {
	rmItem := rentalItem.Item

	i := r.searchIndex(rmItem)
	if i == consts.NotFoundRentalItem {
		return errors.New("삭제할 대여품목을 찾지 못했습니다.")
	}
	r.items = append(r.items[:i], r.items[i+1:]...)
	return nil
}

func (r *RentalCard) addReturnItem(item vo.ReturnItem) {
	r.returnedItems = append(r.returnedItems, item)
}

func (r *RentalCard) checkRentalAvailable() error {
	if r.status == vo.RentalNotAvailable {
		return errors.New("대여가 불가능한 상태입니다")
	}

	if len(r.items) > 5 {
		return errors.New("이미 5권을 대여했습니다")
	}

	return nil
}

func (r *RentalCard) searchIndex(item vo.Item) int {
	for i, ri := range r.items {
		if ri.Item.ID == item.ID {
			return i
		}
	}
	return consts.NotFoundRentalItem
}

func (r *RentalCard) RentItem(item vo.Item) error {
	err := r.checkRentalAvailable()
	if err != nil {
		return err
	}

	rentalItem := vo.CreateRentalItem(item)
	r.addRentalItem(rentalItem)
	return nil
}

func (r *RentalCard) ReturnRentalItem(item vo.Item, returnDate time.Time) error {
	i := r.searchIndex(item)
	if i == consts.NotFoundRentalItem {
		return errors.New("반환할 대여품목을 찾지 못햇습니다.")
	}

	rentalItem := r.items[i]
	returnItem := vo.CreateReturnItem(rentalItem)

	r.calculateLateFee(rentalItem, returnDate)
	r.addReturnItem(returnItem)
	r.removeRentalItem(rentalItem)

	return nil
}

func (r *RentalCard) calculateLateFee(item vo.RentalItem, returnDate time.Time) {
	itemReturnDate := item.GetReturnDate()
	if returnDate.After(itemReturnDate) {
		daysOverdue := int(returnDate.Sub(itemReturnDate).Hours() / 24)
		points := daysOverdue * consts.LateFeePointWeight
		addPoint := r.lateFee.AddPoint(points)
		r.lateFee = *addPoint
	}
}

func (r *RentalCard) overdueItem(item vo.Item) error {
	i := r.searchIndex(item)
	if i == consts.NotFoundRentalItem {
		return errors.New("대여 품목을 찾지 못했습니다")
	}

	rentalItem := r.items[i]
	rentalItem.Overdue = true
	r.status = vo.RentalNotAvailable
	return nil
}

func (r *RentalCard) makeAvailableRental(point int) (error) {
	if len(r.items) != 0 {
		return errors.New("모든 도서가 반납되어야 정지를 해제할 수 있습니다.")
	}

	updateLateFee, err := r.lateFee.RemovePoint(point)
	if err != nil {
		return err
	}

	r.lateFee = *updateLateFee
	if r.lateFee.Point == 0 {
		r.status = vo.RentalNotAvailable
	}

	return nil
}