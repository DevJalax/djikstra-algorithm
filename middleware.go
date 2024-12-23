package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type MiddlewareRestConService interface {
	Send(reqResp *ReqResp) (*ReqResp, error)
}

type MiddlewareRestConServiceImpl struct {
	httpUrl                  string
	middlewareCloseableHttpClient *http.Client
}

func NewMiddlewareRestConServiceImpl(httpUrl string, client *http.Client) *MiddlewareRestConServiceImpl {
	return &MiddlewareRestConServiceImpl{
		httpUrl: httpUrl,
		middlewareCloseableHttpClient: client,
	}
}

func (s *MiddlewareRestConServiceImpl) Send(reqResp *ReqResp) (*ReqResp, error) {
	log := logrus.WithFields(logrus.Fields{
		"txnType": reqResp.TxnType,
		"txnId":   reqResp.TxnId,
		"rrn":     reqResp.Rrn,
	})
	log.Debug("middleware request", reqResp)

	url := fmt.Sprintf("%s%s", s.httpUrl, strings.ToLower(reqResp.TxnType))
	reqBody, err := json.Marshal(reqResp)
	if err != nil {
		return reqResp, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return reqResp, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("cache-control", "no-cache")

	start := time.Now()
	resp, err := s.middlewareCloseableHttpClient.Do(req)
	if err != nil {
		return reqResp, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return reqResp, err
	}

	duration := time.Since(start)
	log.WithField("duration", duration).Info("middleware response")

	var respReqResp ReqResp
	if err := json.Unmarshal(respBody, &respReqResp); err != nil {
		return reqResp, err
	}

	if respReqResp.RespCode == "" {
		respReqResp.RespCode = ConstantI.UKN
		log.Debug("Blank RespCode from Middleware")
		return nil, errors.New("blank response code from middleware")
	}

	return &respReqResp, nil
}

func getHCEErrorRespCode(txnType string) string {
	switch txnType {
	case ConstantI.DEBIT:
		return ConstantI.SWITCHCOM_HCE_DEBIT
	case ConstantI.CREDIT:
		return ConstantI.SWITCHCOM_HCE_CREDIT
	case ConstantI.REVERSAL:
		return ConstantI.SWITCHCOM_HCE_REVERSAL
	default:
		return ConstantI.UKN
	}
}
