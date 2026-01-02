package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

/*
<?xml version="1.0" encoding="utf-8"?>
<servers version="1">
    <server>
      <serverName>shanghai_VPN</serverName>
      <serverIP>127.0.0.1</serverIP>
    </server>
    <server>
      <serverName>shenzhen_VPN</serverName>
      <serverIP>127.0.0.2</serverIP>
    </server>
</servers>
*/

type Servers struct {
	Name    xml.Name "xml:\"servers\""
	Version string   "xml:\"version,attr\""
	Servers []Server "xml:\"server\""
}

type Server struct {
	ServerName string "xml:\"serverName\""
	ServerIP   string "xml:\"serverIP\""
}

func main() {

	data, err := ioutil.ReadFile("E:\\GO\\GOWORK\\src\\golandProject\\036XML\\config.xml")
	if err != nil {
		fmt.Printf("read err:%v\n", err)
		return
	}

	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Printf("unmarshal failed,err:%v\n", err)
		return
	}
	fmt.Printf("xml:%#v\n", servers)

}
