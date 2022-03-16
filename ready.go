package gateway

func (gw *Gateway) Ready() bool {
	return !gw.readiness || gw.Controller.HasSynced()
}
