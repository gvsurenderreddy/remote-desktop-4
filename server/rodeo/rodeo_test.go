package rodeo

import (
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func TestRodeoServer(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	log.Println("Test Rodeo Server")
	RodeoServer()
	time.Sleep(time.Millisecond * 100)
	go addHost("HostFromGoTest#1", t)
	time.Sleep(time.Millisecond * 100)
	go addHost("HostFromGoTest#2", t)
	time.Sleep(time.Millisecond * 200)
	go addClient("CLIENT 1", t)
	time.Sleep(time.Second * 1)
	go addClient("CLIENT 2", t)
	time.Sleep(time.Second * 1)
	go addClient("CLIENT 3", t)
	time.Sleep(time.Second * 1)
}

func addHost(login string, t *testing.T) {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:9595", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	action := Action{"HOST_REGISTER", login}

	err = c.WriteJSON(&action)
	if err != nil {
		log.Println("write:", err)
		return
	}

	var msg_from_client bool = false
	for {
		log.Println(login, ": wait Action...")
		err = c.ReadJSON(&action)
		if err != nil {
			log.Println("readJson:", err)
			return
		}
		log.Println("host:\n", login, ":\n", action.Action, action.Data)
		switch action.Action {
		case "CLIENT_CONNECT":
			sendAction(&Action{"CLIENT_ACCESS", ""}, c)
			go func() {
				time.Sleep(time.Millisecond * 100)
				sendAction(&Action{"MSG_FROM_SERVER", ""}, c)
			}()
			go func() {
				time.Sleep(time.Millisecond * 200)
				if msg_from_client == false {
					t.Error("Host handler expect message from client, but don't get it")
				}
			}()
		case "MSG_FROM_CLIENT":
			msg_from_client = true
		}
	}
}

func addClient(login string, t *testing.T) {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:9595", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	action := Action{"GET_HOSTS", ""}

	err = c.WriteJSON(&action)
	if err != nil {
		log.Println("write:", err)
		return
	}
	var msg_from_server bool = false
	for {
		log.Println(login, ": wait Action...")
		err = c.ReadJSON(&action)
		if err != nil {
			log.Println("readJson:", err)
			return
		}
		log.Println("client:\n", login, ":\n", action.Action, action.Data)
		switch action.Action {
		case "AVALIABLE_HOSTS":
			name := ((action.Data).([]interface{})[0]).(map[string]interface{})["login"]
			sendAction(&Action{"SELECT_HOST", name}, c)
		case "SELECT_SUCCESS":
			go func() {
				time.Sleep(time.Millisecond * 100)
				sendAction(&Action{"MSG_FROM_CLIENT", ""}, c)
			}()
			go func() {
				time.Sleep(time.Millisecond * 200)
				if msg_from_server == false {
					t.Error("Client handler expect message from server, but don't get it")
				}
			}()
		case "SELECT_FAIL":
			data := (action.Data).(string)
			t.Error("Select fail for client:  ", login, ". Err: ", data)
		case "MSG_FROM_SERVER":
			msg_from_server = true
			time.Sleep(time.Millisecond * 300)
			return
		}
	}
}

func sendAction(action *Action, conn *websocket.Conn) {
	if err := conn.WriteJSON(action); err != nil {
		log.Println("sendAction[", action.Action, "] WriteJSON: ", err)
	}
}
