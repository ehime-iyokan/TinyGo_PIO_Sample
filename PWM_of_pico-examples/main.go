package main

import (
	"machine"
	"time"

	pio "github.com/tinygo-org/pio/rp2-pio"
)

// Write `period` to the input shift register
func pioPwmSetPeriod(sm pio.StateMachine, period uint32) {
	asm := pio.AssemblerV0{}

	sm.SetEnabled(false)
	smPutBlocking(sm, period)
	sm.Exec(asm.Pull(false, false).Encode())
	sm.Exec(asm.Out(pio.OutDestISR, 32).Encode())
	sm.SetEnabled(true)
}

// Write `level` to TX FIFO. State machine will copy this into X.
func pioPwmSetLevel(sm pio.StateMachine, level uint32) {
	smPutBlocking(sm, level)
}

// TxPut with blocking.
func smPutBlocking(sm pio.StateMachine, data uint32) {
	for sm.IsTxFIFOFull() {
	}
	sm.TxPut(data)
}

func waitSerial() {
	time.Sleep(1 * time.Second)
}

func main() {
	waitSerial()

	Pio := pio.PIO0
	sm := Pio.StateMachine(0)
	offset, err := Pio.AddProgram(pwmInstructions, pwmOrigin)
	if err != nil {
		panic(err.Error())
	}
	println("Loaded program at", offset)

	pwmProgramInit(sm, offset, machine.LED)
	pioPwmSetPeriod(sm, (1<<16)-1)

	level := uint32(0)
	for {
		println("Level =", level)
		pioPwmSetLevel(sm, level*level)
		level = (level + 1) % 256
		time.Sleep(10 * time.Millisecond)
	}
}
