package parser

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/promql"
)

func Parser() {
	// ルールファイルを読み込む
	ruleFile, err := ioutil.ReadFile("rules.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// ルールファイルをパースしてルールを取得する
	rules, err := promql.ParseRules(string(ruleFile), model.Now(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// ルールごとにメトリック名を取得する
	for _, rule := range rules {
		// ルールの式を取得する
		expr, err := rule.Expr()
		if err != nil {
			log.Fatal(err)
		}

		// ルールの式からメトリック名を取得する
		metrics := promql.AggregateMetrics(promql.Inspect(expr), nil)
		for _, metric := range metrics {
			fmt.Printf("Metric name: %s\n", metric.Name)
		}
	}
}
