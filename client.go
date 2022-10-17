package squarecloud

type Client struct {
	APIToken string
	Logging  bool
}

func NewClient(APIToken string, logging bool) *Client {
	return &Client{
		APIToken: APIToken,
		Logging:  logging,
	}
}
