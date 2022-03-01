package common

type ResourceStatus struct {
	Running     int `json:"running"`
	Pending     int `json:"pending"`
	Failed      int `json:"failed"`
	Succeeded   int `json:"successed"`
	Unknown     int `json:"unknown"`
	Terminating int `json:"terminating"`
}
