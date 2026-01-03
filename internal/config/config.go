package config

import (
	"errors"
	"fmt"
)

// Duration represents the time range for trending repositories
type Duration string

const (
	DurationDay   Duration = "day"
	DurationWeek  Duration = "week"
	DurationMonth Duration = "month"
	DurationYear  Duration = "year"
)

// Config holds the CLI configuration
type Config struct {
	Duration Duration
	Limit    int
}

// Validate checks if configuration is valid
func (c *Config) Validate() error {
	if c.Limit <= 0 {
		return errors.New("limit must be greater than 0")
	}
	if c.Limit > 100 {
		return errors.New("limit cannot exceed 100")
	}

	switch c.Duration {
	case DurationDay, DurationWeek, DurationMonth, DurationYear:
		return nil
	default:
		return fmt.Errorf("invalid duration: %s (must be day, week, month, or year)", c.Duration)
	}
}

// GetDurationInDays returns the number of days for the duration
func (c *Config) GetDurationInDays() int {
	switch c.Duration {
	case DurationDay:
		return 1
	case DurationWeek:
		return 7
	case DurationMonth:
		return 30
	case DurationYear:
		return 365
	default:
		return 7
	}
}
