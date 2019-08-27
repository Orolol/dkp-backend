package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/orolol/dkp-backend/utils"

	"github.com/google/uuid"
	"github.com/zeromq/goczmq"

	"github.com/mitchellh/mapstructure"
)

func createGame() utils.Game {

	var game = utils.GenerateNewGame()
	return game
}

func GameEvent(queue chan utils.GameMsg, game *utils.Game, player *utils.PlayerInGame) {

	for msg := range queue {
		// var p utils.PlayerInGame

		// p = game.Player
		//Apply Effects, cost and special action

		fmt.Println("RECIEVE MSG", msg)

		switch a := msg.Action; a {
		case "StartDungeon":
			var result utils.StartDungeonAPI
			mapstructure.Decode(msg.Values[0], &result)
			utils.StartDungeonFromAPI(result, game)
		}

	}

}

func runGame(game utils.Game, queue chan utils.GameMsg, queueGameOut chan utils.Game) {
	player := &game.Player
	go GameEvent(queue, &game, player)

	var minuteCount = 5

	//fmt.Println("Start game ", player.Nick, " vs ", player2.Nick)
	game.State = "Running"
	queueGameOut <- game
	time.Sleep(time.Second)

	for game.CurrentTurn < 9999 {

		timer1 := time.NewTimer(time.Second * 1)
		game.CurrentTurn++
		if game.Hour < 24 && game.Minute+minuteCount >= 60 {
			game.Minute = game.Minute + minuteCount - 60
			game.Hour += 1
		} else if game.Hour < 24 {
			game.Minute += minuteCount
		} else {
			game.Hour = 0
			game.Minute += minuteCount
		}

		//TEST DUNGEON
		// if game.Hour == 15 && len(game.OngoingDungeonRuns) == 0 {
		// 	var party []*utils.OtherPlayer
		// 	for _, x := range game.Player.Guild.Guildies {
		// 		party = append(party, &x.Player)
		// 	}
		// 	d := utils.GetDungeon("TestDungeon")
		// 	utils.InitDungeon(party, &game, d)
		// }

		fmt.Println("Turn : ", game.CurrentTurn, "Hour : ", game.Hour)

		for i, _ := range game.Player.Guild.Guildies {
			player := &game.Player.Guild.Guildies[i].Player

			utils.PlayerSimulation(player, &game)
			utils.AlgoRollTurnEvent(player, &game)
		}
		for i, _ := range game.OngoingDungeonRuns {
			d := &game.OngoingDungeonRuns[i]
			d.Run(&game)
		}
		fmt.Println("END OF TURN")
		<-timer1.C
		queueGameOut <- game

	}

}

func GameManagerF(queueGameOut chan utils.Game, queueCreation chan [][]byte) {

	var GameList = make(map[uuid.UUID]chan utils.GameMsg)

	for msg := range queueCreation {
		switch string(msg[1]) {
		case "CREATE":

			queueGameInc := make(chan utils.GameMsg, 100)
			game := createGame()
			go runGame(game, queueGameInc, queueGameOut)
			GameList[game.GameID] = queueGameInc
		case "MSG":
			var gs utils.GameMsg
			json.Unmarshal(msg[2], &gs)
			if val, ok := GameList[gs.GameID]; ok {
				val <- gs
			}

		}
	}

}

func ZMQReader(queueCreation chan [][]byte) {
	//fmt.Printf("Init Reader")
	pull := goczmq.NewRouterChanneler("tcp://127.0.0.1:31337")
	for msg := range pull.RecvChan {
		//fmt.Println("Recieving new game msg in ZMQ !! TYPE : ", string(msg[1]))
		queueCreation <- msg
	}
}
func ZMQPusher() *goczmq.Channeler {
	//fmt.Printf("Init Pusher")
	push := goczmq.NewDealerChanneler("tcp://127.0.0.1:31338")

	return push
}

func FromChanToZMQ(queue chan utils.Game) {
	pushSock := ZMQPusher()
	for msg := range queue {
		//fmt.Println("Read Game from Queue and send to ZMQ")
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			//
			//fmt.Println(err)
		}

		pushSock.SendChan <- [][]byte{[]byte(msg.GameID.String()), []byte(jsonMsg)}
		//fmt.Println("SENT : ", msg)
	}
}

func main() {
	utils.SetBaseValueTags()
	utils.SetBaseValueActions()
	utils.SetBaseValueEvents()
	utils.SetBaseValueDungeons()
	utils.SetBaseValueClass()
	//fmt.Printf("Enter Main")
	queueGameOut := make(chan utils.Game)
	queueGameIn := make(chan [][]byte)

	go GameManagerF(queueGameOut, queueGameIn)
	go ZMQReader(queueGameIn)
	go FromChanToZMQ(queueGameOut)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
