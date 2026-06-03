// Package audit provides a tiny reusable behavior — created/updated tracking —
// that other types gain by EMBEDDING it. Embedding is Go's composition tool:
// the outer type gets Timestamps' fields and methods promoted onto it, without
// any inheritance.
package audit

import "time"

// Timestamps records when something was created and last changed.
type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Touch stamps UpdatedAt (and CreatedAt the first time). A pointer receiver is
// required because it mutates the struct.
func (t *Timestamps) Touch(now time.Time) {
	if t.CreatedAt.IsZero() {
		t.CreatedAt = now
	}
	t.UpdatedAt = now
}
