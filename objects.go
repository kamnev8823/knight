package knight

import "time"

// TODO Add comments

type isOk struct {
	Ok bool `json:"ok"`
}

type Error struct {
	Message string `json:"error"`
}

//AccountEmail https://lichess.org/api#operation/accountEmail
type AccountEmail struct {
	Email string `json:"email"`
}

//KidMode https://lichess.org/api#operation/accountKid
type KidMode struct {
	Kid bool `json:"kid"`
}

//UserStatus https://lichess.org/api#operation/apiUsersStatus
type UserStatus struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Online    bool   `json:"online"`
	Playing   bool   `json:"playing"`
	Streaming bool   `json:"streaming"`
	Patron    bool   `json:"patron"`
}

//Account https://lichess.org/api#tag/Account
type Account struct {
	Id             string       `json:"id"`
	Username       string       `json:"username"`
	Online         bool         `json:"online"`
	Perfs          AccountPerfs `json:"perfs"`
	CreatedAt      int64        `json:"createdAt"`
	Disabled       bool         `json:"disabled"`
	TosViolation   bool         `json:"tosViolation"`
	Profile        Profile      `json:"profile"`
	SeenAt         int64        `json:"seenAt"`
	Patron         bool         `json:"patron"`
	PlayTime       PlayTime     `json:"playTime"`
	Language       string       `json:"language"`
	Title          string       `json:"title"`
	Url            string       `json:"url"`
	Playing        string       `json:"playing"`
	NbFollowing    int          `json:"nbFollowing"`
	NbFollowers    int          `json:"nbFollowers"`
	CompletionRate int          `json:"completionRate"`
	Count          Count        `json:"count"`
	Streaming      bool         `json:"streaming"`
	Followable     bool         `json:"followable"`
	Following      bool         `json:"following"`
	Blocking       bool         `json:"blocking"`
	FollowsYou     bool         `json:"followsYou"`
}

//AccountPerfs https://lichess.org/api#tag/Account
type AccountPerfs struct {
	Chess960       PerfStat `json:"chess960"`
	Atomic         PerfStat `json:"atomic"`
	RacingKings    PerfStat `json:"racingKings"`
	UltraBullet    PerfStat `json:"ultraBullet"`
	Blitz          PerfStat `json:"blitz"`
	KingOfTheHill  PerfStat `json:"kingOfTheHill"`
	Bullet         PerfStat `json:"bullet"`
	Correspondence PerfStat `json:"correspondence"`
	Horde          PerfStat `json:"horde"`
	Puzzle         PerfStat `json:"puzzle"`
	Classical      PerfStat `json:"classical"`
	Rapid          PerfStat `json:"rapid"`
	Storm          PerfStat `json:"storm"`
}

//PerfStat https://lichess.org/api#tag/Account
type PerfStat struct {
	Games  int  `json:"games"`
	Rating int  `json:"rating"`
	Rd     int  `json:"rd"`
	Prog   int  `json:"prog"`
	Prov   bool `json:"prov"`
}

