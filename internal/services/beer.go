package services

import (
	"SeminarioGo/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

//BeerService ...
type BeerService interface {
	Insert(Beer) error
	FindByID(int) *Beer
	Update(int, Beer) int
	Delete(int, error) int
	FindAll() []*Beer
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
func (s beerService) FindByID(ID int) *Beer {
	query := "SELECT * FROM beer WHERE id = :id"
	smtCreate, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ID)
	var id int
	var name string
	var alcoholContent, price float32
	err = smtCreate.QueryRow(ID).Scan(&id, &name, &alcoholContent, &price)
	if err != nil {
		fmt.Println(err.Error())
	}
	if id > 0 {
		return &Beer{id, name, alcoholContent, price}
	}
	return nil
}

//Update one beer ...
func (s beerService) Update(ID int, b Beer) int {
	query := "UPDATE beer SET name = ?, alcohol_content = ?, price = ? WHERE id = :id"
	stmtCreate, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmtCreate.Exec(b.Name, b.AlcoholContent, b.Price, ID)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ID
}

//Find all beers ...
func (s beerService) FindAll() []*Beer {
	var beers []*Beer
	query := "SELECT * FROM beer"
	smtCreate, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	list, err := smtCreate.Query()
	if err != nil {
		fmt.Println(err.Error())
	}
	for list.Next() {
		var id int
		var name string
		var alcoholContent, price float32
		list.Scan(&id, &name, &alcoholContent, &price)
		beers = append(beers, &Beer{id, name, alcoholContent, price})
	}
	return beers
}

//Delete beer ...
func (s beerService) Delete(ID int, e error) int {
	query := "DELETE FROM beer WHERE id = :id"
	stmtCreate, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmtCreate.Exec(ID)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ID
}
