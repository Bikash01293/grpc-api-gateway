# grpc-api-gateway

Step1: For calling any of the services you have to first start the client by running command as:

make client

![Screenshot 2022-05-14 at 5 15 15 PM](https://user-images.githubusercontent.com/48493235/168424275-e57943a1-1cfc-4872-a83b-bfc6835e6d50.png)


Step2: In a new terminal send the request as to register:

  curl --request POST \
  --url http://localhost:3000/auth/register \
  --header 'Content-Type: application/json' \
  --data '{
 "email": "bikash@gmail.com",
 "password": "hey@123" 
}'

![Screenshot 2022-05-14 at 5 07 46 PM](https://user-images.githubusercontent.com/48493235/168424316-cbd28558-3c0d-445b-ab7e-646e2d4f2182.png)

Step3: Hence to login and recieve a token send the request as:

  curl --request POST \
  --url http://localhost:3000/auth/login \
  --header 'Content-Type: application/json' \
  --data '{
 "email": "bikash@gmail.com",
 "password": "hey@123"
}'

![Screenshot 2022-05-14 at 5 19 34 PM](https://user-images.githubusercontent.com/48493235/168424496-bf8315de-658d-4da9-8c80-54041d0881c3.png)



