package schemas

import (
	"../../config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Security struct {
	Symbol               string `json:"symbol" bson:"symbol"`
	Quantity             int    `json:"quantity" bson:"quantity"`
	AllocationPercentage int    `json:"allocation_percentage" bson:"allocation_percentage"`
}

type Algorithm struct {
	Id                      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	NumPartitions           int           `json:"num_partitions" bson:"num_partitions,omitempty"`
	MaxAllocationPercentage int           `json:"max_allocation_percentage" bson:"max_allocation_percentage,omitempty"`
	DataPeriod              int           `json:"data_period" bson:"data_period,omitempty"`
	WaitingPeriod           int           `json:"waiting_period" bson:"waiting_period,omitempty"`
	TrendStrength           int           `json:"trend_strength" bson:"trend_strength,omitempty"`
	RateOfChange            int           `json:"rate_of_change" bson:"rate_of_change,omitempty"`
	Volatility              int           `json:"volatility" bson:"volatility,omitempty"`
	SecurityIds             []string      `json:"security_ids" bson:"security_ids,omitempty"`
	UserId                  string        `json:"user_id" bson:"user_id,omitempty"`
}

type AlgorithmSchema struct {
	Session *mgo.Session
}

func (s *AlgorithmSchema) GetCollection() *mgo.Collection {
	return s.Session.DB(config.DBName).C("Algorithms")
}

// TODO: paginated list
// TODO: query strings
func (s *AlgorithmSchema) GetAll(query interface{}) ([]Algorithm, error) {
	collection := s.GetCollection()

	result := []Algorithm{}
	err := collection.Find(query).All(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *AlgorithmSchema) GetById(id string) (Algorithm, error) {
	collection := s.GetCollection()

	result := Algorithm{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *AlgorithmSchema) Create(newAlgorithm *Algorithm) error {
	collection := s.GetCollection()

	return collection.Insert(newAlgorithm)
}

func (s *AlgorithmSchema) Update(id string, updateAlgorithm *Algorithm) error {
	collection := s.GetCollection()

	colQuerier := bson.M{"_id": bson.ObjectIdHex(id)}
	change := bson.M{"$set": updateAlgorithm}

	return collection.Update(colQuerier, change)
}

// TODO: delete?
