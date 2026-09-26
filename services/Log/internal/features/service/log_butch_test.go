package log_service

import (
	log_domains "Log_service/internal/core/domains"
	logs_mocks "Log_service/internal/features/service/mocks"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/mock/gomock"
)

func TestLogButch(t *testing.T) {
	tests := []struct {
		name              string
		events            []log_domains.LogEvent
		waitRightMessages int
		prepareLogger     func(m *logs_mocks.MockLogger)
	}{
		{
			name: "success test",
			events: []log_domains.LogEvent{
				log_domains.NewAuthLogEvent(uuid.New(), uuid.NewString(), "", nil, kafka.Message{}),
				log_domains.NewFailMessage(uuid.New(), errors.New("fail message"), kafka.Message{}),
			},
			waitRightMessages: 1,
			prepareLogger: func(m *logs_mocks.MockLogger) {
				m.EXPECT().LogMesage(gomock.Any()).Return().Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := logs_mocks.NewMockLogger(ctrl)
			tt.prepareLogger(mockLogger)

			service := NewLogService(nil, mockLogger)

			num, _ := service.LogBatch(t.Context(), tt.events)

			if num != tt.waitRightMessages {
				t.Errorf("waitRightMesssage=%v, but returned nu,=%v", tt.waitRightMessages, num)
			}
		})
	}
}
