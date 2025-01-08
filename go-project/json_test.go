package go_project_test

import (
	"encoding/json"
	"testing"
)

type Settings struct {
	Enabled bool `json:"enabled"`
}

func (s *Settings) FromBytes(bytes []byte) error {
	if len(bytes) == 0 || string(bytes) == "{}" {
		s.Enabled = true
		return nil
	}
	var tmpSettings struct {
		Enabled *bool `json:"enabled"`
	}
	err := json.Unmarshal(bytes, &tmpSettings)
	if err != nil {
		return err
	}
	if tmpSettings.Enabled != nil {
		s.Enabled = *tmpSettings.Enabled
	} else {
		s.Enabled = true
	}
	return nil
}

func TestFromBytes(t *testing.T) {
	tests := map[string]struct {
		input string
		want  Settings
	}{
		"with enabled set to true": {
			input: `{"enabled": true}`,
			want:  Settings{Enabled: true},
		},
		"with enabled set to false": {
			input: `{"enabled": false}`,
			want:  Settings{Enabled: false},
		},
		"with enabled set to true and other fields": {
			input: `{"enabled": true, "other": "value"}`,
			want:  Settings{Enabled: true},
		},
		"with enabled set to false and other fields": {
			input: `{"enabled": false, "other": "value"}`,
			want:  Settings{Enabled: false},
		},
		"with enabled not set - defaults to true": {
			input: `{}`,
			want:  Settings{Enabled: true},
		},
		"with enabled not set - defaults to true and other fields": {
			input: `{"other": "value"}`,
			want:  Settings{Enabled: true},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			settings := Settings{}
			err := settings.FromBytes([]byte(test.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if settings != test.want {
				t.Errorf("expected %v, got %v", test.want, settings)
			}
		})
	}

}
