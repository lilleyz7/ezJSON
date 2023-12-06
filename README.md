# ezJSON

Functions to easily handle JSON in Golang.

### JSONPayLoad

```Golang
data := "transferred data"
payload := JSONPayLoad{
    Data: data
    Error: false
    Message: "this message came from the ezJSON package"
}
```

### ReadRequest
Takes parameters for 
1. http.ResponseWriter
2. *http.Request
3. destination for read data of type any
4. an int64 for the max number of bytes to be read. Enter 0 or lower to use default

```Golang
var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

err := ReadRequest(w, r, &requestPayload, 0)
if err != nil{
    //handle err
}
```
### WriteResponse
Takes parameters for 
1. http.ResponseWriter
2. status code int
3. data to be sent over the response of type JSONPayLoad
4. list of headers for the response  ---- ("Content-Type", "application/json") by default

```Golang
func DoAction(w *http.ResponseWriter)error{
    headers := http.Header{
        "Host": {"www.host.com"},
        "Authorization": {"Bearer Token"},
    }
    
    payload := JSONPayLoad{
            Error:   false,
            Message: "Message to the user",
            Data:    user,
        }
    
    return WriteResponse(w, http.StatusAccepted, payload, headers)
}
```

### WriteErrorResponse
Takes parameters for 
1. http.ResponseWriter
2. error message
3. status codes

```Golang
func LetsCheck(w *http.ResponseWriter)error{
    var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

    err := ReadRequest(w, r, &requestPayload, 0)
    if err != nil{
        return WriteErrorResponse(w, err, 400)
    } 
    //handle successfully response
}
```