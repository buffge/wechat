package utils

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostXML(url string, obj interface{}) ([]byte, error) {
	xmlData, err := xml.Marshal(obj)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(xmlData)
	resp, err := http.Post(url, "application/xml;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : url=%v , statusCode=%v", url, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
