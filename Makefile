compile: streamreader notified

streamreader:
	GOOS=linux go build -o bin/streamreader functions/streamreader/*.go

notified:
	GOOS=linux go build -o bin/notified functions/notified/*.go

deploy:
	sam package --template-file template.yml --output-template-file packaged.yml --s3-bucket sampack-97068
	sam deploy --template-file ./packaged.yml --stack-name streamwrecks3 --capabilities CAPABILITY_IAM

delete:
	aws cloudformation delete-stack --stack-name streamwrecks3
