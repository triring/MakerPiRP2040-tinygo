// > tinygo build -o AnalogInput_light.uf2 -target=pico -size short --monitor .
//    code    data     bss |   flash     ram
//    8232     108    3176 |    8340    3284
// > tinygo flash -target=pico -size short -monitor .

package main

import ( // 必要なパッケージの読み込み
	"fmt"
	"machine"
	"time"
)

/*
Analog Input
ADCの先に以下のようにLEDを直結。

GP26	アノード	足の長い方
GND 	カソード	足の短い方

LEDに光が当てると発生する微弱な電圧を測定する。
*/
func main() {
	var sensorValue uint16 = 0
	// LEDをのGPIO26/GNDに接続,ADC0に設定
	machine.InitADC() // ADCの初期化
	light_sensor := machine.ADC{Pin: machine.ADC0}
	light_sensor.Configure(machine.ADCConfig{})
	sw := machine.GP20 // ボタンの接続ピンを入力モードに設定
	sw.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	for {
		if !sw.Get() { // ボタンの状態を読込み、状態を判別
			sensorValue = light_sensor.Get() // ADC値を読み取る
			// 読み取った値を表示
			fmt.Printf("%v,%6d\n", time.Now().Format(time.DateTime), sensorValue)
		}
		time.Sleep(time.Millisecond * 125)
	}
}
