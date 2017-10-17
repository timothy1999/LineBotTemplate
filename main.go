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
	"time"
	"math/rand"
	"github.com/line/line-bot-sdk-go/linebot"
)

var story [3]string
int storynum = 3;

var bot *linebot.Client
var startedd bool = false
func main() {
	story[0] = "有一天YEH表走在路上，他遇到一個在裝弱的人，葉表不開心，心想自己才是最會裝弱的人。上前想去了解那個人到底是誰。突然，看到那個人擺出一個醜臉和一支中指。YEH表心想:我輸了。原來他在照鏡子啊！";
	story[1] = "從前從前，有個飛行義大利麵怪物在一次嚴重酒醉後創造了宇宙。它的朋友海盜們是「絕對神性的體現者」，負面形象都是神學家抹黑；死後的天堂裡有一座噴啤酒的火山和一間脫衣舞俱樂部，地獄和天堂差不多，不過啤酒是過期的，脫衣舞孃都有性病。";
	story[2] = "莊圓是站着跳舞而拿斧頭的唯一的人。他身材不太高大；青白臉色，皺紋間時常夾些傷痕；一部亂蓬蓬的花白的鬍子。拿的雖然是斧頭，可是又鈍又破，似乎十多年沒有磨，也沒有洗。他對人說話，總是滿口謗佛者死，教人半懂不懂的。因爲他姓莊，別人便從描紅紙上的「美男子佛陀莊圓」這半懂不懂的話裏，替他取下一個綽號，叫作莊圓大師。莊圓一到店，所有精怪便都看着他笑，有的叫道，「莊圓，你信徒又喝強胃散了！」他不回答，對櫃裏說，「溫兩碗烏醋，要一碟強胃散。」便排出九文大錢。他們又故意的高聲嚷道，「你一定又粉碎了人家的靈體了！」莊圓睜大眼睛說，「你怎麼這樣憑空汚人清白……」「什麼清白？我前天親眼見你爆了何家的屁，弔着打。」莊圓便漲紅了臉，額上的青筋條條綻出，爭辯道，「爆屁不能算die……爆屁！……佛陀的事，能算die麼？」接連便是難懂的話，什麼「清除體內負能量」，什麼「重新大翻新」之類，引得衆人都鬨笑起來：店內外充滿了快活的空氣。 （感謝duang提供）";
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
 						linebot.NewImageMessage("https://truth.bahamut.com.tw/s01/201307/3e55f59e1d2e01285be1a9ebe7a83401.JPG","https://truth.bahamut.com.tw/s01/201307/3e55f59e1d2e01285be1a9ebe7a83401.JPG"),
 						linebot.NewTextMessage("我很遺憾") ).Do(); err != nil {
					log.Print(err)
					}
 				}else if message.Text=="尻"||strings.Contains(message.Text,"尻尻"){
 					if _, err = bot.ReplyMessage(event.ReplyToken,
 					linebot.NewTextMessage("聽說阿麵禁尻考爆好，還不+但可樂我很想尻欸")).Do(); err != nil{
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
 				}else if message.Text=="可樂"||message.Text=="可樂南海紙張"{
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("嗨!想補充體內的ATC?")).Do(); err != nil {
					log.Print(err)
					}
 				}else if message.Text=="help"{
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("沒有抽卡指令啦放棄吧！")).Do(); err != nil {
					log.Print(err)
					}
 				}else if message.Text=="嗨"{
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("嗨屁嗨，你以為有人會理你喔??")).Do(); err != nil {
					log.Print(err)
					}
 				}else if message.Text=="哈囉"{
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哈囉婊子")).Do(); err != nil {
					log.Print(err)
					}
 				}else if message.Text=="講個故事"||message.Text=="故事"{
 					r := rand.New(rand.NewSource(time.Now().UnixNano()));
 					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(story[r.Intn(storynum)])).Do(); err != nil {
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
