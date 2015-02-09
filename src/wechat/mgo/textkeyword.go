package mgo

/*
插入关键字的回复
*/
import (
	"net/http"

	"fmt"
	"labix.org/v2/mgo"

	"labix.org/v2/mgo/bson"
)

type Key struct {
	Keyword string `xml:"keyword"`
	Content string `xml:"content"`
}

func MongoinsertKeyword(w http.ResponseWriter, r *http.Request) {
	// data := whttp.Parserequest(r)
	// v := Key{}
	// wxml.Decodebytes(data, &v)
	// realmap := wmap.Struct2map(v)
	r.ParseForm()
	mapdata := map[string]string{}
	mapdata["keyword"] = r.FormValue("keyword")
	mapdata["content"] = r.FormValue("content")
	mapdata["flag"] = r.FormValue("flag")

	MongoInsert(mapdata, "keyword")
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域,解决跨域问题

	fmt.Fprintf(w, "ok")

}
func MongoFindKeyword(w http.ResponseWriter, r *http.Request) (Key, error) {
	r.ParseForm()
	message := r.FormValue("keyword")
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("wechat").C("keyword")

	result := Key{}
	err = c.Find(bson.M{"keyword": message}).One(&result)
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域,解决跨域问题

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, result)
	}

	return result, err
}
func MongoFindKeywordLocal(message string) (Key, error) {

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("wechat").C("keyword")

	result := Key{}
	err = c.Find(bson.M{"keyword": message}).One(&result)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	return result, err
}
