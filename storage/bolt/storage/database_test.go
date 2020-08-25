package storage

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

var TestDatabase = func() (*DB, func()) {
	// create database
	os.Remove("test.db")
	db := &DB{}
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
	user := &TestUser{
		Login:    []byte("TestController_Delete"),
		Password: "",
	}
	// create user
	EP(db.Create(user))
	exists, err := db.Exists(user)
	EP(err)
	// check exists
	if !exists {
		t.Error()
	}
	// delete
	EP(db.Delete(user))

	// test
	b, err := db.Get(user)
	EP(err)
	if len(b) > 0 {
		t.Error()
	}
	// exists
	exists, err = db.Exists(user)
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
	user := &TestUser{
		Login:    []byte("TestController_Get"),
		Password: "",
	}
	//
	// empty
	b, err := db.Get(user)
	EP(err)
	if len(b) > 0 {
		t.Error()
	}
}

func TestController_Exists(t *testing.T) {
	// setup test
	db, def := TestDatabase()
	defer def()
	user := &TestUser{
		Login:    []byte("TestController_Exists"),
		Password: "",
	}
	//
	// create user & check
	EP(db.Create(user))
	exists, err := db.Exists(user)
	EP(err)
	if !exists {
		t.Error()
	}
	// check false
	user.Login = []byte("TestController_ExistsTestController_Exists")
	exists, err = db.Exists(user)
	EP(err)
	if exists {
		t.Error()
	}
}

func TestBoltDatabase_Init(t *testing.T) {
	// setup test
	db, def := TestDatabase()
	defer def()
	user := &TestUser{
		Login:    []byte("TestBoltDatabase_Init"),
		Password: "TestBoltDatabase_Init",
	}
	//

	// insert user
	EP(db.Create(user))

	// get user bytes
	b, err := db.Get(user)
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
	EP(db.Create(user))

	// get
	b, err = db.Get(user)
	EP(json.Unmarshal(b, &user2))

	// check
	if !reflect.DeepEqual(user, user2) {
		t.Error(user, user2)
	}
	// check count
	stats, err := db.Stats(user)
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

func TestDB_GetAll(t *testing.T) {
	// setup test
	db, def := TestDatabase()
	defer def()

	user := &TestUser{
		Login:    []byte("login"),
		Password: "passwd",
	}
	// insert user
	EP(db.Create(user))
	user2 := &TestUser{
		Login:    []byte("test"),
		Password: "test",
	}
	EP(db.Create(user2))
	// get all users
	users, err := db.GetAll(user)
	EP(err)
	if len(users) != 2 {
		t.Error(users)
	}
}

func TestDB_Locked(t *testing.T) {
	// setup test
	db, def := TestDatabase()
	defer def()

	user := &TestUser{
		Login:    []byte("login"),
		Password: "passwd",
	}
	if err := db.CreateLocked(user); err != nil {
		t.Error(err)
	}
	if v, err := db.ExistsLocked(user); err != nil {
		t.Error(err)
	} else if !v{
		t.Error(v)
	}
}