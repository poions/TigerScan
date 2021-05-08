package common

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type DataStore struct {
	MgoUrl  string
	Session *mgo.Session
	DB      string
}

func (ds *DataStore) MgoInit() {
	session, err := mgo.Dial(ds.MgoUrl)

	if err != nil {
		fmt.Println(err)
	}
	ds.Session = session
}

func (ds *DataStore) MgoConnect(collection string) (*mgo.Session, *mgo.Collection) {
	if ds.Session != nil {
		ms := ds.Session.Copy()
		c := ms.DB(ds.DB).C(collection)
		ms.SetMode(mgo.Monotonic, true)
		return ms, c
	} else {
		return nil, nil
	}
}

//FindOne : 获取单条记录
func (ds *DataStore) FindOne(collection string, query, result interface{}) error {
	ms, c := ds.MgoConnect(collection)
	if c != nil {
		defer ms.Close()
		return c.Find(query).One(result)
	} else {
		return fmt.Errorf("connect mongodb error")
	}
}

//FindAll : 获取所有记录
func (ds *DataStore) FindAll(collection string, query, result interface{}) error {
	ms, c := ds.MgoConnect(collection)
	if c != nil {
		defer ms.Close()
		return c.Find(query).All(result)
	} else {
		return fmt.Errorf("connect mongodb error")
	}
}
