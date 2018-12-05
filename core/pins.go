package core

import "machine"

const (
	OUTPUT byte = iota
	INPUT
	INPUT_PULLUP
)

// DigitalWrite sets pin high or low.
func DigitalWrite(pin byte, val bool) {
	gpio := machine.GPIO{pin}
	switch val {
	case true:
		gpio.High()
	case false:
		gpio.Low()
	}
}

// PinMode sets pin to mode. If mode > INPUT_PULLUP,
// no-op.
func PinMode(pin, mode byte) {
	gpio := machine.GPIO{pin}
	switch mode {
	case OUTPUT:
		gpio.Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})
	case INPUT:
		gpio.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT})
	case INPUT_PULLUP:
		gpio.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT})
		gpio.High()
	}
}
