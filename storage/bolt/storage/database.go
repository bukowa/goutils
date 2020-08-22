package storage

import (
	"encoding/binary"
	"encoding/json"
	"github.com/boltdb/bolt"
	"reflect"
)

type DB struct {
	*bolt.DB
}

func (db *DB) Get(m Model) (b []byte, err error) {
	if err = checkKey(m); err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		bucket, err := db.BucketFor(m, tx)
		if err != nil {
			return err
		}
		b = bucket.Get(m.Key())
		return nil
	})
	return
}

func (db *DB) Create(m Model) (err error) {
	if err = checkKey(m); err != nil {
		return
	}
	var b []byte
	b, err = json.Marshal(m)
	if err != nil {
		return
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := db.BucketFor(m, tx)
		if err != nil {
			return err
		}
		return bucket.Put(m.Key(), b)
	})
	return
}

func (db *DB) Delete(m Model) (err error) {
	if err = checkKey(m); err != nil {
		return
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := db.BucketFor(m, tx)
		if err != nil {
			return err
		}
		return bucket.Put(m.Key(), nil)
	})
	return
}

func (db *DB) Exists(m Model) (t bool, err error) {
	if err = checkKey(m); err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) (err error) {
		bucket, err := db.BucketFor(m, tx)
		if err != nil {
			return err
		}
		b := bucket.Get(m.Key())
		if len(b) > 0 {
			t = true
		}
		return
	})
	return
}

func (db *DB) Stats(m Model) (bs bolt.BucketStats, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		bucket, err := db.BucketFor(m, tx)
		if err != nil {
			return err
		}
		bs = bucket.Stats()
		return nil
	})
	return
}

func (db *DB) Init(opts *bolt.Options, path string, types ...Model) (err error) {
	// open database
	db.DB, err = bolt.Open(path, 0600, opts)
	if err != nil {
		return
	}
	// create buckets
	err = db.Update(func(tx *bolt.Tx) error {
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

func (db *DB) NextID(bucket *bolt.Bucket) (b []byte, err error) {
	seq, err := bucket.NextSequence()
	if err != nil {
		return
	}
	return BigEndian(seq), nil
}

func (db *DB) BucketFor(m Model, tx *bolt.Tx) (*bolt.Bucket, error) {
	name := []byte(getType(m))
	bucket := tx.Bucket(name)
	if bucket == nil {
		return nil, ErrorBucketDoesNotExists(name)
	}
	return bucket, nil
}

func BigEndian(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func getType(x interface{}) string {
	var t = reflect.TypeOf(x)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}

func checkKey(m Model) error {
	if len(m.Key()) < 1 {
		return ErrorEmptyKey(getType(m))
	}
	return nil
}
