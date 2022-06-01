package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"server/internal/domein/event"
	"strconv"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var event event.Event
		err := json.NewDecoder(request.Body).Decode(&event)
		if err != nil {
			fmt.Printf("json.NewDecoder:", err)
		}
		defer request.Body.Close()

		eventUpdate, err := (*c.service).Update(&event)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = internalServerError(writer, err)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
			return
		}

		err = success(writer, eventUpdate)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
		}

	}
}

func (c *EventController) Insert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var event event.Event
		err := json.NewDecoder(request.Body).Decode(&event)
		if err != nil {
			fmt.Printf("json.NewDecoder:", err)
		}
		defer request.Body.Close()

		eventInsert, err := (*c.service).Insert(&event)
		if err != nil {
			fmt.Printf("EventController.Insert(): %s", err)
			err = internalServerError(writer, err)
			if err != nil {
				fmt.Printf("EventController.Insert(): %s", err)
			}
			return
		}

		err = success(writer, eventInsert)
		if err != nil {
			fmt.Printf("EventController.Insert(): %s", err)
		}
	}
}

func (c *EventController) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var event event.Event
		err := json.NewDecoder(request.Body).Decode(&event)
		if err != nil {
			fmt.Printf("json.NewDecoder:", err)
		}
		defer request.Body.Close()

		eventDelete, err := (*c.service).Delete(&event)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
			err = internalServerError(writer, err)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
			return
		}

		err = success(writer, eventDelete)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
		}
	}
}
