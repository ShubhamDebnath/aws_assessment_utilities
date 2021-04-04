package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
)

// Creating a new ElasticSearchService domain
// Not using this for now, need go through detailed
// configuration for elastic search domain
// for now will create a new domain using aws console ui
func createEsDomain(sess *session.Session) {

	svc := elasticsearchservice.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	input := &elasticsearchservice.CreateElasticsearchDomainInput{
		DomainName: aws.String("cupcake-updates-01"),
	}

	result, err := svc.CreateElasticsearchDomain(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
