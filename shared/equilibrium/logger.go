package equilibrium

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	events "github.com/ChainSafe/chainbridge-substrate-events"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

const loggerPrefix = "Equilibrium Bridge (Relay)"

var logger *Logger

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
//      ResourceId   ResourceId
//      Payload      []interface{} // data associated with event sequence
// }
func Message(text string, m msg.Message) {
	if logger == nil {
		_, _ = fmt.Fprintf(os.Stderr, "Graylog writing is disabled")
		return
	}
	ctx := make([]interface{}, 0)
	ctx = append(ctx, "source_chain", m.Source)
	ctx = append(ctx, "destination_chain", m.Destination)
	ctx = append(ctx, "action", m.Type)
	ctx = append(ctx, "nonce", m.DepositNonce)
	if m.Type == msg.FungibleTransfer {
		if len(m.Payload) > 0 {
			ctx = append(ctx, "value", fmt.Sprintf("%v", m.Payload[0]))
		}
		if len(m.Payload) > 1 {
			recipient, ok := m.Payload[1].([]byte)
			if ok {
				ctx = append(ctx, "recipient", hex.EncodeToString(recipient))
			} else {
				ctx = append(ctx, "recipient", fmt.Sprintf("%v", m.Payload[1]))
			}
		}
	}
	Info(text, ctx...)
}

// type EventFungibleTransfer struct {
//  	Destination  types.U8
//  	DepositNonce types.U64
//      ResourceId   types.Bytes32
//  	Amount       types.U256
//  	Recipient    types.Bytes
// }
func EventFungibleTransfer(text string, e events.EventFungibleTransfer) {
	if logger == nil {
		_, _ = fmt.Fprintf(os.Stderr, "Graylog writing is disabled")
		return
	}
	ctx := make([]interface{}, 0)
	ctx = append(ctx, "action", msg.FungibleTransfer)
	ctx = append(ctx, "destination_chain", e.Destination)
	ctx = append(ctx, "nonce", e.DepositNonce)
	ctx = append(ctx, "value", e.Amount.Int.String())
	ctx = append(ctx, "recipient", hex.EncodeToString(e.Recipient))
	Info(text, ctx...)
}

func Info(text string, ctx ...interface{}) {
	if logger == nil {
		_, _ = fmt.Fprintf(os.Stderr, "Graylog writing is disabled")
		return
	}
	message := newMessage(text, ctx...)
	message.Level = gelf.LOG_INFO
	_, _ = fmt.Fprintf(os.Stdout, "WriteMessage: %v\n", message)
	err := logger.gelf.WriteMessage(message)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "WriteMessage error: %s\n", err.Error())
	}
}

func Warn(text string, ctx ...interface{}) {
	if logger == nil {
		_, _ = fmt.Fprintf(os.Stderr, "Graylog writing is disabled")
		return
	}
	message := newMessage(text, ctx...)
	message.Level = gelf.LOG_WARNING
	err := logger.gelf.WriteMessage(message)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "WriteMessage error: %s\n", err.Error())
	}
}

func Error(text string, ctx ...interface{}) {
	if logger == nil {
		_, _ = fmt.Fprintf(os.Stderr, "Graylog writing is disabled")
		return
	}
	message := newMessage(text, ctx...)
	message.Level = gelf.LOG_ERR
	err := logger.gelf.WriteMessage(message)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "WriteMessage error: %s\n", err.Error())
	}
}

func newMessage(text string, ctx ...interface{}) *gelf.Message {
	hostname, err := os.Hostname()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Hostname error: %s\n", err.Error())
	}

	attrs := newAttributes(ctx...)

	return &gelf.Message{
		Version:  "1.1",
		Host:     hostname,
		Short:    loggerPrefix + " " + text,
		Full:     text,
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
	if N%2 != 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Uneven set of attributes '%v'\n", ctx)
		return attrs
	}

	for i := 0; i < N-1; i += 2 {
		name := fmt.Sprintf("%v", ctx[i])
		value := ctx[i+1]
		//_, _ = fmt.Fprintf(os.Stdout, "i=%v, name='%s', value='%v'\n", i, name, value)
		_, supported := attrs[name]
		if supported {
			attrs[name] = value
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "Unsupported attribute '%s=%v'\n", name, value)
		}
	}

	return attrs
}
