package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
)

func listStreams(sess *session.Session) string {
	svc := dynamodbstreams.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	input := &dynamodbstreams.ListStreamsInput{
		TableName: aws.String("Cupcake_Data"),
	}

	result, err := svc.ListStreams(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodbstreams.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodbstreams.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodbstreams.ErrCodeInternalServerError:
				fmt.Println(dynamodbstreams.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return ""
	}

	/*
		fmt.Println(result)
	*/
	if len(result.Streams) > 0 {
		return *result.Streams[0].StreamArn
	}

	return ""
}

func describeStream(sess *session.Session, streamArn string) []string {
	svc := dynamodbstreams.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	var shardIds []string

	input := &dynamodbstreams.DescribeStreamInput{
		StreamArn: aws.String(streamArn),
	}

	result, err := svc.DescribeStream(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodbstreams.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodbstreams.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodbstreams.ErrCodeInternalServerError:
				fmt.Println(dynamodbstreams.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil
	}

	/*
		fmt.Println(result)
	*/

	for _, shard := range result.StreamDescription.Shards {

		shardIds = append(shardIds, *shard.ShardId)
	}

	return shardIds
}

func getShardIterator(sess *session.Session, streamArn string, shardId string) *string {
	svc := dynamodbstreams.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	input := &dynamodbstreams.GetShardIteratorInput{
		ShardId: aws.String(shardId),

		// Determines how the shard iterator is used to start reading stream records
		// from the shard:
		//
		//    * AT_SEQUENCE_NUMBER - Start reading exactly from the position denoted
		//    by a specific sequence number.
		//
		//    * AFTER_SEQUENCE_NUMBER - Start reading right after the position denoted
		//    by a specific sequence number.
		//
		//    * TRIM_HORIZON - Start reading at the last (untrimmed) stream record,
		//    which is the oldest record in the shard. In DynamoDB Streams, there is
		//    a 24 hour limit on data retention. Stream records whose age exceeds this
		//    limit are subject to removal (trimming) from the stream.
		//
		//    * LATEST - Start reading just after the most recent stream record in the
		//    shard, so that you always read the most recent data in the shard.
		//
		// ShardIteratorType is a required field

		ShardIteratorType: aws.String("TRIM_HORIZON"),
		StreamArn:         aws.String(streamArn),
	}

	result, err := svc.GetShardIterator(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodbstreams.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodbstreams.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodbstreams.ErrCodeInternalServerError:
				fmt.Println(dynamodbstreams.ErrCodeInternalServerError, aerr.Error())
			case dynamodbstreams.ErrCodeTrimmedDataAccessException:
				fmt.Println(dynamodbstreams.ErrCodeTrimmedDataAccessException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil
	}

	fmt.Println(result)

	return result.ShardIterator

}

func getRecords(sess *session.Session, shardIterator string) {
	svc := dynamodbstreams.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	input := &dynamodbstreams.GetRecordsInput{
		ShardIterator: aws.String(shardIterator),
	}

	result, err := svc.GetRecords(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodbstreams.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodbstreams.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodbstreams.ErrCodeLimitExceededException:
				fmt.Println(dynamodbstreams.ErrCodeLimitExceededException, aerr.Error())
			case dynamodbstreams.ErrCodeInternalServerError:
				fmt.Println(dynamodbstreams.ErrCodeInternalServerError, aerr.Error())
			case dynamodbstreams.ErrCodeExpiredIteratorException:
				fmt.Println(dynamodbstreams.ErrCodeExpiredIteratorException, aerr.Error())
			case dynamodbstreams.ErrCodeTrimmedDataAccessException:
				fmt.Println(dynamodbstreams.ErrCodeTrimmedDataAccessException, aerr.Error())
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
