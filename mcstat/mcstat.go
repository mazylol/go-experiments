package mcstat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Ip    string `json:"ip"`
	Port  int    `json:"port"`
	Debug struct {
		Ping          bool `json:"ping"`
		Query         bool `json:"query"`
		Srv           bool `json:"srv"`
		Querymismatch bool `json:"querymismatch"`
		Ipinsrv       bool `json:"ipinsrv"`
		Cnameinsrv    bool `json:"cnameinsrv"`
		Animatedmotd  bool `json:"animatedmotd"`
		Cachetime     int  `json:"cachetime"`
		Apiversion    int  `json:"apiversion"`
		Dns           struct {
			Error struct {
				Srv struct {
					Hostname string `json:"hostname"`
					Message  string `json:"message"`
				} `json:"srv"`
			} `json:"error"`
			A []struct {
				Name     string `json:"name"`
				Type     string `json:"type"`
				Class    string `json:"class"`
				Ttl      int    `json:"ttl"`
				Rdlength int    `json:"rdlength"`
				Rdata    string `json:"rdata"`
				Cname    string `json:"cname,omitempty"`
				Address  string `json:"address,omitempty"`
			} `json:"a"`
		} `json:"dns"`
		Error struct {
			Query string `json:"query"`
		} `json:"error"`
	} `json:"debug"`
	Motd struct {
		Raw   []string `json:"raw"`
		Clean []string `json:"clean"`
		Html  []string `json:"html"`
	} `json:"motd"`
	Players struct {
		Online int `json:"online"`
		Max    int `json:"max"`
	} `json:"players"`
	Version  string `json:"version"`
	Online   bool   `json:"online"`
	Protocol int    `json:"protocol"`
	Hostname string `json:"hostname"`
	Icon     string `json:"icon"`
	Software string `json:"software"`
}

func Mcstat() {
	var srvurl string

	fmt.Print("Minecraft server URL: ")
	_, err := fmt.Scanln(&srvurl)
	if err != nil {
		return
	}

	url := "https://api.mcsrvstat.us/2/" + srvurl
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}

	serverJson := string(body)
	var server Data
	err = json.Unmarshal([]byte(serverJson), &server)
	if err != nil {
		return
	}
	fmt.Printf("Online: %t\n", server.Online)
	fmt.Printf("Hostname: %s\n", server.Hostname)
	fmt.Printf("IP: %s\n", server.Ip)
	fmt.Printf("Port: %d\n", server.Port)
	fmt.Printf("Version: %s\n", server.Version)
	fmt.Printf("Software: %s\n", server.Software)
	fmt.Printf("Players: %d of %d\n", server.Players.Online, server.Players.Max)
}
