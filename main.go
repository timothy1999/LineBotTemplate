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
				
 				}else if message.Text=="幹"{
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("欲求不滿的自己回家撸")).Do(); err != nil {
					log.Print(err)
					}
 				}else if message.Text=="你認識小雷嗎"||message.Text=="你認識小雷嗎?"||message.Text=="抽"{
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你除了抽你還會幹嘛，我比他強多了")).Do(); err != nil {
					log.Print(err)
					}
 				}else{
 					if startedd{
 						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("考完北模還不去喝可樂")).Do(); err != nil {
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
