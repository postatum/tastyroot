package resources

import (
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"strconv"
)

type MongoEngine struct {
	ConnectionAddress string
	DatabaseName      string
	CollectionName    string
	PerPage           int
}

func (eng *MongoEngine) HandleGETData(r *http.Request, requestedId string) (interface{}, error) {
	session, err := mgo.Dial(eng.ConnectionAddress)
	if err != nil {
		return nil, errors.New("Failed to connect to database.")
	}
	defer session.Close()
	collection := session.DB(eng.DatabaseName).C(eng.CollectionName)

	if len(requestedId) > 0 {
		singleObject := make(map[string]interface{})
		requestedId, err := strconv.Atoi(requestedId)
		if err != nil {
			return nil, errors.New("Failed to parse provided object ID.")
		}
		objId := bson.M{"id": requestedId}
		if err := collection.Find(objId).One(&singleObject); err != nil {
			return nil, errors.New("Object with provided ID not found.")
		}
		return singleObject, nil
	} else {
		objectList := make([]map[string]interface{}, eng.PerPage)
		if err := collection.Find(bson.M{}).Sort("-_id").Limit(eng.PerPage).All(&objectList); err != nil {
			return nil, errors.New("Failed to fetch data from collection")
		}
		return objectList, nil
	}
	return nil, nil
}

func (eng *MongoEngine) HandlePOSTData(r *http.Request) error {
	return nil
}

func MongoResource(BaseUrl, ConnectionAddress, DatabaseName, CollectionName string, PerPage int) ResourceInterface {
	engine := MongoEngine{
		ConnectionAddress,
		DatabaseName,
		CollectionName,
		PerPage,
	}
	resource := EngineResource{
		BaseUrl,
		&engine,
	}
	return &resource
}
