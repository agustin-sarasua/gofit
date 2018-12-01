First, remember to create the bucket "gofit-lambda-functions" to store the lambda functions.

Package and Deploy the lambdas

1. api-lambda 
2. Custom Authorizer Lambda


DynamoDB Companies Table

¿How this table will be accesed?
User 
 - Get his companies and the information related to it (staff, services, etc.)
    - This is why we have a GSI by UserSub
 


- Composite key: Partition Key + Range Key

PartitionKey   |   SortKey   | Attributes
Comapny-[CompanyID] | Service-[ServiceID]#Class-[ClassID] | [UserSub]
[CompanyID] | Staff-[StaffID] | [UserSub]

Table
Company-company1 | Company-company1 | user1 | owner
Company-company1 | Staff-user2 | user2 | staff
Company-company1 | Staff-user3 | user3 | staff
Company-company2 | Company-company2 | user2 | owner
Company-company2 | Staff-user1 | user1 | staff

Global Secondary index
user1 | company-company1 
user2 | staff-user2 
user3 | staff-user3 
user2 | company-company2 
user1 | staff-user1 



Global Seconday Index
[UserSub] | [CompanyID]

GET /companies
Query GSI by UserSub

GET /companies/:id



aws dynamodb query \
    --table-name Companies \
    --key-condition-expression "UserSub = :name" \
    --expression-attribute-values  '{":name":{"S":"776d21e0-3b27-49df-a878-e0c7458c3100"}}'