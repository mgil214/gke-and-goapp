package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"
)

type echoRequest struct {
	Echo string `json:"echo"`
}

type echo struct {
	YourName string `json:"your_name"`
	YouSaid  string `json:"you_said"`
	Me       string `json:"me"`
}

// echoHandler requires a POST request with body:
//
// {
//   "echo": "what you sent"
// }
//
// Returns:
//
// {
//   "your_name": "asdfasdfasdfasdfasdf",
//   "you_said": "what you sent"
//   "me": "what you sent"
// }

func echoHandler(w http.ResponseWriter, r *http.Request) {
	req := echoRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Request body error: "+err.Error(), 400)
		return
	}
	e := echo{
		YourName: r.UserAgent(),
		YouSaid:  req.Echo,
		Me:       req.Echo,
	}

	log.Printf("echo: %v", e)
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	json.NewEncoder(w).Encode(e)
}

type datetime struct {
	RFC3339 string `json:"rfc3339"`
	Unix    int64  `json:"unix"`
}

type now struct {
	YourName string   `json:"your_name"`
	Datetime datetime `json:"datetime"`
}

// nowHandler requires a GET request, returns:
//
// {
//   "your-name": "asdfasdfasdfasdfasdf",
//   "datetime":
//   {
//     "rfc3339": "rfc string",
//     "unix": 12312341235123
//   }
// }
func nowHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	n := now{

		YourName: r.UserAgent(),
		Datetime: datetime{
			RFC3339: t.Format(time.RFC3339),
			Unix:    t.Unix(),
		},
	}

	log.Printf("server time: %v", n)
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	json.NewEncoder(w).Encode(n)
}

func main() {
	port := flag.Int("port", 8080, "port to listen from")
	flag.Parse()
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/now", nowHandler)
	log.Printf("Running on port %d", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
