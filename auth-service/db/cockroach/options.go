package cockroach

import "time"

// Option -.
type Option func(*Cockroach)

// MaxPoolSize -.
func MaxPoolSize(size int) Option {
	return func(c *Cockroach) {
		c.maxPoolSize = size
	}
}

// ConnAttempts -.
func ConnAttempts(attempts int) Option {
	return func(c *Cockroach) {
		c.connAttempts = attempts
	}
}

// ConnTimeout -.
func ConnTimeout(timeout time.Duration) Option {
	return func(c *Cockroach) {
		c.connTimeout = timeout
	}
}
