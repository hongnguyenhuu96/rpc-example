package main

import (
	"log"
	"net/rpc"

	"github.com/hongnguyenhuu96/rpc/types"
)

func main() {
	var err error
	var reply types.ToDo

	// Create a TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	finishApp := types.ToDo{Title: "Finish App", Status: "Started"}
	makeDinner := types.ToDo{Title: "Make Dinner", Status: "Not Started"}
	walkDog := types.ToDo{Title: "Walk the dog", Status: "Not Started"}

	client.Call("Task.MakeToDo", finishApp, &reply)
	client.Call("Task.MakeToDo", makeDinner, &reply)
	client.Call("Task.MakeToDo", walkDog, &reply)

	client.Call("Task.DeleteToDo", makeDinner, &reply)

	client.Call("Task.MakeToDo", makeDinner, &reply)

	client.Call("Task.GetToDo", "Finish App", &reply)
	log.Println("Finish App: ", reply)

	err = client.Call("Task.EditToDo", types.EditToDo{Title: "Finish App", NewTitle: "Finish App", NewStatus: "Completed"}, &reply)
	if err != nil {
		log.Fatal("Problem editing ToDo: ", err)
	}
}