//Profile https://lichess.org/api#tag/Account
type Profile struct {
	Country    string `json:"country"`
	Location   string `json:"location"`
	Bio        string `json:"bio"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	FideRating int    `json:"fideRating"`
	UscfRating int    `json:"uscfRating"`
	EcfRating  int    `json:"ecfRating"`
	Links      string `json:"links"`
}

//PlayTime https://lichess.org/api#tag/Account
type PlayTime struct {
	Total int `json:"total"`
	Tv    int `json:"tv"`
}

//Count https://lichess.org/api#tag/Account
type Count struct {
	All         int `json:"all"`
	Rated       int `json:"rated"`
	Ai          int `json:"ai"`
	Draw        int `json:"draw"`
	DrawH       int `json:"drawH"`
	Loss        int `json:"loss"`
	LossH       int `json:"lossH"`
	Win         int `json:"win"`
	WinH        int `json:"winH"`
	Bookmark    int `json:"bookmark"`
	Playing     int `json:"playing"`
	ImportGames int `json:"import"`
	Me          int `json:"me"`
}

//Preferences https://lichess.org/api#operation/account
type Preferences struct {
	Dark          bool   `json:"dark"`
	Transp        bool   `json:"transp"`
	BgImg         string `json:"bgImg"`
	Is3d          bool   `json:"is3d"`
	Theme         string `json:"theme"`
	PieceSet      string `json:"pieceSet"`
	Theme3d       string `json:"theme3d"`
	PieceSet3d    string `json:"pieceSet3d"`
	SoundSet      string `json:"soundSet"`
	Blindfold     int    `json:"blindfold"`
	AutoQueen     int    `json:"autoQueen"`
	AutoThreefold int    `json:"autoThreefold"`
	Takeback      int    `json:"takeback"`
	Moretime      int    `json:"moretime"`
	ClockTenths   int    `json:"clockTenths"`
	ClockBar      bool   `json:"clockBar"`
	ClockSound    bool   `json:"clockSound"`
	Premove       bool   `json:"premove"`
	Animation     int    `json:"animation"`
	Captured      bool   `json:"captured"`
	Follow        bool   `json:"follow"`
	Highlight     bool   `json:"highlight"`
	Destination   bool   `json:"destination"`
	Coords        int    `json:"coords"`
	Replay        int    `json:"replay"`
	Challenge     int    `json:"challenge"`
	Message       int    `json:"message"`
	CoordColor    int    `json:"coordColor"`
	SubmitMove    int    `json:"submitMove"`
	ConfirmResign int    `json:"confirmResign"`
	InsightShare  int    `json:"insightShare"`
	KeyboardMove  int    `json:"keyboardMove"`
	Zen           int    `json:"zen"`
	MoveEvent     int    `json:"moveEvent"`
	RookCastle    int    `json:"rookCastle"`
}

//History https://lichess.org/api#operation/apiUserRatingHistory
type History struct {
	Name   string  `json:"name"`
	Points [][]int `json:"points"` // Format of an entry is [year, month, day, rating]
}

type Performance struct {
	Perf       ShortPerf `json:"perf"`
	Rank       int       `json:"rank"`
	Percentile float32   `json:"percentile"`
	Stat       Stat      `json:"stat"`
}

type ShortPerf struct {
	Glicko   Glicko `json:"glicko"`
	Nb       int    `json:"nb"`
	Progress int    `json:"progress"`
}

type Glicko struct {
	Rating      float32 `json:"rating"`
	Deviation   float32 `json:"deviation"`
	Provisional bool    `json:"provisional"`
}

type Stat struct {
	PerfType     PerfType       `json:"perfType"`
	Highest      ReachedScore   `json:"highest"`
	Lowest       ReachedScore   `json:"lowest"`
	BestWins     ReachedResults `json:"bestWins"`
	WorstLosses  ReachedResults `json:"worstLosses"`
	Count        ShortCount     `json:"count"`
	ResultStreak ResultStreak   `json:"resultStreak"`
	PlayStreak   PlayStreak     `json:"playStreak"`
}

type PerfType struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type ReachedScore struct {
	Int    int       `json:"int"`
	At     time.Time `json:"at"`
	GameId string    `json:"gameId"`
}

type ReachedResults struct {
	Results []ShortGameStat `json:"results"`
}

type ShortGameStat struct {
	OpInt  int       `json:"opInt"`
	OpId   OpId      `json:"opId"`
	At     time.Time `json:"at"`
	GameId string    `json:"gameId"`
}

type OpId struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type ShortCount struct {
	All         int     `json:"all"`
	Rated       int     `json:"rated"`
	Draw        int     `json:"draw"`
	Loss        int     `json:"loss"`
	Win         int     `json:"win"`
	Tour        int     `json:"tour"`
	Berserk     int     `json:"berserk"`
	OpAvg       float32 `json:"opAvg"`
	Seconds     int     `json:"seconds"`
	Disconnects int     `json:"disconnects"`
}

type ResultStreak struct {
	Win  Streak `json:"win"`
	Loss Streak `json:"loss"`
}

type Streak struct {
	Cur StreakStat `json:"cur"`
	Max StreakStat `json:"max"`
}

type StreakStat struct {
	V    int          `json:"v"`
	From StreakPeriod `json:"from"`
	To   StreakPeriod `json:"To"`
}

type StreakPeriod struct {
	At     time.Time `json:"at"`
	GameId string    `json:"gameId"`
}

type PlayStreak struct {
	Nb       Streak    `json:"nb"`
	Time     Streak    `json:"time"`
	LastDate time.Time `json:"lastDate"`
}

type Activity struct {
	Interval            Interval                    `json:"interval"`
	Games               GamesActivity               `json:"games"`
	Puzzles             PuzzlesActivity             `json:"puzzles"`
	Tournaments         TournamentsActivity         `json:"tournaments"`
	Practice            []Practice                  `json:"practice"`
	CorrespondenceMoves CorrespondenceMovesActivity `json:"correspondenceMoves"`
	CorrespondenceEnds  CorrespondenceEnds          `json:"correspondenceEnds"`
	Follows             Follows                     `json:"follows"`
	Teams               []Team                      `json:"teams"`
	Posts               []PostActivity              `json:"posts"`
}

type Interval struct {
	Start int `json:"start"` //timestamp
	End   int `json:"end"`   //timestamp
}

type GamesActivity struct {
	Chess960       Score `json:"chess960"`
	Atomic         Score `json:"atomic"`
	RacingKings    Score `json:"racingKings"`
	UltraBullet    Score `json:"ultraBullet"`
	Blitz          Score `json:"blitz"`
	KingOfTheHill  Score `json:"kingOfTheHill"`
	Bullet         Score `json:"bullet"`
	Correspondence Score `json:"correspondence"`
	Horde          Score `json:"horde"`
	Classical      Score `json:"classical"`
	Rapid          Score `json:"rapid"`
	Storm          Score `json:"storm"`
}

type Score struct {
	Win  int `json:"win"`
	Loss int `json:"loss"`
	Draw int `json:"draw"`
	Rp   Rp  `json:"rp"`
}

type Rp struct {
	Before int `json:"before"`
	After  int `json:"after"`
}

type PuzzlesActivity struct {
	Score Score `json:"score"`
}

type TournamentsActivity struct {
	Nb   int                       `json:"nb"`
	Best []BestTournamentsActivity `json:"best"`
}

type BestTournamentsActivity struct {
	Tournament  ShortTournamentInfo `json:"tournament"`
	NbGames     int                 `json:"nbGames"`
	Score       int                 `json:"score"`
	Rank        int                 `json:"rank"`
	RankPercent int                 `json:"rankPercent"`
}

//TODO maybe refact
type ShortTournamentInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Practice struct {
	Url         string `json:"url"`
	Name        string `json:"name"`
	NbPositions int    `json:"nbPositions"`
}

type CorrespondenceMovesActivity struct {
	Nb    int                           `json:"nb"`
	Games []CorrespondenceGamesActivity `json:"games"`
}

type CorrespondenceGamesActivity struct {
	Id       string   `json:"id"`
	Color    string   `json:"color"`
	Url      string   `json:"url"`
	Variant  string   `json:"variant"`
	Speed    string   `json:"speed"`
	Perf     string   `json:"perf"`
	Rated    bool     `json:"rated"`
	Opponent Opponent `json:"opponent"`
}

// TODO maybe need refactor
type Opponent struct {
	User   string `json:"user	"`
	Rating int    `json:"rating"`
}

type CorrespondenceEnds struct {
	Score Score                         `json:"score"`
	Games []CorrespondenceGamesActivity `json:"games"`
}

// TODO maybe need refactor
type Follows struct {
	In  FollowsIds `json:"in"`
	Out FollowsIds `json:"out"`
}

type FollowsIds struct {
	Ids []string `json:"ids"`
}

// TODO maybe need refactor
type Team struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type PostActivity struct {
	TopicUrl  string `json:"topicUrl"`
	TopicName string `json:"topicName"`
	Posts     []Post `json:"posts"`
}

type Post struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}

// TODO add to Account
type User struct {
	Id           string       `json:"id"`
	Username     string       `json:"username"`
	Online       bool         `json:"online"`
	Perfs        AccountPerfs `json:"perfs"`
	CreatedAt    int          `json:"createdAt"`
	Disabled     bool         `json:"disabled"`
	TosViolation bool         `json:"tosViolation"`
	Profile      Profile      `json:"profile"`
	SeenAt       int          `json:"seenAt"`
	Patron       bool         `json:"patron"`
	PlayTime     PlayTime     `json:"playTime"`
	Language     string       `json:"language"`
	Title        string       `json:"title"`
}

type Streamer struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Playing bool   `json:"playing"`
	Patron  bool   `json:"patron"`
}

//TODO refactor
type Crosstable struct {
	Users   map[string]float64 `json:"users"`
	NbGames int                `json:"nbGames"`
	Matchup CrosstableMatchup  `json:"matchup"`
}

type CrosstableMatchup struct {
	Users   map[string]float64 `json:"users"`
	NbGames int                `json:"nbGames"`
}

type Event struct {
	Type      string    `json:"type"`
	Game      Game      `json:"game"`
	Challenge Challenge `json:"challenge"`
}

type Challenge struct {
	Id          string      `json:"id"`
	Status      string      `json:"status"`
	Compat      Compat      `json:"compat"`
	Challenger  Challenger  `json:"challenger"`
	DestUser    DestUser    `json:"destUser"`
	Variant     Variant     `json:"variant"`
	Rated       bool        `json:"rated"`
	TimeControl TimeControl `json:"time_control"`
	Color       string      `json:"color"`
	Perf        Perf        `json:"perf"`
}

type Perf struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
}

type Game struct {
	Id     string `json:"id"`
	Compat Compat `json:"compat"`
}

type Compat struct {
	Bot   bool `json:"bot"`
	Board bool `json:"board"`
}

type Variant struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Short string `json:"short"`
}

type Challenger struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Rating int    `json:"rating"`
	Patron bool   `json:"patron"`
	Online bool   `json:"online"`
	Lag    int    `json:"lag"`
}

type DestUser struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Rating      int    `json:"rating"`
	Provisional bool   `json:"provisional"`
	Online      bool   `json:"online"`
	Lag         int    `json:"lag"`
}

type TimeControl struct {
	Type      string `json:"type"`
	Limit     int    `json:"limit"`
	Increment int    `json:"increment"`
	Show      string `json:"show"`
}

type TV struct {
	Bot           TVChannel `json:"Bot"`
	Blitz         TVChannel `json:"Blitz"`
	RacingKings   TVChannel `json:"Racing Kings"`
	UltraBullet   TVChannel `json:"UltraBullet"`
	Bullet        TVChannel `json:"Bullet"`
	Classical     TVChannel `json:"Classical"`
	ThreeCheck    TVChannel `json:"Three-check"`
	Antichess     TVChannel `json:"Antichess"`
	Computer      TVChannel `json:"Computer"`
	Horde         TVChannel `json:"Horde"`
	Atomic        TVChannel `json:"Atomic"`
	Crazyhouse    TVChannel `json:"Crazyhouse"`
	Chess960      TVChannel `json:"Chess960"`
	KingOfTheHill TVChannel `json:"King of the Hill"`
	TopRated      TVChannel `json:"Top Rated"`
}

type TVChannel struct {
	User   TVUser `json:"user"`
	Rating int    `json:"rating"`
	GameId string `json:"gameId"`
}

type TVUser struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Patron bool   `json:"patron"`
}

type TVStream struct {
	T string `json:"t"`
	D TVD    `json:"d"`
}

type TVD struct {
	Id          string      `json:"id"`
	Orientation string      `json:"orientation"`
	Players     [2]TVPlayer `json:"players"` // TODO if lichess add chess for three players, then to be an error
	Fen         string      `json:"fen"`
}

type TVPlayer struct {
	Color  string `json:"color"`
	User   TVUser `json:"user"`
	Rating int    `json:"rating"`
}

type TVBest struct {
	Id          string    `json:"id"`
	Rated       bool      `json:"rated"`
	Variant     string    `json:"variant"`
	Speed       string    `json:"speed"`
	Perf        string    `json:"perf"`
	CreatedAt   int       `json:"createdAt"`
	LastMoveAt  int       `json:"lastMoveAt"`
	Status      string    `json:"status"`
	Players     TVPlayers `json:"players"`
	InitialFen  string    `json:"initialFen"`
	Winner      string    `json:"winner"`
	Opening     Opening   `json:"opening"`
	Moves       string    `json:"moves"`
	Clock       Clock     `json:"clock"`
	Pgn         string    `json:"pgn"`
	DaysPerTurn string    `json:"daysPerTurn"`
	Analysis    []Analyze `json:"analysis"`
	Tournament  string    `json:"tournament"`
	Swiss       string    `json:"swiss"`
}

type TVPlayers struct {
	White TVOpponent `json:"white"`
	Black TVOpponent `json:"black"`
}

type TVOpponent struct {
	User       TVUser `json:"user"`
	Rating     int    `json:"rating"`
	RatingDiff int    `json:"ratingDiff"`
}

type Opening struct {
	Eco  string `json:"eco"`
	Name string `json:"name"`
	Ply  int    `json:"ply"`
}

type Clock struct {
	Initial   int `json:"initial"`
	Increment int `json:"increment"`
	TotalTime int `json:"totalTime"`
}

type Analyze struct {
	Eval      int             `json:"eval"`
	Best      string          `json:"best"`
	Variation string          `json:"variation"`
	Judgment  AnalyzeJudgment `json:"judgment"`
}

type AnalyzeJudgment struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}
