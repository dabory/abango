package abango

import (
	"encoding/json"
	"os"
	"strings"

	"sync"

	_ "github.com/go-sql-driver/mysql"

	e "github.com/dabory/abango/etc"
)

// type Controller struct {
// }

type EnvConf struct { //Kangan only
	AppName      string
	HttpProtocol string
	HttpAddr     string
	HttpPort     string
	SiteName     string

	DbType     string
	DbHost     string
	DbUser     string
	DbPassword string
	DbPort     string
	DbName     string
	DbPrefix   string
	DbTimezone string

	DbStr string
}

type RunConf struct {
	RunMode     string
	DevPrefix   string
	ProdPrefix  string
	ConfPostFix string
}

func init() {
	// e.OkLog("Abango Initialized")
}

func RunServicePoint(KafkaHandler func(ask *AbangoAsk), GrpcHandler func(), RestHandler func(ask *AbangoAsk)) {

	var wg sync.WaitGroup

	e.AokLog("Abango Clustered Framework Started !")
	if err := GetXConfig(); err == nil {
		if XConfig["KafkaOn"] == "Yes" {
			wg.Add(1)
			go func() {
				KafkaSvcStandBy(KafkaHandler)
				wg.Done()
			}()
		}
		if XConfig["gRpcOn"] == "Yes" {
			// e.AokLog("gRpc API StandBy !")
			wg.Add(1)
			go func() {
				GrpcSvcStandBy(GrpcHandler)
				wg.Done()
			}()
		}
		if XConfig["RestOn"] == "Yes" {
			// e.AokLog("RESTful API StandBy !")
			wg.Add(1)
			go func() {
				RestSvcStandBy(RestHandler)
				wg.Done()
			}()
		}

	} else {
		e.Atp("Error running RunServicePoint")
	}

	wg.Wait()
}

func RunEndRequest(docroot string, params string, body string) string {

	// return docroot + "//" + params + "//" + body
	testModeYes := false
	homeroot := ""
	// e.Tp(devdir)
	if docroot != "" {
		homeroot = e.ParentDir(docroot) + "/"
	} else {
		testModeYes = true
	}

	// f, _ := os.OpenFile(homeroot+"abango.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	// defer f.Close()

	// w := bufio.NewWriter(f)
	// w.WriteString("This is 1" + "\n")
	// w.Flush()

	if err := GetXConfig(homeroot); err == nil {
		if docroot == "" { // golang test mode
			testModeYes = true
			askfile := e.GetAskName()
			arrask := strings.Split(askfile, "@") // login@post 앞의 문자를 askname으로 설정
			askname := arrask[0]

			jsonsend := homeroot + XConfig["JsonSendDir"] + askname + ".json"

			var err error
			if body, err = e.FileToStr(jsonsend); err != nil {
				return e.MyErr("WERZDSVCZSRE-JsonSendFile Not Found: ", err, true).Error()
			}
		}

		if XConfig["ApiType"] == "Kafka" {
			e.MyLog(homeroot+"abango.log", "homeroot="+homeroot)
			return RunRequest(KafkaRequest, &homeroot, &params, &body, testModeYes)
			// } else if XConfig["ApiType"] == "gRpc" {
			// 	return RunRequest(GrpcRequest)
			// } else if XConfig["ApiType"] == "Rest" {
			// 	return RunRequest(RestRequest)
		} else {
			e.MyLog(homeroot+"abango.log", "A-B")
			return e.MyErr("QREWFGARTEGF-Wrong ApiType in RunEndRequest()", nil, true).Error()
		}
	} else {
		return e.MyErr("XCVZDSFGQWERDZ-Unable to get GetXConfig()", nil, true).Error()
	}
	return "Reached to end of RunEndRequest !"
}

