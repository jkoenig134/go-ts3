package go_ts3

import (
	"bufio"
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/jkoenig134/schema"
	"net"
	"net/url"
	"strings"
	"time"
)

type EventClient struct {
	conn     net.Conn
	user     string
	password string
	bus      EventBus.Bus
	decoder  *schema.Decoder
}

func newEventClient(host string, user string, password string, bus EventBus.Bus) (*EventClient, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	evc := &EventClient{
		conn:     conn,
		user:     user,
		password: password,
		bus:      bus,
		decoder:  schema.NewDecoder(),
	}

	go evc.keepAlive()
	go evc.HandlePayloads()

	return evc, nil
}

func (cl *EventClient) HandlePayloads() {
	scanner := bufio.NewScanner(cl.conn)
	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), "\n\r ")
		length := len(text)

		if length != 0 {
			if strings.HasPrefix(text, "Welcome to the TeamSpeak 3 ServerQuery interface") {
				_ = cl.login()
			} else if strings.HasPrefix(text, "notify") {
				cl.handleEvent(text)
			}
		}
	}
}

func (cl *EventClient) Stop() error {
	err := cl.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

func (cl *EventClient) login() error {
	err := cl.write(fmt.Sprintf("login %s %s", cl.user, cl.password))
	if err != nil {
		return err
	}

	return nil
}

func (cl *EventClient) handleEvent(event string) {
	eventName, parameters := splitEvent(event)

	switch eventName {
	case "notifycliententerview":
		cl.publishEvent(eventName, parameters, &ClientEnterViewEvent{})
	case "notifyclientleftview":
		cl.publishEvent(eventName, parameters, &ClientLeftViewEvent{})
	case "notifyserveredited":
		cl.publishEvent(eventName, parameters, &ServerEditedEvent{})
	case "notifychannelpasswordchanged":
		cl.publishEvent(eventName, parameters, &ChannelPasswordChangedEvent{})
	case "notifychanneldescriptionchanged":
		cl.publishEvent(eventName, parameters, &ChannelDescriptionChangedEvent{})
	case "notifychannelmoved":
		cl.publishEvent(eventName, parameters, &ChannelMovedEvent{})
	case "notifychanneledited":
		cl.publishEvent(eventName, parameters, &ChannelEditedEvent{})
	case "notifychannelcreated":
		cl.publishEvent(eventName, parameters, &ChannelCreatedEvent{})
	case "notifychanneldeleted":
		cl.publishEvent(eventName, parameters, &ChannelDeletedEvent{})
	case "notifyclientmoved":
		cl.publishEvent(eventName, parameters, &ClientMovedEvent{})
	case "notifytextmessage":
		cl.publishEvent(eventName, parameters, &TextMessageEvent{})
	case "notifytokenused":
		cl.publishEvent(eventName, parameters, &TokenUsedEvent{})
	default:
		fmt.Printf("Event %s not registered.\nPayload %s\n\n", eventName, parameters)
	}

}

func (cl *EventClient) publishEvent(eventName string, parameters url.Values, v interface{}) {
	err := cl.decoder.Decode(v, parameters)
	if err != nil {
		fmt.Println(err)
		return
	}

	cl.bus.Publish(eventName, v)
}

func splitEvent(e string) (eventName string, values url.Values) {
	split := strings.Split(e, " ")

	eventName = split[0]
	parameters := split[1:]

	values = url.Values{}
	applyParameters(parameters, &values)

	return
}

func applyParameters(parameters []string, val *url.Values) {
	for _, parameter := range parameters {
		if strings.Contains(parameter, "|") {
			splitArray := strings.Split(parameter, "|")
			applyParameters(splitArray, val)
		} else {
			applyParameter(parameter, val)
		}
	}
}

func applyParameter(parameter string, val *url.Values) {
	splitParam := strings.SplitN(parameter, "=", 2)
	if len(splitParam) == 2 {
		val.Add(splitParam[0], splitParam[1])
	}
}

func (cl *EventClient) keepAlive() {
	for {
		err := cl.write(" \n")
		if err != nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func (cl *EventClient) write(payload string) error {
	command := fmt.Sprintf("%s\n", payload)
	_, err := cl.conn.Write([]byte(command))
	return err
}

func (cl *EventClient) SwitchServer(server int) error {
	err := cl.write(fmt.Sprintf("use %d", server))
	if err != nil {
		return err
	}

	err = cl.registerNotifications()
	if err != nil {
		return err
	}

	return nil
}

func (cl *EventClient) registerNotifications() error {
	err := cl.write("servernotifyregister event=channel id=0")
	if err != nil {
		return err
	}

	err = cl.write("servernotifyregister event=server")
	if err != nil {
		return err
	}

	err = cl.write("servernotifyregister event=textserver")
	if err != nil {
		return err
	}

	err = cl.write("servernotifyregister event=textchannel")
	if err != nil {
		return err
	}

	err = cl.write("servernotifyregister event=textprivate")
	if err != nil {
		return err
	}

	return nil
}
