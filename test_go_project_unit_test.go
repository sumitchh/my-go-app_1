// main_test.go
package main
import "testing"
func TestMain(t *testing.T) {
    got := "Hello, CI!"
    want := "Hello, CI!"
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}

