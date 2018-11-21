package main

import (
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

func main() {
	connString := "root:example@localhost"
	conn, err := mgo.Dial(connString)
	if err != nil {
		log.Printf("dial mongodb server with connection string %q: %v", connString, err)
		return
	}

	h := &handler{
		mongo: conn,
		db:    "document",
		col:   "todo",
	}

	e := echo.New()

	e.GET("/todos", h.list)
	e.POST("/todos", h.create)
	e.GET("/todos/:id", h.view)
	e.PUT("/todos/:id", h.done)
	e.DELETE("/todos/:id", h.remove)

	e.Logger.Fatal(e.Start(":8080"))
}

type todo struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Topic string        `json:"topic" bson:"topic"`
	Done  bool          `json:"done" bson:"done"`
}

type handler struct {
	mongo *mgo.Session
	db    string
	col   string
}

func (h *handler) list(c echo.Context) error {
	conn := h.mongo.Copy()
	defer conn.Close()
	var ts []todo
	if err := conn.DB(h.db).C(h.col).Find(nil).All(&ts); err != nil {
		return err
	}
	c.JSON(http.StatusOK, ts)
	return nil
}

func (h *handler) view(c echo.Context) error {
	return nil
}

func (h *handler) create(c echo.Context) error {
	id := bson.NewObjectId()
	var t todo
	if err := c.Bind(&t); err != nil {
		return err
	}
	t.ID = id
	t.Done = false

	conn := h.mongo.Copy()
	defer conn.Close()
	if err := conn.DB(h.db).C(h.col).Insert(t); err != nil {
		return err
	}

	c.JSON(http.StatusOK, t)
	return nil
}

func (h *handler) done(c echo.Context) error {
	return nil
}

func (h *handler) remove(c echo.Context) error {
	return nil
}
