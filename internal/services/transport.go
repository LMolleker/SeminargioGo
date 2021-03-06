package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

//HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

//NewHTTPTransport ...
func NewHTTPTransport(s BeerService) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s BeerService) []*endpoint {
	list := []*endpoint{}
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/beer",
		function: insert(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/beers",
		function: findAll(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/beer/:id",
		function: delete(s),
	})

	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/beer/:id",
		function: update(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/beer/:id",
		function: findByID(s),
	})
	return list
}

func findByID(s BeerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		beer, err := s.FindByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Beer could not be found",
			})
			fmt.Println(err)
			os.Exit(1)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Beer": beer,
			})
		}
	}
}

func update(s BeerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var beer Beer
		if err = json.Unmarshal(data, &beer); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err = json.Unmarshal(data, &beer); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		id, err = s.Update(id, beer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Beer could not be updated",
			})
			fmt.Println(err)
			os.Exit(1)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Beer updated",
			})
		}
	}
}

func delete(s BeerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		_, err = s.Delete(id)
		//ID, err := s.Delete(strconv.Atoi(c.Param("id")))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Beer could not be deleted",
			})
			fmt.Println(err)
			os.Exit(1)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Beer deleted",
			})
		}
	}
}

func findAll(s BeerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		beers, err := s.FindAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Beers could not be found",
			})
			fmt.Println(err)
			os.Exit(1)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Data": beers,
			})
		}
	}
}

func insert(s BeerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var beer Beer
		if err = json.Unmarshal(data, &beer); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = s.Insert(beer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Beers could not be inserted",
			})
			fmt.Println(err)
			os.Exit(1)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Beer inserted",
			})
		}
	}
}

type httpService struct {
	endpoints []*endpoint
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

//Register ...
func (h httpService) Register(gin *gin.Engine) {
	for _, e := range h.endpoints {
		gin.Handle(e.method, e.path, e.function)
	}
}
