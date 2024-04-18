package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"time"
)

var db *sql.DB

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
			http.NotFound(w, r)
			return
	}

	var err error
	//w.Header().Set("Set-Cookie", "session_id=test")
	r.ParseForm()
	user := r.Form.Get("name")

	if user == "" {
		http.Error(w, "Please enter a Name!", 403)
		return
	}

	//get existing task of user:
	query := fmt.Sprintf("select t.id from tasks as t left outer join users as u on t.id=u.task_id WHERE task_id IS NULL ORDER BY RANDOM() LIMIT 1;")

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Internal Error.", 500)
		return
	}

	var new_id int = 0
	rows.Next()
	rows.Scan(&new_id)

	if new_id == 0 {
		http.Error(w, "Sorry. No more tasks", 403)
		return
	}

	query = fmt.Sprintf("insert into users(name, task_id) VALUES ('%s', (select t.id from tasks as t left outer join users as u on t.id=u.task_id WHERE task_id IS NULL ORDER BY RANDOM() LIMIT 1)) RETURNING session;", user)

	rows, err = db.Query(query)
	if err != nil {
		http.Error(w, "Name is already in use.", 403)
		return
	}

	var session_id string = ""
	rows.Next()
	rows.Scan(&session_id)

	expiration := time.Now().Add(1 * time.Hour * 24 * 7)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   session_id,
		Expires: expiration}
	http.SetCookie(w, &cookie)

	defer rows.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
	}
	read_cookie, cookie_err := r.Cookie("session_id")

	if cookie_err == nil {

		query := fmt.Sprintf("select u.name, t.task from tasks as t left outer join users as u on t.id=u.task_id where u.session = '%s';", read_cookie.Value)

		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "Die Cookies waren zu lang im Ofen.", 403)
			return
		}

		var name string = ""
		var task string = ""
		rows.Next()
		rows.Scan(&name, &task)
		if name != "" {
			w.Write([]byte("<!DOCTYPE html>\n"))
			w.Write([]byte("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />"))
			w.Write([]byte("<h1>Chaos Tasks</h1>\n"))
			w.Write([]byte(fmt.Sprintf("<p><u>Name</u>: %s</p>", name)))
			w.Write([]byte(fmt.Sprintf("<p><u>%s</u>: %s</p>", lang_trans(r, "task"), task)))
			w.Write([]byte("</html>\n"))
			return
		}
	}
	w.Write([]byte("<!DOCTYPE html>\n"))
	w.Write([]byte("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />"))
	w.Write([]byte("<h1>Chaos Tasks</h1>\n"))
	w.Write([]byte(fmt.Sprintf("<p>%s</p>\n", lang_trans(r, "descr"))))
	w.Write([]byte("<form action=\"/post\" method=\"post\">\n"))
	w.Write([]byte(fmt.Sprintf("<label for=\"name\">%s</label><br>\n", lang_trans(r, "y_name"))))
	w.Write([]byte("<input type=\"text\" id=\"name\" name=\"name\">\n"))
	w.Write([]byte(fmt.Sprintf("<input type=\"submit\" value=\"%s\"/>\n", lang_trans(r, "button"))))
	w.Write([]byte("</form>\n"))
	w.Write([]byte("</html>\n"))
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {

	var db_conf DBConfAll

	db_conf = getDBconfig()

	port := getEnv("PORT", "3000")

	mux := http.NewServeMux()
	var err error
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		db_conf.conf.db_host,
		db_conf.conf.db_user,
		db_conf.conf.db_name,
		db_conf.db_password)
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		fmt.Printf("Database connection not working:")
		os.Exit(1)
	}

	rows, err := db.Query("SELECT name FROM users")

	if err != nil {
		fmt.Printf("Database connection not working with following credentials:\n")
		fmt.Printf("%+v\n db_password:XXX\n", db_conf.conf)
		_ = rows
		os.Exit(1)
	}

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/post", postHandler)

	http.ListenAndServe(":"+port, mux)
}
