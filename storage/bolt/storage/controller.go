package storage

import (
	"github.com/boltdb/bolt"
)

type Controller interface {
	Create(Model) error
	Delete(Model) error
	Get(Model) ([]byte, error)
	Exists(Model) (bool, error)
	Stats(Model) (bolt.BucketStats, error)
}
