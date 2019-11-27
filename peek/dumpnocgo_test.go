// +build !cgo !testcgo

package peek

func addCgoDumpTests() {
	// Don't add any tests for cgo since this file is only compiled when
	// there should not be any cgo tests.
}
