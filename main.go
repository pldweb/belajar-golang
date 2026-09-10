package main

import (
	"fmt" 
	"html/template"
	"net/http"
	"net/url"
	"github.com/resend/resend-go/v3"
	"context"
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
			var title = "Hubungi Kami"
			var tmpl = template.Must(template.New("form").ParseFiles(
				"views/hubungi-kami.html",
				"views/_header.html",
			))
			var err = tmpl.Execute(w, M{"title": title})

			if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		http.Error(w, "", http.StatusBadRequest)
	}

	func routeSubmitPost(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST" {
        	var tmpl = template.Must(template.New("form").ParseFiles(
				"views/hubungi-kami.html",
				"views/_header.html",
			))
			
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
        	}

			var name = r.FormValue("nama")
			var phone = r.FormValue("nomorTelepon")
			var email = r.FormValue("email")
			var report = r.FormValue("laporan")
			var userID = r.FormValue("userid")
			
			sendEmail(name, phone, email, report)
			sendTelegram("Nama: " + name + "\nNomor Telepon: " + phone + "\nEmail: " + email + "\nLaporan: " + report, userID)
			sendTelegram("Nama: " + name + "\nNomor Telepon: " + phone + "\nEmail: " + email + "\nLaporan: " + report, "851200267")

        	var data = "Berhasil mengirim laporan. Terima kasih atas laporan Anda"

			var err = tmpl.Execute(w, M{"data": data})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		http.Error(w, "", http.StatusBadRequest)
		
	}

	func sendTelegram(text string, userID string) {
		botToken :="8940303475:AAHJOXXZLfh8Jcru4mieOUCbF4Nx-0_WYP4"
		chatID := userID // Replace with the actual chat ID
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

	func sendEmail(nama string, nomorTelepon string, email string, laporan string) {
		ctx := context.TODO()
  		client := resend.NewClient("re_VZVK5xtV_GNbce2jf7E4PyojgheyPdaym")
		params := &resend.SendEmailRequest{
			From:        "Support Rinkweb <support@rinkwebstudio.my.id>",
			To:          []string{email},
			Subject:     "Report Golang",
			Html:        "<p>Nama: " + nama + "</p><br><p>Nomor Telepon: " + nomorTelepon + "</p><br><p>Email: " + email + "</p><br><p>Laporan: " + laporan + "</p>",
  		}

		sent, err := client.Emails.SendWithContext(ctx, params)

		if err != nil {
			panic(err)
		}
  		fmt.Println(sent.Id)
	}

func main() {

	// index
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var title = "Home"
		var tmpl = template.Must(template.New("index").ParseFiles(
			"views/index.html",
			"views/_header.html",
			))
		var err = tmpl.Execute(w, M{"title": title})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}	
	})

	http.HandleFunc("/tentang", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("tentang").ParseFiles(
			"views/tentang.html",
			"views/_header.html",
			))
		var title = "Tentang Kami"
		var err = tmpl.Execute(w, M{"title": title})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}	
	})

	http.HandleFunc("/testimoni", func(w http.ResponseWriter, r *http.Request) {
		var title = "Testimoni"
		var tmpl = template.Must(template.New("testimoni").ParseFiles(
			"views/testimoni.html",
			"views/_header.html",
			))
		var err = tmpl.Execute(w, M{"title": title})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}	
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