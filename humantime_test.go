package humantime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testcases = []struct {
	duration string
	expect   string
}{
	{
		duration: "1s",
		expect:   "just now",
	},
	{
		duration: "30s",
		expect:   "seconds ago",
	},
	{
		duration: "60s",
		expect:   "a minute ago",
	},
	{
		duration: "4m",
		expect:   "4 minutes ago",
	},
	{
		duration: "1h",
		expect:   "an hour ago",
	},
	{
		duration: "1.34h",
		expect:   "an hour ago",
	},
	{
		duration: "2h",
		expect:   "2 hours ago",
	},
	{
		duration: "23h",
		expect:   "23 hours ago",
	},
	{
		duration: "24h",
		expect:   "a day ago",
	},
	{
		duration: "32h",
		expect:   "a day ago",
	},
	{
		duration: "48h",
		expect:   "2 days ago",
	},
	{
		duration: "72h",
		expect:   "3 days ago",
	},
	{
		duration: "168h",
		expect:   "a week ago",
	},
	{
		duration: "200h",
		expect:   "a week ago",
	},
	{
		duration: "336h",
		expect:   "2 weeks ago",
	},
	{
		duration: "401h",
		expect:   "2 weeks ago",
	},
	{
		duration: "672h",
		expect:   "4 weeks ago",
	},
	{
		duration: "720h",
		expect:   "4 weeks ago",
	},
	{
		duration: "730h",
		expect:   "about a month ago",
	},
	{
		duration: "900h",
		expect:   "about a month ago",
	},
	{
		duration: "1300h",
		expect:   "about a month ago",
	},
	{
		duration: "1460h",
		expect:   "about 2 months ago",
	},
	{
		duration: "2190h",
		expect:   "about 3 months ago",
	},
	{
		duration: "8030h",
		expect:   "about 11 months ago",
	},
	{
		duration: "8760h",
		expect:   "a year ago",
	},
	{
		duration: "16000h",
		expect:   "a year ago",
	},
	{
		duration: "17520h",
		expect:   "2 years ago",
	},
	{
		duration: "271560h",
		expect:   "31 years ago",
	},
	{
		duration: "274560h",
		expect:   "31 years ago",
	},
	{
		duration: "876000h",
		expect:   "100 years ago",
	},
}

func TestDuration(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.duration, func(t *testing.T) {
			d, err := time.ParseDuration(tc.duration)
			require.NoError(t, err, "bug in TEST CASE - parse duration failed")

			timeStr := Duration(d)
			assert.Equal(t, tc.expect, timeStr, "output does not match expect")
		})
	}
}

func TestSince(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.duration, func(t *testing.T) {
			d, err := time.ParseDuration(tc.duration)
			require.NoError(t, err, "bug in TEST CASE - parse duration failed")

			nowLessDuration := time.Now().Add(-d)

			timeStr := Since(nowLessDuration)
			assert.Equal(t, tc.expect, timeStr, "output does not match expect")
		})
	}
}
