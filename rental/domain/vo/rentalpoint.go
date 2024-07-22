package vo

import "fmt"

type LateFee struct {
	Point int
}

func CreateLateFee() LateFee {
	return LateFee{
		Point: 0,
	}
}

func (l *LateFee) AddPoint(point int) *LateFee {
	return &LateFee{
		Point: l.Point + point,
	}
}

func (l *LateFee) RemovePoint(point int) (*LateFee, error) {
	if l.Point > point {
		return nil, fmt.Errorf("연체를 삭제할 포인트가 충분하지 않습니다")
	}

	return &LateFee{ Point: l.Point - point }, nil
}


