package dstvmoz

import (
	"context"
	"errors"

	"github.com/getopends/libprovider"
)

func New(_ context.Context) (libprovider.Provider, error) {
	newProvider := &dstvProvider{}

	return &libprovider.DefaultProvider{
		TransactionCreator: libprovider.TransactionCreatorFunc(newProvider.CreateTransaction),
	}, nil
}

type dstvProvider struct{}

func (d dstvProvider) CreateTransaction(ctx context.Context, t *libprovider.Transaction) (*libprovider.CreateTransactionResult, error) {
	return nil, errors.ErrUnsupported
}
