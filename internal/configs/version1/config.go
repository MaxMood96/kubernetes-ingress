package version1

// IngressNginxConfig describes an NGINX configuration.
type IngressNginxConfig struct {
	Upstreams []Upstream
	Servers   []Server
	Keepalive string
	Ingress   Ingress
}

// Ingress holds information about an Ingress resource.
type Ingress struct {
	Name        string
	Namespace   string
	Annotations map[string]string
}

// Upstream describes an NGINX upstream.
type Upstream struct {
	Name             string
	UpstreamServers  []UpstreamServer
	StickyCookie     string
	LBMethod         string
	Queue            int64
	QueueTimeout     int64
	UpstreamZoneSize string
}

// UpstreamServer describes a server in an NGINX upstream.
type UpstreamServer struct {
	Address     string
	Port        string
	MaxFails    int
	MaxConns    int
	FailTimeout string
	SlowStart   string
	Resolve     bool
}

// HealthCheck describes an active HTTP health check.
type HealthCheck struct {
	UpstreamName   string
	URI            string
	Interval       int32
	Fails          int32
	Passes         int32
	Scheme         string
	Mandatory      bool
	Headers        map[string]string
	TimeoutSeconds int64
}

// Server describes an NGINX server.
type Server struct {
	ServerSnippets        []string
	Name                  string
	ServerTokens          string
	Locations             []Location
	SSL                   bool
	SSLCertificate        string
	SSLCertificateKey     string
	SSLCiphers            string
	TLSPassthrough        bool
	GRPCOnly              bool
	StatusZone            string
	HTTP2                 bool
	RedirectToHTTPS       bool
	SSLRedirect           bool
	ProxyProtocol         bool
	HSTS                  bool
	HSTSMaxAge            int64
	HSTSIncludeSubdomains bool
	HSTSBehindProxy       bool
	ProxyHideHeaders      []string
	ProxyPassHeaders      []string

	HealthChecks map[string]HealthCheck

	RealIPHeader    string
	SetRealIPFrom   []string
	RealIPRecursive bool

	JWTAuth              *JWTAuth
	JWTRedirectLocations []JWTRedirectLocation

	Ports    []int
	SSLPorts []int
}

// JWTRedirectLocation describes a location for redirecting client requests to a login URL for JWT Authentication.
type JWTRedirectLocation struct {
	Name     string
	LoginURL string
}

// JWTAuth holds JWT authentication configuration.
type JWTAuth struct {
	Key                  string
	Realm                string
	Token                string
	RedirectLocationName string
}

// Location describes an NGINX location.
type Location struct {
	LocationSnippets     []string
	Path                 string
	Upstream             Upstream
	ProxyConnectTimeout  string
	ProxyReadTimeout     string
	ProxySendTimeout     string
	ClientMaxBodySize    string
	Websocket            bool
	Rewrite              string
	SSL                  bool
	GRPC                 bool
	ProxyBuffering       bool
	ProxyBuffers         string
	ProxyBufferSize      string
	ProxyMaxTempFileSize string
	JWTAuth              *JWTAuth

	MinionIngress *Ingress
}

// MainConfig describe the main NGINX configuration file.
type MainConfig struct {
	AccessLogOff                   bool
	DefaultServerAccessLogOff      bool
	ErrorLogLevel                  string
	HealthStatus                   bool
	HealthStatusURI                string
	HTTP2                          bool
	HTTPSnippets                   []string
	KeepaliveRequests              int64
	KeepaliveTimeout               string
	LogFormat                      []string
	LogFormatEscaping              string
	MainSnippets                   []string
	NginxStatus                    bool
	NginxStatusAllowCIDRs          []string
	NginxStatusPort                int
	OpenTracingEnabled             bool
	OpenTracingLoadModule          bool
	OpenTracingTracer              string
	OpenTracingTracerConfig        string
	ProxyProtocol                  bool
	ResolverAddresses              []string
	ResolverIPV6                   bool
	ResolverTimeout                string
	ResolverValid                  string
	ServerNamesHashBucketSize      string
	ServerNamesHashMaxSize         string
	ServerTokens                   string
	SSLCiphers                     string
	SSLDHParam                     string
	SSLPreferServerCiphers         bool
	SSLProtocols                   string
	StreamLogFormat                []string
	StreamLogFormatEscaping        string
	StreamSnippets                 []string
	StubStatusOverUnixSocketForOSS bool
	TLSPassthrough                 bool
	VariablesHashBucketSize        uint64
	VariablesHashMaxSize           uint64
	WorkerConnections              string
	WorkerCPUAffinity              string
	WorkerProcesses                string
	WorkerRlimitNofile             string
	WorkerShutdownTimeout          string
}

// NewUpstreamWithDefaultServer creates an upstream with the default server.
// proxy_pass to an upstream with the default server returns 502.
// We use it for services that have no endpoints.
func NewUpstreamWithDefaultServer(name string) Upstream {
	return Upstream{
		Name:             name,
		UpstreamZoneSize: "256k",
		UpstreamServers: []UpstreamServer{
			{
				Address:     "127.0.0.1",
				Port:        "8181",
				MaxFails:    1,
				MaxConns:    0,
				FailTimeout: "10s",
			},
		},
	}
}
