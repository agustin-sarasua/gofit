###############################
####### Facebook Configuration
###############################

FB AppID: 353981615348312
FB App Secret: xxx

FB App Domain:

Valid OAuth Redirect URIs:
https://gofit-app.auth.us-east-1.amazoncognito.com/oauth2/idpresponse

App Domains:
https://gofit-app.auth.us-east-1.amazoncognito.com/

Website Site URL:
https://gofit-app.auth.us-east-1.amazoncognito.com/oauth2/idpresponse

###############################
####### Cognito Configuration
###############################

> Identity Providers
1. Configure Facebook
2. Add attribute mappings

> App Client Settings
1. Enable Facebook and Cognito
2. Callback URL: https://www.example.com (for now)
3. https://www.example.com/signout.html (for now)
4. Allowed OAuth Flows  > Auth code grant and Implicit grant 
5. Allowed OAuth Scopes > email, openid (for now)

> Domain name
1. https://gofit-app.auth.us-east-1.amazoncognito.com/


ClientID = 4mfq5inrdfhfoepq88uullo9u2
Authorization Code Flow
https://gofit-app.auth.us-east-1.amazoncognito.com/login?response_type=code&client_id=2ivcpvtbr2ddeu7ccmu4usqbkm&redirect_uri=https://www.example.com
Then, you will have to exchange the code for an auth token using the auth endpoint.

Authorization Token Flow
https://gofit-app.auth.us-east-1.amazoncognito.com/login?response_type=token&client_id=2ivcpvtbr2ddeu7ccmu4usqbkm&redirect_uri=https://www.example.com

Remember that you need to map the attributes from the provider to the user pool
