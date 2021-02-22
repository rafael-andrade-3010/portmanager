package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	portgrpc "portapi/proto"
	"portapi/service"
	"strconv"
)

var BATCH_SIZE = 10

func SetupRoutes() {
	http.HandleFunc("/client-api/port", handle)
}

func parseJson(reader io.Reader) ([]*portgrpc.Port, error) {
	ports := make([]*portgrpc.Port, 0)
	dec := json.NewDecoder(reader)
	_, err := dec.Token()
	if err != nil {
		return nil, err
	}
	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			return nil, err
		}
		var p *portgrpc.Port
		err = dec.Decode(&p)
		p.Key = key.(string)
		if err != nil {
			return nil, err
		}

		if len(ports) == BATCH_SIZE {
			log.Printf("flushing %v", ports)
			err := service.SavePorts(ports)
			if err != nil {
				return nil, err
			}
			ports = make([]*portgrpc.Port, 0)
		}

		ports = append(ports, p)
	}

	return ports, nil
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	ports, err := parseJson(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("%v", ports)
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	start, _ := strconv.Atoi(req.URL.Query().Get("start"))
	limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))
	ports, err := service.GetPorts(int32(start), int32(limit))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ports)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

func handle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		handleGet(w, req)
	case "POST":
		handlePost(w, req)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}
