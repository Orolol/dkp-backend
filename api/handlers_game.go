package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/orolol/dkp-backend/utils"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Welcome")

}

func ChangePolicy(c *gin.Context) {

}

func GetTranslations(c *gin.Context) {

	var translations []utils.Translation
	language := c.Param("language")
	translations = utils.GetTranslationsByLanguage(language)
	c.JSON(http.StatusOK, translations)
}

func GetInfos(c *gin.Context) {
	var translations *[]utils.DisplayInfoCat

	// translations = utils.GetInfos()
	c.JSON(http.StatusOK, translations)
}

func GetServerInfos(c *gin.Context) {
	var infos = utils.ServerInfos{
		IsOnline:       true,
		Region:         "EU",
		PlayersOnline:  1,
		PlayersWaiting: len(poolPendingPlayer),
		OnGoingGames:   len(onGoingGames),
	}
	fmt.Println("GO SERVER INFOS", infos)
	c.JSON(http.StatusOK, infos)
}

func GetHistory(c *gin.Context) {

	// var acc utils.Account
	// var accList []utils.Account
	// var list []utils.GameHistory
	// var apiList []utils.GameHistoryApi

	// db, err := gorm.Open("mysql", ConnexionString)

	// claims := jwt.ExtractClaims(c)
	// db.Where("Login = ?", claims["id"]).First(&acc)
	// db.Where(&acc).First(&acc)
	// db.Find(&accList)
	// fmt.Println(acc.ID)
	// db.Where("winner_id = ? OR loser_id = ?", acc.ID, acc.ID).Find(&list).Joins("JOIN accounts ON winner_id = accounts.ID OR loser_id = accounts.ID")
	// rows, err := db.Table("game_histories").Where("winner_id = ? OR loser_id = ?", acc.ID, acc.ID).Select("game_histories.created_at, game_histories.elo_diff, winner.Name, loser.Name").Joins("JOIN accounts as winner ON winner_id = winner.ID").Joins("JOIN accounts as loser ON loser_id = loser.ID").Rows()
	// fmt.Println("list", list)
	// fmt.Println("rows", rows)
	// fmt.Println("err", err)
	// for rows.Next() {
	// 	var apiHist utils.GameHistoryApi
	// 	rows.Scan(&apiHist.Created_at, &apiHist.ELODiff, &apiHist.WinnerNick, &apiHist.LoserNick)
	// 	fmt.Println(apiHist)
	// 	if apiHist.WinnerNick == acc.Name || apiHist.LoserNick == acc.Name {
	// 		apiList = append(apiList, apiHist)
	// 	}

	// }

	// // rows, err := db.Table("game_histories").Select("game_histories.created_at, game_histories.elo_diff, winner.Name, loser.Name").Where("game_histories.winner_id = ? OR game_histories.loser_id = ?", acc.ID, acc.ID).Joins("JOIN accounts as winner ON winner_id = winner.ID OR winner_id = 0").Joins("JOIN accounts as loser ON loser_id = loser.ID OR loser_id = 0").Rows()
	// // rows, err := db.Table("game_histories").Select("game_histories.created_at, game_histories.elo_diff, game_histories.winner_nick, game_histories.loser_id").Where("game_histories.winner_id = ? OR game_histories.loser_id = ?", acc.ID, acc.ID).Rows()
	// fmt.Println("list", list, rows, err)
	// for rows.Next() {
	// 	var apiHist utils.GameHistoryApi
	// 	rows.Scan(&apiHist.Created_at, &apiHist.ELODiff, &apiHist.WinnerNick, &apiHist.LoserNick)
	// 	fmt.Println(apiHist)
	// 	apiList = append(apiList, apiHist)
	// }
	// c.JSON(http.StatusOK, apiList)
}

func GetLeaderBoard(c *gin.Context) {

	// var accs []utils.Account
	// var accsApi []utils.AccountLeaderboardApi
	// db, _ := gorm.Open("mysql", ConnexionString)
	// db.Order("ELO desc, Name").Find(&accs)

	// for _, i := range accs {
	// 	accsApi = append(accsApi, utils.AccountLeaderboardApi{
	// 		Name: i.Name,
	// 		ELO:  i.ELO,
	// 	})

	// }
	// c.JSON(http.StatusOK, accsApi)
}

