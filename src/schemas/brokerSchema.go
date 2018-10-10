package schemas

import (
	"../../config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Broker struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name,omitempty"`
	Default   bool          `json:"default" bson:"default,omitempty"`
	UserId    string        `json:"user_id" bson:"user_id,omitempty"`
	UserToken string        `json:"user_token" bson:"user_token,omitempty"`
	User      bson.ObjectId `json:"user" bson:"user,omitempty"`
}

type BrokerSchema struct {
	Session *mgo.Session
}

func (s *BrokerSchema) GetCollection() *mgo.Collection {
	collection := s.Session.DB(config.DBName).C("Brokers")

	index := mgo.Index{
		Key:      []string{"name", "user"},
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
func (s *BrokerSchema) GetAll(query interface{}) ([]Broker, error) {
	collection := s.GetCollection()

	result := []Broker{}
	err := collection.Find(query).All(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *BrokerSchema) GetById(id string) (Broker, error) {
	collection := s.GetCollection()

	result := Broker{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *BrokerSchema) Create(newBroker *Broker) error {
	collection := s.GetCollection()
	newBroker.Id = bson.NewObjectId()

	return collection.Insert(newBroker)
}

func (s *BrokerSchema) Update(id string, updateBroker *Broker) error {
	collection := s.GetCollection()

	query := bson.M{"_id": bson.ObjectIdHex(id)}
	change := bson.M{"$set": updateBroker}

	return collection.Update(query, change)
}

func (s *BrokerSchema) Delete(id string) error {
	collection := s.GetCollection()

	query := bson.M{"_id": bson.ObjectIdHex(id)}

	return collection.Remove(query)
}
