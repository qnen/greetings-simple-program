package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Brecci"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Brecci")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Brecci") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
