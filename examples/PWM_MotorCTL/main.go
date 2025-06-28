// > tinygo build -o PWM_MotorCTL.uf2 -target=pico -size short .
//   code    data     bss |   flash     ram
//  10120     108    5240 |   10228    5348
//
// > tinygo flash -target=pico -size short -monitor .

package main

import (
	"machine"
	"time"
)

// 1パルスを送り、モータを回す。単位は、ミリ秒で設定
func one_pulse(motorA machine.Pin, motorB machine.Pin, on_time time.Duration, off_time time.Duration) {
	motorA.High()
	motorB.Low()
	time.Sleep(on_time * time.Millisecond)
	motorA.Low()
	motorB.Low()
	time.Sleep(off_time * time.Millisecond)
}

func main() {
	Motor1A := machine.GP8 // M1A : GP8 モータ制御ポートA
	Motor1A.Configure(machine.PinConfig{Mode: machine.PinOutput})
	Motor1B := machine.GP9 // M1A : GP9 モータ制御ポートB
	Motor1B.Configure(machine.PinConfig{Mode: machine.PinOutput})
	Motor1B.Low()

	for i := 0; i < 6; i++ {
		for i := 0; i < 25; i++ {
			one_pulse(Motor1A, Motor1B, 10, 90)
		}
		time.Sleep(1000 * time.Millisecond)
		for i := 0; i < 25; i++ {
			one_pulse(Motor1A, Motor1B, 30, 70)
		}
		time.Sleep(1000 * time.Millisecond)
		for i := 0; i < 25; i++ {
			one_pulse(Motor1A, Motor1B, 50, 50)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
