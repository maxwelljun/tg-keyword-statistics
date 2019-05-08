package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strconv"
	"strings"
)

var bot *tgbotapi.BotAPI
var TELEGRAM_TOKEN = "741657632:AAHvcdm18EySo9uJ-ftHXEHl2ISxZvnFbXI"

func main() {
	dbopen()
	args := os.Args
	if args != nil && len(args) == 2 {
		dbinit(args[1])
		TELEGRAM_TOKEN = args[1]
	}
	dbread()
	dbcron()
	start()
}

func start() {
	bott, err := tgbotapi.NewBotAPI(TELEGRAM_TOKEN)
	if err != nil {
		log.Panic(err)
	}
	bot = bott

	bot.Debug = true

	log.Printf("Authorized on account: %s  ID: %d", bot.Self.UserName, bot.Self.ID)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		//判断是否是服务群,如果不是则退出
		gid := update.Message.Chat.ID
		if gid > 0 && (gid != int64(667918518) && gid != int64(779814472) && gid != 731400898) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "本机器人不想跟你说话")
			sendmsg(msg)
			continue
		}

		if update.Message.IsCommand() {
			if update.Message.Chat.IsPrivate() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				order := update.Message.CommandArguments()
				num, numerr := strconv.Atoi(order)
				switch update.Message.Command() {
				case "start", "help":
					msg.Text = "本机器人专为羊毛群设计,用以统计大家都喜欢买啥东西"
					sendmsg(msg)
				case "hour":
					if numerr == nil {
						msg.Text = topKey("hour", num)
					} else {
						msg.Text = topKey("hour", 20)
					}
					sendmsg(msg)
				case "day":
					if numerr == nil {
						msg.Text = topKey("day", num)
					} else {
						msg.Text = topKey("day", 20)
					}
					sendmsg(msg)
				case "week":
					if numerr == nil {
						msg.Text = topKey("week", num)
					} else {
						msg.Text = topKey("week", 20)
					}
					sendmsg(msg)
				case "month":
					if numerr == nil {
						msg.Text = topKey("month", num)
					} else {
						msg.Text = topKey("month", 20)
					}
					sendmsg(msg)
				case "year":
					if numerr == nil {
						msg.Text = topKey("year", num)
					} else {
						msg.Text = topKey("year", 20)
					}
					sendmsg(msg)
				case "sum":
					if numerr == nil {
						msg.Text = topKey("sum", num)
					} else {
						msg.Text = topKey("sum", 20)
					}
					sendmsg(msg)
				default:

				}
			}
		} else {
			text := update.Message.Text
			if strings.HasPrefix(text, string('找')) {
				isKeyword(string([]rune(text)[1:]))
			}
		}
	}
}

func topKey(unit string, limit int) string {
	var kvs []kc
	sumquery := dbCountQuery(unit)
	if limit != 0 {
		kvs = dbFetchTopKeyword(unit, limit)
	} else {
		kvs = dbFetchTopKeyword(unit, 20)
	}
	var unitstr string
	switch unit {
	case "year":
		unitstr = "今年"
	case "month":
		unitstr = "这个月"
	case "week":
		unitstr = "这个周"
	case "day":
		unitstr = "今天"
	case "sum":
		unitstr = "这辈子"
	case "hour":
		unitstr = "这个小时"
	}
	str := unitstr + "查询总次数: " + strconv.Itoa(sumquery) + "次\r\n"
	for _, v := range kvs {
		str = str + "\r\n" + strconv.Itoa(v.count) + "\t" + v.keyword
	}
	return str
}

func isKeyword(word string) {
	if word != "" {
		isold := hasWord(word)
		if isold {
			oldkeyWord(word)
		} else {
			newkeyWord(word)
		}
	}
}

func hasWord(word string) bool {
	return dbHasKeyword(word)
}

func newkeyWord(word string) {
	dbNewKeyword(word)
}

func oldkeyWord(word string) {
	dbKeywordInc(word)
}

func checkAdmin(admins []tgbotapi.ChatMember, who tgbotapi.User) bool {
	for _, user := range admins {
		if who == *user.User {
			return true
		}
	}
	return false
}

func pushDayly() {
	msg := tgbotapi.NewMessage(667918518, "")
	msg.Text = topKey("month", 20)
	sendmsg(msg)
	msg.Text = topKey("week", 20)
	sendmsg(msg)
	msg.Text = topKey("day", 20)
	sendmsg(msg)
}

func pushHourly() {
	msg := tgbotapi.NewMessage(667918518, "")
	msg.Text = topKey("hour", 20)
	sendmsg(msg)
}

func sendmsg(msg tgbotapi.MessageConfig) tgbotapi.Message {
	if msg.Text == "" {
		msg.Text = "出现错误,请联系 @veezer"
	}
	mmsg, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
	return mmsg
}
