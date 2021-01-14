package equilibrium

import (
	"fmt"
	"os"
	"time"

	"github.com/ChainSafe/chainbridge-utils/msg"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

var logger *Logger

const LoggerPrefix = "Equilibrium Bridge (Relay) "

type Logger struct {
	gelf *gelf.TCPWriter
}

func CreateGrayLogger(addr string) error {
	if logger != nil {
		return fmt.Errorf("equilibrium logger already created")
	}

	gelfWriter, err := gelf.NewTCPWriter(addr)
	if err != nil {
		return err
	}

	logger = &Logger{gelfWriter}

	return nil
}

// type Message struct {
//  	Source       ChainId      // Source where message was initiated
//  	Destination  ChainId      // Destination chain of message
//  	Type         TransferType // type of bridge transfer
//  	DepositNonce Nonce        // Nonce for the deposit
// }
func Message(msg string, m msg.Message) {
	if logger == nil {
		return
	}
	ctx := make([]interface{}, 0, 8)
	ctx = append(ctx, "source_chain", m.Source)
	ctx = append(ctx, "destination_chain", m.Destination)
	ctx = append(ctx, "action", m.Type)
	ctx = append(ctx, "nonce", m.DepositNonce)
	Info(msg, ctx)
}

func Info(msg string, ctx ...interface{}) {
	if logger == nil {
		return
	}
	message := newMessage(msg, ctx...)
	message.Level = gelf.LOG_INFO
	err := logger.gelf.WriteMessage(message)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func Warn(msg string, ctx ...interface{}) {
	if logger == nil {
		return
	}
	message := newMessage(msg, ctx...)
	message.Level = gelf.LOG_WARNING
	err := logger.gelf.WriteMessage(message)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func Error(msg string, ctx ...interface{}) {
	if logger == nil {
		return
	}
	message := newMessage(msg, ctx...)
	message.Level = gelf.LOG_ERR
	err := logger.gelf.WriteMessage(message)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func newMessage(msg string, ctx ...interface{}) *gelf.Message {
	hostname, err := os.Hostname()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
	attrs := newAttributes(ctx)
	return &gelf.Message{
		Version:  "1.1",
		Host:     hostname,
		Short:    msg,
		Full:     LoggerPrefix + msg,
		TimeUnix: float64(time.Now().UnixNano()) / float64(time.Second),
		Level:    gelf.LOG_INFO,
		Facility: "Equilibrium",
		Extra:    attrs,
	}
}

func newAttributes(ctx ...interface{}) map[string]interface{} {
	attrs := map[string]interface{}{
		"action":            nil,
		"environment":       nil,
		"source_chain":      nil,
		"destination_chain": nil,
		"sender":            nil,
		"recipient":         nil,
		"value":             nil,
		"token":             nil,
		"nonce":             nil,
		"block":             nil,
	}
	N := len(ctx)
	for i := 0; i < N; i++ {
		name := fmt.Sprintf("%v", ctx[i])
		j := i + 1
		if j >= N {
			_, _ = fmt.Fprintf(os.Stderr, "Uneven set of attributes '%v'\n", ctx)
			break
		}
		value := ctx[j]
		if name == "action" {
			attrs[name] = value
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "Unsupported attribute '%s=%v'\n", name, value)
		}
	}
	return attrs
}

// type Message struct {
//  	Version  string                 `json:"version"`
//  	Host     string                 `json:"host"`
// 	    Short    string                 `json:"short_message"`
//  	Full     string                 `json:"full_message,omitempty"`
//  	TimeUnix float64                `json:"timestamp"`
// 	    Level    int32                  `json:"level,omitempty"`
// 	    Facility string                 `json:"facility,omitempty"`
// 	    Extra    map[string]interface{} `json:"-"`
// 	    RawExtra json.RawMessage        `json:"-"`
// }
