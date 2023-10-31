package add

type Adder interface {
	Achieve() error
	Adjust() error
	Store() error
}

type AddOptions struct {
	Host       string
	Port       string
	Username   string
	Password   string
	PrivateKey string
	SourcePath string
	AcpUrl     string
	Cluster    string
	Name       string
	Namespace  string
	token      string
}
