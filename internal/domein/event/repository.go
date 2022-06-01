package event

import (
	"fmt"
	"github.com/upper/db/v4"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Update(event *Event) (*Event, error)
	Insert(event *Event) ([]Event, error)
	Delete(event *Event) ([]Event, error)
}

type repository struct {
	sess db.Session
}

func NewRepository(sees db.Session) Repository {
	return &repository{sess: sees}
}

func (r *repository) FindAll() ([]Event, error) {
	var event []Event
	err := r.sess.Collection("event").Find().All(&event)
	if err != nil {
		fmt.Printf("r.sess.Collection.All:", err)
		return nil, err
	}
	return event, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	var event Event
	err := r.sess.Collection("event").Find(id).One(&event)
	if err != nil {
		fmt.Println("r.sess.Collection.One:", err)
		return nil, err
	}
	return &event, nil
}

func (r *repository) Update(event *Event) (*Event, error) {
	err := r.sess.Collection("event").Find(event.EventId).Update(event)
	if err != nil {
		fmt.Printf("r.sess.Collection:", err)
		return nil, err
	}
	return event, nil
}

func (r *repository) Insert(event *Event) ([]Event, error) {
	err := r.sess.Collection("event").InsertReturning(event)

	if err != nil {
		fmt.Println("r.sess.Collection.Insert:", err)
		return nil, err
	}

	var eventAll []Event
	err2 := r.sess.Collection("event").Find().All(&eventAll)
	if err2 != nil {
		fmt.Printf("r.sess.Collection:", err)
		return nil, err2
	}
	return eventAll, nil
}

func (r *repository) Delete(event *Event) ([]Event, error) {

	err := r.sess.Collection("event").Find(event.EventId).Delete()
	if err != nil {
		fmt.Printf("r.sess.Collection.Delete:", err)
		return nil, err
	}

	var eventAll []Event
	err2 := r.sess.Collection("event").Find().All(&eventAll)
	if err2 != nil {
		fmt.Printf("r.sess.Collection:", err)
		return nil, err2
	}
	return eventAll, nil
}
