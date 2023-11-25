package postgres

import (
	"github.com/andrepriyanto10/favaa_mitra/config/environment"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestOpenConnection(t *testing.T) {

	cfg, _ := env.env.LoadEnv()

	tests := []struct {
		name    string
		args    *viper.Viper
		want    *gorm.DB
		wantErr bool
	}{
		{
			name:    "Test Open Connection",
			args:    cfg,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := OpenConnection(tt.args)

			assert.Nil(t, err)

			//if (err != nil) != tt.wantErr {
			//	t.Errorf("OpenConnection() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}

			//if got != tt.want {
			//	t.Errorf("OpenConnection() got = %v, want %v", got, tt.want)
			//}
		})
	}

}
