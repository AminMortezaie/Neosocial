package neo

import (
	// "fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var (
	Driver neo4j.Driver
)

func InitNeo4j() error {
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		return err
	}

	Driver = driver
	return nil
}

func CloseNeo4j() {
	if Driver != nil {
		_ = Driver.Close()
	}
}
