package profiles

type Infrastructure struct {
	FruitsKVS         KVS           `json:"fruits_kvs"`
	ReputationAPI     ReputationAPI `json:"reputation_api"`
	ProposalStorageOS ObjectStorage `json:"proposal_storage_os"`
}

type KVS struct {
	Container    string `json:"container"`
	ReadRetries  int    `json:"read_retries"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteRetries int    `json:"write_retries"`
	WriteTimeout int    `json:"write_timeout"`
}

type ReputationAPI struct {
	HTTPClientTimeout                   int    `json:"http_client_timeout"`
	SellerReputationURL                 string `json:"seller_reputation_url"`
	SellerReputationCallerScopesQuery   string `json:"seller_reputation_caller_scopes_query:"`
	SellerReputationExcludedStatusCodes []int  `json:"seller_reputation_excluded_status_codes"`
}

type ObjectStorage struct {
	OSName  string `json:"os_name"`
	Retries int    `json:"retries"`
	Timeout int    `json:"timeout"`
}
