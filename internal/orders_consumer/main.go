package main

import (
	"encoding/json"
	"fmt"

	"github.com/CarlosEduardoNop/rabbitmq-test/internal/repository"
	"github.com/CarlosEduardoNop/rabbitmq-test/pkg/rabbitmq"
)

func main() {
	ch, _ := rabbitmq.OpenChannel()

	rabbitmq.Consume(ch, func(msg []byte) {
		var input map[string]interface{}
		err := json.Unmarshal(msg, &input)

		if err != nil {
			fmt.Println(err)
		}

		repository.Create("orders", map[string]interface{}{
			"id":    input["id"],
			"name":  input["name"],
			"price": input["price"],
		})
	}, "post-orders")
}


