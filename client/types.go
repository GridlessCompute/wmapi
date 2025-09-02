package client

type Pool struct {
	URL      string
	Worker   string
	Password string
}

type CommandResponse struct {
	STATUS      string `json:"STATUS"`
	When        int    `json:"When"`
	Code        int    `json:"Code"`
	Msg         any    `json:"Msg"`
	Description string `json:"Description"`
}

type ErrorResponse struct {
	STATUS      string `json:"STATUS"`
	When        int    `json:"When"`
	Code        int    `json:"Code"`
	Msg         any    `json:"Msg"`
	Description string `json:"Description"`
}

type StatusResponse struct {
	Btmineroff      string `json:"btmineroff"`
	FirmwareVersion string `json:"Firmware Version"`
	PowerMode       string `json:"power_mode"`
	PowerLimitSet   string `json:"power_limit_set"`
	HashPercent     string `json:"hash_percent"`
}

type VersionResponse struct {
	STATUS string `json:"STATUS"`
	When   int    `json:"When"`
	Code   int    `json:"Code"`
	Msg    struct {
		APIVer   string `json:"api_ver"`
		FwVer    string `json:"fw_ver"`
		Platform string `json:"platform"`
		Chip     string `json:"chip"`
	} `json:"Msg"`
	Description string `json:"Description"`
}

type PSUResponse struct {
	STATUS string `json:"STATUS"`
	When   int    `json:"When"`
	Code   int    `json:"Code"`
	Msg    struct {
		Name      string `json:"name"`
		HwVersion string `json:"hw_version"`
		SwVersion string `json:"sw_version"`
		Model     string `json:"model"`
		Iin       string `json:"iin"`
		Vin       string `json:"vin"`
		Pin       string `json:"pin"`
		FanSpeed  string `json:"fan_speed"`
		Version   string `json:"version"`
		SerialNo  string `json:"serial_no"`
		Vendor    string `json:"vendor"`
		Temp0     string `json:"temp0"`
	} `json:"Msg"`
	Description string `json:"Description"`
}

type DevdetailsResponse struct {
	STATUS []struct {
		STATUS      string `json:"STATUS"`
		When        int    `json:"When"`
		Code        int    `json:"Code"`
		Msg         string `json:"Msg"`
		Description string `json:"Description"`
	} `json:"STATUS"`
	DEVDETAILS []struct {
		DEVDETAILS int    `json:"DEVDETAILS"`
		Name       string `json:"Name"`
		ID         int    `json:"ID"`
		Driver     string `json:"Driver"`
		Kernel     string `json:"Kernel"`
		Model      string `json:"Model"`
	} `json:"DEVDETAILS"`
}

type EdevsResponse struct {
	STATUS []struct {
		STATUS string `json:"STATUS"`
		Msg    string `json:"Msg"`
	} `json:"STATUS"`
	DEVS []struct {
		ASC            int     `json:"ASC"`
		Slot           int     `json:"Slot"`
		Enabled        string  `json:"Enabled"`
		Status         string  `json:"Status"`
		Temperature    int     `json:"Temperature"`
		ChipFrequency  int     `json:"Chip Frequency"`
		MHSAv          float64 `json:"MHS av"`
		MHS5S          float64 `json:"MHS 5s"`
		MHS1M          float64 `json:"MHS 1m"`
		MHS5M          float64 `json:"MHS 5m"`
		MHS15M         float64 `json:"MHS 15m"`
		HSRT           float64 `json:"HS RT"`
		HSFactory      int     `json:"HS Factory,omitempty"`
		Accepted       int     `json:"Accepted"`
		Rejected       int     `json:"Rejected"`
		LastValidWork  int     `json:"Last Valid Work"`
		UpfreqComplete int     `json:"Upfreq Complete"`
		EffectiveChips int     `json:"Effective Chips"`
		PCBSN          string  `json:"PCB SN"`
		ChipData       string  `json:"Chip Data"`
		ChipTempMin    float64 `json:"Chip Temp Min"`
		ChipTempMax    int     `json:"Chip Temp Max"`
		ChipTempAvg    float64 `json:"Chip Temp Avg"`
		ChipVolDiff    int     `json:"chip_vol_diff"`
	} `json:"DEVS"`
}

