// Package docker
// @program: compose-yml
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-10-22 16:06
package docker

import "testing"

func TestUnmarshalYaml(t *testing.T) {
	const yaml = `
version: '3.6'

services:
  mongo:
    image: fairmarket.casdc.cn/dataspace/mongo:v1.0.0
    volumes:
      - "/mnt/data/fairman/DataSpace/mongo:/data/db"
  mysql:
    image: fairmarket.casdc.cn/dataspace/mysql:v1.0.0
    volumes:
      - type: "bind"
        source: "/mnt/data/fairman/DataSpace/mongo1"
        target: "/data/db1"
      - "/mnt/data/fairman/DataSpace/mongo:/data/db"
  api:
    image: fairmarket.casdc.cn/dataspace/api:v1.0.0
    ports:
      - target: 5193
        published: 5193
        protocol: tcp
      - target: 8080
        published: 8193
        protocol: tcp
    volumes:
      - "/mnt/data/fairman/DataSpace/disk:/data/disk"
      - "/mnt/data/fairman/DataSpace/ftpDir:/data/ftpDir"
    expose:
      - 5193
    environment:
      ACC_HOST: "127.0.0.1"
      ACC_PORT: "8193"
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

	t.Log(dockerCompose)

	marshalYaml, err := MarshalYaml(dockerCompose)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(marshalYaml)
}
