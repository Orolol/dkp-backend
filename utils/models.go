package utils

import "github.com/google/uuid"

//GameManager list of current game.
type GameManager struct {
	GameList map[uuid.UUID]Game
}

//GameConf Game conf

//Game Game state
type Game struct {
	GameID             uuid.UUID
	State              string
	CurrentTurn        float64
	Phase              int
	Hour               int
	Minute             int
	Player             PlayerInGame
	PendingEvent       []PlayerEvent
	PendingQuest       []PlayerQuest
	OngoingQuest       []PlayerQuest
	OngoingDungeonRuns []DungeonRun
}

//PlayerInGame player ig

// type Status struct {
// 	Description map[string]Translation
// 	TimeSpent   int
// 	Name        string
// }

type PlayerInformation struct {
	Type        string
	SubType     string
	Name        string
	Description string
	Value       float64
}

type PlayerLog struct {
	Turn       int
	ActionName string
}

//PlayerModifier modifeirs

type PlayerLastOrders struct {
	OrderID  string
	Cooldown int
}

//GameMsg msg send to the routine
type GameMsg struct {
	GameID   uuid.UUID
	Text     string
	PlayerID int
	Action   string
	Type     string
	Values   []interface{}
}

type Translation struct {
	ActionName  string
	Language    string
	ShortName   string
	LongName    string
	Description string
	ToolTip     string
}

type DisplayInfoCat struct {
	Category string
	Infos    []DisplayInfos
}

type DisplayInfos struct {
	Name         string
	Type         string
	LowAlert     float64
	VeryLowAlert float64
	GrowthName   string
	GrowthType   string
	Display      int
}
type ServerInfos struct {
	Region         string
	PlayersOnline  int
	PlayersWaiting int
	IsOnline       bool
	OnGoingGames   int
}

type News struct {
	Title     string
	Date      string
	Paragrahs []string
}

type Configuration struct {
	Connection_String string
}

type PlayerQuest struct {
	From        OtherPlayer
	Description map[string]Translation
	Name        string
	Constraints []Constraint
	Success     []Effect
	Failure     []Effect
	Ignore      []Effect
}
