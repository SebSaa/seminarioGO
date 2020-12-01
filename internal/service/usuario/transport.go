package usuario

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
		path:     "/usuarios",
		function: getAll(s),
	}, &endpoint{
		method:   "GET",
		path:     "/usuarios/:id",
		function: getById(s),
	}, &endpoint{
		method:   "POST",
		path:     "/usuarios",
		function: addUsuarios(s),
	}, &endpoint{
		method:   "PUT",
		path:     "/usuarios",
		function: updateUsuarios(s),
	}, &endpoint{
		method:   "DELETE",
		path:     "/usuarios/:id",
		function: delUsuarios(s),
	})

	return list
}

func updateUsuarios(s Service) gin.HandlerFunc {
	var u Usuario
	return func(c *gin.Context) {
		c.BindJSON(&u)
		result, err := s.UpdateUsuarios(u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"usuario": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"usuario": result,
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
				"usuario": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"usuario": result,
			})
		}
	}
}

func delUsuarios(s Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, err := s.RemoveByID(ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"usuario": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"usuario": result,
			})
		}
	}
}

func addUsuarios(s Service) gin.HandlerFunc {
	var u Usuario
	return func(c *gin.Context) {
		c.BindJSON(&u)
		result, err := s.AddUsuarios(u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"usuarios": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"usuarios": result,
			})
		}
	}
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"usuarios": s.FindAll(),
		})
	}

}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
