package gimmeProxy

import()

type GimmeProxyResponse struct {
    Get             bool    `json:"get"`
    Post            bool    `json:"post"`
    Cookies         bool    `json:"cookies"`
    Referer         bool    `json:"referer"`
    UserAgent       bool    `json:"user-agent"`
    AnonymityLevel  int     `json:"anonymityLevel"`
    SupportsHttps   bool    `json:"supportsHttps"`
    Protocol        string  `json:"protocol"`
    IP              string  `json:"ip"`
    Port            string  `json:"port"`
    Websites        struct {
        Example bool `json:"example"`
        Google  bool `json:"google"`
        Amazon  bool `json:"amazon"`
    } `json:"websites"`
    Country         string      `json:"country"`
    TsChecked       int         `json:"tsChecked"` // EPOCH the IP was last verified as working
    Curl            string      `json:"curl"`
    IPPort          string      `json:"ipPort"`
    Type            string      `json:"type"`
    Speed           float64     `json:"speed"`
    OtherProtocols  interface{} `json:"otherProtocols"`
    Error           string      `json:"error"`
}

type GimmeProxyRequestParams struct {
    ApiKey          string  `param:"api_key"`
    Get             bool    `param:"get"`
    Post            bool    `param:"post"`
    Cookies         bool    `param:"cookies"`
    Referer         bool    `param:"referer"`
    UserAgent       bool    `param:"user-agent"`
    AnonymityLevel  int     `param:"anonymityLevel"`
    SupportsHttps   bool    `param:"supportsHttps"`
    Protocol        string  `param:"protocol"`
    Port            string  `param:"port"`
    Country         string  `param:"country"`
    Websites        string  `param:"websites"`
    MinSpeed        int     `param:"minSpeed"`
    NotCountry      string  `param:"notCountry"`
}
