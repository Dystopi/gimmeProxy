package gimmeProxy

import(
    "encoding/json"
    "net/http"
    "strconv"
    "reflect"
)

var (
    baseUrl = "http://gimmeproxy.com/api/getProxy"
)

type Api struct {
    httpClient *http.Client
}

func NewApi() (*Api, error){
    return &Api{&http.Client{}}, nil
}

func (a *Api) GetProxy(requestModel *GimmeProxyRequestParams) (*GimmeProxyResponse, error) {
    req, err := a.buildRequest(requestModel)
    if err != nil {
        return nil, err
    }

    resp, err := a.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var resStruct GimmeProxyResponse
    json.NewDecoder(resp.Body).Decode(&resStruct)
    return &resStruct, nil
}

func (a *Api) buildRequest(requestModel *GimmeProxyRequestParams) (*http.Request, error) {
    req, err := http.NewRequest("GET", baseUrl, nil)
    if err != nil {
        return nil, err
    }

    queryParams := req.URL.Query()

    reflectedValue := reflect.ValueOf(*requestModel)
    reflectedType  := reflect.TypeOf(*requestModel)
    for i := 0; i < reflectedValue.NumField(); i++ {
        value        := reflectedValue.Field(i)
        fieldName    := reflectedType.Field(i).Tag.Get("param")
        valInterface := value.Interface()

        switch v := valInterface.(type) {
            case string:
                if valInterface.(string) == "" {
                    continue
                }
                queryParams.Add(fieldName, v)
            case bool:
                queryParams.Add(fieldName, strconv.FormatBool(v))
            case int:
                queryParams.Add(fieldName, strconv.Itoa(v))
            default:
                break
        }
    }

    req.URL.RawQuery = queryParams.Encode()
    return req, nil
}
