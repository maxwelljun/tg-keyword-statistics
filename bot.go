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
		if gid != 123456666 {
			//退出群组或者停止聊天
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start", "help":
				msg.Text = "本机器人专为羊毛群设计,用以统计大家都喜欢买啥东西"
				sendmsg(msg)
			case "day":
				msg.Text = topKey("day")
				sendmsg(msg)
			case "week":
				msg.Text = topKey("week")
				sendmsg(msg)
			case "month":
				msg.Text = topKey("month")
				sendmsg(msg)
			case "year":
				msg.Text = topKey("year")
				sendmsg(msg)
			default:

			}
		} else {
			text := update.Message.Text
			if strings.HasPrefix(text,string('找')) {
				isKeyword(string([]rune(text)[1:]))
			}
		}
	}
}


func topKey(unit string) string{
	kvs := dbFetchTopKeyword(unit, 20)
	str := ""
	for _,v := range kvs {
		str = str + "\r\n" +strconv.Itoa(v.count)  +"\t"+v.keyword
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
	for _,user := range admins{
		if who==*user.User {
			return true
		}
	}
	return false
}


func sendmsg(msg tgbotapi.MessageConfig) tgbotapi.Message {
	if msg.Text==""{
		msg.Text = "出现错误,请联系 @veezer"
	}
	mmsg, err := bot.Send(msg);
	if  err != nil {
		log.Println(err)
	}
	return mmsg
}






