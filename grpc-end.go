package abango

import (
	"context"
	"encoding/json"
	"time"

	grp1 "github.com/dabory/abango/protos"

	e "github.com/dabory/abango/etc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func GrpcRequest(v *AbangoAsk) (string, string, error) {

	dial, err := grpc.Dial(
		XConfig["gRpcAddr"]+":"+XConfig["gRpcPort"],
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Millisecond,
			Timeout:             time.Millisecond,
			PermitWithoutStream: true,
		}),
	)

	if err != nil {
		e.MyErr("WOEIURQPWERQ", nil, true)
	}
	defer dial.Close()
	c := grp1.NewGrp1Client(dial)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	
	askstr, _ := json.Marshal(&v)
	if r, err := c.StdRpc(ctx, &grp1.StdAsk{AskMsg: []byte(askstr)}); err == nil {
		return string(r.RetMsg), string(r.RetSta), nil
	} else {
		return "", "", e.MyErr("WERVCZWERTGS", err, true)
	}

}
