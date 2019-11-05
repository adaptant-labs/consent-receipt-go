package api

type DataController struct {
	ControllerName		string `json:"piiController"`
	OnBehalf			bool `json:"on_behalf"`
	ControllerUrl		string `json:"piiControllerUrl"`
}

func NewDataController() *DataController {
	return &DataController{
		ControllerName: "Data Controller, Inc.",
		OnBehalf:       false,
		ControllerUrl:  "https://wwww.example.com",
	}
}
