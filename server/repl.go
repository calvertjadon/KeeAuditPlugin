package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/calvertjadon/KeeAuditPlugin/internal/pubsub"
	"github.com/calvertjadon/KeeAuditPlugin/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func getInput() []string {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return nil
	}
	line := scanner.Text()
	line = strings.TrimSpace(line)
	return strings.Fields(line)
}

type payload struct {
	Message string `json:"message"`
}

func enterRepl(ch *amqp.Channel) {
	running := true
	for running {
		words := getInput()
		if len(words) == 0 {
			continue
		}

		var err error
		switch words[0] {
		case "test":
			log.Println("sending pause message")
			err = pubsub.PublishJSON(
				ch,
				routing.ExchangeAuditCommands,
				routing.RunAuditPrefix+".all",
				payload{Message: "hello from go"},
			)
		default:
			log.Println("unknown command")
		}

		if err != nil {
			log.Printf("Error: %s", err)
		}
	}
}
