# WeatherBFF

Personal project to learn Go. In facts, it's a normal server.
But, it's part of learning how to write BFF.

## What have done

2024/1/8:

1. add a get api to fetch data from GaoDe map.
   (params: ad_code--ad_code for city; extensions--way of forecast)
2. add a post api to fetch 24 hours forecast. (params: same as 1)
3. add test

2024/1/9:

1. Added benchmark testing – efficiency tested through extensive loop iterations.

2. Debugged and resolved an issue related to accessing a POST API through the frontend: Encountered connection refusal
   error when attempting to access a completed backend using Retrofit on Android. Forgot to include the expected
   response body.

2024/1/10:

Implemented grpc server and implemented communication between grpc server and client

2024/1/11:

1. Database added. Now bff can access grpc, and grpcsever obtains data from the database and gives it to bff.
2. Optimized project structure
3. Add test example 