package sandbox

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/cbroglie/mustache"
	"github.com/codebyshennan/mockapi/domain"
	"github.com/codebyshennan/mockapi/internal"
	"github.com/codebyshennan/mockapi/internal/common/h"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct that represents the parsed request URL
type sandboxUri struct {
	mockServerId primitive.ObjectID
	method       string
	endPt        string
	provider     *internal.Provider
	valuesMap    map[string]any
}

// Parses the request url
func newSandboxUri(r *http.Request, p *internal.Provider) (*sandboxUri, error) {
	method := r.Method
	if method == "" {
		method = "GET"
	}

	url := r.URL.String()
	urlTokens := strings.SplitN(url, "/", 4)

	mockServerId, err := primitive.ObjectIDFromHex(urlTokens[2])
	if err != nil {
		return nil, err
	}

	var endpt string
	if len(urlTokens) >= 4 {
		endpt = "/" + urlTokens[3]
	}

	valuesMap := map[string]any{}
	_ = json.NewDecoder(r.Body).Decode(&valuesMap)

	return &sandboxUri{
		mockServerId: mockServerId,
		method:       method,
		endPt:        endpt,
		provider:     p,
		valuesMap:    valuesMap,
	}, nil
}

// Set the response
func (o sandboxUri) setResponse(w *http.ResponseWriter) {
	endpts, err := o.provider.MockEndPtRepo.GetMockEndPts(o.mockServerId)
	if err != nil {
		h.JsonRes(*w, http.StatusNotImplemented, h.Json{
			"code":    "ERR_SB_TODO",
			"message": "This route has not been implemented yet",
		}, o.provider.Logger)
	}

	for _, endpt := range endpts {
		if endpt.Method != o.method {
			continue
		}

		if match, _ := regexp.MatchString(endpt.EndpointRegex, o.endPt); match {
			res, _ := o.getResBody(*endpt.ResponseBody)

			if endpt.Timeout != nil && *endpt.Timeout != 0 {
				time.Sleep(time.Duration(*endpt.Timeout) * time.Second)
			}

			if endpt.Writes != nil {
				o.provider.MockServerStoreRepo.CreateRecord(o.mockServerId, &domain.MockServerStoreCreateData{
					Data: endpt.Writes.Data,
					Dest: "default",
				})
			}

			h.JsonRes(*w, endpt.ResponseCode, res, o.provider.Logger)
			return
		}
	}

	h.JsonRes(*w, http.StatusNotImplemented, h.Json{
		"code":    "ERR_SB_TODO",
		"message": "This route has not been implemented yet",
	}, o.provider.Logger)
}

func (o sandboxUri) getResBody(data string) (string, error) {
	o.provider.Logger.Debugln("INFO_SB_PARSING", "raw", data)

	template, err := mustache.ParseString(data)
	if err != nil {
		o.provider.Logger.Errorln("ERR_SB_PARSING", err.Error())
		return data, err
	}
	o.provider.Logger.Debugln("INFO_SB_PARSING", "template", template)

	if err = o.getValues(template.Tags()); err != nil {
		return data, err
	}

	output, err := template.Render(o.valuesMap)
	if err != nil {
		o.provider.Logger.Errorln("ERR_SB_PARSING", err.Error())
		return data, err
	}
	o.provider.Logger.Debugln("INFO_SB_PARSING", "output", output)

	return output, nil
}

func (o sandboxUri) getValues(tags []mustache.Tag) (err error) {
	for _, tag := range tags {
		key := tag.Name()
		o.provider.Logger.Debugln("INFO_SB_PARSING", "key", key)

		toks := strings.SplitN(key, ".", 2)
		prefix := toks[0]
		o.provider.Logger.Debugln("INFO_SB_PARSING", "toks", toks)

		switch prefix {
		case "store":
			storeRes, err := o.provider.MockServerStoreRepo.GetRecord(o.mockServerId, "default")
			if err != nil {
				o.provider.Logger.Debugln("INFO_SB_PARSING", "getValuesErr", err.Error())
				o.valuesMap[key] = key
			} else {
				o.provider.Logger.Debugln("INFO_SB_PARSING", "getValuesOk", storeRes)
				o.valuesMap["store"] = map[string]any{
					"default": storeRes,
				}
				// strings.Join(storeRes, ",")
			}
		default:
			o.valuesMap[key] = key
		}
	}

	o.provider.Logger.Debugln("INFO_SB_PARSING", "valuesMap", o.valuesMap)
	return nil
}
