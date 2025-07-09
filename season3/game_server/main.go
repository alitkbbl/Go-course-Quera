package main

import "sync"

type Player struct {
	name         string
	currentMapId int
	receiveChan  chan string
	game         *Game
	lock         sync.Mutex
}

type Map struct {
	id          int
	mapChatChan chan string
	players     map[string]*Player
	lock        sync.RWMutex
}

type Game struct {
	maps    map[string]*Map
	players map[string]*Player
	lock    sync.RWMutex
}

func NewGame(mapIds []int) (*Game, error) {
	return nil, nil
}

func (g *Game) ConnectPlayer(name string) error {
	return nil
}

func (g *Game) SwitchPlayerMap(name string, mapId int) error {
	return nil
}

func (g *Game) GetPlayer(name string) (*Player, error) {
	return nil, nil
}

func (g *Game) GetMap(mapId int) (*Map, error) {
	return nil, nil
}

func (m *Map) FanOutMessages() {
}

func (p *Player) GetChannel() <-chan string {
	return nil
}

func (p *Player) SendMessage(msg string) error {
	return nil
}

func (p *Player) GetName() string {
	return ""
}
