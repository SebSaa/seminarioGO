package usuario

import (
	"fmt"

	"github.com/SeminarioGo/seminarioGO/internal/config"
	"github.com/jmoiron/sqlx"
)

//Message ...
type Usuario struct {
	ID     int64
	Nombre string
	Dni    int64
}

//Service ...
type Service interface {
	AddUsuarios(Usuario) (Usuario, error)
	FindByID(int) (*Usuario, error)
	FindAll() []*Usuario
	RemoveByID(int) (bool, error)
	UpdateUsuarios(Usuario) (bool, error)
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// Update ...
func (s service) UpdateUsuarios(u Usuario) (bool, error) {

	sqlStatement := "UPDATE usuarios SET nombre = ?, dni = ? WHERE ID = ?"

	_, err := s.db.Exec(sqlStatement, u.Nombre, u.Dni, u.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

// FindById ...
func (s service) FindByID(ID int) (*Usuario, error) {
	var u Usuario
	sqlStatement := "SELECT * FROM usuarios WHERE ID=?"
	if err := s.db.Get(&u, sqlStatement, ID); err != nil {
		return nil, err
	}

	return &u, nil

}

// RemoveById
func (s service) RemoveByID(ID int) (bool, error) {
	sqlStatement := "DELETE FROM usuarios WHERE ID = ?"
	_, err := s.db.Exec(sqlStatement, ID)
	if err != nil {
		return false, err
	}

	return true, nil

}

// AddUsuario
func (s service) AddUsuarios(u Usuario) (Usuario, error) {
	sqlStatement := "INSERT INTO usuarios (nombre, dni) VALUES (?,?)"

	res, err := s.db.Exec(sqlStatement, u.Nombre, u.Dni)
	if err != nil {
		return u, err
	}

	u.ID, _ = res.LastInsertId()
	fmt.Println(res.LastInsertId())

	return u, err
}

// FindAll ...
func (s service) FindAll() []*Usuario {
	var list []*Usuario
	if err := s.db.Select(&list, "SELECT * FROM usuarios"); err != nil {
		panic(err)
	}
	return list
}
