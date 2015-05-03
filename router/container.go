package router

import (
	"strings"
	"net/http"
    "encoding/json"

	"github.com/cosmos-io/cosmos/service"
	"github.com/cosmos-io/cosmos/util"
    
	"github.com/go-martini/martini"
)

func GetContainers(w http.ResponseWriter, r *http.Request, cosmos *service.CosmosService) {
    w.Header().Set("Content-Type", "application/json")
    
	containers, err := cosmos.GetContainers()
	if err != nil {
        res := map[string]string { "error": err.Error() }
        js, _ := json.Marshal(res)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(js)
		return
	}

    js, err := json.Marshal(containers)
    if err != nil {
        res := map[string]string { "error": err.Error() }
        js, _ := json.Marshal(res)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(js)
		return
    }
    
    w.Write(js)
}

func AddContainersOfPlanet(w http.ResponseWriter, params martini.Params, r *http.Request, cosmos *service.CosmosService) {
    w.Header().Set("Content-Type", "application/json")
    
	r.ParseForm()
	planet := params["planetName"]
	body, err := util.GetBodyFromRequest(r)

	if err != nil {
        res := map[string]string { "error": err.Error() }
        js, _ := json.Marshal(res)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(js)
		return
	}

	err = cosmos.AddContainersOfPlanet(planet, body)
	if err != nil {
        res := map[string]string { "error": err.Error() }
        js, _ := json.Marshal(res)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(js)
		return
	}

    w.Write([]byte(""))
}

func GetContainersOfPlanet(w http.ResponseWriter, params martini.Params, r *http.Request, cosmos *service.CosmosService) {
    w.Header().Set("Content-Type", "application/json")
    
	planet := params["planetName"]

	containers, err := cosmos.GetContainersOfPlanet(planet, true)
	if err != nil {
        res := map[string]string { "error": err.Error() }
        js, _ := json.Marshal(res)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(js)
		return
	}

    js, err := json.Marshal(containers)
    if err != nil {
        res := map[string]string { "error": err.Error() }
        js, _ := json.Marshal(res)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(js)
		return
    }

    w.Write(js)
}

func GetContainerMetrics(w http.ResponseWriter, params martini.Params, r *http.Request, cosmos *service.CosmosService) {
    w.Header().Set("Content-Type", "application/json")
    
	r.ParseForm()

	metric := strings.Split(util.GetQueryParam(r, "metric", "all"), ",")
	period := util.GetQueryParam(r, "period", "10m")

	planetName := params["planetName"]
	containerName := strings.Replace(params["containerName"], ".", "_", -1)

	metrics, err := cosmos.GetContainerMetrics(planetName, containerName, metric, period)
	if err != nil {
        res := map[string]string { "error": err.Error() }
        js, _ := json.Marshal(res)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(js)
		return
	}

    js, err := json.Marshal(metrics)
    if err != nil {
        res := map[string]string { "error": err.Error() }
        js, _ := json.Marshal(res)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(js)
		return
    }

    w.Write(js)
}
