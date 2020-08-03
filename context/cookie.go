package main

import (
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", StatusHandler)
	mux.HandleFunc("/login", LoginHandler)
	mux.HandleFunc("/logout", LogoutHandler)
	contextedMux := AddContextSupport(mux)
	log.Fatal(http.ListenAndServe(":8080", contextedMux))
}

func AddContextSupport(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("username")
		if cookie != nil {
			ctx := context.WithValue(r.Context(), "username", cookie.Value)
			// WithContext returns a shallow copy of r with its context changed
			// to ctx. The provided ctx must be non-nil.
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	expitation := time.Now().Add(24 * time.Hour)
	var username string
	if username = r.URL.Query().Get("username"); username == "" {
		username = "guest"
	}
	cookie := http.Cookie{Name: "username", Value: username, Expires: expitation}
	http.SetCookie(w, &cookie)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().AddDate(0, 0, -1)
	cookie := http.Cookie{Name: "username", Value: "alice_cooper@gmail.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {

	if username := r.Context().Value("username"); username != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hi username:" + username.(string) + "\n"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}
func longRunningCalculation(timeCost int) chan string {

	result := make(chan string)
	go func() {
		time.Sleep(time.Second * (time.Duration(timeCost)))
		result <- "Done"
	}()
	return result
}
func jobWithCancelHandler(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	var cancel context.CancelFunc
	var jobtime string
	if jobtime = r.URL.Query().Get("jobtime"); jobtime == "" {
		jobtime = "10"
	}
	timecost, err := strconv.Atoi(jobtime)
	if err != nil {
		timecost = 10
	}
	log.Println("Job will cost : " + jobtime + "s")
	ctx, cancel = context.WithCancel(r.Context())
	defer cancel()

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return
	case result := <-longRunningCalculation(timecost):
		io.WriteString(w, result)
	}
	return
}
