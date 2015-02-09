from websocket import create_connection
import json
ws = create_connection("ws://127.0.0.1:8080/ws")
payload = {
	"ClientId": "RUSSELL2",
	"Mode": 2
}

print "Sending..."
ws.send(json.dumps(payload))
print "Sent"
print "Reeiving..."
while True:
	result =  ws.recv()
	print "Received '%s'" % result
ws.close()
