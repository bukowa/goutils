package pkg

import (
	"encoding/json"
	"github.com/boltdb/bolt"
)

var ErrorEmptyKey = NewError("value of %s.Key() for model %s is empty")

type ControllerInterface interface {
	Create(Model) error
	Delete(Model) error
	Get(Model) ([]byte, error)
	Exists(Model) (bool, error)
	Stats(Model) (bolt.BucketStats, error)
}

type Controller struct {
	*Database
}

func (c *Controller) Get(m Model) (b []byte, err error) {
	if err = checkKey(m); err != nil {
		return
	}
	err = c.View(func(tx *bolt.Tx) error {
		bucket, err := c.BucketFor(m, tx)
		if err != nil {
			return err
		}
		b = bucket.Get(m.Key())
		return nil
	})
	return
}

func (c *Controller) Create(m Model) (err error) {
	if err = checkKey(m); err != nil {
		return
	}
	var b []byte
	b, err = json.Marshal(m)
	if err != nil {
		return
	}
	err = c.Update(func(tx *bolt.Tx) error {
		bucket, err := c.BucketFor(m, tx)
		if err != nil {
			return err
		}
		return bucket.Put(m.Key(), b)
	})
	return
}

func (c *Controller) Delete(m Model) (err error) {
	if err = checkKey(m); err != nil {
		return
	}
	err = c.Update(func(tx *bolt.Tx) error {
		bucket, err := c.BucketFor(m, tx)
		if err != nil {
			return err
		}
		return bucket.Put(m.Key(), nil)
	})
	return
}

func (c *Controller) Exists(m Model) (t bool, err error) {
	if err = checkKey(m); err != nil {
		return
	}
	err = c.View(func(tx *bolt.Tx) (err error) {
		bucket, err := c.BucketFor(m, tx)
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

func (c *Controller) Stats(m Model) (bs bolt.BucketStats, err error) {
	err = c.View(func(tx *bolt.Tx) error {
		bucket, err := c.BucketFor(m, tx)
		if err != nil {
			return err
		}
		bs = bucket.Stats()
		return nil
	})
	return
}

func checkKey(m Model) error {
	if len(m.Key()) < 1 {
		return ErrorEmptyKey.ForModel(m)
	}
	return nil
}
