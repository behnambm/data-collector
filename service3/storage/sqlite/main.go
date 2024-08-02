package sqlite

import (
	"context"
	"github.com/behnambm/data-collector/common/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteStorage struct {
	db *gorm.DB
}

func New(cfg *Config) (*SqliteStorage, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DBName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return &SqliteStorage{
		db: db,
	}, nil
}

func (ss *SqliteStorage) Store(ctx context.Context, entry *types.ServiceResultEntry) error {
	dbEntry := &ServiceResultModel{
		Status:      entry.Status,
		Svc1Latency: entry.Svc1Latency,
		Svc2Latency: entry.Svc2Latency,
	}
	err := ss.db.WithContext(ctx).Create(&dbEntry).Error
	if err != nil {
		return err
	}

	// populate the entry after being saved to DB
	entry.ID = dbEntry.ID
	entry.DateTime = dbEntry.CreatedAt

	return nil
}

func (ss *SqliteStorage) SetupModels() error {
	err := ss.db.AutoMigrate(&ServiceResultModel{})
	if err != nil {
		return err
	}

	return nil
}

func (ss *SqliteStorage) Close() {
	ss.Close()
}
