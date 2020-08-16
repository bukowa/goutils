package pkg

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

var TestDatabase = func() (*BoltDatabase, func()) {
	// create database
	os.Remove("test.db")
	db := &BoltDatabase{}
	EP(db.Init(nil, "test.db", TestUser{}))
	return db, func() {
		db.Close()
		os.Remove("test.db")
	}
}

type TestUser struct {
	Login    []byte
	Password string
}

func (t TestUser) Key() []byte {
	return t.Login
}

func TestController_Delete(t *testing.T) {
	// setup test
	db, def := TestDatabase()
	defer def()
	c := &Controller{db}
	user := &TestUser{
		Login:    []byte("TestController_Delete"),
		Password: "",
	}
	// create user
	EP(c.Create(user))
	exists, err := c.Exists(user)
	EP(err)
	// check exists
	if !exists {
		t.Error()
	}
	// delete
	EP(c.Delete(user))

	// test
	b, err := c.Get(user)
	EP(err)
	if len(b) > 0 {
		t.Error()
	}
	// exists
	exists, err = c.Exists(user)
	EP(err)
	// check exists
	if exists {
		t.Error()
	}
}

func TestController_Get(t *testing.T) {
	// setup test
	db, def := TestDatabase()
	defer def()
	c := &Controller{db}
	user := &TestUser{
		Login:    []byte("TestController_Get"),
		Password: "",
	}
	//
	// empty
	b, err := c.Get(user)
	EP(err)
	if len(b) > 0 {
		t.Error()
	}
}

func TestController_Exists(t *testing.T) {
	// setup test
	db, def := TestDatabase()
	defer def()
	c := &Controller{db}
	user := &TestUser{
		Login:    []byte("TestController_Exists"),
		Password: "",
	}
	//
	// create user & check
	EP(c.Create(user))
	exists, err := c.Exists(user)
	EP(err)
	if !exists {
		t.Error()
	}
	// check false
	user.Login = []byte("TestController_ExistsTestController_Exists")
	exists, err = c.Exists(user)
	EP(err)
	if exists {
		t.Error()
	}
}

func TestBoltDatabase_Init(t *testing.T) {
	// setup test
	db, def := TestDatabase()
	defer def()
	c := &Controller{db}
	user := &TestUser{
		Login:    []byte("TestBoltDatabase_Init"),
		Password: "TestBoltDatabase_Init",
	}
	//

	// insert user
	EP(c.Create(user))

	// get user bytes
	b, err := c.Get(user)
	if len(b) < 1 {
		EP(err)
	}
	// unmarshal
	var user2 *TestUser
	EP(json.Unmarshal(b, &user2))

	// check
	if !reflect.DeepEqual(user, user2) {
		t.Error(user, user2)
	}

	// insert again
	user.Login = []byte("TestBoltDatabase_InitTestBoltDatabase_Init")
	user.Password = user.Password + "TestBoltDatabase_Init"
	EP(c.Create(user))

	// get
	b, err = c.Get(user)
	EP(json.Unmarshal(b, &user2))

	// check
	if !reflect.DeepEqual(user, user2) {
		t.Error(user, user2)
	}
	// check count
	stats, err := c.Stats(user)
	EP(err)

	if stats.KeyN != 2 {
		t.Error(stats)
	}

}

func EP(err error) {
	if err != nil {
		panic(err)
	}
}
