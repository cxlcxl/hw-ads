package curl

import (
	"bs.mobgi.cc/app/vars"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	AllowMethodGet  = "GET"
	AllowMethodPost = "POST"
	AllowMethodPut  = "PUT"
)

type Curl struct {
	host        string
	request     *http.Request
	data        io.Reader
	method      string
	debug       bool
	showResBody bool // 打印响应报文(debug 时生效)
}

func New(host string) *Curl {
	return &Curl{
		host:   host,
		method: AllowMethodGet,
		debug:  vars.YmlConfig.GetBool("Debug"),
	}
}

func (c *Curl) Request(v interface{}, headers ...HeaderHandlers) (err error) {
	c.request, err = http.NewRequest(c.method, c.host, c.data)
	if err != nil {
		return err
	}
	for _, header := range headers {
		header(c)
	}
	res, err := (&http.Client{}).Do(c.request)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var bs []byte
	bs, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if c.debug && c.showResBody {
		fmt.Println()
		fmt.Println("===> 响应报文：", string(bs))
		fmt.Println()
	}
	err = json.Unmarshal(bs, v)
	if err != nil {
		if res.StatusCode != 200 {
			return errors.New(res.Status)
		}
		return err
	}
	return nil
}

func (c *Curl) Get() *Curl {
	c.method = AllowMethodGet
	return c
}

func (c *Curl) Post() *Curl {
	c.method = AllowMethodPost
	return c
}

func (c *Curl) Put() *Curl {
	c.method = AllowMethodPut
	return c
}

func (c *Curl) Debug(b bool) *Curl {
	c.debug = b
	return c
}

func (c *Curl) ResBody(b bool) *Curl {
	c.showResBody = b
	return c
}