type PoolsResponse struct {
	STATUS []struct {
		STATUS string `json:"STATUS"`
		Msg    string `json:"Msg"`
	} `json:"STATUS"`
	POOLS []struct {
		POOL                int    `json:"POOL"`
		URL                 string `json:"URL"`
		Status              string `json:"Status"`
		Priority            int    `json:"Priority"`
		Quota               int    `json:"Quota"`
		LongPoll            string `json:"Long Poll"`
		Getworks            int    `json:"Getworks"`
		Accepted            int    `json:"Accepted"`
		Rejected            int    `json:"Rejected"`
		Works               int    `json:"Works"`
		Discarded           int    `json:"Discarded"`
		Stale               int    `json:"Stale"`
		GetFailures         int    `json:"Get Failures"`
		RemoteFailures      int    `json:"Remote Failures"`
		User                string `json:"User"`
		LastShareTime       int    `json:"Last Share Time"`
		Diff1Shares         int    `json:"Diff1 Shares"`
		ProxyType           string `json:"Proxy Type"`
		Proxy               string `json:"Proxy"`
		DifficultyAccepted  int    `json:"Difficulty Accepted"`
		DifficultyRejected  int    `json:"Difficulty Rejected"`
		DifficultyStale     int    `json:"Difficulty Stale"`
		LastShareDifficulty int    `json:"Last Share Difficulty"`
		WorkDifficulty      int    `json:"Work Difficulty"`
		HasStratum          int    `json:"Has Stratum"`
		StratumActive       bool   `json:"Stratum Active"`
		StratumURL          string `json:"Stratum URL"`
		StratumDifficulty   int    `json:"Stratum Difficulty"`
		BestShare           int    `json:"Best Share"`
		PoolRejected        int    `json:"Pool Rejected%"`
		PoolStale           int    `json:"Pool Stale%"`
		BadWork             int    `json:"Bad Work"`
		CurrentBlockHeight  int    `json:"Current Block Height"`
		CurrentBlockVersion int    `json:"Current Block Version"`
	} `json:"POOLS"`
}

type SummaryResponse struct {
	STATUS []struct {
		STATUS string `json:"STATUS"`
		Msg    string `json:"Msg"`
	} `json:"STATUS"`
	SUMMARY []struct {
		Elapsed               int     `json:"Elapsed"`
		MHSAv                 float64 `json:"MHS av"`
		MHS5S                 float64 `json:"MHS 5s"`
		MHS1M                 float64 `json:"MHS 1m"`
		MHS5M                 float64 `json:"MHS 5m"`
		MHS15M                float64 `json:"MHS 15m"`
		HSRT                  float64 `json:"HS RT"`
		Accepted              int     `json:"Accepted"`
		Rejected              int     `json:"Rejected"`
		TotalMH               int64   `json:"Total MH"`
		Temperature           int     `json:"Temperature"`
		FreqAvg               int     `json:"freq_avg"`
		FanSpeedIn            int     `json:"Fan Speed In"`
		FanSpeedOut           int     `json:"Fan Speed Out"`
		Power                 int     `json:"Power"`
		PowerRate             float64 `json:"Power Rate"`
		PoolRejected          int     `json:"Pool Rejected%"`
		PoolStale             int     `json:"Pool Stale%"`
		LastGetwork           int     `json:"Last getwork"`
		Uptime                int     `json:"Uptime"`
		SecurityMode          int     `json:"Security Mode"`
		HashStable            bool    `json:"Hash Stable"`
		HashStableCostSeconds int     `json:"Hash Stable Cost Seconds"`
		HashDeviation         float64 `json:"Hash Deviation%"`
		TargetFreq            int     `json:"Target Freq"`
		TargetMHS             int     `json:"Target MHS"`
		EnvTemp               int     `json:"Env Temp"`
		PowerMode             string  `json:"Power Mode"`
		FactoryGHS            int     `json:"Factory GHS"`
		PowerLimit            int     `json:"Power Limit"`
		ChipTempMin           float64 `json:"Chip Temp Min"`
		ChipTempMax           float64 `json:"Chip Temp Max"`
		ChipTempAvg           float64 `json:"Chip Temp Avg"`
		Debug                 string  `json:"Debug"`
		BtminerFastBoot       string  `json:"Btminer Fast Boot"`
	} `json:"SUMMARY"`
}