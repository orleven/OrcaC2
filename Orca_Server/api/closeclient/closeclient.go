package closeclient

import (
	"Orca_Server/api"
	"Orca_Server/define/retcode"
	"Orca_Server/servers"
	"encoding/json"
	"net/http"
)

type Controller struct {
}

type inputData struct {
	ClientId string `json:"clientId" validate:"required"`
}

func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	var inputData inputData
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := api.Validate(inputData)
	if err != nil {
		api.Render(w, retcode.FAIL, err.Error(), []string{})
		return
	}

	systemId := r.Header.Get("SystemId")

	//发送信息
	servers.CloseClient(inputData.ClientId, systemId)

	api.Render(w, retcode.SUCCESS, "success", map[string]string{})
	return
}
