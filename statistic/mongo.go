package main

import (
	mgo "v1/mgo.v2"
	"v1/mgo.v2/bson"
)

type Mongo struct {
	Session *mgo.Session
}
type MgoStu struct {
	SelectBson bson.M
	UpdateBson bson.M
	DB         string
	C          string
}

func (m *Mongo) Connection(url string) error {
	var err error
	if url == "" {
		url = "127.0.0.1:27017"
	}
	m.Session, err = mgo.Dial(url)
	if err != nil {
		Error(err)
		return err
	}
	return nil
}

//mgo query
func (m *Mongo) Query(queryStu MgoStu) ([]interface{}, error) {
	var result []interface{}
	session := m.Session.Copy()
	c := session.DB(queryStu.DB).C(queryStu.C)
	err := c.Find(queryStu.SelectBson).All(&result)
	if err != nil {
		Error(err)
		return nil, err
	}
	return result, nil

}

//mgo delete
func (m *Mongo) Delete(deleteStu MgoStu) error {
	session := m.Session.Copy()
	c := session.DB(deleteStu.DB).C(deleteStu.C)
	query := c.Find(deleteStu.SelectBson)
	if query == nil {
		Debug("delete condition not satisfy")
		return nil
	}
	err := c.Remove(deleteStu.SelectBson)
	if err != nil {
		Error(err)
		return err
	}
	return nil
}

func (m *Mongo) Update(updateStu MgoStu) error {
	session := m.Session.Copy()
	c := session.DB(updateStu.DB).C(updateStu.C)

	err := c.Update(updateStu.SelectBson, updateStu.UpdateBson)
	if err != nil {
		Error(err)
		return err
	}
	return nil
}

func (m *Mongo) Insert(insertStu MgoStu) error {
	session := m.Session.Copy()
	c := session.DB(insertStu.DB).C(insertStu.C)
	err := c.Insert(insertStu.SelectBson)
	if err != nil {
		Error(err)
		return err
	}
	return nil
}

//todo  update和 insert 的区别和联系呢  如果没有update的 那么要insert吗
//todo mgo find 可以使用作为nil吗
//todo 能不能约定俗成的那种?update 的两个bson在一个bson中
