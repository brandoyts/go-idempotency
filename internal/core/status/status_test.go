package status

import (
	"fmt"
	"testing"
)

func TestStatus_ToString(t *testing.T) {
	tests := []struct {
		name     string
		status   Status
		expected string
	}{
		{
			name:     "Processing",
			status:   Processing,
			expected: "Processing",
		},
		{
			name:     "Done",
			status:   Done,
			expected: "Done",
		},
		{
			name:     "Unknown status",
			status:   Status(99),
			expected: fmt.Sprintf("unknown status (%v)", Status(99)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.status.ToString()
			if got != tt.expected {
				t.Errorf("ToString() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
