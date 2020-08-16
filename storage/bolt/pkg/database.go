package pkg

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
	"reflect"
)

var ErrorBucketDoesNotExist = NewError("bucket for %s does not exists")

type Database struct {
	*bolt.DB
}

func (d *Database) Init(opts *bolt.Options, path string, types ...Model) (err error) {
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

func (d *Database) NextID(bucket *bolt.Bucket) (b []byte, err error) {
	seq, err := bucket.NextSequence()
	if err != nil {
		return
	}
	return BigEndian(seq), nil
}

func (d *Database) BucketFor(m Model, tx *bolt.Tx) (*bolt.Bucket, error) {
	name := []byte(getType(m))
	bucket := tx.Bucket(name)
	if bucket == nil {
		return nil, ErrorBucketDoesNotExist.ForString(string(name))
	}
	return bucket, nil
}

func getType(x interface{}) string {
	var t = reflect.TypeOf(x)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}

func BigEndian(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
