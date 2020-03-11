package abango

import (
	"encoding/json"
	"strings"

	e "github.com/dabory/abango/etc"
)

//////////// Kafka EndPoint /////////////
func RestRequest(v *AbangoAsk) (string, string, error) {

	svars := make(map[string]string)
	for _, p := range v.ServerParams {
		svars[p.Key] = p.Value
	}

	apiMethod := strings.ToUpper(svars["api_method"])
	if apiMethod == "" { // Default is POST
		apiMethod = "POST"
	}
	askBytes, _ := json.Marshal(&v)

	if apiMethod == "POST" {
		if retstr, retsta, err := e.GetHttpResponse(apiMethod, XConfig["RestConnect"], askBytes); err == nil {
			return string(retstr), string(retsta), nil
		} else {
			return "", "", err
		}
	} else if apiMethod == "UPLOAD" {
		// 	httpRtn, statusCode, err = getUploadHttpResponse(apiMethod, _apiUrlPrefix+apiUri, postBytes)
		// 	if err != nil {
		// 		return C.CString(sndrDisconnect(statusCode, err))
		// 		// return sndrDisconnect(statusCode, err)
		// 	}

		// }
		return "", "", e.MyErr("WERZDDERUVE-api_method not available UPLOAD-["+apiMethod+"]", nil, true)
	} else {
		return "", "", e.MyErr("WERZDDERUVE-api_method not available ["+apiMethod+"]", nil, true)
	}

}
