// > tinygo build -o SwitchToConsole.uf2 -target=pico -size short -monitor .
//    code    data     bss |   flash     ram
//    8232     108    3176 |    8340    3284
// > tinygo flash -target=pico -size short -monitor .

package main

import ( // 必要なパッケージの読み込み
	"fmt"
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
			fmt.Printf("0, \n")
			led.Low() // LED を消灯する。
		} else {
			fmt.Printf("1, ###########################\n")
			led.High() // LED を点灯する。
		}
		time.Sleep(time.Millisecond * 100)
	}
}
