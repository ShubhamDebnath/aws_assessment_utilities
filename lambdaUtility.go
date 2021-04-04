package main

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// I have already created a simple lambda function
// placed its zip file in same folder as this package

// Did not use this for now
// some error related to signature validation
// used aws cli command instead
func createLambdaFunction(sess *session.Session) {

	svc := lambda.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	dat, err := ioutil.ReadFile("./main.zip")

	if err != nil {
		fmt.Println("Error reading zip file")
		panic(err)
	}

	input := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{
			ZipFile: dat,
		},
		FunctionName: aws.String("publishCupcakeUpdates"),
		Handler:      aws.String("main"),
		MemorySize:   aws.Int64(256),
		Publish:      aws.Bool(true),
		Role:         aws.String("arn:aws:iam::234825976347:role/service-role/892905-lambda-role"),
		Runtime:      aws.String("go1.x"),
		Timeout:      aws.Int64(15),
		TracingConfig: &lambda.TracingConfig{
			Mode: aws.String("Active"),
		},
	}

	result, err := svc.CreateFunction(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			case lambda.ErrCodeResourceNotFoundException:
				fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
			case lambda.ErrCodeResourceConflictException:
				fmt.Println(lambda.ErrCodeResourceConflictException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeCodeStorageExceededException:
				fmt.Println(lambda.ErrCodeCodeStorageExceededException, aerr.Error())
			case lambda.ErrCodeCodeVerificationFailedException:
				fmt.Println(lambda.ErrCodeCodeVerificationFailedException, aerr.Error())
			case lambda.ErrCodeInvalidCodeSignatureException:
				fmt.Println(lambda.ErrCodeInvalidCodeSignatureException, aerr.Error())
			case lambda.ErrCodeCodeSigningConfigNotFoundException:
				fmt.Println(lambda.ErrCodeCodeSigningConfigNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}

// basically create a new event source mapping
// by specifying the latest stream arn for our
// dynamo db table
func createLambdaTrigger(sess *session.Session, streamArn string) {

	svc := lambda.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	input := &lambda.CreateEventSourceMappingInput{
		BatchSize:        aws.Int64(5),
		EventSourceArn:   aws.String(streamArn),
		FunctionName:     aws.String("publishCupcakeUpdates"),
		StartingPosition: aws.String("TRIM_HORIZON"),
	}

	result, err := svc.CreateEventSourceMapping(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			case lambda.ErrCodeResourceConflictException:
				fmt.Println(lambda.ErrCodeResourceConflictException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeResourceNotFoundException:
				fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
