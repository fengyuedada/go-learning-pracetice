package panic

import (
	"awesomeProject/env"
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

const url = "http://localhost:8888/report-panic"

// 为了避免造成 panic report 服务被打挂，降低发送 http 请求频率，进程生命周期内只发一次
var panicReportOnce sync.Once

type PanicReq struct {
	Service   string `json:"service"`
	ErrorInfo string `json:"error_info"`
	Stack     string `json:"stack"`
	LogId     string `json:"log_id"`
	FuncName  string `json:"func_name"`
	Host      string `json:"host"`
	PodName   string `json:"pod_name"`
}

func ReportPanic(errInfo, funcName, stack string) (err error) {
	panicReportOnce.Do(func() {
		defer func() { recover() }()

		go func() {
			panicReq := &PanicReq{
				Service:   env.Service(),
				ErrorInfo: errInfo,
				Stack:     stack,
				FuncName:  funcName,
				Host:      env.HostIP(),
				PodName:   env.PodName(),
			}

			var jsonBytes []byte
			jsonBytes, err = json.Marshal(panicReq)
			if err != nil {
				return
			}

			var req *http.Request
			req, err = http.NewRequest("GET", url, bytes.NewBuffer(jsonBytes))
			if err != nil {
				return
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{Timeout: 5 * time.Second}
			var resp *http.Response
			resp, err = client.Do(req)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			return
		}()
	})

	return
}
