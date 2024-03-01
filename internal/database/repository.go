package db

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UserRepository struct {
	Session neo4j.Session
}

func NewUserRepository(session neo4j.Session) *UserRepository {
	return &UserRepository{Session: session}
}

func (r *UserRepository) CreateUser(user *User) error {
	_, err := r.Session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			"CREATE (u:User {age: $age, name: $name, password: $password})",
			map[string]interface{}{"age": user.Age, "name": user.Name, "password": user.Password},
		)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}
