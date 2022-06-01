package event

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Update(event *Event) (*Event, error)
	Insert(event *Event) ([]Event, error)
	Delete(event *Event) ([]Event, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*Event, error) {
	return (*s.repo).FindOne(id)
}

func (s *service) Update(event *Event) (*Event, error) {
	return (*s.repo).Update(event)
}

func (s *service) Insert(event *Event) ([]Event, error) {
	return (*s.repo).Insert(event)
}

func (s *service) Delete(event *Event) ([]Event, error) {
	return (*s.repo).Delete(event)
}
