package scheduler

import (
	"context"
	"errors"
	"time"
)

type Scheduler interface {
	Start(ctx context.Context) error
	Stop()
}

type BaseScheduler struct {
	cron *CronCore
}

func NewBaseScheduler() *BaseScheduler {
	return &BaseScheduler{
		cron: NewCronCore(),
	}
}

func (s *BaseScheduler) Start(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		s.Stop()
	}()
	s.cron.Start()
	return nil
}

func (s *BaseScheduler) Stop() {
	s.cron.Stop()
}

type ScheduleManager struct {
	schedulers []Scheduler
}

func NewSchedulerManager() *ScheduleManager {
	return &ScheduleManager{
		schedulers: make([]Scheduler, 0),
	}
}

func (sm *ScheduleManager) AddScheduler(s Scheduler) {
	sm.schedulers = append(sm.schedulers, s)
}

func (sm *ScheduleManager) StartAll(ctx context.Context) error {
	for _, s := range sm.schedulers {
		if err := s.Start(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (sm *ScheduleManager) StopAll() {
	for _, s := range sm.schedulers {
		s.Stop()
	}
}

type Spec int

const (
	Hourly Spec = iota
	Daily
	MidNight
	Minutely // 로컬 테스트용으로만 사용하길 바람
)

type (
	timeSpecParser struct {
		spec Spec
	}

	TimeSpecParser interface {
		ParseSpec() (string, error)
	}
)

func NewTimeSpecParser(spec Spec) TimeSpecParser {
	return &timeSpecParser{
		spec: spec,
	}
}

func (s timeSpecParser) ParseSpec() (string, error) {
	if s.spec == Hourly {
		return "@every 1h", nil
	} else if s.spec == Daily {
		return "@every 24h", nil
	} else if s.spec == MidNight {
		localTimeZone, _ := time.Now().Zone()
		if localTimeZone == "KST" { // KST 자정에 실행
			return "0 0 0 * * *", nil
		} else { // UTC 기준 KST 자정에서 9시간 전인 (즉, UTC 15:00)에 실행
			return "0 0 15 * * *", nil
		}
	} else if s.spec == Minutely {
		return "@every 1m", nil
	} else {
		return "", errors.New("미구현된 스펙입니다")
	}
}
