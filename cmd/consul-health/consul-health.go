package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

// Sarama configuration options
var (
	host    = "127.0.0.1"
	port    = 8500
	ssl     = false
	service = "consul"
)

func init() {

	flag.StringVar(&host, "host", "127.0.0.1", "Consul Api Server. The default value is 127.0.0.1")
	flag.BoolVar(&ssl, "ssl", false, "Use Https. The default value is false")
	flag.IntVar(&port, "port", 8500, "Consul Api Port")
	flag.StringVar(&service, "service", "node-exporter", "Consul Service")
	flag.Parse()

}
func main() {

	address := fmt.Sprintf("%s:%d", host, port)

	var scheme string

	if ssl {
		scheme = "https"
	} else {
		scheme = "http"
	}

	config := &api.Config{
		Address:    address,
		Scheme:     scheme,
		Datacenter: "dc1",
	}
	// Get a new client
	client, err := api.NewClient(config)

	if err != nil {
		panic(err)
	}
	health := client.Health()
	checks, meta, err := health.Service(service, "", false, nil)

	if err != nil {
		panic(err)
	}
	if meta.LastIndex == 0 {
		log.Fatalf("bad: %v", meta)
	}
	if len(checks) == 0 {
		log.Fatalf("Bad: %v", checks)
	}
	if _, ok := checks[0].Node.TaggedAddresses["wan"]; !ok {
		log.Fatalf("Bad: %v", checks[0].Node)
	}
	if checks[0].Node.Datacenter != "dc1" {
		log.Fatalf("Bad datacenter: %v", checks[0].Node)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "IP", "Port", "Tags", "Meta"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	for _, item := range checks {
		v := []string{item.Service.ID, item.Service.Address, fmt.Sprintf("%d", item.Service.Port), fmt.Sprintf("%s", item.Service.Tags), fmt.Sprintf("%s", item.Service.Meta)}
		table.Append(v)

	}

	table.Render()
}
