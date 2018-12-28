package apiv1

type Response struct {
    Success     bool        `json:"success"`
    ErrMsg      string      `json:"err_msg"`
    ApiEndpoint string      `json:"api_endpoint"`
    Payload     interface{} `json:"payload"`
}

type AccountTypeRequest struct {
    Label       string  `json:"label"`
    IsGlobal    bool    `json:"is_global"`
}
