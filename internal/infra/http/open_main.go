package http

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"server/internal/domein/event"
	"strconv"
)

type UseWeb struct {
	service *event.Service
}

func NewUseWeb(s *event.Service) *UseWeb {
	return &UseWeb{
		service: s,
	}
}

func (c *UseWeb) UpdateInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		temp, err := template.ParseFiles("internal/web/Main.html")
		if err != nil {
			log.Fatal("template.ParseFiles:", err)
		}

		if request.FormValue("id") == "" {
			fmt.Println("No values specified!")
		} else {
			id, err := strconv.Atoi(request.FormValue("id"))
			if err != nil {
				log.Fatal("strconv.Atoi:", err)
			}
			event2 := event.Event{
				EventId:          int64(id),
				Title:            request.FormValue("title"),
				ShortDescription: request.FormValue("shortDes"),
				Description:      request.FormValue("Des"),
				LongLat:          request.FormValue("long_lat"),
				Images:           request.FormValue("images"),
				Preview:          request.FormValue("preview"),
			}

			err3 := (*c.service).CheckMove(&event2)
			if err3 != nil {
				log.Fatal("CheckMove:", err3)
			}
		}
		err2 := temp.Execute(writer, "lol")
		if err2 != nil {
			log.Fatal("Problem:", err2)
		}
	}
}
