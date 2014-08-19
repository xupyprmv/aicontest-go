package structs

type BotRich struct {
	Bot
	CMP_BOTNAME  string
	CMP_USERNAME string
	CMP_WINS     int
	CMP_TIES     int
	CMP_LOSES    int
	CMP_GAMES    int
	CMP_RATING   float32
}
