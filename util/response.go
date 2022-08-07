package util

type Response struct {
	//	status        int
	statusMessage string
	timestamp     string
	data          interface{}
}

/*
func (res Response) GetStatus() int {
	return res.status
}

func (res *Response) SetStatus(status int) {
	res.status = status
}
*/

func (res *Response) SetTimestamp(timestamp string) {
	res.timestamp = timestamp
}

func (res Response) GetStatusMessage() string {
	return res.statusMessage
}

func (res *Response) SetStatusMessage(statusMessage string) {
	res.statusMessage = statusMessage
}

func (res Response) GetData() interface{} {
	return res.data
}

func (res Response) SetData(data interface{}) {
	res.data = data
}

func GenerateResponse(res *Response, err error) {
	res.SetTimestamp(GetTimestamp())

	if err != nil {
		res.SetStatusMessage(err.Error())
	} else {
		res.SetStatusMessage(SuccessMessage)
	}
}
