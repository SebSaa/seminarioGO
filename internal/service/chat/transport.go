package chat

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/messages",
		function: getAll(s),
	}, &endpoint{
		method:   "GET",
		path:     "/messages/:id",
		function: getById(s),
	}, &endpoint{
		method:   "POST",
		path:     "/messages",
		function: addMessage(s),
	}, &endpoint{
		method:   "PUT",
		path:     "/messages",
		function: updateMessage(s),
	}, &endpoint{
		method:   "DELETE",
		path:     "/messages/:id",
		function: delMessage(s),
	})

	return list
}

func updateMessage(s Service) gin.HandlerFunc {
	var m Message
	return func(c *gin.Context) {
		c.BindJSON(&m)
		result, err := s.UpdateMensaje(m)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": result,
			})
		}
	}
}

func getById(s Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, err := s.FindByID(ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": result,
			})
		}
	}
}

func delMessage(s Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, err := s.RemoveByID(ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": result,
			})
		}
	}
}

func addMessage(s Service) gin.HandlerFunc {
	var m Message
	return func(c *gin.Context) {
		c.BindJSON(&m)
		result, err := s.AddMessage(m)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": result,
			})
		}
	}
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": s.FindAll(),
		})
	}

}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
