// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"strings"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var startedd bool = false
func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text=="go"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我只是想說說話阿，你不知道這很難寫嗎")).Do(); err != nil {
					log.Print(err)
				
				}
				if startedd{
					startedd=false
				
				}else{
					startedd=true
				
				}
				
 				}else if message.Text=="幹"||strings.Contains(message.Text,"幹你")||strings.Contains(message.Text,"幹幹"){
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("好喔"),linebot.NewImageMessage("https://fsp.youthwant.com/images/580_785386238269148.jpg","https://fsp.youthwant.com/images/580_785386238269148.jpg")).Do(); err != nil {
					log.Print(err)
					}
 				}else if strings.Contains(message.Text,"造物者")||strings.Contains(message.Text, "造化者")||strings.Contains(message.Text, "天帝")||strings.Contains(message.Text, "萬王之王")||strings.Contains(message.Text, "烏醋")||strings.Contains(message.Text, "強胃散")||strings.Contains(message.Text, "殺")||strings.Contains(message.Text, "爆屁")||strings.Contains(message.Text, "莊園"){
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("想清除體內負能量?"),
 						linebot.NewImageMessage("http://www.ghostisland.com.tw/upload/2014/11/20/file_546ccc179b288.JPG","http://www.ghostisland.com.tw/upload/2014/11/20/file_546ccc179b288.JPG"),
 						linebot.NewTextMessage("我很遺憾") ).Do(); err != nil {
					log.Print(err)
					}
 				}else if message.Text=="尻"||strings.Contains(message.Text,"尻尻"){
 					if _, err = bot.ReplyMessage(event.ReplyToken,
 					linebot.NewTextMessage("聽說阿麵禁尻考爆好，還不+1")).Do(); err != nil{
					log.Print(err)
					}
 				}else if strings.Contains(message.Text,"昶"){

 //					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("想體驗永日火暴!?可是我還不會選照片耶，你要推薦嗎")).Do(); err != nil {
//=======
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("想體驗永日火暴!?可是我還沒選照片耶，你要推薦嗎"),linebot.NewImageMessage("https://scontent-tpe1-1.xx.fbcdn.net/v/t1.0-9/10402927_1555832194689772_5563271631371161517_n.jpg?oh=354c07e94fcb41da0cb25e843cc96fe4&oe=5A4F8B79","https://scontent-tpe1-1.xx.fbcdn.net/v/t1.0-9/10402927_1555832194689772_5563271631371161517_n.jpg?oh=354c07e94fcb41da0cb25e843cc96fe4&oe=5A4F8B79")).Do(); err != nil {
					log.Print(err)
					}
 				}else if strings.Contains(message.Text,"早安"){
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("早安，特別提醒你 早上一瓶可樂，可以維持你一整天的精神喔")).Do(); err != nil {
					log.Print(err)
					}
 				}else if strings.Contains(message.Text,"yee")||strings.Contains(message.Text,"Yee"){
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("都幾歲了，還yee")).Do(); err != nil {
					log.Print(err)
					}
 				}else if message.Text=="你認識小雷嗎"||message.Text=="你認識小雷嗎?"||message.Text=="你認識小雷嗎？"||strings.Contains(message.Text,"小雷"){
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("它跟俊能一樣肥，我比他強多了!")).Do(); err != nil {
					log.Print(err)
					}
 				}else if message.Text=="抽"||message.Text=="抽卡"{
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你除了抽你還會幹嘛?")).Do(); err != nil {
					log.Print(err)
					}
 				}else{
 					if startedd{
 						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("喔是喔邱俊熊")).Do(); err != nil {
							log.Print(err)
				
						}

 					}
 				}
			case *linebot.ImageMessage:
				if startedd{
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("胖")).Do(); err != nil {
					log.Print(err)
				
				}
				}
			}

			
		}
	}
}
