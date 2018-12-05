package core

import (
	"device/avr"
	"machine"
)

const clockCyclesPerMicrosecond = machine.CPU_FREQUENCY / 1000000
const microsecondsPerOverflow = 64 * 256 / clockCyclesPerMicrosecond
const milliInc = microsecondsPerOverflow / 1000
const fractInt = microsecondsPerOverflow % 3
const fractMax = 125

//go:volatile
type timing uint32

//go:volatile
type fraction uint8

var overflowCount timing
var timerMillis timing
var timerFract fraction

//go:interrupt TIMER0_OVF_vect
func timer0() {
	overflowCount++
	m := uint32(timerMillis)
	f := uint8(timerFract)
	m += milliInc
	f += fractInt
	if f >= fractMax {
		f -= fractMax
		m++
	}
	timerMillis = timing(m)
	timerFract = fraction(f)
}

func setup() {
	avr.Asm("sei")
	machine.InitPWM()
	*avr.TIMSK0 = avr.TIMSK0_TOIE0
}

func init() {
	setup()
}

// Millis returns milliseconds since program start.
func Millis() uint32 {
	var m uint32
	sr := byte(*avr.SREG)
	avr.Asm("cli")
	m = uint32(timerMillis)
	*avr.SREG = avr.RegValue(sr)
	avr.Asm("sei")
	return m
}

// Micros returns microseconds since program start.
func Micros() uint32 {
	var m uint32
	sr := byte(*avr.SREG)
	avr.Asm("cli")
	t := uint32(*avr.TCNT0)
	m = uint32(overflowCount)
	*avr.SREG = avr.RegValue(sr)
	avr.Asm("sei")
	return (m<<8 + t) * (64 / clockCyclesPerMicrosecond)
}

// Delay busy-loops for ms milliseconds.
func Delay(ms uint32) {
	start := Millis()
	for ms > 0 {
		avr.Asm("nop")
		for start == Millis() {
			avr.Asm("nop")
		}
		ms--
		start = Millis()
	}
}

// DelayMicroseconds busy-loops for us microseconds.
func DelayMicroseconds(us uint32) {
	if us < 1 {
		return
	}
	us = us << 2
	us -= 5
	for us > 0 {
		avr.Asm("nop")
		us--
	}
}
