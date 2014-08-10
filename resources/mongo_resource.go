package resources

import (
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strconv"
)

type MongoEngine struct {
	ConnectionAddress string
	DatabaseName      string
	CollectionName    string
	PerPage           int
}

// HandleGETData gets data from a given db & collection.
// If requestedId is not an empty string, single object is fetched. Object is looked up
// by looking for object that has "id"=requestedId field.
// If requestedId is an empty string, eng.PerPage results from collection are returned.
func (eng *MongoEngine) HandleGETData(requestedId string) (interface{}, error) {
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
		if err := collection.Find(bson.M{}).Sort("-id").Limit(eng.PerPage).All(&objectList); err != nil {
			return nil, errors.New("Failed to fetch data from collection")
		}
		return objectList, nil
	}
	return nil, nil
}

// HandlePOSTData accepts map of data, parsed from request and saves new document to collection.
// Data provided must have an `id` attribute.
func (eng *MongoEngine) HandlePOSTData(requestData map[string]interface{}) error {
	if _, ok := requestData["id"]; !ok {
		return errors.New("Request must contain an `id` field.")
	}
	session, err := mgo.Dial(eng.ConnectionAddress)
	if err != nil {
		return errors.New("Failed to connect to database.")
	}
	defer session.Close()
	collection := session.DB(eng.DatabaseName).C(eng.CollectionName)

	if err := collection.Insert(requestData); err != nil {
		return errors.New("Failed to insert new document to the collection.")
	}
	return nil
}

// MongoResource constructs an EngineResource using MongoEngine and
// attributes provided.
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
