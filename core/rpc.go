package core

import (
	"bytes"
	j "encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/raven0520/btc/app"
)

// RPCConfig Configure of RPC
type RPCConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Debug    bool   `json:"debug"`
}

// RPCRequest Request of BTC RPC
type RPCRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int64       `json:"id"`
	JSONRpc string      `json:"jsonrpc"`
}

// RPCResponse Response of BTC RPC
type RPCResponse struct {
	ID     int64            `json:"id"`
	Result j.RawMessage     `json:"result"`
	Error  RPCResponseError `json:"error"`
}

// RPCResponseError Error of response
type RPCResponseError struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

//	Params:
//	node: chain node
//	method: rpc method
//	params: rpc params
//	result: recive result

//	Return:

// CallJSON Call BTC RPC for json response
func CallJSON(node, method string, params, result interface{}) (res []byte, err error) {
	var rpc RPCResponse
	config, ok := app.GetBTCNodeConfig(node)
	if !ok {
		err = errors.New("BTC Config Error")
		return
	}
	host := fmt.Sprintf("%s:%d", config.Host, config.Port)
	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}
	b, err := json.Marshal(RPCRequest{method, params, time.Now().UnixNano(), "1.0"})
	if err != nil {
		return
	}
	bodyReader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPost, host, bodyReader)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	if config.User != "" {
		req.SetBasicAuth(config.User, config.Pass)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if config.Debug {
		log.Println("[BTC Debug] rpc return", string(body))
	}
	if err = json.Unmarshal(body, &rpc); err != nil {
		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf("%d, %s", resp.StatusCode, string(body))
		}
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Code : %d \n Error : %s", rpc.Error.Code, rpc.Error.Message)
		return
	}

	if result != nil {
		if err = json.Unmarshal(rpc.Result, &result); err != nil {
			return nil, err
		}
	}
	return rpc.Result, err
}
