// > tinygo build -o hand-cranked_Motor.uf2 -target=pico -size short .
//
//	 code    data     bss |   flash     ram
//	16640     180    3184 |   16820    3364
//
// > tinygo flash -target=pico -size short .
package main

import ( // 必要なパッケージの読み込み
	"fmt"
	"machine"
	"time"
)

func main() {
	fmt.Printf("hand-cranked_Motor\n")

	switchPin := machine.GP20 // ボード上のスイッチポート
	Motor1APin := machine.GP8 // M1A : GP8 モータ制御ポートA
	Motor1BPin := machine.GP9 // M1B : GP9 モータ制御ポートB

	// 入出力の設定
	switchPin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	Motor1APin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	Motor1BPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for { // 無限ループ
		// スイッチはが Pull-Up されていて離している時は High
		if !switchPin.Get() {
			Motor1APin.Low() // モータを回す
			Motor1BPin.High()
		}
		time.Sleep(time.Millisecond * 25)
		Motor1APin.Low() // モータを止める
		Motor1BPin.Low()
		time.Sleep(time.Millisecond * 25)
	}
}
