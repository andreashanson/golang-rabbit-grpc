package scheduler

import (
	"errors"
	"fmt"

	"github.com/andreashanson/golang-rabbit-grpc/internal/config"
	"github.com/andreashanson/golang-rabbit-grpc/internal/publisher"
	"github.com/robfig/cron/v3"
)

type Service struct {
	pub *publisher.Service
	cfg *config.RabbitConfig
}

func NewService(ps *publisher.Service, cfg *config.RabbitConfig) *Service {
	return &Service{pub: ps, cfg: cfg}
}

func (s *Service) Start() chan (error) {
	ss := cron.New()
	ss.Start()

	c := make(chan error)

	jobs := []job{
		{
			name:         "test1",
			cronSchedule: "@every 5s",
			jobFunc: func() {
				fmt.Println("TEST@!")
				err := s.pub.Publish(s.cfg.Queue, "TEST1")
				if err != nil {
					c <- err
				}
			},
		},
		{
			name:         "test2",
			cronSchedule: "@every 2s",
			jobFunc: func() {
				fmt.Println("TEST4")
				err := s.pub.Publish(s.cfg.Queue, "TEST1")
				err = errors.New("TEST ERROR")
				if err != nil {
					c <- err
				}
			},
		},
	}

	for _, j := range jobs {
		_, err := ss.AddFunc(j.cronSchedule, j.jobFunc)
		if err != nil {
			c <- err
		}
	}
	fmt.Println("HEJ")
	return c
}

type job struct {
	name         string
	cronSchedule string
	jobFunc      func()
}
