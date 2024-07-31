package mock

import (
	"context"
	"github.com/behnambm/data-collector/common/types"
	"github.com/stretchr/testify/mock"
)

type DummyDB struct {
	mock.Mock
}

func (d DummyDB) Store(ctx context.Context, entry *types.ServiceResultEntry) error {
	args := d.Called(ctx, entry)
	return args.Error(0)
}
