package options

type UriOptions struct {
	Endpoint            string
	Header              map[string]string
	TickerUri           string
	DepthUri            string
	KlineUri            string
	GetOrderUri         string
	GetPendingOrdersUri string
	GetHistoryOrdersUri string
	CancelOrderUri      string
	ClosePositionUri    string
	NewOrderUri         string
	GetAccountUri       string
	GetPositionsUri     string
	GetExchangeInfoUri  string
}

type UriOption func(*UriOptions)

func WithEndpoint(endpoint string) UriOption {
	return func(c *UriOptions) {
		c.Endpoint = endpoint
	}
}

func WithTickerUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.TickerUri = uri
	}
}

func WithDepthUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.DepthUri = uri
	}
}

func WithKlineUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.KlineUri = uri
	}
}

func WithGetOrderUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.GetOrderUri = uri
	}
}

func WithGetPendingOrdersUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.GetPendingOrdersUri = uri
	}
}

func WithHeader(header map[string]string) UriOption {
	return func(c *UriOptions) {
		c.Header = header
	}
}

func WithCancelOrderUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.CancelOrderUri = uri
	}
}

func WithClosePositionUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.ClosePositionUri = uri
	}
}

func WithNewOrderUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.NewOrderUri = uri
	}
}

func WithGetHistoryOrdersUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.GetHistoryOrdersUri = uri
	}
}

func WithGetAccountUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.GetAccountUri = uri
	}
}

func WithGetPositionsUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.GetPositionsUri = uri
	}
}

func WithGetExchangeUri(uri string) UriOption {
	return func(c *UriOptions) {
		c.GetExchangeInfoUri = uri
	}
}
