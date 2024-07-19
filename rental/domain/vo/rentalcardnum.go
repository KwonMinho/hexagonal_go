package vo

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type RentalCardNum struct {
	No string
}

func CreateRentalCardNum() RentalCardNum {
	uuid := uuid.New().String()
	now := time.Now()
	no := fmt.Sprintf("%s-%s", uuid, now.Format("20060102150405"))

	return RentalCardNum{No: no}
}