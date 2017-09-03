package main
import (
	"net/http"
	"log"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"stu/model"
	"fmt"
	"strconv"
)
var o orm.Ormer
func main() {

	http.HandleFunc("/stu/add", addHandler)
	http.HandleFunc("/stu/del", delHandler)
	http.HandleFunc("/stu/list", listHandler)
	http.Handle("/", http.FileServer(http.Dir("/opt/www")))
	log.Print("www server 8888 running.")
	http.ListenAndServe(":8888", nil)
}
func addHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	name:=string(req.Form["name"][0])
	namesx:=string(req.Form["namesx"][0])
	grade,_:= strconv.Atoi(string(req.Form["grade"][0]))
	class:= string(req.Form["class"][0])

	stu:=model.Student{Name:name,Namesx:namesx,Grade:grade,Class:class}

	id, err := o.Insert(&stu)
	if err == nil {
		fmt.Println(id)
	}
}
func listHandler(w http.ResponseWriter, req *http.Request) {
	stu := model.Student{}
	err := o.Read(&stu)
	fmt.Printf("ERR: %v\n", err)
}
func delHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	id,_:= strconv.Atoi(string(req.Form["id"][0]))
	if num, err := o.Delete(&model.Student{Id: id}); err == nil {
		fmt.Println(num)
	}
}

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:cloudos@/my_db?charset=utf8", 30)

	// 注册定义的 model
	orm.RegisterModel(new(model.Student))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	// 创建 table
	orm.RunSyncdb("default", false, true)

	o = orm.NewOrm()
}

