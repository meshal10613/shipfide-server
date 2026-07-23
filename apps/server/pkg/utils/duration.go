package utils

import (
	"fmt"
	"server/internal/config"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(jwtDuration *string) (time.Duration, error) {
	var duration string

	if jwtDuration != nil {
		duration = *jwtDuration
	} else {
		cfg, err := config.LoadEnv()
		if err != nil {
			return 0, fmt.Errorf("failed to load environment variables: %w", err)
		}

		duration = cfg.JwtDuration
	}

	duration = strings.TrimSpace(strings.ToLower(duration))

	if strings.HasSuffix(duration, "d") {
		daysStr := strings.TrimSuffix(duration, "d")

		days, err := strconv.Atoi(daysStr)
		if err != nil {
			return 0, fmt.Errorf("invalid day duration: %s", duration)
		}

		return time.Duration(days) * 24 * time.Hour, nil
	}

	return time.ParseDuration(duration)
}

// GetBoundedTokenDuration parses a duration string and clamps it between 7 days and 30 days.
func GetBoundedTokenDuration(cfgDuration string) time.Duration {
	dur, err := ParseDuration(&cfgDuration)
	if err != nil {
		dur = 7 * 24 * time.Hour
	}
	minDur := 7 * 24 * time.Hour
	maxDur := 30 * 24 * time.Hour
	if dur < minDur {
		dur = minDur
	}
	if dur > maxDur {
		dur = maxDur
	}
	return dur
}
