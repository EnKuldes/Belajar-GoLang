// Membuat unit testing terhadap module
package greetings

import (
	// Import Module Testing
	"testing"
	// Import Module Regexp
	"regexp"
)

// Test Func dg parameter ada
func TestHelloName(t *testing.T) {
	name := "Farhan"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := HelloV3("Farhan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`HelloV3("Farhan") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Test Func dg parameter ga ada
func TestHelloEmpty(t *testing.T) {
	msg, err := HelloV3("")
	if msg != "" || err == nil {
		t.Fatalf(`HelloV3("") = %q, %v, want "", error`, msg, err)
	}
}