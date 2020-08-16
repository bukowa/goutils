package pkg

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"reflect"
)

type BoltDatabase struct {
	*bolt.DB
}

func (d *BoltDatabase) Init(opts *bolt.Options, path string, types ...Model) (err error) {
	// open database
	d.DB, err = bolt.Open(path, 0600, opts)
	if err != nil {
		return
	}
	// create buckets
	err = d.Update(func(tx *bolt.Tx) error {
		for _, each := range types {
			name := getType(each)
			if _, err := tx.CreateBucketIfNotExists([]byte(name)); err != nil {
				return err
			}
		}
		return nil
	})
	return
}

func (d *BoltDatabase) NextId(bucket *bolt.Bucket) (b []byte, err error) {
	seq, err := bucket.NextSequence()
	if err != nil {
		return
	}
	return BigEndian(seq), nil
}

func (d *BoltDatabase) BucketFor(m Model, tx *bolt.Tx) *bolt.Bucket {
	name := []byte(getType(m))
	bucket := tx.Bucket(name)
	if bucket == nil {
		panic(errors.New(fmt.Sprintf("bucket %v does not exist", name)))
	}
	return bucket
}

func getType(t interface{}) string {
	if t := reflect.TypeOf(t); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func BigEndian(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
