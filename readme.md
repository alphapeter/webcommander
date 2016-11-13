#Webcommander
##Summary
Simple web application written in go that allows you to exeute commands, scripts and make http requests using an http rest API.
The result of the executed command, script or http request will be returned as a result to the GET requests.

*The server does currently not implement https. If it would be used publicly, it should be placed behind an SSL termination proxy*

##Dependencies
To build the application you will need to install the [go framework](https://golang.org/)
 
##Building the application
type `go build ` to create the binary. It is also possible to cross compile for other platforms, see the go language documentation for details.

##Running the application
**From source**  
type `go run *.go` or `go run *.go --settings path-to-settings/settings.json`
**Compiled binary**  
By default the application will look for a `settings.json` in the execution path, to use settings file located in an other directory, supply the parameter `--settings path-to-settings/settings.json`, ex. `./webcommander --settings /etc/webcommanded/settings.json`

##Configuration
The configuration is done using a JSON file to define address, an optional api token, commands and http poxy requests.
By default the application will look for a `settings.json` in the execution path 

* address: ex. `":8080"` or `"192.168.1.1:8080"`
* apiToken: ex. `"9zk1HT7027716xk8z4PY08L5MiyZP6qi"`
* commands: ex. a json list of commands ex `[command]` where commands is composed of path, command, arguments and description
    * command: ex. `"ls"` - the command to be executed
    * arguments: ex. `["/", "-la"]` - the arguments for the command
    * path: ex. `"/ls"` - the path for the server, the result for the example would be `http://localhost:8080/ls`
    * description: a description for the endpoint, accessible from the root path `http://localhost:8080` ex. `"lists all files"`
* proxyRequests:
    * path: ex. `"/ip"`
    * url: ex. `"https://api.ipify.org?format=json"`
    * description: ex. `"shows the server external ip"`

### Sample configuration file
*settings.json*
```json
{
  "address": ":8080",
  "apiToken":"",
  "commands": [{
    "path": "/restart-vpn",
    "command": "/etc/init.d/vpn",
    "arguments": ["restart"],
    "description": "restart vpn"
  },
    {
      "path": "/start-vpn",
      "command": "/etc/init.d/vpn",
      "arguments": ["start"],
      "description": "start vpn"
    },
    {
      "path": "/stop-vpn",
      "command": "/etc/init.d/vpn",
      "arguments": ["stop"],
      "description": "stop vpn"
    }],
  "proxyRequests":[{
    "path": "/ip",
    "url": "https://api.ipify.org?format=json",
    "description": "shows the server external ip"
  }]
}

```

##Making requests

**List available commands**    
Make a http GET request to the root path ex. `http://localhost:8080`

**Execute a command**   
Make an http GET request to the corresponding path ex. `http://localhost:8080/restart-server` where `/restart-server` is the path

**Authorization token**   
The token can be added to the request using either URL parameter or header field, read the Authorzation section below.

##Authorization
 
**No authorization**  
The authorization is optional, if the configuration property if left empty, no apiToken is required.

**Required authorization**  
Any string could be used as apiToken, ex. password or generated string.
*Be aware that when using unencrypted http requests anyone could potentially read the token in plain text.*
 
**Adding token to the GET request**  

* Using url parameter  
    Append the token to the url using apiToken=<token>, ex. `http://localhost/restart-server?apiToken=9zk1HT7027716xk8z4PY08L5MiyZP6qi`
* Using header field   
    Add the header field `apiToken` with the token as a value

