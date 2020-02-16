package abango

import (
	"encoding/json"
	"os"
	"strings"
	time "time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)
import "sync"
import e "github.com/dabory/abango/etc"

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
	// func RunServicePoint(KafkaHandler func(ask *AbangoAsk), GrpcHandler func(), RestHandler func(ask *AbangoAsk)) {

	// KafkaSvcStandBy(KafkaHandler)
	// GrpcSvcStandBy(GrpcHandler)

	var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	GrpcSvcStandBy(GrpcHandler)
	// 	wg.Done()
	// }()
	// wg.Add(1)
	// go func() {
	// 	RestSvcStandBy(RestHandler)
	// 	wg.Done()
	// }()

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

func RunEndRequest(params ...string) {
	if err := GetXConfig(); err == nil {
		if XConfig["ApiType"] == "Kafka" {
			RunRequest(KafkaRequest)
		} else if XConfig["ApiType"] == "gRpc" {
			RunRequest(GrpcRequest)
		} else if XConfig["ApiType"] == "Rest" {
			RunRequest(RestRequest)
		} else {
			e.Atp("Error running RunEndPoint")
		}
	} else {

	}
	// e.Atp(XConfig["Dummy"])
}

func RunRequest(MsgHandler func(v *AbangoAsk) (string, string, error)) error {

	unique_id := e.RandString(20)

	askfile := e.GetAskName()
	arrask := strings.Split(askfile, "@") // @앞의 문자를 askname으로 설정
	askname := arrask[0]
	apimethod := ""
	if len(arrask) >= 2 {
		apimethod = arrask[1]
	}
	jsonsend := XConfig["JsonSendDir"] + askname + ".json"
	jsonreceive := XConfig["JsonReceiveDir"] + askname + ".json"
	jsonsvrparams := XConfig["JsonServerParamsPath"]

	if file, err := os.Open(jsonsvrparams); err == nil {
		var v AbangoAsk
		if err = json.NewDecoder(file).Decode(&v.ServerParams); err == nil {
			if askstr, err := e.FileToStr(jsonsend); err == nil {
				v.ApiType = XConfig["ApiType"]
				v.AskName = askname
				v.UniqueId = unique_id
				v.Body = []byte(askstr)

				for i := 0; i < len(v.ServerParams); i++ {
					if v.ServerParams[i].Key == "api_method" {
						v.ServerParams[i].Value = apimethod
					}
				}
				if retstr, retsta, err := MsgHandler(&v); err == nil {
					e.Tp("ReturnStatus: " + retsta + "  ReturnJsonFile: " + jsonreceive)
					e.StrToFile(jsonreceive, retstr)
					if XConfig["ShowReceivedJson"] == "Yes" {
						e.Tp(retstr)
					}
				} else {
					e.MyErr("QWERDSFAERQRDA-MsgHandler", err, true)
				}
			} else {
				e.MyErr("WERZDSVCZSRE-JsonSendFile", err, true)
			}
		} else {
			return e.MyErr("LAAFDFDFERHYWE", err, true)
		}
	} else {
		return e.MyErr("LAAFDFDWDERHYWE-"+jsonsvrparams+" file not found", err, true)
	}

	return nil
}

func Run(params ...string) { //Kangan only

	if err := GetEnvConf(); err == nil {
		db, err := xorm.NewEngine(XEnv.DbType, XEnv.DbStr)
		// db, err := xorm.NewEngine(XEnv.DbType, "root:root@tcp(127.0.0.1:3306)/kangan?charset=utf8&parseTime=True")

		dbStr := XEnv.DbType + ":(" + XEnv.DbHost + ":" + XEnv.DbPort + ")->[" + XEnv.DbPrefix + XEnv.DbName + "] DB Schema"
		if err == nil {
			e.OkLog(dbStr)
		} else {
			e.MyErr(dbStr, err, true)
			// panic(fmt.Errorf("Database open error: %s \n", err))
		}

		db.ShowSQL(false)
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(20)
		db.SetConnMaxLifetime(60 * time.Second)
		if _, err := db.IsTableExist("aaa"); err != nil { //Connect Check
			e.MyErr("DATABASE DISCONNECTED", err, true)
		} else {
			e.OkLog("DATABASE CONNECTED")
		}
		XDB = db
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
