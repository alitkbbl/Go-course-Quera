package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Message struct {
	sender string
	text   string
}

type Player struct {
	name         string
	lowerName    string
	currentMapId int
	pointerToMap *Map
	receiveChan  chan string
	game         *Game
	lock         sync.Mutex
}

type Map struct {
	id          int
	mapChatChan chan Message
	players     map[string]*Player
	lock        sync.RWMutex
}

type Game struct {
	maps    map[int]*Map
	players map[string]*Player
	lock    sync.RWMutex
}

func NewGame(mapIds []int) (*Game, error) {
	maps := make(map[int]*Map)
	for _, v := range mapIds {
		if v <= 0 {
			return nil, errors.New("invalid mapId, must be > 0")
		}
		maps[v] = &Map{
			id:          v,
			mapChatChan: make(chan Message, 100),
			players:     make(map[string]*Player),
		}
		go maps[v].FanOutMessages()
	}
	return &Game{
		maps:    maps,
		players: make(map[string]*Player),
	}, nil
}

func (g *Game) ConnectPlayer(name string) error {
	lowerName := strings.ToLower(name)
	g.lock.Lock()
	defer g.lock.Unlock()

	if _, ok := g.players[lowerName]; ok {
		return errors.New("player already exists")
	}
	receiveChan := make(chan string, 100)

	player := &Player{
		name:         formatName(name),
		lowerName:    lowerName,
		receiveChan:  receiveChan,
		game:         g,
		currentMapId: 0,
		pointerToMap: nil,
	}
	g.players[lowerName] = player
	return nil
}

func (g *Game) SwitchPlayerMap(name string, mapId int) error {
	lowerName := strings.ToLower(name)

	g.lock.Lock()
	defer g.lock.Unlock()

	player, ok := g.players[lowerName]
	if !ok {
		return errors.New("player not found")
	}
	if player.currentMapId == mapId {
		return errors.New("player already in this map")
	}

	newMap, ok := g.maps[mapId]
	if !ok {
		return errors.New("map not found")
	}

	if player.pointerToMap != nil {
		player.pointerToMap.lock.Lock()
		delete(player.pointerToMap.players, lowerName)
		player.pointerToMap.lock.Unlock()
	}

	newMap.lock.Lock()
	newMap.players[lowerName] = player
	newMap.lock.Unlock()

	player.lock.Lock()
	player.currentMapId = mapId
	player.pointerToMap = newMap
	player.lock.Unlock()

	return nil
}

func (g *Game) GetPlayer(name string) (*Player, error) {
	lowerName := strings.ToLower(name)

	g.lock.RLock()
	defer g.lock.RUnlock()

	player, ok := g.players[lowerName]
	if !ok {
		return nil, errors.New("player not found")
	}
	return player, nil
}

func (g *Game) GetMap(mapId int) (*Map, error) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	m, ok := g.maps[mapId]
	if !ok {
		return nil, errors.New("map not found")
	}
	return m, nil
}

func (m *Map) FanOutMessages() {
	for {
		msg := <-m.mapChatChan

		m.lock.RLock()
		for key, player := range m.players {
			if key == msg.sender {
				continue
			}
			select {
			case player.receiveChan <- msg.text:
			default:

			}
		}
		m.lock.RUnlock()
	}
}

func (p *Player) GetChannel() <-chan string {
	return p.receiveChan
}

func (p *Player) SendMessage(msg string) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.currentMapId == 0 || p.pointerToMap == nil {
		return errors.New("player is not in any map")
	}
	formattedMsg := fmt.Sprintf("%s says: %s", p.GetName(), msg)

	rawMsg := Message{
		sender: p.lowerName,
		text:   formattedMsg,
	}
	select {
	case p.pointerToMap.mapChatChan <- rawMsg:
		return nil
	default:
		return errors.New("failed to send message")
	}
}

func formatName(name string) string {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return ""
	}
	return strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
}

func (p *Player) GetName() string {
	return p.name
}
