// Package docker
// @program: compose-yml
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-10-22 16:06
package docker

import (
	"fmt"
	"strings"
	"testing"
)

func TestUnmarshalYaml(t *testing.T) {
	const yaml = `
version: '3.6'

services:
  mongo:
    image: fairmarket.casdc.cn/dataspace/mongo:v1.0.0
    volumes:
      - type: "bind"
        source: "/mnt/data/fairman/DataSpace/mongo"
        target: "/data/db"
        desc: "mongo1"
  mysql:
    image: fairmarket.casdc.cn/dataspace/mysql:v1.0.0
    volumes:
      - type: "bind"
        source: "/mnt/data/fairman/DataSpace/mongo1"
        target: "/data/db1"
        desc: "mongo1"
      - type: "bind"
        source: "/mnt/data/fairman/DataSpace/mongo"
        target: "/data/db"
        desc: "mongo"
  api:
    image: fairmarket.casdc.cn/dataspace/api:v1.0.0
    ports:
      - target: 5193-10000
        published: 5193-10000
        protocol: tcp
        desc: "后台端口"
      - target: 8080
        published: 8193
        protocol: tcp
        desc: "首页端口"
    volumes:
      - type: "bind"
        source: "/mnt/data/fairman/DataSpace/disk"
        target: "/data/disk"
        desc: "静态文件"
      - type: "bind"
        source: "/mnt/data/fairman/DataSpace/ftpDir"
        target: "/data/ftpDir"
        desc: "ftp文件"
    environment:
      - key: "ACC_HOST"
        val: "127.0.0.1"
        desc: "首页IP"
      - key: "ACC_PORT"
        val: "8193"
        desc: "首页端口"
    depends_on:
      - mongo
      - mysql
`

	var dockerCompose Yml
	err := UnmarshalYaml(yaml, &dockerCompose)
	if err != nil {
		t.Error(err)
		return
	}

	a := "123:345"
	split := strings.Split(a, ":")
	fmt.Println(split)

	for _, service := range dockerCompose.Services {
		var aaa = "111111111"
		service.HostName = aaa

		service.Ports = append(service.Ports, Port{
			Target:    "1111",
			Published: "1111",
			Protocol:  "tcp",
		})
	}

	fmt.Println(dockerCompose.Services["api"].HostName)
	for _, port := range dockerCompose.Services["api"].Ports {
		fmt.Println(port)
	}

	t.Log(dockerCompose)

	marshalYaml, err := MarshalYaml(dockerCompose)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(marshalYaml)
}
