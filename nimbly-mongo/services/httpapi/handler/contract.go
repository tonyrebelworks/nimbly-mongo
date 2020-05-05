package handler

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"chi-rest/bootstrap"
	"chi-rest/lib/utils"

	validator "gopkg.in/go-playground/validator.v9"
)

const (
	// XSignature custom header to hold signature string
	XSignature = "X-SIGNATURE"

	// XTimestamp custom header to hold timestamp that used for signature
	XTimestamp = "X-TIMESTAMPT"

	// XPlayer is a token that we get from OneSignal Push notification
	XPlayer = "X-PLAYER"

	// MsgSuccess ...
	MsgSuccess = "APP:SUCCESS"

	// MsgErrValidation ...
	MsgErrValidation = "ERR:VALIDATION"

	// MsgEmptyData Data not found ...
	MsgEmptyData = "ERR:EMPTY_DATA"

	// MsgErrParam error parameter argument or anything in query string
	MsgErrParam = "ERR:INVALID_PARAM"

	// MsgBadReq for general bad request
	MsgBadReq = "ERR:BAD_REQUEST"
)

// Contract ...
type Contract struct {
	*bootstrap.App
}

// Bind bind the API request payload (body) into request struct.
func (h Contract) Bind(r *http.Request, input interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&input)

	return err
}

// EmptyJSONArr ...
func (h Contract) EmptyJSONArr() []map[string]interface{} {
	return []map[string]interface{}{}
}

// SendSuccess send success into response with 200 http code.
func (h Contract) SendSuccess(w http.ResponseWriter, payload interface{}, pagination interface{}) {
	if pagination == nil {
		pagination = h.EmptyJSONArr()
	}
	h.RespondWithJSON(w, 200, MsgSuccess, "Success", payload, pagination)
}

// SendBadRequest send bad request into response with 400 http code.
func (h Contract) SendBadRequest(w http.ResponseWriter, message string) {
	h.RespondWithJSON(w, 400, MsgBadReq, message, h.EmptyJSONArr(), h.EmptyJSONArr())
}

// SendRequestValidationError Send validation error response to consumers.
func (h Contract) SendRequestValidationError(w http.ResponseWriter, validationErrors validator.ValidationErrors) {
	errorResponse := map[string][]string{}
	errorTranslation := validationErrors.Translate(h.Validator.Translator)
	for _, err := range validationErrors {
		errKey := utils.Underscore(err.StructField())
		errorResponse[errKey] = append(
			errorResponse[errKey],
			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), "[]", -1),
		)
	}

	h.RespondWithJSON(w, 400, MsgErrValidation, "validation error", errorResponse, h.EmptyJSONArr())
}

// RespondWithJSON write json response format
func (h Contract) RespondWithJSON(w http.ResponseWriter, httpCode int, statCode string, message string, payload interface{}, pagination interface{}) {
	respPayload := map[string]interface{}{
		"stat_code":  statCode,
		"stat_msg":   message,
		"pagination": pagination,
		"data":       payload,
	}

	response, _ := json.Marshal(respPayload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}

// requestIDFromContext ...
func requestIDFromContext(ctx context.Context, key string) string {
	return ctx.Value(key).(string)
}

func generateSig(s reflect.Value, timestamp string) string {
	typeOfT := s.Type()
	combine := ""
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		value := fmt.Sprintf("%v", f.Interface())
		vb := []byte(value)
		bs := sha1.Sum(vb)
		hVal := hex.EncodeToString(bs[:])

		combine += fmt.Sprintf("%s%s", typeOfT.Field(i).Tag.Get("json"), hVal)
	}
	// create complete sha1
	bSum := sha1.Sum([]byte(combine + timestamp))
	bSumVal := hex.EncodeToString(bSum[:])

	return bSumVal
}

// isValidSignature ...
func isValidSignature(obj reflect.Value, timestamp, comparator string) bool {
	return generateSig(obj, timestamp) == comparator
}

// isValidSettingSignature ...
func isValidSettingSignature(r *http.Request, key string) bool {
	sig := r.Header.Get(XSignature)
	ts := r.Header.Get(XTimestamp)

	bSum := sha1.Sum([]byte(key + ts))
	bSumVal := hex.EncodeToString(bSum[:])

	return sig == bSumVal
}

// SaveBrwXPlayer ...
// func SaveBrwXPlayer(db *sqlx.DB, brwID uint, token string) {
// 	if token != "" {
// 		err := model.NewXPlayerModel(db).BrwStore(brwID, token)
// 		if err != nil {
// 			fmt.Println("Error Save X-Player[add new loan]: " + err.Error())
// 		}
// 	} else {
// 		fmt.Println("X-Player token empty")
// 	}
// }
