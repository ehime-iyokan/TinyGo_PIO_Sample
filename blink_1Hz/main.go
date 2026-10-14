package main

import (
	"machine"

	pio "github.com/tinygo-org/pio/rp2-pio"
)

func main() {
	Pio := pio.PIO0

	offset, err := Pio.AddProgram(blinkInstructions, blinkOrigin)
	if err != nil {
		panic(err.Error())
	}

	sm := Pio.StateMachine(0)
	machine.LED.Configure(machine.PinConfig{Mode: Pio.PinMode()})
	sm.SetPindirsConsecutive(machine.LED, 1, true)

	cfg := blinkProgramDefaultConfig(offset)
	cfg.SetSetPins(machine.LED, 1)
	sm.Init(offset, cfg)

	sm.TxPut(machine.CPUFrequency())

	sm.SetEnabled(true)

	select {}
}
