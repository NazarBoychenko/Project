package event

import (
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	CheckMove(event *Event) error
}

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("postgresql.Open: ", err)
	}
	defer sess.Close()

	events := sess.Collection("event")

	eve := []Event{}
	err = events.Find().All(&eve)
	if err != nil {
		log.Fatal("events.Find: ", err)
	}

	return eve, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("postgresql.Open: ", err)
	}
	defer sess.Close()

	eventCollection := sess.Collection("event")
	res := eventCollection.Find()
	count, _ := res.Count()
	if id <= int64(count) {
		q := sess.SQL().SelectFrom("event").Where("\"event_id\" =?", id)
		var event Event
		if err := q.One(&event); err != nil {
			log.Fatal("q.One:", err)
			return nil, err
		}
		return &event, nil
	} else {
		return nil, nil
	}
}

func (r *repository) CheckMove(event *Event) error {
	if event.EventId == 0 {
		Insert(event)
		return nil
	} else {
		Update(event)
		return nil
	}
}

func Update(event *Event) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("postgresql.Open: ", err)
	}
	defer sess.Close()

	sess.SQL().Update("event").Set("title = ?", event.Title, "shortDescription = ?", event.ShortDescription, "description = ?", event.Description, "long_lat = ?", event.LongLat, "images = ?", event.Images, "preview = ?", event.Preview).
		Where("\"event_id\" =?", event.EventId).Exec()

	if err != nil {
		log.Fatal("sess.Update:", err)
	}
}

func Insert(event *Event) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("postgresql.Open: ", err)
	}
	defer sess.Close()

	sess.SQL().InsertInto("event").Columns("title", "shortDescription", "description", "long_lat", "images", "preview").
		Values(event.Title, event.ShortDescription, event.Description, event.LongLat, event.Images, event.Preview).Exec()

	if err != nil {
		log.Fatal("sess.Insert:", err)
	}
}
