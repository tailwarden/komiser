package k8s

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OpenCostResponse struct {
	Code   int                   `json:"code"`
	Status string                `json:"status"`
	Data   []map[string]costData `json:"data"`
}

type costData struct {
	Properties struct {
		Cluster string `json:"cluster"`
		Node    string `json:"node"`
	} `json:"properties"`
	Start              time.Time         `json:"start"`
	End                time.Time         `json:"end"`
	Minutes            int               `json:"minutes"`
	CPUCores           float64           `json:"cpuCores"`
	CPUCoreRequestAvg  float64           `json:"cpuCoreRequestAverage"`
	CPUCoreUsageAvg    float64           `json:"cpuCoreUsageAverage"`
	CPUCoreHours       float64           `json:"cpuCoreHours"`
	CPUCost            float64           `json:"cpuCost"`
	CPUCostAdjustment  float64           `json:"cpuCostAdjustment"`
	CPUEfficiency      float64           `json:"cpuEfficiency"`
	GPUCount           int               `json:"gpuCount"`
	GPUHours           int               `json:"gpuHours"`
	GPUCost            float64           `json:"gpuCost"`
	GPUCostAdjustment  float64           `json:"gpuCostAdjustment"`
	NetworkTransfer    int64             `json:"networkTransferBytes"`
	NetworkReceive     int64             `json:"networkReceiveBytes"`
	NetworkCost        float64           `json:"networkCost"`
	NetworkCrossZone   float64           `json:"networkCrossZoneCost"`
	NetworkCrossRegion float64           `json:"networkCrossRegionCost"`
	NetworkInternet    float64           `json:"networkInternetCost"`
	NetworkAdjustment  float64           `json:"networkCostAdjustment"`
	LoadBalancerCost   float64           `json:"loadBalancerCost"`
	LBCostAdjustment   float64           `json:"loadBalancerCostAdjustment"`
	PVBytes            int64             `json:"pvBytes"`
	PVByteHours        float64           `json:"pvByteHours"`
	PVCost             float64           `json:"pvCost"`
	PVs                map[string]pvInfo `json:"pvs"`
	PVCostAdjustment   float64           `json:"pvCostAdjustment"`
	RAMBytes           int64             `json:"ramBytes"`
	RAMByteRequestAvg  int64             `json:"ramByteRequestAverage"`
	RAMByteUsageAvg    int64             `json:"ramByteUsageAverage"`
	RAMByteHours       float64           `json:"ramByteHours"`
	RAMCost            float64           `json:"ramCost"`
	RAMCostAdjustment  float64           `json:"ramCostAdjustment"`
	RAMEfficiency      float64           `json:"ramEfficiency"`
	ExternalCost       float64           `json:"externalCost"`
	SharedCost         float64           `json:"sharedCost"`
	TotalCost          float64           `json:"totalCost"`
	TotalEfficiency    float64           `json:"totalEfficiency"`
}

type pvInfo struct {
	ByteHours float64 `json:"byteHours"`
	Cost      float64 `json:"cost"`
}

func GetOpencostInfo(aggregate string) (map[string]costData, error) {
	apiURL := fmt.Sprintf("http://%s:%d/allocation/compute", "127.0.0.1", 9003)

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Couldn't build request: %v", err)
	}

	query := req.URL.Query()
	query.Add("window", "month")
	query.Add("aggregate", aggregate)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Couldn't send request: %v", err)
	}
	defer resp.Body.Close()

	var ocResp OpenCostResponse
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("ERROR: Couldn't read response body: %v", err)
		}
		err = json.Unmarshal(bodyBytes, &ocResp)
		if err != nil {
			return nil, fmt.Errorf("ERROR: Couldn't unmarshal json: %v", err)
		}
	}

	return FlattenMapSlice(ocResp.Data), nil
}

func FlattenMapSlice(mapSlice []map[string]costData) map[string]costData {
	res := make(map[string]costData)
	for _, data := range mapSlice {
		for key, value := range data {
			res[key] = value
		}
	}
	return res
}
