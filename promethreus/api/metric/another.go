// package another
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// )
//
// func queryMetricValue(metricName string) (float64, error) {
// 	query := fmt.Sprintf("query=%s", url.QueryEscape(metricName))
// 	resp, err := http.Get(fmt.Sprintf("http://localhost:9090/api/v1/query?%s", query))
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer resp.Body.Close()
//
// 	var data QueryResult
// 	err = json.NewDecoder(resp.Body).Decode(&data)
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	if data.Status != "success" {
// 		return 0, fmt.Errorf("failed to query metric %s", metricName)
// 	}
//
// 	result := data.Data.Result
// 	if len(result) == 0 {
// 		return 0, fmt.Errorf("no data found for metric %s", metricName)
// 	}
//
// 	value := result[0].Value[1]
// 	floatValue, err := strconv.ParseFloat(value, 64)
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	return floatValue, nil
// }
//
// func main() {
// 	// /api/v1/label/__name__/valuesからすべてのメトリック名を取得するコード
//
// 	missingMetrics := make([]string, 0)
// 	for _, metric := range metrics {
// 		value, err := queryMetricValue(metric)
// 		if err != nil {
// 			missingMetrics = append(missingMetrics, metric)
// 		} else {
// 			fmt.Printf("%s: %f\n", metric, value)
// 		}
// 	}
//
// 	fmt.Printf("Missing metrics: %v\n", missingMetrics)
// }
