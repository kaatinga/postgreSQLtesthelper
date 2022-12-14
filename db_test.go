package postgreSQLtesthelper

import (
	"context"
	"testing"
)

func TestSetupPostgreDatabase(t *testing.T) {
	tests := []struct {
		login    string
		password string
	}{
		{"kaatinga", "12345"},
	}
	for _, tt := range tests {
		t.Run(tt.login+":"+tt.password, func(t *testing.T) {
			dbContainer, pool := SetupPostgreDatabase(tt.login, tt.password, t)
			defer pool.Close()
			defer dbContainer.Terminate(context.Background())
			err := pool.Ping(context.Background())
			if err != nil {
				t.Error("ping failed:", err)
			}
		})
	}
}
