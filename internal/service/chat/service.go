package chat

import (
	"fmt"

	"github.com/SeminarioGo/seminarioGO/internal/config"
	"github.com/jmoiron/sqlx"
)

//Message ...
type Message struct {
	ID   int64
	Text string
}

//Service ...
type Service interface {
	AddMessage(Message) (Message, error)
	FindByID(int) *Message
	FindAll() []*Message
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddMessage(m Message) (Message, error) {
	sqlStatement := "INSERT INTO messages (text) VALUES (?)"

	res, err := s.db.Exec(sqlStatement, m.Text)
	if err != nil {
		return m, err
	}

	m.ID, _ = res.LastInsertId()
	fmt.Println(res.LastInsertId())

	return m, err
}

func (s service) FindByID(ID int) *Message {
	return nil
}

func (s service) FindAll() []*Message {
	var list []*Message
	if err := s.db.Select(&list, "SELECT * FROM messages"); err != nil {
		panic(err)
	}
	return list
}
