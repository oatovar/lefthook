package log

import (
	"fmt"
	"testing"
)

func TestSkipSetting(t *testing.T) {
	for i, tt := range [...]struct {
		settings []string
		results  map[string]bool
	}{
		{
			settings: []string{},
			results:  map[string]bool{},
		},
		{
			settings: []string{"failure", "execution"},
			results: map[string]bool{
				"failure":   true,
				"execution": true,
			},
		},
		{
			settings: []string{"meta", "summary", "success", "failure", "execution"},
			results: map[string]bool{
				"meta":      true,
				"summary":   true,
				"success":   true,
				"failure":   true,
				"execution": true,
			},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			var settings SkipSettings

			for _, option := range tt.settings {
				(&settings).ApplySetting(option)
			}

			if settings.SkipMeta() != tt.results["meta"] {
				t.Errorf("expected SkipMeta to be %v", tt.results["meta"])
			}

			if settings.SkipSuccess() != tt.results["success"] {
				t.Errorf("expected SkipSuccess to be %v", tt.results["success"])
			}

			if settings.SkipFailure() != tt.results["failure"] {
				t.Errorf("expected SkipFailure to be %v", tt.results["failure"])
			}

			if settings.SkipSummary() != tt.results["summary"] {
				t.Errorf("expected SkipSummary to be %v", tt.results["summary"])
			}

			if settings.SkipExecution() != tt.results["execution"] {
				t.Errorf("expected SkipExecution to be %v", tt.results["execution"])
			}
		})
	}
}
