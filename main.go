package main

import (
	"github.com/takuyaymd/prometheus_tools/promethreus/api/metric"
	"github.com/takuyaymd/prometheus_tools/promethreus/rules/parser"
)

api/check/check

func main() {
	// metricパッケージの関数を呼び出す
	metric.Metric()
	parser.Parser()
}
