package abango

import "github.com/go-xorm/xorm"

var (
	XConfig   map[string]string
	FrontVars map[string]string //Fronrt End Server Variables
)
var (
	XEnv *EnvConf     //Kangan only
	XDB  *xorm.Engine //Kangan only

)

// 1. Receivers /////////////////////////////////////////////////////////////////
type Param struct {
	Key   string
	Value string
}

type Controller struct {
	// Ctx            *context.Context
	Ctx            Context
	controllerName string
	actionName     string
	ConnString     string
	ServerVars     map[string]string //Fronrt End Server Variables
	GlobalVars     map[string]string //Fronrt End Global Variables
	Data           map[interface{}]interface{}
}

type Context struct {
	Ask    AbangoAsk
	Answer AbangoAnswer
}

type AbangoAsk struct {
	ApiType      string
	AuthToken    string
	AskName      string
	UniqueId     string
	Body         []byte
	ServerParams []Param
}

type AbangoAnswer struct {
	Status      []byte
	Body        []byte
	ReturnTopic string
}
