package greeter

import (
	"testing"
	"regexp"
)

func TestGreetName(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` +name+ `\b`)
	msg, err := Greet(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greet("Juan") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGreetEmpty(t *testing.T) {
	msg, err := Greet("")
	if msg != "" || err == nil {
		t.Fatalf(`Greet("") = %q, %v, want "", error`, msg, err)
	}
}
