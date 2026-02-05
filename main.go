package main

import "fmt"

type Ligth struct {
	Name string
	On   bool
}
type Fan struct {
	Speed int
	On    bool
}
type Thermostat struct {
	Temperature int
	On          bool
}
type Device interface {
	TurnOn() string
	TurnOff() string
}
type SmartDevice interface {
	Device
	ConnectToWifi(ssid string) string
}

func main() {

	var (
		light      = Ligth{Name: "Гостиная"}
		fan        = Fan{Speed: 3}
		thermostat = &Thermostat{Temperature: 22}
	)

	// Подключаем «умные» устройства
	smartDevices := [...]SmartDevice{thermostat}
	for _, sd := range smartDevices {
		fmt.Println(sd.ConnectToWifi("HomeNetwork"))
		fmt.Println(sd.TurnOn())
	}

	// Управляем как обычными устройствами
	devices := []Device{&light, &fan, thermostat}
	ControlDevices(devices)
}
func (l *Ligth) TurnOn() string {
	l.On = true
	return "Свет включен"
}
func (l *Ligth) TurnOff() string {
	l.On = false
	return "Свет выключен в гостинице"
}
func (f *Fan) TurnOn() string {
	f.On = true
	return "Вентелятор включен"
}
func (f *Fan) TurnOff() string {
	f.On = false
	return "Вентилятор выключен"
}
func (t *Thermostat) TurnOn() string {
	t.On = true
	return "Термостат работает"
}
func (t *Thermostat) TurnOff() string {
	t.On = false
	return "Термостат выключен"
}
func (t *Thermostat) ConnectToWifi(ssid string) string {
	return fmt.Sprintf("Термостат подключен к сети: %s", ssid)
}
func ControlDevices(devices []Device) {
	for _, device := range devices {
		fmt.Println(device.TurnOn())
		fmt.Println(device.TurnOff())
	}

}

