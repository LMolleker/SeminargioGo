package services

import (
	"SeminarioGo/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

//BeerService ...
type BeerService interface {
	Insert(Beer) error
	FindByID(int) (*Beer, error)
	Update(int, Beer) (int, error)
	Delete(int) (int, error)
	FindAll() ([]*Beer, error)
}

//Beer ...
type Beer struct {
	ID             int
	Name           string
	AlcoholContent float32
	Price          float32
}

type beerService struct {
	db     *sqlx.DB
	config *config.Config
}

//NewService ...
func NewService(db *sqlx.DB, c *config.Config) (BeerService, error) {
	return beerService{db, c}, nil
}

//Insert a beer ...
func (s beerService) Insert(b Beer) error {
	query := "INSERT INTO Beer (name, alcohol_content, price) VALUES (?,?,?)"

	stmtCreate, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println("err 1")
		return err
	}
	fmt.Println(b)
	_, err = stmtCreate.Exec(b.Name, b.AlcoholContent, b.Price)
	if err != nil {
		fmt.Println("err 2")
		return err
	}
	return nil
}

//Find beer by id ...
func (s beerService) FindByID(ID int) (*Beer, error) {
	query := "SELECT * FROM beer WHERE id = :id"
	smtCreate, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(ID)
	var id int
	var name string
	var alcoholContent, price float32
	err = smtCreate.QueryRow(ID).Scan(&id, &name, &alcoholContent, &price)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if id > 0 {
		return &Beer{id, name, alcoholContent, price}, nil
	}
	return nil, nil
}

//Update one beer ...
func (s beerService) Update(ID int, b Beer) (int, error) {
	query := "UPDATE beer SET name = ?, alcohol_content = ?, price = ? WHERE id = :id"
	stmtCreate, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	_, err = stmtCreate.Exec(b.Name, b.AlcoholContent, b.Price, ID)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return ID, nil
}

//Find all beers ...
func (s beerService) FindAll() ([]*Beer, error) {
	var beers []*Beer
	query := "SELECT * FROM beer"
	smtCreate, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	list, err := smtCreate.Query()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	for list.Next() {
		var id int
		var name string
		var alcoholContent, price float32
		list.Scan(&id, &name, &alcoholContent, &price)
		beers = append(beers, &Beer{id, name, alcoholContent, price})
	}
	return beers, nil
}

//Delete beer ...
func (s beerService) Delete(ID int) (int, error) {
	query := "DELETE FROM beer WHERE id = :id"
	stmtCreate, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	_, err = stmtCreate.Exec(ID)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return ID, nil
}
