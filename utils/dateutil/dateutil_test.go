package dateutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsQuarterly(t *testing.T) {
	tests := []struct {
		give time.Time
		want bool
	}{
		{
			give: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			give: time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			give: time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC),
			want: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, IsQuarterly(test.give))
	}
}
