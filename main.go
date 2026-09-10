package main

import (
	"fmt" 
	"html/template"
	"net/http"
	"net/url"
)

type M map[string]interface{}

type Person struct {
	 	Nama string
		NomorTelepon string
		Email string
		Laporan string
	}

	func routeIndexGet(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET" {
			var tmpl = template.Must(template.New("form").ParseFiles("views/hubungi-kami.html"))
			var err = tmpl.Execute(w, nil)

			if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		http.Error(w, "", http.StatusBadRequest)
	}

	func routeSubmitPost(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST" {
        	var tmpl = template.Must(template.New("result").ParseFiles("views/hubungi-kami.html"))
			
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
        	}

			var name = r.FormValue("nama")
			var phone = r.FormValue("nomorTelepon")
			var email = r.FormValue("email")
			var report = r.FormValue("laporan")
			
			sendTelegram("Nama: " + name + "\nNomor Telepon: " + phone + "\nEmail: " + email + "\nLaporan: " + report)

        	var data = map[string]string{"nama": name, "nomorTelepon": phone, "email": email, "laporan": report}

			var err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		http.Error(w, "", http.StatusBadRequest)
		
	}

	func sendTelegram(text string) {
		botToken :="8940303475:AAHJOXXZLfh8Jcru4mieOUCbF4Nx-0_WYP4"
		chatID := "851200267"
		fullURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", botToken, chatID, text)

		params := url.Values{}
		params.Add("chat_id", chatID)
		params.Add("text", text)

		fullURL = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?%s", botToken, params.Encode())

		resp, err := http.Get(fullURL)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}

func main() {

	// index
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
	})

	http.HandleFunc("/hubungi-kami", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))


			// http.Handle("/static", http.FileServer(http.Dir("assets")))


	// http.HandleFunc("/static/style.css", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "/views/static/css/style.css")
	// })

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte("post"))
		case "GET":
			w.Write([]byte("get"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})


	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)


}