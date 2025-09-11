package client

type Pool struct {
	URL      string
	Worker   string
	Password string
}

type CommandResponse struct {
	STATUS      string  `json:"STATUS"`
	When        float64 `json:"When"`
	Code        float64 `json:"Code"`
	Msg         any     `json:"Msg"`
	Description string  `json:"Description"`
}

type ErrorResponse struct {
	STATUS      string  `json:"STATUS"`
	When        float64 `json:"When"`
	Code        float64 `json:"Code"`
	Msg         any     `json:"Msg"`
	Description string  `json:"Description"`
}

type StatusResponse struct {
	Btmineroff      string `json:"btmineroff"`
	FirmwareVersion string `json:"Firmware Version"`
	PowerMode       string `json:"power_mode"`
	PowerLimitSet   string `json:"power_limit_set"`
	HashPercent     string `json:"hash_percent"`
}

type VersionResponse struct {
	STATUS string  `json:"STATUS"`
	When   float64 `json:"When"`
	Code   float64 `json:"Code"`
	Msg    struct {
		APIVer   string `json:"api_ver"`
		FwVer    string `json:"fw_ver"`
		Platform string `json:"platform"`
		Chip     string `json:"chip"`
	} `json:"Msg"`
	Description string `json:"Description"`
}

type PSUResponse struct {
	STATUS string  `json:"STATUS"`
	When   float64 `json:"When"`
	Code   float64 `json:"Code"`
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
	STATUS     ResponseStatus `json:"STATUS"`
	DEVDETAILS []struct {
		DEVDETAILS float64 `json:"DEVDETAILS"`
		Name       string  `json:"Name"`
		ID         float64 `json:"ID"`
		Driver     string  `json:"Driver"`
		Kernel     string  `json:"Kernel"`
		Model      string  `json:"Model"`
	} `json:"DEVDETAILS"`
}

type EdevsResponse struct {
	STATUS ResponseStatus `json:"STATUS"`
	DEVS   []struct {
		ASC            float64 `json:"ASC"`
		Slot           float64 `json:"Slot"`
		Enabled        string  `json:"Enabled"`
		Status         string  `json:"Status"`
		Temperature    float64 `json:"Temperature"`
		ChipFrequency  float64 `json:"Chip Frequency"`
		MHSAv          float64 `json:"MHS av"`
		MHS5S          float64 `json:"MHS 5s"`
		MHS1M          float64 `json:"MHS 1m"`
		MHS5M          float64 `json:"MHS 5m"`
		MHS15M         float64 `json:"MHS 15m"`
		HSRT           float64 `json:"HS RT"`
		HSFactory      float64 `json:"HS Factory,omitempty"`
		Accepted       float64 `json:"Accepted"`
		Rejected       float64 `json:"Rejected"`
		LastValidWork  float64 `json:"Last Valid Work"`
		UpfreqComplete float64 `json:"Upfreq Complete"`
		EffectiveChips float64 `json:"Effective Chips"`
		PCBSN          string  `json:"PCB SN"`
		ChipData       string  `json:"Chip Data"`
		ChipTempMin    float64 `json:"Chip Temp Min"`
		ChipTempMax    float64 `json:"Chip Temp Max"`
		ChipTempAvg    float64 `json:"Chip Temp Avg"`
		ChipVolDiff    float64 `json:"chip_vol_diff"`
	} `json:"DEVS"`
}

type MinerInfoResponse struct {
	STATUS string  `json:"STATUS"`
	When   float64 `json:"When"`
	Code   float64 `json:"Code"`
	Msg    struct {
		IP      string `json:"ip"`
		Proto   string `json:"proto"`
		Netmask string `json:"netmask"`
		DNS     string `json:"dns"`
		Mac     string `json:"mac"`
		Ledstat string `json:"ledstat"`
		Gateway string `json:"gateway"`
	} `json:"Msg"`
	Description string `json:"Description"`
}

type ResponseStatus []struct {
	// STATUS      string  `json:"STATUS"`
	When        float64 `json:"When"`
	Code        float64 `json:"Code"`
	Msg         string  `json:"Msg"`
	Description string  `json:"Description"`
}

