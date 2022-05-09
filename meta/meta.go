package meta

import (
	"github.com/ethereum/go-ethereum/common"
	"net/http"
	"strings"
	"time"
)

const (
	personPointer      = ":cic.person"
	phonePointer       = ":cic.phone"
	preferencesPointer = ":cic.preferences"
	customPointer      = ":cic.custom"
)

type CicMeta struct {
	httpClient *http.Client
	baseUrl    string
}

type CustomResponse struct {
	Tags []string `json:"tags"`
}

type PreferencesResponse struct {
	PreferredLanguage string `json:"preferred_language"`
}

type PersonResponse struct {
	DateRegistered int         `json:"date_registered"`
	VCard          string      `json:"vcard"`
	Gender         string      `json:"gender"`
	Location       Location    `json:"location"`
	Products       []string    `json:"products"`
	DateOfBirth    DateOfBirth `json:"date_of_birth"`
}

type Location struct {
	AreaName string `json:"area_name"`
}

type DateOfBirth struct {
	Year int `json:"year"`
}

func NewCicMeta(metaEndpoint string) *CicMeta {
	return &CicMeta{
		httpClient: &http.Client{
			Timeout: time.Second * 3,
		},
		baseUrl: metaEndpoint,
	}
}

func (c *CicMeta) GetPhonePointer(phone string) (string, error) {
	hashedKey := mergeSha256Key([]byte(phone), []byte(phonePointer))

	r, err := requestHandler(c, hashedKey)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(r), "\""), nil
}

func (c *CicMeta) GetPersonMetadata(address string) (PersonResponse, error) {
	hashedKey := mergeSha256Key(common.HexToAddress(address).Bytes(), []byte(personPointer))

	respData, err := requestHandler(c, hashedKey)
	if err != nil {
		return PersonResponse{}, err
	}

	metadata, err := jsonUnmarshaler(respData, PersonResponse{})
	if err != nil {
		return PersonResponse{}, err
	}

	return metadata, nil
}

func (c *CicMeta) GetPreferencesMetadata(address string) (PreferencesResponse, error) {
	hashedKey := mergeSha256Key(common.HexToAddress(address).Bytes(), []byte(preferencesPointer))

	respData, err := requestHandler(c, hashedKey)
	if err != nil {
		return PreferencesResponse{}, err
	}

	metadata, err := jsonUnmarshaler(respData, PreferencesResponse{})
	if err != nil {
		return PreferencesResponse{}, err
	}

	return metadata, nil
}

func (c *CicMeta) GetCustomMetadata(address string) (CustomResponse, error) {
	hashedKey := mergeSha256Key(common.HexToAddress(address).Bytes(), []byte(customPointer))

	respData, err := requestHandler(c, hashedKey)
	if err != nil {
		return CustomResponse{}, err
	}

	metadata, err := jsonUnmarshaler(respData, CustomResponse{})
	if err != nil {
		return CustomResponse{}, err
	}

	return metadata, nil
}
