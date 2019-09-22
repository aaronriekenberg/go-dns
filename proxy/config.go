package proxy

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
)

// HostAndPort is a host and port.
type HostAndPort struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// JoinHostPort joins the host and port.
func (hostAndPort *HostAndPort) joinHostPort() string {
	return net.JoinHostPort(hostAndPort.Host, hostAndPort.Port)
}

// ForwardNameToAddress is a forward name to IP address mapping.
type ForwardNameToAddress struct {
	Name      string `json:"name"`
	IPAddress string `json:"ipAddress"`
}

// ReverseAddressToName is a reverse address to name mapping.
type ReverseAddressToName struct {
	ReverseAddress string `json:"reverseAddress"`
	Name           string `json:"name"`
}

// Configuration is the DNS proxy configuration.
type Configuration struct {
	ListenAddress           HostAndPort            `json:"listenAddress"`
	RemoteHTTPURL           string                 `json:"remoteHTTPURL"`
	ForwardDomain           string                 `json:"forwardDomain"`
	ForwardNamesToAddresses []ForwardNameToAddress `json:"forwardNamesToAddresses"`
	ReverseDomain           string                 `json:"reverseDomain"`
	ReverseAddressesToNames []ReverseAddressToName `json:"reverseAddressesToNames"`
	MinTTLSeconds           uint32                 `json:"minTTLSeconds"`
	MaxTTLSeconds           uint32                 `json:"maxTTLSeconds"`
	MaxCacheSize            int                    `json:"maxCacheSize"`
	TimerIntervalSeconds    int                    `json:"timerIntervalSeconds"`
	MaxPurgesPerTimerPop    int                    `json:"maxPurgesPerTimerPop"`
}

// ReadConfiguration reads the DNS proxy configuration from a json file.
func ReadConfiguration(configFile string) (*Configuration, error) {
	log.Printf("reading config file %q", configFile)

	source, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config Configuration
	if err = json.Unmarshal(source, &config); err != nil {
		return nil, err
	}

	return &config, nil
}