# Build all the lambdas
cd api-lambda
env GOOS=linux GOARCH=amd64 go build -o /tmp/main
zip -j /tmp/companies-lambda.zip /tmp/main

cd ../staff-lambda
env GOOS=linux GOARCH=amd64 go build -o /tmp/main
zip -j /tmp/company-staff-lambda.zip /tmp/main

cd ../services-lambda
env GOOS=linux GOARCH=amd64 go build -o /tmp/main
zip -j /tmp/company-services-lambda.zip /tmp/main

# Upload zip files to s3
aws s3 cp /tmp/companies-lambda.zip s3://gofit-lambda-functions/
aws s3 cp /tmp/company-staff-lambda.zip s3://gofit-lambda-functions/
aws s3 cp /tmp/company-services-lambda.zip s3://gofit-lambda-functions/

# Update Function Code
aws lambda update-function-code --function-name GoFitCompaniesApi \
--s3-bucket gofit-lambda-functions --s3-key companies-lambda.zip

aws lambda update-function-code --function-name GoFitCompanyServicesApi \
--s3-bucket gofit-lambda-functions --s3-key company-services-lambda.zip

aws lambda update-function-code --function-name GoFitCompanyStaffApi \
--s3-bucket gofit-lambda-functions --s3-key company-staff-lambda.zip