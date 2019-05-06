package statistics

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

var bot *tgbotapi.BotAPI


func statistics() {
	start()
}


func start() {
	bott, err := tgbotapi.NewBotAPI(token)
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
			case "week":
			case "month":
			case "year":
			default:

			}
		} else {
			msg := update.Message.Text
			if strings.HasPrefix(msg,"找") {
				//是关键词, 开始记录
			}
		}
	}
}


func isKeyword(word string) {
	isold := checkWord(word)
	if isold {
		oldkeyWord(word)
	} else {
		newkeyWord(word)
	}
}

func checkWord(word string) bool {
	return false
}

func newkeyWord(word string) {

}

func oldkeyWord(word string) {

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






