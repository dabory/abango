package abango

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	e "github.com/dabory/abango/etc"

	"github.com/go-xorm/xorm"
)

func (c *Controller) Init(ask AbangoAsk) {
	c.ServerVars = make(map[string]string) // 반드시 할당해줘야 함.
	c.Data = make(map[interface{}]interface{})

	c.Ctx.Ask = ask
	c.ConnString = XConfig["KafkaAddr"] + ":" + XConfig["KafkaPort"]
	c.Ctx.ReturnTopic = c.Ctx.Ask.UniqueId

	for _, p := range c.Ctx.Ask.ServerParams {
		c.ServerVars[p.Key] = p.Value
	}

	c.GetAbangoAccessAndDb()
}

func (c *Controller) KafkaAnswer(body string) {

	// c.Ctx.Answer.Body = []byte(body) // 쓸데없는 것 같은데 나중에 지

	if _, _, err := KafkaProducer(body,
		c.Ctx.ReturnTopic, c.ConnString, XConfig["api_method"]); err != nil {
		e.MyErr("WERRWEEWQRFDFHQW", err, false)
	}
}

func (c *Controller) AnswerJson() {

	var ret []byte
	if c.Data["json"] == nil {
		msg := " QVCZEFAQWERQ " + c.Ctx.Ask.AskName + "[\"" + c.Ctx.Ask.ApiType + "\"] Data[json] is empty !"
		e.MyErr(msg, errors.New(""), true)
		return
	}
	if c.ServerVars["indent_answer"] == "yes" {
		ret, _ = json.MarshalIndent(c.Data["json"], "", "  ")
	} else {
		ret, _ = json.Marshal(c.Data["json"])
	}

	if c.Ctx.Ask.ApiType == "Kafka" {
		c.KafkaAnswer(string(ret))
	} else if c.Ctx.Ask.ApiType == "gRpc" {
		// 점진적으로 채워나가자.
	} else if c.Ctx.Ask.ApiType == "Rest" {
		// 점진적으로 채워나가자.
	}

}

func (c *Controller) GetAbangoAccessAndDb() error {

	if err := c.GetAccessAuth(); err == nil {
		var err2 error
		if c.Db, err2 = xorm.NewEngine(c.Access.DbType, c.Access.DbConnStr); err2 == nil {
			// db, err := xorm.NewEngine(XEnc.DbType, "root:root@tcp(127.0.0.1:3306)/kangan?charset=utf8&parseTime=True")

			strArr := strings.Split(c.Access.DbConnStr, "@tcp")
			// if len(strArr) == 2 {
			// 	e.OkLog(strArr[1])
			// } else {
			// 	e.MyErr(strArr[1], err2, true)
			// 	return err2
			// }

			c.Db.ShowSQL(false)
			c.Db.SetMaxOpenConns(100)
			c.Db.SetMaxIdleConns(20)
			c.Db.SetConnMaxLifetime(60 * time.Second)
			if _, err := c.Db.IsTableExist("admin_menu"); err != nil {
				e.MyErr("DB DISconnected in "+strArr[1], err, true)
			} else {
				e.OkLog("DB connect in " + strArr[1])
			}
		} else {
			e.MyErr("xorm.NewEngine", err, true)
		}
	} else {
		e.MyErr("GetAccessAuth", err, true)
	}

	return nil
}

func (c *Controller) GetAccessAuth() error {

	c.Access.UserId = 10
	c.Access.DbType = "mysql"
	c.Access.DbConnStr = "dbr_shop:k99614791@tcp(aurora-biozoa.clhvwtdfc3ov.ap-northeast-2.rds.amazonaws.com:3306)/dbr_shop?charset=utf8"

	return nil
}
