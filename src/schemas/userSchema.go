package schemas

import (
	"../../config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Topic struct {
	Announcements *bool `json:"announcements" bson:"announcements,omitempty"`
	Account       *bool `json:"account" bson:"account,omitempty"`
}

type Method struct {
	Email *bool `json:"email" bson:"email,omitempty"`
	Sms   *bool `json:"sms" bson:"sms,omitempty"`
}

type Notification struct {
	Topics  Topic  `json:"topics" bson:"topics,omitempty,inline"`
	Methods Method `json:"methods" bson:"methods,omitempty,inline"`
}

type User struct {
	Id            bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	EmailAddress  string        `json:"email_address" bson:"email_address,omitempty"`
	Password      string        `json:"password" bson:"password,omitempty"`
	FirstName     string        `json:"first_name" bson:"first_name,omitempty"`
	LastName      string        `json:"last_name" bson:"last_name,omitempty"`
	Notifications Notification  `json:"notifications" bson:"notifications,omitempty,inline"`
}

type UserSchema struct {
	Session *mgo.Session
}

func (s *UserSchema) GetCollection() *mgo.Collection {
	collection := s.Session.DB(config.DBName).C("Users")

	index := mgo.Index{
		Key:      []string{"email_address"},
		Unique:   true,
		DropDups: true,
	}

	err := collection.EnsureIndex(index)

	if err != nil {
		panic(err)
	}

	return collection
}

// TODO: paginated list
func (s *UserSchema) GetAll(query interface{}) ([]User, error) {
	collection := s.GetCollection()

	result := []User{}
	err := collection.Find(query).All(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *UserSchema) GetById(id string) (User, error) {
	collection := s.GetCollection()

	result := User{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *UserSchema) Create(newUser *User) error {
	collection := s.GetCollection()

	return collection.Insert(newUser)
}

func (s *UserSchema) Update(id string, updateUser *User) error {
	collection := s.GetCollection()

	colQuerier := bson.M{"_id": bson.ObjectIdHex(id)}
	change := bson.M{"$set": updateUser}

	return collection.Update(colQuerier, change)
}

// TODO: delete?
