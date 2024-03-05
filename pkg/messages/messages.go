package messages

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	SmartlifeIotSmartbulbLightingservice SmartlifeIotSmartbulbLightingservice `json:"smartlife.iot.smartbulb.lightingservice,omitempty"`
}
type TransitionLightState struct {
	Hue              int `json:"hue,omitempty"`
	Saturation       int `json:"saturation,omitempty"`
	ColorTemp        int `json:"color_temp,omitempty"`
	Brightness       int `json:"brightness,omitempty"`
	TransitionPeriod int `json:"transition_period"`
	OnOff            int `json:"on_off"`
	IgnoreDefault    int `json:"ignore_default"`
}
type SmartlifeIotSmartbulbLightingservice struct {
	TransitionLightState TransitionLightState `json:"transition_light_state,omitempty"`
}

type Encrypter interface {
	EncryptBytes([]byte) ([]byte, error)
}

type MessageGenerator struct {
	Encrypter Encrypter
}

func (mg MessageGenerator) fromInterface(i any) ([]byte, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("Failed to json encode message: %w", err)
	}
	return mg.Encrypter.EncryptBytes(b)
}

func (mg MessageGenerator) Off() []byte {
	b, err := mg.fromInterface(Message{
		SmartlifeIotSmartbulbLightingservice{
			TransitionLightState{
				OnOff:         0,
				IgnoreDefault: 1,
			}}})
	if err != nil {
		panic(err)
	}
	return b
}

func (mg MessageGenerator) On() []byte {
	b, err := mg.fromInterface(Message{
		SmartlifeIotSmartbulbLightingservice{
			TransitionLightState{
				OnOff:         1,
				IgnoreDefault: 1,
			}}})
	if err != nil {
		panic(err)
	}
	return b
}

func (mg MessageGenerator) HSV(hue, saturation, value int) []byte {
	if hue > 360 || hue < 0 {
		panic("Hue must be between 0 and 360")
	}
	if saturation > 100 || saturation < 0 {
		panic("Saturation must be between 0 and 100")
	}
	if value > 100 || value < 0 {
		panic("Value must be between 0 and 100")
	}
	b, err := mg.fromInterface(Message{
		SmartlifeIotSmartbulbLightingservice{
			TransitionLightState{
				OnOff:            1,
				IgnoreDefault:    1,
				Hue:              hue,
				Saturation:       saturation,
				ColorTemp:        0,
				Brightness:       value,
				TransitionPeriod: 100,
			}}})
	if err != nil {
		panic(err)
	}
	return b
}
