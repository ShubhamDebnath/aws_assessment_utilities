This Repo is only responsible for various utility functions like reading from csv, creating DynamoDB table, testing DynamoDB streams etc,
Currently, if you look at main.go, only, sending random updates to DynamoDB part is enabled
which is simply an infinite loop running every 5 mins

I have other repos which do other parts of the functionality

Commands used

GO init

```
go mod init example.com/hello
```

Role Creation

1. create role with name 892905-lambda-role:

```
aws iam create-role --role-name 892905-lambda-role --path "/service-role/"  --assume-role-policy-document file://trust-relationship.json
```

2. update the role policy using role-policy.json file

```
aws iam put-role-policy --role-name 892905-lambda-role --policy-name 892905-lambda-role-policy --policy-document file://role-policy.json
```

for some reason, the above command created a new policy, but did not attatch it to the previously created role,
so atached the policy to the role manually web console ui

Lambda function creation

1. write a lambda function handler, I have a separate repo for just the lambda handler

2. Download the build-lambda-zip tool from GitHub:

```
go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

3. navigate to GOPATH and find the zip tool in bin folder, use that to create zip file for lambda code
   Usually GOPATH is set at %USERPROFILE%

```
set GOOS=linux
go build -o main main.go
%USERPROFILE%\Go\bin\build-lambda-zip.exe -output main.zip main
```

4. Create the lambda function using above zip file

```
aws lambda create-function --function-name publishCupcakeUpdates --runtime go1.x --zip-file fileb://main.zip --handler main --role arn:aws:iam::234825976347:role/service-role/892905-lambda-role
```

5. Create trigger for the above function

```
aws lambda create-event-source-mapping \
    --region ap-south-1 \
    --function-name publishCupcakeUpdates \
    --event-source arn:aws:dynamodb:ap-south-1:234825976347:table/Cupcake_Data/stream/2021-04-02T09:55:17.156  \
    --batch-size 5 \
    --starting-position TRIM_HORIZON
```

Notes:

1. In the instructions it was suggested to use cloudwatch, but I did not create a new cloudwatch event to fire a separate lambda
   Just to prevent cost overrun on my free tier account (scared of paying too much),
   Just a matter of creating a new lambda function and uploading my random update code along with that
   for now just using a local scheduler thread to do the same

   I have created most of the tables, functions etc using AWS SDK, code for the same I have commented out in this project's main.go,
   For now the main.go only sends random 100 updates to the DynamoDB

2. In the 2nd step it was mentioned to update random 100 of the previously uploaded rows, then how will one capture delete or insert(apart from first upload) events
   so assumed only to send update requests to modify the previously uploaded ones only, this will create only modified event in elastic search, can modify the logic if required

3. I have already given commands and methods to create a lambda function, role and trigger for the 1st lamda to send updates events from DynamoDB stream to ElasticSearch service,
   for creating the 2nd lambda function, which is hosting the rest api,

```
aws lambda create-function --function-name cupcakeAPI
--runtime go1.x --zip-file fileb://main.zip --handler main --role arn:aws:iam::234825976347:role/service-role/892905-lambda-role
```
