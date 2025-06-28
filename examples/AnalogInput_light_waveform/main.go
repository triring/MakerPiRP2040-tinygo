// > tinygo build -o AnalogInput_light_waveform.uf2 -target=pico -size short --monitor .
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
	for {
		sensorValue = light_sensor.Get() // ADC値を読み取る
		// 読み取った値を表示
		fmt.Printf("%v,%6d, \"", time.Now().Format(time.DateTime), sensorValue)
		n := int(sensorValue / 200)
		for i := 0; i < n; i++ {
			fmt.Printf("#")
		}
		fmt.Printf("\"\n")
		time.Sleep(time.Millisecond * 125)
	}
}
