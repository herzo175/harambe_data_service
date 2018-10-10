package schemas

import (
	"../../config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Report struct {
	Id                  bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Description         string        `json:"description" bson:"description,omitempty"`
	Date                int64         `json:"date" bson:"date,omitempty"`
	Value               float64       `json:"value" bson:"value,omitempty"`
	Cash                float64       `json:"cash" bson:"cash,omitempty"`
	BuyingPower         float64       `json:"buying_power" bson:"buying_power,omitempty"`
	Broker              bson.ObjectId `json:"broker" bson:"broker,omitempty"`
	BrokerAccountNumber string        `json:"broker_account_number" bson:"broker_account_number,omitempty"`
}

type ReportSchema struct {
	Session *mgo.Session
}

func (s *ReportSchema) GetCollection() *mgo.Collection {
	collection := s.Session.DB(config.DBName).C("Reports")
	return collection
}

// TODO: paginated list
func (s *ReportSchema) GetAll(query interface{}) ([]Report, error) {
	collection := s.GetCollection()

	result := []Report{}
	err := collection.Find(query).All(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *ReportSchema) GetById(id string) (Report, error) {
	collection := s.GetCollection()

	result := Report{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *ReportSchema) Create(newReport *Report) error {
	collection := s.GetCollection()
	newReport.Id = bson.NewObjectId()

	return collection.Insert(newReport)
}

func (s *ReportSchema) Update(id string, updateReport *Report) error {
	collection := s.GetCollection()

	query := bson.M{"_id": bson.ObjectIdHex(id)}
	change := bson.M{"$set": updateReport}

	return collection.Update(query, change)
}
