package response

import (
	"compress/gzip"
	"encoding/json"
	"net/http"

	"github.com/gongcf/gtool/logger"
)

type GResponse struct{}

// Result represents a common-used result struct.
type Result struct {
	Code int         `json:"code"` // return code
	Msg  string      `json:"msg"`  // message
	Data interface{} `json:"data"` // data object
}

// NewResult creates a result with Code=0, Msg="", Data=nil.
func (*GResponse) NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}

// RetResult writes HTTP response with "Content-Type, application/json".
func (*GResponse) RetResult(w http.ResponseWriter, r *http.Request, res *Result) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(res)
	if err != nil {
		logger.Glogger.Error(err)

		return
	}

	w.Write(data)
}

// RetGzResult writes HTTP response with "Content-Type, application/json" and "Content-Encoding, gzip".
func (*GResponse) RetGzResult(w http.ResponseWriter, r *http.Request, res *Result) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")

	gz := gzip.NewWriter(w)
	err := json.NewEncoder(gz).Encode(res)
	if nil != err {
		logger.Glogger.Error(err)

		return
	}

	err = gz.Close()
	if nil != err {
		logger.Glogger.Error(err)

		return
	}
}

// RetJSON writes HTTP response with "Content-Type, application/json".
func (*GResponse) RetJSON(w http.ResponseWriter, r *http.Request, res map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(res)
	if err != nil {
		logger.Glogger.Error(err)

		return
	}

	w.Write(data)
}

// RetGzJSON writes HTTP response with "Content-Type, application/json" and "Content-Encoding, gzip".
func (*GResponse) RetGzJSON(w http.ResponseWriter, r *http.Request, res map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")

	gz := gzip.NewWriter(w)
	err := json.NewEncoder(gz).Encode(res)
	if nil != err {
		logger.Glogger.Error(err)

		return
	}

	err = gz.Close()
	if nil != err {
		logger.Glogger.Error(err)

		return
	}
}
