package main

import (
	"fmt"
	"os"

	v1 "github.com/alewkinr/example-grpc/proto-generate-json/sdk/go/hello/message/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	request := v1.HelloRequest{Name: "John Wick"}
	response := v1.HelloResponse{Message: fmt.Sprintf("Hello %s!", request.GetName())}

	// Convert the message to JSON
	marshalledReq, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(&request)
	if err != nil {
		fmt.Println("Error encoding request to JSON:", err)
		os.Exit(1)
	}

	marshalledResp, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(&response)
	if err != nil {
		fmt.Println("Error encoding response to JSON:", err)
		os.Exit(1)
	}

	fmt.Println("request: ", string(marshalledReq))
	fmt.Println("response: ", string(marshalledResp))
}
