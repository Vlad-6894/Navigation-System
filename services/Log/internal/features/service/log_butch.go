package log_service

import (
	log_domains "Log_service/internal/core/domains"
	"context"
	"sync"
)

func (s LogService) LogBatch(
	ctx context.Context,
	events []log_domains.LogEvent,
) (int, error) {
	num := -1
	wg := &sync.WaitGroup{}

	for i, event := range events {
		num = i
		if event.GetErr() == nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logInter.LogMesage(event)
			}()
		}
	}

	wg.Wait()

	return num, nil
}

func (l *LoggerRepository) LogMesage(event log_domains.LogEvent) {

	logger := event.With(l.log)
	logger.Info("")
}
