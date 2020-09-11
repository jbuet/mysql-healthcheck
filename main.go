package main

import (
	"database/sql"
	"log"
	"net/http"
	"net"
	"os"
	"gopkg.in/gcfg.v1"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Config struct {
	Mysql struct {
		UserName string
		Password string
		Database string
		Port     string
	}
	HealthCheck struct {
		Port             string
	}
}

var cnf Config

func main() {

	err := gcfg.ReadFileInto(&cnf, os.Getenv("MYSQL_HEALTHCHECK_PATH"))
	if err != nil {
		log.Printf("%v", err)
		return
	}

	http.HandleFunc("/", healthcheck)                 
	err = http.ListenAndServe(":"+cnf.HealthCheck.Port, nil) 
    if err != nil {
		log.Printf("%v", err)
		return
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", cnf.Mysql.UserName+":"+cnf.Mysql.Password+"@/"+cnf.Mysql.Database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	ip := getIP(r)


	err = db.Ping()
	if err != nil {
		log.Printf("MySQL Running: %v [false]", ip)
		w.Header().Set("Error", "Unable to connect to mysql")
		w.WriteHeader(500)
		return
	}

	log.Printf("MySQL Running: %v [true]", ip)
	w.Header().Set("Server", "Mysql-Health-Check")

}

func getIP(r *http.Request) (string) {
    //Get IP from the X-REAL-IP header
    ip := r.Header.Get("X-REAL-IP")
    netIP := net.ParseIP(ip)
    if netIP != nil {
        return ip
    }

    //Get IP from X-FORWARDED-FOR header
    ips := r.Header.Get("X-FORWARDED-FOR")
    splitIps := strings.Split(ips, ",")
    for _, ip := range splitIps {
        netIP := net.ParseIP(ip)
        if netIP != nil {
            return ip
        }
    }

    //Get IP from RemoteAddr
    ip, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return ""
    }
    netIP = net.ParseIP(ip)
    if netIP != nil {
        return ip
	}

    return ""
}