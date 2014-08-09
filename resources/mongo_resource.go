package resources

import (
	"net/http"
)

type MongoEngine struct {}

func (eng *MongoEngine) HandleGETData(r *http.Request) (interface{}, error) {
	return nil, nil
}

func (eng *MongoEngine) HandlePOSTData(r *http.Request) error {
	return nil
}

func MongoResource(BaseUrl, ConnectionAddress, DatabaseName, TableName string) ResourceInterface {
	engine := new(MongoEngine)
	resource := EngineResource{
		BaseUrl,
		engine,
		ConnectionAddress,
		DatabaseName,
		TableName,
	}
	return &resource
}
