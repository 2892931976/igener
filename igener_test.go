package igener

import (
	"fmt"
	"hash/crc32"
	"os"
	"testing"
	"time"
)

func BenchmarkIGener(b *testing.B) {
	ig := NewIGener()
	for n := 0; n < b.N; n++ {
		<-ig
	}
}

func TestIGener(t *testing.T) {
	ig := NewIGener()
	var id string
	m := make(map[string]struct{})
	for i := 0; i < 10000; i++ {
		id = <-ig
		m[id] = struct{}{}
	}
	if len(m) != 10000 {
		t.Error("TestIGener Error")
	}
	t.Log("Test IGener PASS")
}

func TestMachinePidEncode(t *testing.T) {
	hostname, _ := os.Hostname()
	hashCode := crc32.ChecksumIEEE([]byte(hostname))
	machineCode := fmt.Sprintf("%06x", hashCode)
	pid := os.Getpid()
	pidCode := fmt.Sprintf("%04x", pid)

	ig := &IGener{
		second: time.Now().Unix(),
		inc:    0,
		idChan: make(chan string),
	}

	machineCodeFromIg, pidCodeFromIg := ig.machinePidEncode()

	machineCodeLen := len(machineCode)
	if machineCode[machineCodeLen-6:machineCodeLen] != fmt.Sprintf("%x", machineCodeFromIg) {
		t.Error("Test MachineCode Error")
	}

	pidCodeLen := len(pidCode)
	if pidCode[pidCodeLen-4:pidCodeLen] != fmt.Sprintf("%x", pidCodeFromIg) {
		t.Error("Test PidCode Error")
	}

	t.Log("Test MachineCode PidCode PASS")
}
