package messages

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	SmartlifeIotSmartbulbLightingservice SmartlifeIotSmartbulbLightingservice `json:"smartlife.iot.smartbulb.lightingservice,omitempty"`
	SmartlifeIotCommonCloud              SmartlifeIotCommonCloud              `json:"smartlife.iot.common.cloud,omitempty"`
}

type SmartlifeIotCommonCloud struct {
	GetInfo any `json:"get_info"`
}
type TransitionLightState struct {
	Hue              *int `json:"hue,omitempty"`
	Saturation       *int `json:"saturation,omitempty"`
	ColorTemp        *int `json:"color_temp,omitempty"`
	Brightness       *int `json:"brightness,omitempty"`
	TransitionPeriod *int `json:"transition_period,omitempty"`
	OnOff            *int `json:"on_off,omitempty"`
	IgnoreDefault    *int `json:"ignore_default,omitempty"`
}
type SmartlifeIotSmartbulbLightingservice struct {
	TransitionLightState TransitionLightState `json:"transition_light_state,omitempty"`
}

type Encrypter interface {
	EncryptBytes([]byte) []byte
}

type MessageGenerator struct {
	Encrypter Encrypter
}

func (mg MessageGenerator) fromInterface(i any) []byte {
	b, err := json.Marshal(i)
	if err != nil {
		panic(fmt.Errorf("Failed to json encode message: %w", err))
	}
	return mg.Encrypter.EncryptBytes(b)
}

func (mg MessageGenerator) Off() []byte {
	return mg.fromInterface(Message{
		SmartlifeIotSmartbulbLightingservice: SmartlifeIotSmartbulbLightingservice{
			TransitionLightState{
				OnOff:         IntPtr(0),
				IgnoreDefault: IntPtr(1),
			}}})
}

func (mg MessageGenerator) On() []byte {
	return mg.fromInterface(Message{
		SmartlifeIotSmartbulbLightingservice: SmartlifeIotSmartbulbLightingservice{
			TransitionLightState{
				OnOff:         IntPtr(1),
				IgnoreDefault: IntPtr(1),
			}}})
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
	return mg.fromInterface(Message{
		SmartlifeIotSmartbulbLightingservice: SmartlifeIotSmartbulbLightingservice{
			TransitionLightState{
				OnOff:            IntPtr(1),
				IgnoreDefault:    IntPtr(1),
				Hue:              IntPtr(hue),
				Saturation:       IntPtr(saturation),
				ColorTemp:        IntPtr(0),
				Brightness:       IntPtr(value),
				TransitionPeriod: IntPtr(100),
			}}})
}
