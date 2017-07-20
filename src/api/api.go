package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type Model struct {
	Players map[string]Player
	Ch      chan int64
}

type Player struct {
	ID        string
	Counter   int64
	SumValue  int64
	LastValue int64
}

type Config struct {
	ELK_HOST string `yaml:"ELK_HOST"`
}

func (model *Model) Update(_ http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	Elkhost := elkHost()

	param := <-model.Ch
	if playerr, ok := model.Players[id]; ok {
		param = playerr.LastValue
	}

	response, err := http.Get(Elkhost + "?n=" + strconv.Itoa(int(param)))
	if err != nil {
		// Missing err log
		return
	}
	defer response.Body.Close()
	// Skipping err check
	body, _ := ioutil.ReadAll(response.Body)
	value, _ := strconv.ParseInt(string(body), 10, 64)

	if player, ok := model.Players[id]; ok {
		model.Players[id] = Player{ID: id, Counter: player.Counter + 1, SumValue: player.SumValue + value, LastValue: value}
	} else {
		model.Players[id] = Player{ID: id, Counter: 1, SumValue: value, LastValue: value}
	}
	fmt.Println("Model:", model)
	model.Ch <- 0
}

func (model *Model) GetStatistic(_ http.ResponseWriter, _ *http.Request) {
	totalRequests := int64(0)
	for id, player := range model.Players {
		totalRequests += player.Counter
		fmt.Println("Player", id)
		fmt.Println("Requests", player.Counter, "|", "SumValue", player.SumValue)
	}
	fmt.Println("-- Total requests", totalRequests)
}

func elkHost() string {
	ymlFile, err := ioutil.ReadFile("src/application.yml")
	var config Config
	if err == nil {
		err = yaml.Unmarshal(ymlFile, &config)
		if err != nil {
			panic(err)
		}
		return config.ELK_HOST
	}
	panic(err)
}
