package main

import (
	"fmt"
	"io/ioutil"

	"github.com/charmbracelet/log"
	"github.com/suryatresna/pglogicalstream"
	"github.com/suryatresna/pglogicalstream/pkg/replication"
	"gopkg.in/yaml.v3"
)

func main() {

	var config pglogicalstream.Config
	yamlFile, err := ioutil.ReadFile("./example/simple/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	pgStream, err := pglogicalstream.NewPgStream(config, log.WithPrefix("pg-cdc"))
	if err != nil {
		panic(err)
	}

	pgStream.OnMessage(func(message replication.Wal2JsonChanges) {
		fmt.Println(message.Changes)
	})
}
