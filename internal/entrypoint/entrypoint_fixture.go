package entrypoint

import "context"

type entrypointForTests struct{}

func NewFixture() Entrypoint {
	return &entrypointForTests{}
}

func (e entrypointForTests) GetDB() Database {
	return nil
}

func (e entrypointForTests) TxWrapper(ctx context.Context, fn func(ctx context.Context, tx Database) error) (err error) {
	return fn(ctx, nil)
}
