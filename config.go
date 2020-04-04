package abango

import (
	"encoding/json"
	"os"

	e "github.com/dabory/abango/etc"

	_ "github.com/go-sql-driver/mysql"
)

func GetXConfig(params ...string) error { // Kafka, gRpc, REST 통합 업그레이드

	conf := "conf/"
	if len(params) != 0 {
		conf = params[0] + conf
	}

	RunFilename := conf + "config_select.json"

	run := struct {
		ConfSelect  string
		ConfPostFix string
	}{}

	if file, err := os.Open(RunFilename); err != nil {
		e.MyErr("WERQRRQERQWERFD", nil, true)
		return err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&run); err != nil {
			e.MyErr("ERTFDFDAFA", err, true)
			return err
		}
	}

	XConfig = make(map[string]string) // Just like malloc
	config := []Param{}

	// var varMap []map[string]interface{}
	filename := conf + run.ConfSelect + run.ConfPostFix
	if file, err := os.Open(filename); err != nil {
		e.MyErr("QERTRRTRRW", err, true)
		return err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&config); err == nil {
			for _, p := range config {
				XConfig[p.Key] = p.Value
			}
		} else {
			e.MyErr("LAAFDFERHWERYTY", err, true)
			return err
		}
	}

	if XConfig["KafkaOn"] == "Yes" || XConfig["ApiType"] == "Kafka" {
		e.Tp("==" + "Config file prefix: " + run.ConfSelect + "== Kafka Connection: " + XConfig["KafkaAddr"] + ":" + XConfig["KafkaPort"] + "==")
	}
	if XConfig["gRpcOn"] == "Yes" || XConfig["ApiType"] == "gRpc" {
		e.Tp("==" + "Config file prefix: " + run.ConfSelect + "== gRpc Connection: " + XConfig["gRpcAddr"] + ":" + XConfig["gRpcPort"] + "==")
	}
	if XConfig["RestOn"] == "Yes" || XConfig["ApiType"] == "Rest" {
		e.Tp("==" + "Config file prefix: " + run.ConfSelect + "== REST Connection: " + XConfig["RestConnect"] + "==")
	}
	return nil
}

// func GetServerVarsInEnd(askname string, unique_id string) (string, error) { // Kafka, gRpc, REST 통합 업그레이드

// 	// unique_id := e.RandString(20)
// 	fvars := []Param{}
// 	// comarr := make(map[string]string)

// 	filename := "conf/server-vars.json"
// 	if file, err := os.Open(filename); err != nil {
// 		e.MyErr("QERTRRTRRWQWRE", err, true)
// 		return "", err
// 	} else {
// 		decoder := json.NewDecoder(file)
// 		if err = decoder.Decode(&fvars); err == nil {
// 			for i := 0; i < len(fvars); i++ {
// 				if fvars[i].Key == "askname" {
// 					fvars[i].Value = askname // 유일키
// 				} else if fvars[i].Key == "unique_id" {
// 					fvars[i].Value = unique_id
// 				} else if fvars[i].Key == "server_addr" {
// 					addrs, _ := net.InterfaceAddrs()
// 					fvars[i].Value = fmt.Sprintf("%v", addrs[0]) // Server IP
// 				}
// 			}
// 		} else {
// 			e.MyErr("LAAFDFERHYWE", err, true)
// 			return "", err
// 		}
// 	}

// 	fstr, _ := json.Marshal(&fvars)
// 	return string(fstr), nil
// }

// func GetServerVarsInSvc(t []byte) error { // Kafka, gRpc, REST 통합 업그레이드

// 	ServerVars = make(map[string]string) // 반드시 = 로 할 것
// 	evars := []Param{}

// 	if err := json.Unmarshal(t, &evars); err == nil {
// 		for _, p := range evars {
// 			ServerVars[p.Key] = p.Value
// 		}
// 	} else {
// 		e.MyErr("QWECVZDFVBXGF", err, true)
// 		return err
// 	}

// 	return nil
// }

// func GetMapVars(t []Param) (map[string]string, error) { // Kafka, gRpc, REST 통합 업그레이드

// 	comarr := make(map[string]string)
// 	if content, err := ioutil.ReadFile("golangcode.txt"); err == nil {
// 		if err := json.Unmarshal(content, &t); err == nil {
// 			for _, p := range t {
// 				comarr[p.Name] = p.Value
// 			}
// 		} else {
// 			e.MyErr("QWECVZDFVBXGF", err, true)
// 			return nil, err
// 		}
// 	} else {
// 		e.MyErr("QERTRRTRRW", err, true)
// 		return nil, err
// 	}

// 	return comarr, nil
// }
