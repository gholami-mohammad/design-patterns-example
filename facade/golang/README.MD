# Problem:
We have a service that send our notification. Everytime that we need to send a notification, we have to follow these steps:
- Load credentials from config
- Make a request to its api server
- Handle the response

# What to do?
Create a Facade service to handle all those steps in a single funtion.
