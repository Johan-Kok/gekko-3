
package log4go

import (
    "testing"
)

// test WhenIsValid()
func TestWhenIsValid(t *testing.T) {
    if ! WhenIsValid("midNIGHT") {
        t.Error("err in WhenIsValid('midNIGHT')")
    }

    if WhenIsValid("mid-night") {
        t.Error("err in WhenIsValid('mid-night')")
    }

    if !WhenIsValid("m") {
        t.Error("err in WhenIsValid('m')")
    }

    if !WhenIsValid("H") {
        t.Error("err in WhenIsValid('H')")
    }

    if !WhenIsValid("d") {
        t.Error("err in WhenIsValid('H')")
    }
}