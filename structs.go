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

	Access AbangoAccess
	Db     *xorm.Engine
}

type Context struct {
	Ask         AbangoAsk
	Answer      AbangoAnswer
	ReturnTopic string
}

type AbangoAsk struct {
	ApiType      string
	AskName      string
	AccessToken  string
	UniqueId     string
	DocRoot      string
	Body         []byte
	ServerParams []Param
}

type AbangoAnswer struct {
	Body []byte
}

type AbangoAccess struct {
	UserId    int64
	UserGuid  string
	UserName  string
	NickName  string
	DbType    string
	DbConnStr string
}
