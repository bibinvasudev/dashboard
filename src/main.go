// websockets.go
package main

import (
    "fmt"
    "net/http"
    "time"
    "github.com/gorilla/websocket"
    "log"
    "bufio"
    "os"


)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
	// Create a simple file server
    fs := http.FileServer(http.Dir("../public"))
    http.Handle("/", fs)
    file, err := os.Open("access.log")
    if err != nil {
            log.Fatal(err)
    }
    defer file.Close()

    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity



        for {
            // Read message from browser
            msgType, msg, err := conn.ReadMessage()
            fmt.Printf(string (msg))
            if err != nil {
                return
            }

            // Print the message to the console
            fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
            msg = []byte("Here is a string....")


            scanner := bufio.NewScanner(file)
            for scanner.Scan() {
                fmt.Println(scanner.Text())
                time.Sleep(2 * time.Second)
                msg = []byte(scanner.Text())
                conn.WriteMessage(msgType, msg)

            }

            if err := scanner.Err(); err != nil {
                log.Fatal(err)
            }

        }
    })

    http.HandleFunc("/param", func(w http.ResponseWriter, r *http.Request) {
        conn1, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity



        for {
            // Read message from browser
            //msgType, msg, err := conn.ReadMessage()
            //fmt.Printf(string (msg))
            //if err != nil {
            //    return
            //}

            // Print the message to the console
            //fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
            //msg = []byte("Here is a string....")


            //scanner := bufio.NewScanner(file)
            //for scanner.Scan() {
            //    fmt.Println(scanner.Text())
            //    time.Sleep(2 * time.Second)
            //    msg = []byte(scanner.Text())
            //    conn.WriteMessage(msgType, msg)

            //}

            //if err := scanner.Err(); err != nil {
            //    log.Fatal(err)
            //}
            
           msgType1, msg1, err1 := conn1.ReadMessage()

           if err1 != nil {
                return
            }
 
          msg1 = []byte(string(55))
           conn1.WriteMessage(msgType1, msg1)
         

        }
    })
    http.ListenAndServe(":8080", nil)
}
