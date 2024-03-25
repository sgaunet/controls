package config_test

import (
	"testing"

	"github.com/sgaunet/controls/internal/config"
	"github.com/sgaunet/controls/internal/httpctl"
)

func TestYamlConfig_IsValid(t *testing.T) {
	tests := []struct {
		name string
		cfg  config.YamlConfig
		want bool
	}{
		{
			name: "Empty config",
			cfg:  config.YamlConfig{},
			want: true,
		},
		{
			name: "zabbix config with empty user",
			cfg: config.YamlConfig{
				ZbxCtl: &config.ZbxCtlConfig{
					APIEndpoint:          "http://localhost",
					User:                 "",
					Password:             "password",
					Since:                0,
					SeverityThreshold:    0,
					FilterProblemsByTags: []config.ZbxTagsFilterProblemConfig{},
				},
			},
			want: false,
		},
		{
			name: "zabbix config with empty password",
			cfg: config.YamlConfig{
				ZbxCtl: &config.ZbxCtlConfig{
					APIEndpoint:          "http://localhost",
					User:                 "user",
					Password:             "",
					Since:                0,
					SeverityThreshold:    0,
					FilterProblemsByTags: []config.ZbxTagsFilterProblemConfig{},
				},
			},
			want: false,
		},
		{
			name: "zabbix config with endpoint",
			cfg: config.YamlConfig{
				ZbxCtl: &config.ZbxCtlConfig{
					APIEndpoint:          "",
					User:                 "user",
					Password:             "****",
					Since:                0,
					SeverityThreshold:    0,
					FilterProblemsByTags: []config.ZbxTagsFilterProblemConfig{},
				},
			},
			want: false,
		},
		{
			name: "zabbix config with FilterProblemsByTags malformated",
			cfg: config.YamlConfig{
				ZbxCtl: &config.ZbxCtlConfig{
					APIEndpoint:       "http://localhost",
					User:              "user",
					Password:          "****",
					Since:             0,
					SeverityThreshold: 0,
					FilterProblemsByTags: []config.ZbxTagsFilterProblemConfig{
						{
							Tag:      "tag",
							Value:    "value",
							Operator: "invalid",
						},
					},
				},
			},
			want: false,
		},
		{
			name: "zabbix config OK",
			cfg: config.YamlConfig{
				ZbxCtl: &config.ZbxCtlConfig{
					APIEndpoint:       "http://localhost",
					User:              "user",
					Password:          "****",
					Since:             0,
					SeverityThreshold: 0,
					FilterProblemsByTags: []config.ZbxTagsFilterProblemConfig{
						{
							Tag:      "tag",
							Value:    "value",
							Operator: "not exists",
						},
					},
				},
			},
			want: true,
		},
		{
			name: "assertsHTTP with empty Host",
			cfg: config.YamlConfig{
				AssertsHTTP: []httpctl.AssertHTTP{
					{
						Title:      "title",
						Host:       "",
						HostHeader: "",
					},
				},
			},
			want: false,
		},
		{
			name: "assertsHTTP with empty Title",
			cfg: config.YamlConfig{
				AssertsHTTP: []httpctl.AssertHTTP{
					{
						Title:      "",
						Host:       "http://endpoint",
						HostHeader: "toto.com",
					},
				},
			},
			want: false,
		},
		{
			name: "assertsHTTP OK",
			cfg: config.YamlConfig{
				AssertsHTTP: []httpctl.AssertHTTP{
					{
						Title:      "Title",
						Host:       "http://endpoint",
						HostHeader: "",
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cfg.IsValid(); got != tt.want {
				t.Errorf("YamlConfig.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
