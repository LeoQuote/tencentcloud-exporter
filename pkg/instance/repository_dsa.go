package instance

func init() {
	registerRepository("QCE/DSA", NewCdnTcInstanceRepository)
}
