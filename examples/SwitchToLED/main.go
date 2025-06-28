// > tinygo build -o SwitchToLED.uf2 -target=pico -size short .
//    code    data     bss |   flash     ram
//    8232     108    3176 |    8340    3284
// > tinygo flash -target=pico -size short .

package main

import ( // 必要なパッケージの読み込み
	"machine"
	"time"
)

func main() {
	sw := machine.GP20 // ボタンの接続ピンを入力モードに設定
	sw.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	led := machine.GP9 //LED の接続ピンを出力モードに設定
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		if sw.Get() { // ボタンの状態を読込み、状態を判別
			led.Low() // LED を消灯する。
		} else {
			led.High() // LED を点灯する。
		}
		time.Sleep(time.Millisecond * 100) // 0.1秒休む
	}
}
