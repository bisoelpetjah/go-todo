package response

import (
    "net/http"
    "encoding/json"
)

func CreateSuccessResponse(res http.ResponseWriter, data interface {})  {
    result, _ := json.Marshal(data)

    res.Header().Set("Content-type", "application/json")
    res.Write(result)
}

func CreateErrorResponse(res http.ResponseWriter, status int, message string)  {
    data := map[string]interface {} {
        "status": status,
        "message": message,
    }

    result, _ := json.Marshal(data)

    res.Header().Set("Content-type", "application/json")
    res.WriteHeader(status)
    res.Write(result)
}
