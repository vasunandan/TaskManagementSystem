package constants

type ResponseStatus int

type APIResponse struct {
    StatusCode string `json:"status_code"`
    Message    string `json:"message"`
    Data       any    `json:"data,omitempty"`
}


const (
	Success       ResponseStatus = 1
	DataNotFound  ResponseStatus = 2
	UnknownError  ResponseStatus = 3
	InvalidRequest ResponseStatus = 4
	Unauthorized  ResponseStatus = 5
)

func (r ResponseStatus) GetResponseStatus() string {
    switch r {
    case Success:
        return "SUCCESS"
    case DataNotFound:
        return "DATA_NOT_FOUND"
    case UnknownError:
        return "UNKNOWN_ERROR"
    case InvalidRequest:
        return "INVALID_REQUEST"
    case Unauthorized:
        return "UNAUTHORIZED"
    default:
        return "UNKNOWN"
    }
}

func (r ResponseStatus) GetResponseMessage() string {
    switch r {
    case Success:
        return "Success"
    case DataNotFound:
        return "Data Not Found"
    case UnknownError:
        return "Unknown Error"
    case InvalidRequest:
        return "Invalid Request"
    case Unauthorized:
        return "Unauthorized"
    default:
        return "Something went wrong"
    }
}
