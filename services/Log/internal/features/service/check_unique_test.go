package log_service

import (
	logs_mocks "Log_service/internal/features/service/mocks"
	"errors"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

var (
	cacheError = errors.New("error from cache")
)

func TestCheckUnique(t *testing.T) {
	tests := []struct {
		name         string
		uuid         uuid.UUID
		isUnique     bool
		wantErr      bool
		waitErr      error
		prepareCache func(m *logs_mocks.MockLogCache)
	}{
		{
			name:     "success test",
			uuid:     uuid.New(),
			isUnique: true,
			wantErr:  false,
			waitErr:  nil,
			prepareCache: func(m *logs_mocks.MockLogCache) {
				m.EXPECT().CheckUnique(gomock.Any(), gomock.Any()).Return(true, nil).Times(1)
			},
		},

		{
			name:     "fail test",
			uuid:     uuid.New(),
			isUnique: false,
			wantErr:  true,
			waitErr:  cacheError,
			prepareCache: func(m *logs_mocks.MockLogCache) {
				m.EXPECT().CheckUnique(gomock.Any(), gomock.Any()).Return(false, cacheError).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCache := logs_mocks.NewMockLogCache(ctrl)
			tt.prepareCache(mockCache)

			service := NewLogService(mockCache, nil)

			isUnique, err := service.CheckUnique(t.Context(), tt.uuid)

			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr=%v, but returned err=%v", tt.wantErr, err)
			}

			if !errors.Is(err, tt.waitErr) {
				t.Errorf("waitErr=%v, but returned err=%v", tt.waitErr, err)
			}
			if isUnique != tt.isUnique {
				t.Errorf("isUnique must be=%v, but returned=%v", tt.isUnique, isUnique)
			}
		})
	}
}
