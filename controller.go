package abango

import e "github.com/dabory/abango/etc"

func (v *Controller) Init(ask AbangoAsk) {

	v.Ctx.Ask = ask

	v.ConnString = XConfig["KafkaAddr"] + ":" + XConfig["KafkaPort"]
	v.Ctx.Answer.ReturnTopic = v.Ctx.Ask.UniqueId

	v.ServerVars = make(map[string]string) // 반드시 할당해줘야 함.
	for _, p := range v.Ctx.Ask.ServerParams {
		v.ServerVars[p.Key] = p.Value
	}
}

func (v *Controller) KafkaAnswer() {
	// e.Tp(XConfig["api_method"])
	if _, _, err := KafkaProducer(string(v.Ctx.Answer.Body),
		v.Ctx.Answer.ReturnTopic, v.ConnString, XConfig["api_method"]); err != nil {
		// 	e.OkLog("Kafka-ReturnTopic: " + v.Ctx.Answer.ReturnTopic)
		// } else {
		e.MyErr("WERRWERFAAFHQW", err, false)
	}
}
