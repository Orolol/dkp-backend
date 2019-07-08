package utils

import "github.com/google/uuid"

//GameManager list of current game.
type GameManager struct {
	GameList map[uuid.UUID]Game
}

//GameConf Game conf

//Game Game state
type Game struct {
	GameID      uuid.UUID
	State       string
	CurrentTurn int
	Player      PlayerInGame
}

//PlayerInGame player ig
type PlayerInGame struct {
	PlayerID           int
	Nick               string
	LastOrders         []PlayerLastOrders
	Modifiers          map[string]float32
	Logs               []PlayerLog
	CallbackEffects    []CallbackEffect
	PlayerInformations map[string]*PlayerInformation
}

type PlayerInformation struct {
	Type        string
	SubType     string
	Name        string
	Description string
	Value       float32
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
	Cooldown int
	Value    float32
	Effects  []Effect
	Costs    []Cost
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
	LowAlert     float32
	VeryLowAlert float32
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

//API type for action
type PlayerActionOrderApi struct {
	ID       string
	Value    float32
	PlayerID int
	GameID   uuid.UUID
}

//SQL type for Actions
type PlayerActionOrder struct {
	Name         string `gorm:"not null;unique"`
	ActionName   string
	Constraints  []Constraint
	Description  string
	Costs        []Cost
	Cooldown     int
	Effects      []Effect
	Type         string
	SubType      string
	Selector     string
	BaseValue    float32
	Restrictions []Restriction
}

type Effect struct {
	ModifierType string
	ModifierName string
	Operator     string
	Value        float32
	Target       string
	ActionName   string
	ToolTipValue float32
	Callbacks    []CallbackEffect
}

type CallbackEffect struct {
	Constraints []Constraint
	Effects     []Effect
}

type Cost struct {
	Type  string
	Value float32
}

type Constraint struct {
	Type     string
	Key      string
	Value    string
	Operator string
}
type Restriction struct {
	Type     string
	Key      string
	Value    string
	Operator string
}

type PlayerEvent struct {
	Type        string
	Description string
	Constraints []Constraint
	Effects     []Effect
	ActionName  string
	Name        string
	Weight      int
}

// //Constraint Json type for constraint
// type Constraint struct {
// 	Tech   []string `json:tech`
// 	Turn   int      `json:turn`
// 	Policy []string `json:policy`
// 	Action []string `json:action`
// }
