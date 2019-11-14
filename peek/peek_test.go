

package peek

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

)

// peekFunc is used to identify which public function of the peek package or
// ConfigState a test applies to.
type peekFunc int

const (
	fCSFdump peekFunc = iota
	fCSFprint
	fCSFprintf
	fCSFprintln
	fCSPrint
	fCSPrintln
	fCSSdump
	fCSSprint
	fCSSprintf
	fCSSprintln
	fCSErrorf
	fCSNewFormatter
	fErrorf
	fFprint
	fFprintln
	fPrint
	fPrintln
	fSdump
	fSprint
	fSprintf
	fSprintln
)

// Map of peekFunc values to names for pretty printing.
var peekFuncStrings = map[peekFunc]string{
	fCSFdump:        "ConfigState.Fdump",
	fCSFprint:       "ConfigState.Fprint",
	fCSFprintf:      "ConfigState.Fprintf",
	fCSFprintln:     "ConfigState.Fprintln",
	fCSSdump:        "ConfigState.Sdump",
	fCSPrint:        "ConfigState.Print",
	fCSPrintln:      "ConfigState.Println",
	fCSSprint:       "ConfigState.Sprint",
	fCSSprintf:      "ConfigState.Sprintf",
	fCSSprintln:     "ConfigState.Sprintln",
	fCSErrorf:       "ConfigState.Errorf",
	fCSNewFormatter: "ConfigState.NewFormatter",
	fErrorf:         "peek.Errorf",
	fFprint:         "peek.Fprint",
	fFprintln:       "peek.Fprintln",
	fPrint:          "peek.Print",
	fPrintln:        "peek.Println",
	fSdump:          "peek.Sdump",
	fSprint:         "peek.Sprint",
	fSprintf:        "peek.Sprintf",
	fSprintln:       "peek.Sprintln",
}

func (f peekFunc) String() string {
	if s, ok := peekFuncStrings[f]; ok {
		return s
	}
	return fmt.Sprintf("Unknown peekFunc (%d)", int(f))
}

// peekTest is used to describe a test to be performed against the public
// functions of the peek package or ConfigState.
type peekTest struct {
	cs     *ConfigState
	f      peekFunc
	format string
	in     interface{}
	want   string
}

// peekTests houses the tests to be performed against the public functions of
// the peek package and ConfigState.
//
// These tests are only intended to ensure the public functions are exercised
// and are intentionally not exhaustive of types.  The exhaustive type
// tests are handled in the dump and format tests.
var peekTests []peekTest

// redirStdout is a helper function to return the standard output from f as a
// byte slice.
func redirStdout(f func()) ([]byte, error) {
	tempFile, err := ioutil.TempFile("", "ss-test")
	if err != nil {
		return nil, err
	}
	fileName := tempFile.Name()
	defer os.Remove(fileName) // Ignore error

	origStdout := os.Stdout
	os.Stdout = tempFile
	f()
	os.Stdout = origStdout
	tempFile.Close()

	return ioutil.ReadFile(fileName)
}

