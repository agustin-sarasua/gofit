# DynamoDB Table Classes
[Class-...] [Company-...#Service-...#Class-] [UserSub] [CompanyID] [ServiceID] ...

# Services

POST /classes


GET /classes?companyID=...&serviceID=...&startTime=...&limit=...
- Returns the classes for that user's companies ordered by startTime asc.
- All parameters are optional.

GET /classes/{classID}
- Return Class Data

