package main

import (
	"log"

	"github.com/huobirdcenter/huobi_golang/cmd/accountclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/commonclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/crossmarginclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/etfclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/isolatedmarginclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/marketclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/orderclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/walletclientexample"
	"github.com/huobirdcenter/huobi_golang/logging/perflogger"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("conf/")
	// viper.AddConfigPath("conf2/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	log.Println(viper.GetString("secret"))
	// log.Println(viper.GetConfigPath())

	runAll()
}

// Run all examples
func runAll() {
	// commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	// orderclientexample.RunAllExamples()
	// marketclientexample.RunAllExamples()
	// isolatedmarginclientexample.RunAllExamples()
	// crossmarginclientexample.RunAllExamples()
	// walletclientexample.RunAllExamples()
	// etfclientexample.RunAllExamples()
	// marketwebsocketclientexample.RunAllExamples()
	// accountwebsocketclientexample.RunAllExamples()
	// orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
