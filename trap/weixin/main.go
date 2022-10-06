package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// Ticket 类型
type Ticket struct {
	Errcode   int    `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
	Ticket    string `json:"ticket,omitempty"`
	ExpiresIn int    `json:"expires_in,omitempty"`
}

// Token 类型
type Token struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

// Sign 签名类型
type Sign struct {
	AppID     string `json:"app_id,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	NonceStr  string `json:"nonce_str,omitempty"`
	Signature string `json:"signature,omitempty"`
}

//IndexTemplate 页面模板
type IndexTemplate struct {
	Title template.HTML
}

var (
	//微信公众号
	wxAppID     = ""
	wxSecret    = ""
	port        = ":8383"
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wxNoncestr := RandStringRunes(32)
	wxURL := "https://testh5.xizhihk.com/"
	//wxURL := "http://192.168.154.222:8080/"
	ticket := "LIKLckvwlJT9cWIhEQTwfL8TQv9UUrc9g1Qd-TeSQBzoGlL0nJ8VJDWQOuVKugvW8XbvLsU2pd9E8M1MxbX_Pw"

	timestamp, signature := GetCanshu(wxNoncestr, wxURL, ticket)
	fmt.Printf("wxNoncestr: %s\nwxURL: %s\ntimestamp: %s\nsignature: %s\nticket: %s", wxNoncestr, wxURL, timestamp, signature, ticket)
}

//signHandler 异步处理微信签名
//func signHandler(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	if r.Method == "POST" {
//		wxNoncestr := RandStringRunes(32)
//		wxURL, _ := url.QueryUnescape(r.FormValue("url"))
//		timestamp, signature := GetCanshu(wxNoncestr, wxURL)
//		var u = Sign{
//			AppID:     wxAppID,
//			Timestamp: timestamp,
//			NonceStr:  wxNoncestr,
//			Signature: signature,
//		}
//		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
//		w.Header().Set("Content-type", "application/json")             //返回数据格式是json
//		b, err := json.Marshal(u)
//		if err != nil {
//			log.Println(err.Error())
//		}
//		w.Write(b)
//	} else if r.Method == "GET" {
//		t, _ := template.ParseFiles("html/index.tpl", "html/foot.tpl")
//		t.Execute(w, &IndexTemplate{Title: template.HTML("生成微信分享签名")})
//	}
//}

//GetCanshu 微信签名算法
func GetCanshu(noncestr, url string, ticket string) (timestamp, signature string) {
	timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	longstr := "jsapi_ticket=" + ticket + "&noncestr=" + noncestr + "&timestamp=" + timestamp + "&url=" + url

	h := sha1.New()
	if _, e := h.Write([]byte(longstr)); e != nil {
		log.Println(e.Error())
	}

	signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}

//GetWeixin 得到微信AccessToken和JSTicket
//func GetWeixin(appid, secret string) {
//	var tk Token
//	var tc Ticket
//	db, err := storm.Open("db/weixin.db")
//	if err != nil {
//		log.Println("Database open err:", err.Error())
//	}
//	defer db.Close()
//
//	gorequest.New().Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret).EndStruct(&tk)
//	gorequest.New().Get("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + tk.AccessToken + "&type=jsapi").EndStruct(&tc)
//
//	if e := db.Set("sessions", "token", &tk); e != nil {
//		log.Println(e.Error())
//	}
//	if e := db.Set("sessions", "ticket", &tc); e != nil {
//		log.Println(e.Error())
//	}
//}

//RandStringRunes 生成随机字符串
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
