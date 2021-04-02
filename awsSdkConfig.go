package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
)

func configureAWS() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials("AKIATNLFYBIN5NZE2OXK", "WclXT3YX5rqQksVNznG5At7IL+haRkak5vD3eMri", ""),
	})

	if err != nil {
		log.Fatal(err)
	}

	// adding logging handler to session
	sess.Handlers.Send.PushFront(func(r *request.Request) {
		// Log every request made and its payload
		fmt.Printf("Request: %s/%s, Payload: %s",
			r.ClientInfo.ServiceName, r.Operation, r.Params)
	})

	// fmt.Println(sess)

	return sess
}
