package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	settings := getSettings()
	endPoints, _ := getEndpoints(settings)
	apiToken = settings.ApiToken

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(endPoints)
	})

	for _, command := range settings.Commands {
		commandRequestHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executeCommand(w, command.Command, command.Arguments)
		})
		http.Handle(command.Path, authorization(commandRequestHandler))
	}

	for _, proxyRequest := range settings.ProxyRequests {
		proxyRequestHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			makeRequest(w, proxyRequest.Url)
			fmt.Println(proxyRequest)
		})
		http.Handle(proxyRequest.Path, authorization(proxyRequestHandler))
	}

	if settings.Port == 0 {
		fmt.Printf("invalid port %v \n", settings.Port)
		os.Exit(2)
	}
	fmt.Printf("starting server at port %v\n", settings.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", settings.Port), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func makeRequest(w http.ResponseWriter, url string) {
	response, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(body)
}

func executeCommand(w http.ResponseWriter, command string, arguments []string) {
	result, err := exec.Command(command, arguments...).CombinedOutput()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(result)
}

func getEndpoints(settings Settings) ([]byte, error) {
	numberFfEndPoints := len(settings.Commands) + len(settings.ProxyRequests)
	endpoints := make([]ApiEndpoint, numberFfEndPoints)
	i := 0
	for _, command := range settings.Commands {
		endpoints[i] = command.GetEndPoint()
		i++
	}

	for _, proxyPass := range settings.ProxyRequests {
		endpoints[i] = proxyPass.GetEndPoint()
		i++
	}

	result, err := json.Marshal(endpoints)
	if err != nil {
		return nil, err
	}
	return result, nil
}