type PoolsResponse struct {
	STATUS ResponseStatus `json:"STATUS"`
	POOLS  []struct {
		POOL                float64 `json:"POOL"`
		URL                 string  `json:"URL"`
		Status              string  `json:"Status"`
		Priority            float64 `json:"Priority"`
		Quota               float64 `json:"Quota"`
		LongPoll            string  `json:"Long Poll"`
		Getworks            float64 `json:"Getworks"`
		Accepted            float64 `json:"Accepted"`
		Rejected            float64 `json:"Rejected"`
		Works               float64 `json:"Works"`
		Discarded           float64 `json:"Discarded"`
		Stale               float64 `json:"Stale"`
		GetFailures         float64 `json:"Get Failures"`
		RemoteFailures      float64 `json:"Remote Failures"`
		User                string  `json:"User"`
		LastShareTime       float64 `json:"Last Share Time"`
		Diff1Shares         float64 `json:"Diff1 Shares"`
		ProxyType           string  `json:"Proxy Type"`
		Proxy               string  `json:"Proxy"`
		DifficultyAccepted  float64 `json:"Difficulty Accepted"`
		DifficultyRejected  float64 `json:"Difficulty Rejected"`
		DifficultyStale     float64 `json:"Difficulty Stale"`
		LastShareDifficulty float64 `json:"Last Share Difficulty"`
		WorkDifficulty      float64 `json:"Work Difficulty"`
		HasStratum          float64 `json:"Has Stratum"`
		StratumActive       bool    `json:"Stratum Active"`
		StratumURL          string  `json:"Stratum URL"`
		StratumDifficulty   float64 `json:"Stratum Difficulty"`
		BestShare           float64 `json:"Best Share"`
		PoolRejected        float64 `json:"Pool Rejected%"`
		PoolStale           float64 `json:"Pool Stale%"`
		BadWork             float64 `json:"Bad Work"`
		CurrentBlockHeight  float64 `json:"Current Block Height"`
		CurrentBlockVersion float64 `json:"Current Block Version"`
	} `json:"POOLS"`
}

type SummaryResponse struct {
	STATUS  ResponseStatus `json:"STATUS"`
	SUMMARY []struct {
		Elapsed               float64 `json:"Elapsed"`
		MHSAv                 float64 `json:"MHS av"`
		MHS5S                 float64 `json:"MHS 5s"`
		MHS1M                 float64 `json:"MHS 1m"`
		MHS5M                 float64 `json:"MHS 5m"`
		MHS15M                float64 `json:"MHS 15m"`
		HSRT                  float64 `json:"HS RT"`
		Accepted              float64 `json:"Accepted"`
		Rejected              float64 `json:"Rejected"`
		TotalMH               float64 `json:"Total MH"`
		Temperature           float64 `json:"Temperature"`
		FreqAvg               float64 `json:"freq_avg"`
		FanSpeedIn            float64 `json:"Fan Speed In"`
		FanSpeedOut           float64 `json:"Fan Speed Out"`
		Power                 float64 `json:"Power"`
		PowerRate             float64 `json:"Power Rate"`
		PoolRejected          float64 `json:"Pool Rejected%"`
		PoolStale             float64 `json:"Pool Stale%"`
		LastGetwork           float64 `json:"Last getwork"`
		Uptime                float64 `json:"Uptime"`
		SecurityMode          float64 `json:"Security Mode"`
		HashStable            bool    `json:"Hash Stable"`
		HashStableCostSeconds float64 `json:"Hash Stable Cost Seconds"`
		HashDeviation         float64 `json:"Hash Deviation%"`
		TargetFreq            float64 `json:"Target Freq"`
		TargetMHS             float64 `json:"Target MHS"`
		EnvTemp               float64 `json:"Env Temp"`
		PowerMode             string  `json:"Power Mode"`
		FactoryGHS            float64 `json:"Factory GHS"`
		PowerLimit            float64 `json:"Power Limit"`
		ChipTempMin           float64 `json:"Chip Temp Min"`
		ChipTempMax           float64 `json:"Chip Temp Max"`
		ChipTempAvg           float64 `json:"Chip Temp Avg"`
		Debug                 string  `json:"Debug"`
		BtminerFastBoot       string  `json:"Btminer Fast Boot"`
	} `json:"SUMMARY"`
}
