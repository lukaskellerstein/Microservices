package main

import (
	"fmt"
	"os"

	"github.com/opentracing/opentracing-go/log"
)

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := Init("01_Simpleasdfasdfasfasfasfasdfasdfasfdasdfasdfasdfasdfasfasdfasdf")
	defer closer.Close()

	helloTo := os.Args[1]

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", helloTo)

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	println(helloStr)
	span.LogKV("event", "println")

	span.Finish()
}
