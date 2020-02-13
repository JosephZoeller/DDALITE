package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"gopkg.in/yaml.v2"
)

//Proxy server + backend tcpServers
type Proxy struct {
	Proxyhost string
	Proxyport string
	Loadhost  string
	Loadport  string
	Loghost   string
	Logport   string
	Servers   []Server
}

//Server Hostname, Port
type Server struct {
	Hostname string
	Port     string
}

//CONFIG is yaml config
const CONFIG string = "./config/blandwallconf/config.yml"

func loadConfig(path string) ([]Server, error) {
	c := Proxy{}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(file), &c)
	if err != nil {
		panic(err)
	}

	s := c.Servers

	return s, nil

}

//Backend Servers are Blocked Ports Unless they have the hostname correct.
//Sets the rules for iptables
func process(c Server) error {
	port := fmt.Sprint(c.Port)
	//Template: sudo iptables -A INPUT -p tcp --dport 7777 -j DROP
	if err := iptables("-A", "INPUT", "-p", "tcp", "--dport", port, "-j", "DROP"); err != nil {
		return err
	}

	host := fmt.Sprint(c.Hostname)
	//Template: sudo iptables -A INPUT -p tcp -s localhost --dport 7777 -j ACCEPT
	if err := iptables("-I", "INPUT", "-s", host, "-p", "tcp", "--dport", port, "-j", "ACCEPT"); err != nil {
		return err
	}
	return nil
}

func iptables(args ...string) error {

	cmd := exec.Command("iptables", args...)

	out, err := cmd.Output()
	if err != nil {
		if bytes.Contains(out, []byte("This doesn't exist in IPTables :(")) {
			return nil
		}

		return err
	}

	return nil
}
func main() {
	path := CONFIG

	configs, err := loadConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range configs {

		if err := process(c); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Allowing %v,%v\nBlocking all other connections on Port:%v\n", c.Hostname, c.Port, c.Port)

	}

}
