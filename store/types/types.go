package types

type Store interface {
	Set(key string, value []byte) error

	Get(key string) ([]byte, error)

	Delete(key string) error

	Close() error
}
