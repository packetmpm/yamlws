package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var fin = flag.String("f", "-", "Input file ('-' for stdin)")

type Annotations struct {
	InitConfigs string `yaml:"rabbitmq.init_configs"`
	CheckNames  string `yaml:"rabbitmq.check_names"`
	Instances   string `yaml:"rabbitmq.instances"`
}

type Thing struct {
	Annotations Annotations `yaml:"podAnnotations"`
}

type Things struct {
	Foo Thing `yaml:"foo"`
	Bar Thing `yaml:"bar"`
}

func main() {
	flag.Parse()
	in := os.Stdin
	if *fin != "-" {
		var err error
		in, err = os.Open(*fin)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
	d := yaml.NewDecoder(in)
	ts := Things{}
	err := d.Decode(&ts)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println("Length Foo.Annotations.Instances is", len(ts.Foo.Annotations.Instances))
	fmt.Println("Length Boo.Annotations.Instances is", len(ts.Bar.Annotations.Instances))
}
