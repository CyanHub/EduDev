package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type SagaStep struct {
	Name       string
	Execute    func() error
	Compensate func() error
}

type Saga struct {
	steps []SagaStep
	redis *redis.Client
}

func NewSaga(redisClient *redis.Client) *Saga {
	return &Saga{
		redis: redisClient,
	}
}

func (s *Saga) AddStep(step SagaStep) {
	s.steps = append(s.steps, step)
}

func (s *Saga) Run(ctx context.Context) error {
	var executedSteps []SagaStep

	for _, step := range s.steps {
		err := step.Execute()
		if err != nil {
			log.Printf("Saga execution failed at step %s, starting compensation", step.Name)
			return s.compensate(executedSteps)
		}
		executedSteps = append(executedSteps, step)
	}
	return nil
}

func (s *Saga) compensate(steps []SagaStep) error {
	for i := len(steps) - 1; i >= 0; i-- {
		step := steps[i]
		if step.Compensate != nil {
			if err := step.Compensate(); err != nil {
				log.Printf("saga补偿失败: 补偿步骤 %s 失败: %v", step.Name, err)
				return errors.New("saga补偿失败")
			}
		}
	}
	return nil
}

func (s *Saga) SaveState(ctx context.Context, key string, state interface{}) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return s.redis.Set(ctx, key, data, 24*time.Hour).Err()
}

func (s *Saga) LoadState(ctx context.Context, key string, out interface{}) error {
	data, err := s.redis.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, out)
}
