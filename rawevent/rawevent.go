package rawevent

import (
	"bufio"
	"fmt"
	"github.com/asaskevich/EventBus"
	"net"
	"strings"
	"time"
)

type EventClient struct {
	conn net.Conn
}

func Start(host string, user string, password string, bus EventBus.Bus) (*EventClient, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			text := strings.Trim(scanner.Text(), "\n\r ")
			length := len(text)

			if length != 0 {
				if strings.HasPrefix(text, "Welcome to the TeamSpeak 3 ServerQuery interface") {
					_, _ = conn.Write([]byte(fmt.Sprintf("login %s %s\n", user, password)))
				} else if strings.HasPrefix(text, "notify") {
					eventName, parameters := parseEvent(text)
					bus.Publish(eventName, parameters)
				}
			}
		}
	}()

	evc := &EventClient{conn: conn}

	go evc.keepAlive()

	return evc, nil
}

func (cl *EventClient) Stop() {
	//TODO option for stopping the event client
}

func parseEvent(e string) (eventName string, parameters string) {
	split := strings.Split(e, " ")

	eventName = split[0]
	parameters = strings.Join(split[1:], "&")

	return
}

func (cl *EventClient) keepAlive() {
	for {
		cl.write(" \n")
		time.Sleep(5 * time.Second)
	}
}

func (cl *EventClient) write(payload string) {
	command := fmt.Sprintf("%s\n", payload)
	_, _ = cl.conn.Write([]byte(command))
}

func (cl *EventClient) SwitchServer(server int) {
	cl.write(fmt.Sprintf("use %d", server))
	cl.registerNotifications()
}

func (cl *EventClient) registerNotifications() {
	cl.write("servernotifyregister event=channel id=0")
	cl.write("servernotifyregister event=server")
	cl.write("servernotifyregister event=textserver")
	cl.write("servernotifyregister event=textchannel")
	cl.write("servernotifyregister event=textprivate")
}
