package client

import (
	"github.com/roson9527/trustsql-go/api/asset"
	"github.com/roson9527/trustsql-go/httprequest"
	"github.com/roson9527/trustsql-go/util"
)

type Client struct {
	//HttpClient IHttpClient
	Signer *util.Signer
	Asset  *asset.Asset
}

func NewClient(prvKey []byte, host string, httpClient httprequest.IHttpRequest) (*Client, error) {
	var err error

	c := new(Client)

	c.Signer, err = util.NewSigner(prvKey)
	if err != nil {
		return nil, err
	}

	c.Asset = asset.New(host, httpClient)

	return c, nil
}
