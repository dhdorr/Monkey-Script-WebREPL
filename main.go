package main

import (
	"encoding/json"
	"fmt"
	"go-interpreter/object"
	"go-interpreter/repl"
	sessionidgenerator "go-interpreter/session_id_generator"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"
)

type TestStruct struct {
	Test string
}

func main() {
	// Web server starts here
	test_map := make(map[string]*object.Environment)
	delete_me_arr := make([]string, 0, 100)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("Welcome!\n")

		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("GET /create", func(w http.ResponseWriter, r *http.Request) {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		session_id := sessionidgenerator.String(6)
		_, ok := test_map[session_id]
		if !ok {
			test_map[session_id] = object.NewEnvironment()
			// add session id to queue, once full, start deleting from test_map
			if len(delete_me_arr) >= 90 {
				s := delete_me_arr[0]
				delete_me_arr = append(delete_me_arr[:0], delete_me_arr[1:]...)
				delete(test_map, s)
				fmt.Println("deleting from slice and map")

			}
			delete_me_arr = append(delete_me_arr, session_id)

			// start a timer, delete from map when timer times out
			timer1 := time.NewTimer(1 * time.Hour)
			go func() {

				<-timer1.C
				fmt.Println("SESSION OVER: Timer 1 fired...")
				s := delete_me_arr[0]
				delete(test_map, s)
				if len(delete_me_arr) == 1 {
					delete_me_arr = nil
				} else {
					delete_me_arr = append(delete_me_arr[:0], delete_me_arr[1:]...)
				}

				fmt.Println("session is over: deleting from slice and map")
				fmt.Printf("new test_map: %v\n", test_map)
				fmt.Printf("new delete_me_arr: %v\n", delete_me_arr)
			}()

			fmt.Printf("test_map: %v\n", test_map)
			fmt.Printf("delete_me_arr: %v\n", delete_me_arr)
		}

		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Your Session ID: %s\n", session_id)
		fmt.Printf("Feel free to type in commands\n")
		// code := `let x = 7;
		// puts(x);`
		// apb := repl.Start(os.Stdin, os.Stdout, test_map[session_id], session_id, code)
		// fmt.Println(apb + "\n")

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    session_id,
			Path:     "/",
			SameSite: http.SameSiteStrictMode,
		})

		tmpl, err := template.ParseFiles("./static/repl.html")
		if err != nil {
			log.Fatal(err)
		}

		tmpl.Execute(w, nil)

		// http.ServeFile(w, r, "static/html/repl.html")
	})

	// http.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Printf("Hello! TESTING!\n")
	// 	// mapD := map[string]int{"apple": 5, "lettuce": 7}
	// 	// mapB, _ := json.Marshal(mapD)

	// 	// w.Write(mapB)
	// 	c, err := r.Cookie("session_id")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	code := "x;"

	// 	for k := range test_map {
	// 		fmt.Printf("key[%s]\n", k)
	// 		//t := test_map[k]
	// 		//k := *t

	// 	}

	// 	fmt.Printf("cookie: %s", c.Value)
	// 	helpme := repl.Start(os.Stdin, os.Stdout, test_map[c.Value], c.Value, code)

	// 	tmpl, err := template.ParseFiles("./static/html/test_response.html")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	data := TestStruct{
	// 		Test: helpme,
	// 	}

	// 	tmpl.Execute(w, data)

	// })

	http.HandleFunc("PUT /repl", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("I am doing text box stuff!\n")

		c, err := r.Cookie("session_id")
		if err != nil {
			log.Fatal(err)
		}
		code := r.FormValue("replBox")
		fmt.Printf("what i got: %s!\n", code)

		helpme := ""

		_, ok := test_map[c.Value]
		if !ok {
			fmt.Printf("Session ID: %s not found...\nMust've been the wind...", c.Value)
			helpme = "Session ID not found...\nMust've been the wind..."
		} else {
			helpme = repl.Start(os.Stdin, os.Stdout, test_map[c.Value], c.Value, code)
		}

		tmpl, err := template.ParseFiles("./static/html/test_response.html")
		if err != nil {
			log.Fatal(err)
		}
		data := TestStruct{
			Test: helpme,
		}
		j := make(map[string]string)
		j["myEvent"] = code
		j["myEvent2"] = helpme
		j_t, _ := json.Marshal(j)
		w.Header().Add("HX-Trigger-After-Swap", string(j_t))

		tmpl.Execute(w, data)

	})

	// Add this line to serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8000", nil)
}
