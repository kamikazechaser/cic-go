package cic_meta

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func generateMetaUrl(metaBaseUrl string, key string) string {
	return fmt.Sprintf("%s/%s", metaBaseUrl, key)
}

func mergeSha256Key(x []byte, y []byte) string {
	h := sha256.New()

	h.Write(x)
	h.Write(y)

	return hex.EncodeToString(h.Sum(nil))
}

func requestHandler(cicMeta *CicMeta, metadataKey string) ([]byte, error) {
	resp, err := cicMeta.httpClient.Get(generateMetaUrl(cicMeta.baseUrl, metadataKey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching metadata for key %s: %s", metadataKey, resp.Status)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

func jsonUnmarshaler[T PersonResponse | PreferencesResponse | CustomResponse](respBody []byte, binding T) (T, error) {
	if err := json.Unmarshal(respBody, &binding); err != nil {
		return binding, err
	}

	return binding, nil
}
