package main

import (
	"github.com/tkmru/cibo/core"

	"encoding/hex"
	"log"
	"syscall/js"
)

func emulate(_ js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}

	debugFlag := true
	emu := cibo.NewEmulator(32, 0x7c00, 9, debugFlag)
	cpu := emu.CPU
	reg := &cpu.X86registers
	reg.Init()
	hexString := args[0].String()
	insn, err := hex.DecodeString(hexString)

	if err != nil {
		log.Fatal(err)
	}

	emu.RAM = insn
	emu.Run()
	ret := map[string]interface{}{
		"EAX":    reg.EAX,
		"ECX":    reg.ECX,
		"EDX":    reg.EDX,
		"EBX":    reg.EBX,
		"ESP":    reg.ESP,
		"EBP":    reg.EBP,
		"ESI":    reg.ESI,
		"EDI":    reg.EDI,
		"EIP":    reg.EIP,
		"EFLAGS": reg.EFLAGS,
	}
	return ret
}

var global = js.Global()

func main() {
	c := make(chan struct{}, 0)
	global.Set("emulate", js.FuncOf(emulate))
	<-c
}
