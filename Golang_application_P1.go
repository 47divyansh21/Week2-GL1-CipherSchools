package handler

//book.go
import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahul10-pu/CIGO0122/database"
	"github.com/rahul10-pu/CIGO0122/models"
)

func (h *Handler) GetBooks(in *gin.Context) {
	books, err := database.GetBooks(h.DB)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books)
}
func (h *Handler) SaveBook(in *gin.Context) {
	book := models.Book{}
	err := in.BindJSON(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SaveBook(h.DB, &book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}
func (h *Handler) GetBookByID(in *gin.Context) {
	id := in.Params.ByName("id")
	book, err := database.GetBookByID(h.DB, id)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}
func (h *Handler) DeleteBookByID(in *gin.Context) {
	id := in.Params.ByName("id")
	err := database.DeleteBookByID(h.DB, id)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, id)
}
func (h *Handler) UpdateBookByID(in *gin.Context) {
	book := models.Book{}
	err := in.BindJSON(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.UpdateBookByID(h.DB, &book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}

/*
h:=Handler{}
h.DB=GetDB()
h.GetBooks()
*/

/

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rahul10-pu/CIGO0122/handler"
)

func BookRouter(router *gin.Engine, api handler.Handler) {
	get the default engine for further customizatikon
	api := handler.Handler{
		DB: database.GetDB(), 
	}

	router.GET("/books", api.GetBooks) 
	
	router.POST("/books", api.SaveBook)
	router.GET("/books/:id", api.GetBookByID)
	router.DELETE("/books/:id", api.DeleteBookByID)
	router.PUT("/books", api.UpdateBookByID)
	router.DELETE("/books/:id", api.DeleteBook)
	return router
}


/*
http://localhost:8080/book/11         - params
http://localhost:8080/book?id=11&name=abc	  - query
*/
