package publisher

type publisher interface {
	Publish(queue, body string) error
	CreateQueue(queue string) error
}

type Service struct {
	publisher publisher
}

func NewService(publisher publisher) *Service {
	return &Service{publisher: publisher}
}

func (s *Service) Publish(queue, body string) error {
	return s.publisher.Publish(queue, body)
}

func (s *Service) CreateQueue(queue string) error {
	return s.publisher.CreateQueue(queue)
}
