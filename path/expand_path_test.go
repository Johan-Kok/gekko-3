package path

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func TestExpand(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	absFoo, _ := filepath.Abs("foo/abc")
	goFoo := os.Getenv("GOPATH") + "/foo/abc"
	cases := []struct {
		Input  string
		Output string
		Err    bool
	}{
		{
			"/foo",
			"/foo",
			false,
		},

		{
			"~/foo",
			filepath.Join(u.HomeDir, "foo"),
			false,
		},

		{
			"",
			"",
			false,
		},

		{
			"~",
			u.HomeDir,
			false,
		},

		{
			"~foo/foo",
			"",
			true,
		},

		{
			"$HOME/foo",
			filepath.Join(u.HomeDir, "foo"),
			false,
		},

		{
			"$HOME",
			u.HomeDir,
			false,
		},

		{
			"$HOMEabc",
			"",
			true,
		},

		{
			"/foo/abc/",
			"/foo/abc",
			false,
		},

		{
			"/foo/abc",
			"/foo/abc",
			false,
		},
		{
			"$null/foo/abc/",
			"",
			true,
		},
		{
			"$GOPATH/foo/abc/",
			goFoo,
			false,
		},
		{
			"./foo/abc/",
			absFoo,
			false,
		},
		{
			"./foo/abc",
			absFoo,
			false,
		},
	}

	for _, tc := range cases {
		actual, err := Expand(tc.Input)
		if (err != nil) != tc.Err {
			t.Fatalf("Input: %#v\n\nErr: %s", tc.Input, err)
		}

		if actual != tc.Output {
			t.Fatalf("Input: %#v\n\nOutput: %#v", tc.Input, actual)
		}
	}

}
