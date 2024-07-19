package vo

import "fmt"

type LateFee struct {
	Point int
}

func (l *LateFee) AddPoint(point int) *LateFee {
	return &LateFee{
		Point: l.Point + point,
	}
}

func (l *LateFee) RemovePoint(point int) (*LateFee, error) {
	if l.Point < point {
		return nil, fmt.Errorf("rental point is not enough")
	}

	return &LateFee{ Point: l.Point - point }, nil
}