func Actions(c *gin.Context) {
	// var actionApi utils.PolicyChange
	// var action utils.PlayerActionOrder
	// var gMsg utils.GameMsg

	// c.ShouldBind(&actionApi)

	// fmt.Println(onGoingGames)
	// var isOkAction bool = true
	// if game, ok := onGoingGames[actionApi.GameID]; ok {
	// 	fmt.Println("GOT THE GAME")
	// 	action = utils.GetAction(actionApi.ID)
	// 	for players := range game.ListPlayers {
	// 		if game.ListPlayers[players].PlayerID == actionApi.PlayerID {
	// 			fmt.Println("GOT THE PLayer")
	// 			if len(game.ListPlayers[players].LastOrders) > 0 {
	// 				for actions := range game.ListPlayers[players].LastOrders {
	// 					if game.ListPlayers[players].LastOrders[actions].OrderID == action.ActionName {
	// 						if game.ListPlayers[players].LastOrders[actions].Cooldown > game.CurrentTurn {
	// 							fmt.Println("CD END ", game.ListPlayers[players].LastOrders[actions].Cooldown)
	// 							isOkAction = false
	// 						}
	// 					}
	// 				}
	// 			}
	// 			if !utils.CheckConstraint(&game.ListPlayers[players], action.Constraints, action.Costs, game, actionApi.Value) {
	// 				fmt.Println("CONSTRAINT FAIL")
	// 				isOkAction = false
	// 			} else {
	// 				fmt.Println("CONSTRAINT OK")
	// 			}
	// 		}
	// 	}

	// }
	// if isOkAction {
	// 	fmt.Println("OK ACTION")
	// 	gMsg.Action = action.ActionName
	// 	gMsg.GameID = actionApi.GameID
	// 	gMsg.PlayerID = actionApi.PlayerID
	// 	gMsg.Text = "Order"
	// 	gMsg.Costs = action.Costs

	// 	gMsg.Effects = action.Effects
	// 	if action.Selector == "range" {
	// 		fmt.Println("RANGE ACTION", actionApi.Value)
	// 		for i, e := range gMsg.Effects {
	// 			fmt.Println("APPLY VALUE ON EFFECT")
	// 			e.Value = actionApi.Value
	// 			gMsg.Effects[i] = e

	// 		}
	// 	}

	// 	gMsg.Cooldown = action.Cooldown
	// 	jsonMsg, err := json.Marshal(gMsg)
	// 	fmt.Println(string(jsonMsg))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(onGoingGames)
	// 	ZMQPusher.SendChan <- [][]byte{[]byte("MSG"), []byte(jsonMsg)}
	// } else {
	// 	fmt.Println("ACTION ON CD")
	// }

}

func LeaveQueue(c *gin.Context) {

	db, _ := gorm.Open("mysql", ConnexionString)
	var acc utils.Account

	claims := jwt.ExtractClaims(c)
	db.Where("Login = ?", claims["id"]).First(&acc)
	c.JSON(http.StatusOK, acc)
	leaveMatchmakingQueue <- acc

}

func JoinGame(c *gin.Context) {
	// db, _ := gorm.Open("mysql", ConnexionString)
	// var joinGame JoinGameApi
	// c.ShouldBind(&joinGame)

	// var acc utils.Account
	// claims := jwt.ExtractClaims(c)
	// db.Where("Login = ?", claims["id"]).First(&acc)
	// fmt.Println("current games", onGoingGames)
	// var isNewGame = true
	// for _, g := range onGoingGames {
	// 	for _, p := range g.ListPlayers {
	// 		if p.Nick == acc.Name {
	// 			isNewGame = false
	// 		}
	// 	}
	// }

	// if isNewGame {
	// 	fmt.Println("NEW GAME")
	// 	matchmakingQueue <- acc
	// }
	// c.JSON(http.StatusOK)
}

func JoinGameAi(c *gin.Context) {
	db, _ := gorm.Open("mysql", ConnexionString)
	var joinGame JoinGameApi
	c.ShouldBind(&joinGame)

	var acc utils.Account
	claims := jwt.ExtractClaims(c)
	db.Where("Login = ?", claims["id"]).First(&acc)

	acc.SelectedCountry = joinGame.SelectedCountry

	fmt.Println("current games", onGoingGames)
	// var isNewGame = true
	// for _, g := range onGoingGames {
	// 	for _, p := range g.ListPlayers {
	// 		if p.Nick == acc.Name {
	// 			isNewGame = false
	// 		}
	// 	}
	// }

	// var m = make(map[string]interface{})

	fmt.Println("NEW GAME")
	matchmakingAiQueue <- acc

	c.JSON(http.StatusOK, "OK")
}

type JoinGameApi struct {
	ID              int
	SelectedCountry string
}

func GetDungeons(c *gin.Context) {

	var dungeons = utils.GetDungeons()
	c.JSON(http.StatusOK, dungeons)

}

func StartDungeon(c *gin.Context) {
	var dungeonApi utils.StartDungeonAPI
	var gMsg utils.GameMsg
	c.ShouldBind(&dungeonApi)

	if _, ok := onGoingGames[dungeonApi.GameID]; ok {
		fmt.Println("GOT THE GAME")

		var vs []interface{}

		vs = append(vs, dungeonApi)

		fmt.Println("GOT THE PLayer")
		//TODO CHECK Players status
		gMsg.Action = "StartDungeon"
		gMsg.GameID = dungeonApi.GameID
		gMsg.PlayerID = dungeonApi.PlayerID
		gMsg.Text = "Order"
		gMsg.Values = vs

		jsonMsg, err := json.Marshal(gMsg)
		fmt.Println(string(jsonMsg))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(onGoingGames)
		ZMQPusher.SendChan <- [][]byte{[]byte("MSG"), []byte(jsonMsg)}

	}

}
