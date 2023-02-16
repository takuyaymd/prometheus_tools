package metric

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Metric() {
	// PrometheusのURLを指定
	url := "http://localhost:9090"

	// ラベル名を指定（ここでは__name__を指定している）
	labelName := "__name__"

	// /api/v1/label/__name__/valuesにリクエストを送信
	response, err := http.Get(fmt.Sprintf("%s/api/v1/label/%s/values", url, labelName))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// レスポンスの読み込み
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// レスポンスのJSONをパースして、値が取得できなかったメトリクスをリストアップ
	var labelValuesResponse struct {
		Data []string `json:"data"`
	}
	if err := json.Unmarshal(body, &labelValuesResponse); err != nil {
		panic(err)
	}

	// メトリクスの値を1つずつ確認し、値が取得できなかったメトリクスをリストアップ
	var failedMetrics []string
	for _, labelValue := range labelValuesResponse.Data {
		// /api/v1/queryにリクエストを送信
		queryResponse, err := http.Get(fmt.Sprintf("%s/api/v1/query?query=%s", url, labelValue))
		if err != nil {
			panic(err)
		}
		defer queryResponse.Body.Close()

		// レスポンスの読み込み
		body, err := ioutil.ReadAll(queryResponse.Body)
		if err != nil {
			panic(err)
		}

		// レスポンスのJSONをパースして、メトリクスの値を取得
		var queryResult struct {
			Status string `json:"status"`
			Data   struct {
				ResultType string `json:"resultType"`
				Result     []struct {
					Value []interface{} `json:"value"`
				} `json:"result"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body, &queryResult); err != nil {
			panic(err)
		}

		// メトリクスの値が取得できなかった場合、リストに追加
		if queryResult.Data.ResultType != "vector" || len(queryResult.Data.Result) == 0 {
			failedMetrics = append(failedMetrics, labelValue)
		}
	}

	// 値が取得できなかったメトリクスを出力
	fmt.Printf("Failed metrics: %v\n", failedMetrics)
}
