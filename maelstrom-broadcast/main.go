package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {

	n := maelstrom.NewNode()

	var messages = make([]float64, 10)
	var topology = make(map[string]any)

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		var message float64 = body["message"].(float64)
		body["type"] = "broadcast_ok"
		messages = append(messages, message)

		return n.Reply(msg, map[string]string{"type": "broadcast_ok"})
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		_ = topology
		topology = body["topology"].(map[string]any)
		body["type"] = "topology_ok"

		return n.Reply(msg, map[string]any{"type": "topology_ok"})

	})

	n.Handle("read", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body["type"] = "read_ok"
		body["messages"] = messages

		return n.Reply(msg, map[string]any{"type": "read_ok", "messages": messages})

	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
