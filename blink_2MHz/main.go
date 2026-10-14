package main

import (
	"machine"

	pio "github.com/tinygo-org/pio/rp2-pio"
)

func main() {
	Pio := pio.PIO0
	machine.LED.Configure(machine.PinConfig{Mode: Pio.PinMode()})

	offset, err := Pio.AddProgram(blinkInstructions, blinkOrigin)
	if err != nil {
		panic(err.Error())
	}

	sm := Pio.StateMachine(0)
	sm.SetPindirsConsecutive(machine.LED, 1, true)

	cfg := blinkProgramDefaultConfig(offset)
	cfg.SetSetPins(machine.LED, 1)
	sm.Init(offset, cfg)

	sm.SetEnabled(true)

	select {}
}