func initPeekTests() {
	// Config states with various settings.
	scsDefault := NewDefaultConfig()
	scsNoMethods := &ConfigState{Indent: " ", DisableMethods: true}
	scsNoPmethods := &ConfigState{Indent: " ", DisablePointerMethods: true}
	scsMaxDepth := &ConfigState{Indent: " ", MaxDepth: 1}
	scsContinue := &ConfigState{Indent: " ", ContinueOnMethod: true}
	scsNoPtrAddr := &ConfigState{DisablePointerAddresses: true}
	scsNoCap := &ConfigState{DisableCapacities: true}

	// Variables for tests on types which implement Stringer interface with and
	// without a pointer receiver.
	ts := stringer("test")
	tps := pstringer("test")

	type ptrTester struct {
		s *struct{}
	}
	tptr := &ptrTester{s: &struct{}{}}

	// depthTester is used to test max depth handling for structs, array, slices
	// and maps.
	type depthTester struct {
		ic    indirCir1
		arr   [1]string
		slice []string
		m     map[string]int
	}
	dt := depthTester{indirCir1{nil}, [1]string{"arr"}, []string{"slice"},
		map[string]int{"one": 1}}

	// Variable for tests on types which implement error interface.
	te := customError(10)

	peekTests = []peekTest{
		{scsDefault, fCSFdump, "", int8(127), "(int8) 127\n"},
		{scsDefault, fCSFprint, "", int16(32767), "32767"},
		{scsDefault, fCSFprintf, "%v", int32(2147483647), "2147483647"},
		{scsDefault, fCSFprintln, "", int(2147483647), "2147483647\n"},
		{scsDefault, fCSPrint, "", int64(9223372036854775807), "9223372036854775807"},
		{scsDefault, fCSPrintln, "", uint8(255), "255\n"},
		{scsDefault, fCSSdump, "", uint8(64), "(uint8) 64\n"},
		{scsDefault, fCSSprint, "", complex(1, 2), "(1+2i)"},
		{scsDefault, fCSSprintf, "%v", complex(float32(3), 4), "(3+4i)"},
		{scsDefault, fCSSprintln, "", complex(float64(5), 6), "(5+6i)\n"},
		{scsDefault, fCSErrorf, "%#v", uint16(65535), "(uint16)65535"},
		{scsDefault, fCSNewFormatter, "%v", uint32(4294967295), "4294967295"},
		{scsDefault, fErrorf, "%v", uint64(18446744073709551615), "18446744073709551615"},
		{scsDefault, fFprint, "", float32(3.14), "3.14"},
		{scsDefault, fFprintln, "", float64(6.28), "6.28\n"},
		{scsDefault, fPrint, "", true, "true"},
		{scsDefault, fPrintln, "", false, "false\n"},
		{scsDefault, fSdump, "", complex(-10, -20), "(complex128) (-10-20i)\n"},
		{scsDefault, fSprint, "", complex(-1, -2), "(-1-2i)"},
		{scsDefault, fSprintf, "%v", complex(float32(-3), -4), "(-3-4i)"},
		{scsDefault, fSprintln, "", complex(float64(-5), -6), "(-5-6i)\n"},
		{scsNoMethods, fCSFprint, "", ts, "test"},
		{scsNoMethods, fCSFprint, "", &ts, "<*>test"},
		{scsNoMethods, fCSFprint, "", tps, "test"},
		{scsNoMethods, fCSFprint, "", &tps, "<*>test"},
		{scsNoPmethods, fCSFprint, "", ts, "stringer test"},
		{scsNoPmethods, fCSFprint, "", &ts, "<*>stringer test"},
		{scsNoPmethods, fCSFprint, "", tps, "test"},
		{scsNoPmethods, fCSFprint, "", &tps, "<*>stringer test"},
		{scsMaxDepth, fCSFprint, "", dt, "{{<max>} [<max>] [<max>] map[<max>]}"},
		{scsMaxDepth, fCSFdump, "", dt, "(peek.depthTester) {\n" +
			" ic: (peek.indirCir1) {\n  <max depth reached>\n },\n" +
			" arr: ([1]string) (len=1 cap=1) {\n  <max depth reached>\n },\n" +
			" slice: ([]string) (len=1 cap=1) {\n  <max depth reached>\n },\n" +
			" m: (map[string]int) (len=1) {\n  <max depth reached>\n }\n}\n"},
		{scsContinue, fCSFprint, "", ts, "(stringer test) test"},
		{scsContinue, fCSFdump, "", ts, "(peek.stringer) " +
			"(len=4) (stringer test) \"test\"\n"},
		{scsContinue, fCSFprint, "", te, "(error: 10) 10"},
		{scsContinue, fCSFdump, "", te, "(peek.customError) " +
			"(error: 10) 10\n"},
		{scsNoPtrAddr, fCSFprint, "", tptr, "<*>{<*>{}}"},
		{scsNoPtrAddr, fCSSdump, "", tptr, "(*peek.ptrTester)({\ns: (*struct {})({\n})\n})\n"},
		{scsNoCap, fCSSdump, "", make([]string, 0, 10), "([]string) {\n}\n"},
		{scsNoCap, fCSSdump, "", make([]string, 1, 10), "([]string) (len=1) {\n(string) \"\"\n}\n"},
	}
}

// TestPeek executes all of the tests described by peekTests.
func TestPeek(t *testing.T) {
	initPeekTests()
	t.Logf("Running %d tests", len(peekTests))
	for i, test := range peekTests {
		buf := new(bytes.Buffer)
		switch test.f {
		case fCSFdump:
			test.cs.Fdump(buf, test.in)

		case fCSFprint:
			test.cs.Fprint(buf, test.in)

		case fCSFprintf:
			test.cs.Fprintf(buf, test.format, test.in)

		case fCSFprintln:
			test.cs.Fprintln(buf, test.in)

		case fCSPrint:
			b, err := redirStdout(func() { test.cs.Print(test.in) })
			if err != nil {
				t.Errorf("%v #%d %v", test.f, i, err)
				continue
			}
			buf.Write(b)

		case fCSPrintln:
			b, err := redirStdout(func() { test.cs.Println(test.in) })
			if err != nil {
				t.Errorf("%v #%d %v", test.f, i, err)
				continue
			}
			buf.Write(b)

		case fCSSdump:
			str := test.cs.Sdump(test.in)
			buf.WriteString(str)

		case fCSSprint:
			str := test.cs.Sprint(test.in)
			buf.WriteString(str)

		case fCSSprintf:
			str := test.cs.Sprintf(test.format, test.in)
			buf.WriteString(str)

		case fCSSprintln:
			str := test.cs.Sprintln(test.in)
			buf.WriteString(str)

		case fCSErrorf:
			err := test.cs.Errorf(test.format, test.in)
			buf.WriteString(err.Error())

		case fCSNewFormatter:
			fmt.Fprintf(buf, test.format, test.cs.NewFormatter(test.in))

		case fErrorf:
			err := Errorf(test.format, test.in)
			buf.WriteString(err.Error())

		case fFprint:
			Fprint(buf, test.in)

		case fFprintln:
			Fprintln(buf, test.in)

		case fPrint:
			b, err := redirStdout(func() { Print(test.in) })
			if err != nil {
				t.Errorf("%v #%d %v", test.f, i, err)
				continue
			}
			buf.Write(b)

		case fPrintln:
			b, err := redirStdout(func() { Println(test.in) })
			if err != nil {
				t.Errorf("%v #%d %v", test.f, i, err)
				continue
			}
			buf.Write(b)

		case fSdump:
			str := Sdump(test.in)
			buf.WriteString(str)

		case fSprint:
			str := Sprint(test.in)
			buf.WriteString(str)

		case fSprintf:
			str := Sprintf(test.format, test.in)
			buf.WriteString(str)

		case fSprintln:
			str := Sprintln(test.in)
			buf.WriteString(str)

		default:
			t.Errorf("%v #%d unrecognized function", test.f, i)
			continue
		}
		s := buf.String()
		if test.want != s {
			t.Errorf("ConfigState #%d\n got: %s want: %s", i, s, test.want)
			continue
		}
	}
}
