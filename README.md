Commands used

Role Creation

1) create role with name 892905-lambda-role:
```
aws iam create-role --role-name 892905-lambda-role --path "/service-role/"  --assume-role-policy-document file://trust-relationship.json
```

2) update the role policy using role-policy.json file
```
aws iam put-role-policy --role-name 892905-lambda-role --policy-name 892905-lambda-role-policy --policy-document file://role-policy.json
```

for some reason, the above command created a new policy, but did not attatch it to the previously created role,
so atached the policy to the role manually web console ui


Lambda function creation
1) write a lambda function handler, I have a separate repo for just the lambda handler

2) Download the build-lambda-zip tool from GitHub:
```
go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

3) navigate to GOPATH and find the zip tool in bin folder, use that to create zip file for lambda code
   Usually GOPATH is set at %USERPROFILE%
```
set GOOS=linux
go build -o main main.go
%USERPROFILE%\Go\bin\build-lambda-zip.exe -output main.zip main
```

4) Create the lambda function using above zip file
```
aws lambda create-function --function-name publishCupcakeUpdates --runtime go1.x --zip-file fileb://main.zip --handler main --role arn:aws:iam::234825976347:role/service-role/892905-lambda-role
```

5) Create trigger for the above function
```
aws lambda create-event-source-mapping \
    --region ap-south-1 \
    --function-name publishCupcakeUpdates \
    --event-source arn:aws:dynamodb:ap-south-1:234825976347:table/Cupcake_Data/stream/2021-04-02T09:55:17.156  \
    --batch-size 5 \
    --starting-position TRIM_HORIZON
```

Notes:
1) In the instructions it was suggested to use cloudwatch, but I did not create a new cloudwatch event to fire a separate lambda
   Just to prevent cost overrun on my free trier account (scared of paying too much),
   Just a matter of creating a new lambda function and uploading my random update code along with that
   for now just using a local scheduler thread to do the same

2) In the 2nd step it was mentioned to update random 100 of the previously uploaded rows, then how will one capture delete or insert(apart from first upload) events
   so assumed only to send update requests to modify the previously uploaded ones only, this will create only modified event in elastic search, can modify the logic if required