func RunRequest(MsgHandler func(v *AbangoAsk) (string, string, error), homeroot *string, params *string, body *string, testModeYes bool) string {

	var v AbangoAsk
	v.UniqueId = e.RandString(20)
	v.Body = []byte(*body)
	v.HomeRoot = *homeroot

	jsonsvrparams := *homeroot + XConfig["JsonServerParamsPath"]
	if file, err := os.Open(jsonsvrparams); err == nil {
		if err = json.NewDecoder(file).Decode(&v.ServerParams); err != nil {
			return e.MyErr("LAAFDFDFERHYWE", err, true).Error()
		}
	} else {
		return e.MyErr("LAAFDFDWDERHYWE-"+jsonsvrparams+" File not found", err, true).Error()
	}

	if *params != "" { //User Params 있을 경우 해당을 가져온다.
		var askparmas []Param
		if err := json.Unmarshal([]byte(*params), &askparmas); err == nil {
			for _, j := range askparmas {
				for _, s := range v.ServerParams {
					if s.Key == j.Key {
						s.Value = j.Value
					} // 여기서 api-method 도 처리됨.
				}
				if j.Key == "ApiType" { // Ask Params 에 ApiType 이 지정되어 있다면
					v.ApiType = j.Value
				}
				if j.Key == "AskName" {
					v.AskName = j.Value
				}
			}
		} else {
			return e.MyErr("WERITOGFSERFDH-AskParams Format mismatched:", nil, true).Error()
		}

	} else {
		askfile := e.GetAskName()
		arrask := strings.Split(askfile, "@") // @앞의 문자를 askname으로 설정
		askname := arrask[0]
		apimethod := ""
		if len(arrask) >= 2 { //만약 argv[1] 이 login@Kafka 형태라면
			apimethod = arrask[1]
		}
		for i := 0; i < len(v.ServerParams); i++ {
			if v.ServerParams[i].Key == "api_method" { //GET, POST
				v.ServerParams[i].Value = apimethod
			}
		}

		v.ApiType = XConfig["ApiType"]
		v.AskName = askname
	}

	if v.ApiType == "" || v.AskName == "" {
		return e.MyErr("QWERDSFAERQRDA-ApiType or AskName was not specified:", nil, true).Error()
	}

	if retstr, retsta, err := MsgHandler(&v); err == nil {
		ZZ
		if testModeYes == true {
			jsonreceive := XConfig["JsonReceiveDir"] + v.AskName + ".json"
			if XConfig["SaveReceivedJson"] == "Yes" {
				e.StrToFile(jsonreceive, retstr)
			}
			if XConfig["ShowReceivedJson"] == "Yes" {
				e.Tp("Status: " + retsta + "  ReturnJsonFile: " + jsonreceive)
				e.Tp(retstr)
			}
		}
		return retstr
	} else {
		return e.MyErr("QWERDSFAERQRDA-MsgHandler", err, true).Error()
	}
}

func GetEnvConf() error { // Kangan only

	conf := "conf/"
	RunFilename := conf + "run_conf.json"

	var run RunConf

	if file, err := os.Open(RunFilename); err != nil {
		e.MyErr("SDFLJDSAFJA", nil, true)
		return err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&run); err != nil {
			e.MyErr("LASJLDFJASFJ", err, true)
			return err
		}
	}

	filename := conf + run.RunMode + run.ConfPostFix
	if file, err := os.Open(filename); err != nil {
		e.MyErr("QERTRRTRRW", err, true)
		return err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&XEnv); err != nil {
			e.MyErr("LAAFDFERHY", err, true)
			return err
		}
	}

	if XEnv.DbType == "mysql" {
		XEnv.DbStr = XEnv.DbUser + ":" + XEnv.DbPassword + "@tcp(" + XEnv.DbHost + ":" + XEnv.DbPort + ")/" + XEnv.DbPrefix + XEnv.DbName + "?charset=utf8"
	} else if XEnv.DbType == "mssql" {
		// Add on more DbStr of Db types
	}

	return nil
}
