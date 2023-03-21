package h

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/codebyshennan/mockapi/domain"
)

// Checks if a file exists and is not a directory
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Sets the code and payload in the response
// Similar to Express's res.json and res.status
func JsonRes(w http.ResponseWriter, code int, payload any, logger domain.ILogger) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("X-Sandbox", os.Getenv("SANDBOX_MODE"))
	// w.Header().Set("Access-Control-Expose-Headers", "X-Sandbox")

	response, err := json.Marshal(payload)
	if err != nil {
		logger.Errorln(EncodeError, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{code: " + EncodeError + "}"))
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}

// A type alias for a generic json object
type Json map[string]any

// Binds a request's body to the output
func BindJson[T interface{}](r *http.Request, output *T) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, output)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Get limit and skip from request query params
func GetLimitSkip(r *http.Request) (limit int64, skip int64) {
	if l, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64); err != nil {
		limit = l
	}
	if s, err := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64); err != nil {
		skip = s
	}
	return
}
