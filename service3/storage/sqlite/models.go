package sqlite

import (
	"github.com/behnambm/data-collector/common/types"
	"gorm.io/gorm"
)

type ServiceResultModel struct {
	gorm.Model
	ID          uint64             `gorm:"primaryKey;autoIncrement"`
	Status      types.ResultStatus `gorm:"column:status"`
	Svc1Latency int64              `gorm:"column:svc1_latency"`
	Svc2Latency int64              `gorm:"column:svc2_latency"`
}

func (ServiceResultModel) TableName() string {
	return "service_results"
}
