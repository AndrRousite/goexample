package main

import "fmt"
import (
	"math/rand"
	"net/http"
	"strings"
	"log"
	"html/template"
	"time"
	"crypto/md5"
	"strconv"
	"encoding/hex"
	"os"
	"io"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"os/signal"
	"syscall"
	"encoding/base64"
	"github.com/astaxie/beego/session"
	"io/ioutil"
	"encoding/xml"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"regexp"
	"golang.org/x/net/websocket"
	"./hello"
)

func main() {
	//fmt.Println("Hello World.")
	//num := 0
	//for i := 0; i < 10; i++ {
	//	fmt.Println("i = ", i, "----", rand.Intn(10))
	//	num += i
	//	if num%3 == 0 {
	//		fmt.Println("======================")
	//	} else if num%3 == 1 {
	//		fmt.Println("**********************")
	//	} else {
	//		fmt.Println("######################")
	//	}
	//}
	//fmt.Println(math.Pi)
	//
	//v := Person{}
	//v.Name = "张三"
	//
	//fmt.Println(v.Name)
	//
	//s := &v
	//
	//fmt.Println(s)
	//
	//a := make([]string, 10)
	//fmt.Println(a)
	//
	//var pow = []string{"刘枫", "张飒", "李斯", "王屋"}
	//
	//for v := range pow {
	//	fmt.Printf("s = %s\n", pow[v])
	//}
	//
	//t := time.Date(1994, time.July, 14, 2, 0, 0, 0, time.Local)
	//
	//fmt.Printf("您的出生日期为： %s ", t)
	//
	////openMYSQL()
	////test, err := get("test")
	////fmt.Println(test, err)
	//
	////parseXML()
	//parseJson()
	//
	//mould()
	//
	////makeFile()
	//
	//strconvString()
	//
	//http.HandleFunc("/hello", hello)
	//http.HandleFunc("/login", login)
	//http.HandleFunc("/upload", upload)
	//
	//http.HandleFunc("/cookie", setCookie)
	//http.HandleFunc("/unique", unique)
	//http.HandleFunc("/count", count)
	//http.HandleFunc("/regexp", match)
	//http.HandleFunc("/reptile", reptile)
	//
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	log.Fatal("Error:", err)
	//}

	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
	//	os.Exit(1)
	//}
	//name := os.Args[1]
	//addr := net.ParseIP(name)
	//if addr == nil {
	//	fmt.Println("Invalid address")
	//} else {
	//	fmt.Println("The address is ", addr.String())
	//}
	//os.Exit(0)

	//http.Handle("/", websocket.Handler(receiver))
	//if err := http.ListenAndServe(":1234", nil); err != nil {
	//	log.Fatal("ListenAndServe:", err)
	//}

	hh.Hello()

	timeParse()
}

type Person struct {
	Name   string
	Avatar string
	Email  string
}

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Printf("key: %s \t value: %s", k, strings.Join(v, ""))
	}
	fmt.Fprintln(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.gtpl")
		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(w, template.JS("alert('hehe')"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("username:%s \t password:%s \t password:%s", r.FormValue("username"), r.Form["password"][0], strings.Join(r.Form["password"], ""))

		if r.FormValue("username") == "liufeng" && r.FormValue("password") == "123456" {
			fmt.Fprintf(w, "你好：%s", r.FormValue("username"))
		} else {
			t, err := template.ParseFiles("token.gtpl")
			if err != nil {
				log.Fatal(err)
			}
			// md5 加密
			h := md5.New()
			h.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
			err = t.Execute(w, hex.EncodeToString(h.Sum(nil)))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("upload.gtpl")
		if err != nil {
			log.Fatal(err)
		}
		h := md5.New()
		h.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
		err = t.Execute(w, hex.EncodeToString(h.Sum(nil)))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()
		// 创建文件
		f, err := os.OpenFile("./file/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		defer fmt.Fprint(w, "上传成功。")
	}
}

func openMYSQL() {
	db, err := sql.Open("mysql", "root:126515@/go?charset=utf8")
	checkError(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT USER SET user_name=?,password=?,create_time=?")
	checkError(err)

	// 执行
	h := md5.New()
	h.Write([]byte("123456"))
	t := time.Now().Unix()
	result, err := stmt.Exec("liu-feng", hex.EncodeToString(h.Sum(nil)), t)
	checkError(err)

	id, err := result.LastInsertId()
	checkError(err)
	fmt.Printf("插入成功，插入的用户ID为%d", id)
}

var Pool *redis.Pool
var gloablSessions *session.Manager

func init() {
	//beego()
	//Pool = pool()
	//close()
	// session
	gloablSessions, _ = session.NewManager("memory", &session.ManagerConfig{CookieName: "gosessionid",
		EnableSetCookie: true,
		Gclifetime: 3600})
	go gloablSessions.GC()
}

func beego() {
	orm.NewOrm()
}

func pool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func get(key string) ([]byte, error) {

	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		for _, cookie := range r.Cookies() {
			fmt.Fprintf(w, "%s->%s\n", cookie.Name, cookie.Value)
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "username", Value: "astaxie"})
	}
}

func unique(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Fprint(w, "")
		return
	}
	fmt.Fprint(w, base64.URLEncoding.EncodeToString(b))
}

func randString() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func count(w http.ResponseWriter, r *http.Request) {
	sess, _ := gloablSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	fmt.Println(sess.SessionID())
	ct := sess.Get("count")
	if ct == nil {
		sess.Set("count", 1)
	} else {
		sess.Set("count", ct.(int)+1)
	}
	t, err := template.ParseFiles("count.gtpl")
	checkError(err)
	//w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Max-Age", "0")
	err = t.Execute(w, sess.Get("count"))
	checkError(err)
}

type SERVERXML struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName xml.Name `xml:"server"`
	NAME    string   `xml:"serverName"`
	IP      string   `xml:"serverIP"`
}

