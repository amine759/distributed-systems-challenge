import json
import requests
import xdrlib

# Define the server address and port
SERVER_ADDRESS = 'http://127.0.0.1:1234'

# Define the data structure for the sum operation
class SumRequest:
    def __init__(self, num1, num2):
        self.num1 = num1
        self.num2 = num2

# Create a new XDR packer
packer = xdrlib.Packer()

# Pack the SumRequest data structure into XDR format
def pack_sum_request(packer, request):
    packer.pack_int(request.num1)
    packer.pack_int(request.num2)

# Create a SumRequest instance with the numbers to be added
request = SumRequest(10, 5)

# Pack the SumRequest instance into XDR format
pack_sum_request(packer, request)

# Get the packed data
packed_data = packer.get_buffer()

# Construct the JSON-RPC request
payload = {
    "jsonrpc": "2.0",
    "method": "add",
    "params": packed_data.decode('latin1'),  # Convert bytes to string for JSON serialization
    "id": 1
}

# Send the JSON-RPC request to the server
response = requests.post(SERVER_ADDRESS, json=payload)

# Print the response from the server
print("Sum:", response.json()["result"])

