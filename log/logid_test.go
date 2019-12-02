package log

import (
	"testing"
)

func Test_CalcLogId_Case1(t *testing.T) {
	pid := 7899
	sec := int64(1000)
	usec := 2000
	network := "tcp"
	client_ip := "202.113.12.9"
	local_ip := "10.10.1.100"

	logid := calcLogId(network, client_ip, local_ip, pid, sec, usec)
	if logid != 2862844380434513055 {
		t.Errorf("Test_CreateLogId test failed. got %d", logid)
	}
}

func Test_CalcLogId_Case2(t *testing.T) {
	pid := 3401
	sec := int64(1234567890)
	usec := 19191
	network := "tcp"
	client_ip := "220.100.20.188"
	local_ip := "10.20.2.210"

	logid := calcLogId(network, client_ip, local_ip, pid, sec, usec)
	if logid != 5866159772736062243 {
		t.Errorf("Test_CreateLogId test failed. got %d", logid)
	}
}

func Test_CalcLogId_Case3(t *testing.T) {
	pid := 1998
	sec := int64(98765522)
	usec := 3333
	network := "tcp"
	client_ip := "211.88.1.123"
	local_ip := "10.59.23.45"

	logid := calcLogId(network, client_ip, local_ip, pid, sec, usec)
	if logid != 12555628646781081521 {
		t.Errorf("Test_CreateLogId test failed. got %d", logid)
	}
}