func parseXML() {
	file, err := os.Open("test.xml")
	checkError(err)
	defer file.Close()
	// 读取文件的内容
	data, err := ioutil.ReadAll(file)
	checkError(err)
	// 解析xml
	v := SERVERXML{}
	err = xml.Unmarshal(data, &v)
	checkError(err)
	for _, value := range v.Svs {
		fmt.Printf("\n%s->%s", value.NAME, value.IP)
	}

	b, err := xml.Marshal(v)
	checkError(err)
	// new file
	f, err := os.OpenFile("./file/"+randString()+".xml", os.O_CREATE|os.O_WRONLY, 0666)
	checkError(err)
	defer f.Close()
	total, _ := f.Write(b)
	defer fmt.Println(total, "\t上传成功")

}

type Serverslice struct {
	Servers []Server
}

type Server struct {
	ServerName string
	ServerIP   string
}

func parseJson() {
	str := `{"Servers":[{"ServerName":"Shanghai_VPN","ServerIP":"127.0.0.1"},{"ServerName":"Beijing_VPN",
"ServerIP":"127.0.0.2"}]}`
	v := Serverslice{}
	err := json.Unmarshal([]byte(str), &v)
	checkError(err)
	fmt.Println(v)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b, &f)
	checkError(err)

	js, err := simplejson.NewJson(b)
	checkError(err)
	fmt.Println(js.Get("Name1").MustString())
}

func match(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("value")
	var arg string
	if s == "" {
		arg = ""
	} else if m, _ := regexp.MatchString("^[0-9]+$", s); m {
		arg = "数字"
	} else if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", s); m {
		arg = "IP"
	}
	t, err := template.ParseFiles("regexp.gtpl")
	checkError(err)
	t.Execute(w, arg)
}

/**
爬一爬百度
 */
func reptile(w http.ResponseWriter, r *http.Request) {
	var url string
	if len(r.FormValue("url")) > 0 {
		url = r.FormValue("url")
	} else {
		url = "http://www.baidu.com"
	}
	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()
	// 读取数据
	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	html := string(body)

	// 获取到html标签，全部转换为小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllStringFunc(html, strings.ToLower)

	// 去除style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	html = re.ReplaceAllString(html, "")

	// 去除script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	html = re.ReplaceAllString(html, "")

	// 将display:none 去除
	//html = strings.Replace(html, "display:none;", "", -1)

	// 去除所有尖括号中的代码,并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllString(html, "\n")

	// 去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	html = re.ReplaceAllString(html, "\n")

	// 试试错误的表达式
	re = regexp.MustCompile("\\$")
	html = re.ReplaceAllString(html, "XXXXXXXXXXXXXXXX")

	fmt.Fprint(w, html)
}

type Friend struct {
	Fname string
}

type User struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func mould() {
	//t := template.New("test")
	//t.Parse("Hello {{.Name}} {{.Email}}")
	//p := Person{Name: "liufeng", Email: "123456@qq.com"}
	//t.Execute(os.Stdout, p)

	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
			{{end}}
			`)
	p := User{UserName: "Astaxie",
		Emails: []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

func makeFile() {
	err := os.Mkdir("dir", 0777)
	checkError(err)
	err = os.MkdirAll("dir/test1/test2", 0777)
	checkError(err)
	fmt.Println("创建成功")
	//os.RemoveAll("dir")
}

func strconvString() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}

func receiver(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		fmt.Println("Sending to client:" + reply)

		if err = websocket.Message.Send(ws, reply); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func timeParse() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t := time.Now()
	t = t.In(loc)
	// 这个是什么鬼？
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
