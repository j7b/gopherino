package pin

import "machine"

func init() {
	machine.InitADC()
	machine.InitPWM()
}

// Digital is a GPIO pin.
type Digital byte

// Set sets d high or low.
func (d Digital) Set(hi bool) {
	if hi {
		d.High()
	} else {
		d.Low()
	}
}

// High sets d to source.
func (d Digital) High() {
	pin := machine.GPIO{byte(d)}
	pin.High()
}

// Low sets d to sink.
func (d Digital) Low() {
	pin := machine.GPIO{byte(d)}
	pin.Low()
}

// Get returns high or low input.
func (d Digital) Get() bool {
	pin := machine.GPIO{byte(d)}
	return pin.Get()
}

// Output sets d to output.
func (d Digital) Output() {
	pin := machine.GPIO{byte(d)}
	pin.Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})
}

// Input set d to input. If pullup is true,
// pulls d high.
func (d Digital) Input(pullup bool) {
	pin := machine.GPIO{byte(d)}
	pin.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT})
	if pullup {
		d.High()
	}
}

// GPIO pins.
const (
	D0 Digital = iota
	D1
	D2
	D3
	D4
	D5
	D6
	D7
	D8
	D9
	D10
	D11
	D12
	D13
)

// Analog is an ADC pin.
type Analog byte

// Get returns analog to digital conversion of a.
func (a Analog) Get() uint16 {
	adc := machine.ADC{byte(a)}
	return adc.Get()
}

// ADC pins.
const (
	A0 Analog = 14
	A1
	A2
	A3
	A4
	A5
)

// PWM is a pin capable of pulse width modulation.
type PWM byte

// Set sets PWM output. Undefined behavior
// may result if dutycycle > 1023.
func (p PWM) Set(dutycycle uint16) {
	pwm := machine.PWM{byte(p)}
	pwm.Set(dutycycle)
}

// PWM output pins.
const (
	PWM3  = 3
	PWM5  = 5
	PWM6  = 6
	PWM10 = 10
	PWM11 = 11
)
