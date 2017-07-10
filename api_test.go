package gimmeProxy

import(
    "encoding/json"
    "net/http"
    "testing"

    "github.com/Dystopi/sslMock"
)

var sslMockClient *sslMock.Mock
var mockResponse  *GimmeProxyResponse
var mockRequest   *GimmeProxyRequestParams

func init() {
    initMockResponse()
    initMockRequest()
    initSslMockClient()
}

func TestNewApi(t *testing.T) {
    _, err := NewApi()
    if err != nil {
        t.Errorf("Recieved an unexpected error from NewApi() : %v", err.Error())
    }
}

func TestGetProxy(t *testing.T) {
    api := getTestApi()
    proxy, err := api.GetProxy(mockRequest)
    if err != nil {
        t.Errorf("Recieved an unexpected error from GetProxy() : %v", err.Error())
    }

    if proxy.IP != "107.151.136.202" {
        t.Error("Failed to recieve the expected response")
    }
}

func TestBuildRequest(t *testing.T) {
    api := getTestApi()
    req, err := api.buildRequest(mockRequest)
    if err != nil {
        t.Errorf("Recieved an unexpected error from buildRequest() : %v", err.Error())
    }

    if req.URL.Query().Get("protocol") != "http" {
        t.Error("Failed to recieve the expected Query Param Value")
    }
}

func getTestApi() (*Api) {
    api, _ := NewApi()
    api.httpClient = sslMockClient.Client
    return api
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    json.NewEncoder(w).Encode(mockResponse)
}

func initSslMockClient() {
    var err error
    sslMockClient, err = sslMock.NewMock(reqHandler)
    if err != nil {
        panic(err)
    }
}

func initMockResponse() {
    mockResponse = &GimmeProxyResponse{
        Get:            true,
        Post:           true,
        Cookies:        true,
        Referer:        true,
        UserAgent:      true,
        AnonymityLevel: 1,
        SupportsHttps:  true,
        Protocol:       "http",
        IP:             "107.151.136.202",
        Port:           "80",
        Country:        "US",
        TsChecked:      1460536850,
        Curl:           "//107.151.136.202:80",
        IPPort:         "107.151.136.202:80",
        Type:           "http",
        Speed:          0,
        Error:          "",
    }
}

func initMockRequest() {
    mockRequest = &GimmeProxyRequestParams{
        Get:            true,
        Post:           true,
        Cookies:        true,
        Referer:        true,
        UserAgent:      true,
        AnonymityLevel: 1,
        SupportsHttps:  true,
        Protocol:       "http",
        Country:        "US",
        MinSpeed:       0,
    }
}
