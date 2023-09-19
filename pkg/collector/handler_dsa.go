package collector

const (
	DsaNamespace = "QCE/DSA"
)

func init() {
	registerHandler(DsaNamespace, defaultHandlerEnabled, NewCdnHandler)
}
