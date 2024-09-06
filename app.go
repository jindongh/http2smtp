package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"

    "github.com/go-mail/mail"
)

type config struct {
    port string
    smtp_host string
    smtp_port int
    smtp_user string
    smtp_password string
}

func main() {
    conf := loadConfig()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Error parsing form data", http.StatusBadRequest)
            return
        }
        to := r.FormValue("to")
        subject := r.FormValue("subject")
        content := r.FormValue("content")
        if to == "" {
            http.Error(w, "parameter to can't be empty", http.StatusBadRequest)
            return
        }
        if subject == "" {
            http.Error(w, "parameter subject can't be empty", http.StatusBadRequest)
            return
        }
        if content == "" {
            http.Error(w, "parameter content can't be empty", http.StatusBadRequest)
            return
        }
        if err := sendEmail(conf, to, subject, content); err != nil {
            http.Error(w, fmt.Sprintf("Failed to send email %v", err), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Email was sent to %s with subject: %s content: %s", to, subject, content)
    })

    fmt.Println("start listen at " + conf.port)
    err := http.ListenAndServe(":" + conf.port, nil)
    if err != nil {
        log.Fatalf("failed to start server with %v", err)
    }
}

func loadConfig() *config {
    smtpPortStr := os.Getenv("SMTP_PORT")
    smtpPort, err := strconv.Atoi(smtpPortStr)
    if err != nil {
        log.Fatalf("failed to parse SMTP_PORT=%s as int %v", smtpPortStr, err)
    }
    return &config{
        port: os.Getenv("PORT"),
        smtp_host: os.Getenv("SMTP_HOST"),
        smtp_port: smtpPort,
        smtp_user: os.Getenv("SMTP_USER"),
        smtp_password: os.Getenv("SMTP_PASSWORD"),
    }
}

func sendEmail(conf *config, to, subject, content string) error {
    m := mail.NewMessage()
    m.SetHeader("From", conf.smtp_user)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text", content)
    d := mail.NewDialer(conf.smtp_host, conf.smtp_port, conf.smtp_user, conf.smtp_password)
    return d.DialAndSend(m)
}
