# Realtime Chat WebSocket Application - Go Fiber
This application demonstrates a simple realtime chat functionality using Go (Golang) with the Fiber web framework and the WebSocket library. The application allows users to exchange messages in real-time through a web browser.
###
<div align = center> 
<img src="https://i.hizliresim.com/nm5sx32.png" height=600 />
</div>

## Application Description
The application creates a WebSocket endpoint using the Fiber framework, allowing users to establish WebSocket connections to /ws. Additionally, it serves the index.html file and other client-side resources at the / endpoint.
###
When a WebSocket connection is established, the handleWebSocket function is executed. This function adds new connections to a map and broadcasts incoming messages to all connected clients.

## Running the Application
To run the application, enter the following command in your terminal:
```bash
go run main.go
```
Your Fiber application should now be running at `http://localhost:3000`
## How to Use
1. Open your web browser and visit `http://localhost:3000` The page will display a simple HTML form.

2. Enter your username in the input field and click the "Connect" button. The page will establish a WebSocket connection.

3. You will then see a message box with an input field and a "Send" button.

4. Type a message in the input field and click the "Send" button. The message will be instantly sent to all other users.

5. You can establish additional connections using different browser tabs or devices, using the same username, and engage in real-time chat with other users.

Your application should now be working as a basic realtime chat application!
