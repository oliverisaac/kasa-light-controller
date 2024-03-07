package messages

type Discovery struct {
	Countdown                Countdown                `json:"countdown,omitempty"`
	SmartlifeIotCommonCloud  SmartlifeIotCommonCloud  `json:"smartlife.iot.common.cloud,omitempty"`
	SmartlifeIotCommonEmeter SmartlifeIotCommonEmeter `json:"smartlife.iot.common.emeter,omitempty"`
}
type Countdown struct {
	ErrCode *int    `json:"err_code,omitempty"`
	ErrMsg  *string `json:"err_msg,omitempty"`
}
type GetInfo struct {
	Binded        *int    `json:"binded,omitempty"`
	CldConnection *int    `json:"cld_connection,omitempty"`
	ErrCode       *int    `json:"err_code,omitempty"`
	FwDlPage      *string `json:"fwDlPage,omitempty"`
	FwNotifyType  *int    `json:"fwNotifyType,omitempty"`
	IllegalType   *int    `json:"illegalType,omitempty"`
	Server        *string `json:"server,omitempty"`
	StopConnect   *int    `json:"stopConnect,omitempty"`
	TcspInfo      *string `json:"tcspInfo,omitempty"`
	TcspStatus    *int    `json:"tcspStatus,omitempty"`
	Username      *string `json:"username,omitempty"`
}

type DayList struct {
	Day      *int `json:"day,omitempty"`
	EnergyWh *int `json:"energy_wh,omitempty"`
	Month    *int `json:"month,omitempty"`
	Year     *int `json:"year,omitempty"`
}
type GetDaystat struct {
	DayList []DayList `json:"day_list,omitempty"`
	ErrCode *int      `json:"err_code,omitempty"`
}
type MonthList struct {
	EnergyWh *int `json:"energy_wh,omitempty"`
	Month    *int `json:"month,omitempty"`
	Year     *int `json:"year,omitempty"`
}
type GetMonthstat struct {
	ErrCode   *int        `json:"err_code,omitempty"`
	MonthList []MonthList `json:"month_list,omitempty"`
}
type GetRealtime struct {
	ErrCode *int `json:"err_code,omitempty"`
	PowerMw *int `json:"power_mw,omitempty"`
	TotalWh *int `json:"total_wh,omitempty"`
}
type SmartlifeIotCommonEmeter struct {
	GetDaystat   GetDaystat   `json:"get_daystat,omitempty"`
	GetMonthstat GetMonthstat `json:"get_monthstat,omitempty"`
	GetRealtime  GetRealtime  `json:"get_realtime,omitempty"`
}

func (mg MessageGenerator) GetInfo() []byte {
	message := `{"system":{"get_sysinfo":null}}`
	msg := mg.Encrypter.EncryptBytes([]byte(message))
	return msg[4:]
}
