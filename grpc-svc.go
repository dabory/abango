package abango

func GrpcSvcStandBy(RouterHandler func()) {
	RouterHandler()
}
