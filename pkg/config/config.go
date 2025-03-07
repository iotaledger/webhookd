package config

// Config contain global configuration
type Config struct {
	ListenAddr      string `flag:"listen-addr" desc:"HTTP listen address" default:":8080"`
	TLS             bool   `flag:"tls" desc:"Activate TLS" default:"false"`
	TLSCertFile     string `flag:"tls-cert-file" desc:"TLS certificate file" default:"server.pem"`
	TLSKeyFile      string `flag:"tls-key-file" desc:"TLS key file" default:"server.key"`
	TLSDomain       string `flag:"tls-domain" desc:"TLS domain name used by ACME"`
	NbWorkers       int    `flag:"nb-workers" desc:"Number of workers to start" default:"2"`
	HookTimeout     int    `flag:"hook-timeout" desc:"Maximum hook execution time in second" default:"10"`
	HookLogDir      string `flag:"hook-log-dir" desc:"Hook execution logs location" default:""`
	HookLogOutput   bool   `flag:"hook-log-output" desc:"Output hook execution logs to server logs" default:"false"`
	ScriptDir       string `flag:"scripts" desc:"Scripts location" default:"scripts"`
	PasswdFile      string `flag:"passwd-file" desc:"Password file for basic HTTP authentication" default:".htpasswd"`
	LogLevel        string `flag:"log-level" desc:"Log level (debug, info, warn, error)" default:"info"`
	StaticDir       string `flag:"static-dir" desc:"Static file directory to serve on /static path" default:""`
	NotificationURI string `flag:"notification-uri" desc:"Notification URI"`
	TrustStoreFile  string `flag:"trust-store-file" desc:"Trust store used by HTTP signature verifier (.pem or .p12)"`
}
