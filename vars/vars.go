package vars

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type NmapStruct struct {
	Protocol     string   `json:"protocol"`
	Probename    string   `json:"probename"`
	Probestring  string   `json:"probestring"`
	Ports        []string `json:"ports"`
	Sslports     []string `json:"sslports"`
	Totalwaitms  string   `json:"totalwaitms"`
	Tcpwrappedms string   `json:"tcpwrappedms"`
	Rarity       string   `json:"rarity"`
	Fallback     string   `json:"fallback"`
	Matches      []struct {
		Pattern     string `json:"pattern"`
		Name        string `json:"name"`
		PatternFlag string `json:"pattern_flag"`
		Versioninfo struct {
			Cpename           string `json:"cpename"`
			Devicetype        string `json:"devicetype"`
			Hostname          string `json:"hostname"`
			Info              string `json:"info"`
			Operatingsystem   string `json:"operatingsystem"`
			Vendorproductname string `json:"vendorproductname"`
			Version           string `json:"version"`
		} `json:"versioninfo"`
	} `json:"matches"`
}

type NodeInformation struct {
	Ip          string
	Port        int
	Information string
}

var (
	ThreadNum = 10
	ConnResult    *sync.Map
	Host    string
	Port    = "22,23,53,80-139,445,8080"
	Mode    = ""
	Timeout = 4
	InformationResult map[string]NodeInformation
	NmapProbes        []NmapStruct
)


func InitNamp() (error){
	f, err := ioutil.ReadFile("vars/nmap.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(f, &NmapProbes)
	return err
}

func init() {
	ConnResult = &sync.Map{}
	InformationResult = make(map[string]NodeInformation)
}
