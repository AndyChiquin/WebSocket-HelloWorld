package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

// Configuración del WebSocket
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

// Manejador de las conexiones WebSocket
func handleConnection(w http.ResponseWriter, r *http.Request) {
    // Actualiza la conexión HTTP a WebSocket
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    // Enviar "Hola Mundo" al cliente
    err = conn.WriteMessage(websocket.TextMessage, []byte("Hello World - WebSocket"))
    if err != nil {
        log.Println(err)
        return
    }

    // Mantener la conexión abierta
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        fmt.Printf("Recibido: %s\n", msg)
    }
}

func main() {
    // Definir la ruta y el manejador
    http.HandleFunc("/ws", handleConnection)

    // Iniciar el servidor en el puerto 8080
    fmt.Println("Servidor WebSocket corriendo en ws://localhost:8080/ws")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
