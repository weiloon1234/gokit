package utils

import "time"

// BoolValue safely dereferences a *bool, returning false if the pointer is nil.
func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// IntValue safely dereferences a *int, returning 0 if the pointer is nil.
func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// Int64Value safely dereferences a *int64, returning 0 if the pointer is nil.
func Int64Value(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// StringValue safely dereferences a *string, returning an empty string if the pointer is nil.
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Float64Value safely dereferences a *float64, returning 0.0 if the pointer is nil.
func Float64Value(f *float64) float64 {
	if f == nil {
		return 0.0
	}
	return *f
}

// TimeValue safely dereferences a *time.Time, returning the zero time if the pointer is nil.
func TimeValue(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

// DurationValue safely dereferences a *time.Duration, returning 0 if the pointer is nil.
func DurationValue(d *time.Duration) time.Duration {
	if d == nil {
		return 0
	}
	return *d
}
