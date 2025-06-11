package fortniteapi

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

const (
	Version = "1.0.0"
	BaseURL = "https://fortnite-api.com"
)

var (
	ErrEmptyID  = fmt.Errorf("ID parameter cannot be empty")
	ErrNoAPIKey = fmt.Errorf("an API key is required for this request")
)

type Client struct {
	HTTPClient *http.Client
	Language
	APIKey string
}

type APIResponse struct {
	Status int             `json:"status"`
	Data   json.RawMessage `json:"data,omitempty"`
	Error  string          `json:"error,omitempty"`
}

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api error: %d - %s", e.Status, e.Message)
}

func NewClient(language Language, apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		Language:   language,
		APIKey:     apiKey,
	}
}

func (c *Client) fetch(ctx context.Context, method, endpoint string, query, body, result any) error {
	u, err := url.Parse(BaseURL + endpoint)
	if err != nil {
		return fmt.Errorf("parse URL: %w", err)
	}

	params := structToQuery(query, c.Language)
	if len(params) > 0 {
		u.RawQuery = params.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}

		bodyReader = bytes.NewReader(jsonBytes)
	}

	req, err := http.NewRequestWithContext(cmp.Or(ctx, context.Background()), method, u.String(), bodyReader)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("User-Agent", "go-fortnite-api-com/"+Version)

	if c.APIKey != "" {
		req.Header.Set("Authorization", c.APIKey)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("http do: %w", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}

	if len(bodyBytes) == 0 {
		return fmt.Errorf("empty response body, status: %d", resp.StatusCode)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return fmt.Errorf("unmarshal wrapper: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &APIError{
			Status:  resp.StatusCode,
			Message: apiResp.Error,
		}
	}

	if result != nil && len(apiResp.Data) > 0 {
		if err := json.Unmarshal(apiResp.Data, result); err != nil {
			return fmt.Errorf("unmarshal data: %w", err)
		}
	}

	return nil
}

func (c *Client) get(ctx context.Context, endpoint string, params, result any) error {
	return c.fetch(ctx, "GET", endpoint, params, nil, result)
}

func (c *Client) checkAPIKey() error {
	if c.APIKey == "" {
		return ErrNoAPIKey
	}

	return nil
}

func (c *Client) GetAESKey(ctx context.Context, params AESKeyParams) (AESKeyResponse, error) {
	var result AESKeyResponse
	err := c.get(ctx, "/v2/aes", params, &result)
	return result, err
}

func (c *Client) GetBanners(ctx context.Context, params BannersParams) (BannersResponse, error) {
	var result BannersResponse
	err := c.get(ctx, "/v1/banners", params, &result)
	return result, err
}

func (c *Client) GetBannerColors(ctx context.Context) (BannerColorsResponse, error) {
	var result BannerColorsResponse
	err := c.get(ctx, "/v1/banners/colors", nil, &result)
	return result, err
}

func (c *Client) GetAllCosmetics(ctx context.Context, params AllCosmeticsParams) (AllCosmeticsResponse, error) {
	var result AllCosmeticsResponse
	err := c.get(ctx, "/v2/cosmetics", params, &result)
	return result, err
}

func (c *Client) GetNewCosmetics(ctx context.Context, params NewCosmeticsParams) (NewCosmeticsResponse, error) {
	var result NewCosmeticsResponse
	err := c.get(ctx, "/v2/cosmetics/new", params, &result)
	return result, err
}

func (c *Client) GetBRCosmeticsList(ctx context.Context, params BRCosmeticsListParams) (BRCosmeticsListResponse, error) {
	var result BRCosmeticsListResponse
	err := c.get(ctx, "/v2/cosmetics/br", params, &result)
	return result, err
}

func (c *Client) GetTrackCosmeticsList(ctx context.Context, params TrackCosmeticsListParams) (TrackCosmeticsListResponse, error) {
	var result TrackCosmeticsListResponse
	err := c.get(ctx, "/v2/cosmetics/tracks", params, &result)
	return result, err
}

func (c *Client) GetInstrumentCosmeticsList(ctx context.Context, params InstrumentCosmeticsListParams) (InstrumentCosmeticsListResponse, error) {
	var result InstrumentCosmeticsListResponse
	err := c.get(ctx, "/v2/cosmetics/instruments", params, &result)
	return result, err
}

func (c *Client) GetCarCosmeticsList(ctx context.Context, params CarCosmeticsListParams) (CarCosmeticsListResponse, error) {
	var result CarCosmeticsListResponse
	err := c.get(ctx, "/v2/cosmetics/cars", params, &result)
	return result, err
}

func (c *Client) GetLegoCosmeticsList(ctx context.Context, params LegoCosmeticsListParams) (LegoCosmeticsListResponse, error) {
	var result LegoCosmeticsListResponse
	err := c.get(ctx, "/v2/cosmetics/lego", params, &result)
	return result, err
}

func (c *Client) GetLegoKitCosmeticsList(ctx context.Context, params LegoKitCosmeticsListParams) (LegoKitCosmeticsListResponse, error) {
	var result LegoKitCosmeticsListResponse
	err := c.get(ctx, "/v2/cosmetics/lego/kits", params, &result)
	return result, err
}

func (c *Client) GetBeanCosmeticsList(ctx context.Context, params BeanCosmeticsListParams) (BeanCosmeticsListResponse, error) {
	var result BeanCosmeticsListResponse
	err := c.get(ctx, "/v2/cosmetics/beans", params, &result)
	return result, err
}

func (c *Client) GetBRCosmeticByID(ctx context.Context, cosmeticID string, params BRCosmeticByIDParams) (BRCosmeticByIDResponse, error) {
	var result BRCosmeticByIDResponse
	if cosmeticID == "" {
		return result, ErrEmptyID
	}

	err := c.get(ctx, fmt.Sprintf("/v2/cosmetics/br/%s", cosmeticID), params, &result)
	return result, err
}

func (c *Client) SearchBRCosmetic(ctx context.Context, params BRCosmeticSearchParams) (BRCosmeticSearchResponse, error) {
	var result BRCosmeticSearchResponse
	err := c.get(ctx, "/v2/cosmetics/br/search", params, &result)
	return result, err
}

func (c *Client) SearchBRCosmetics(ctx context.Context, params BRCosmeticSearchAllParams) (BRCosmeticSearchAllResponse, error) {
	var result BRCosmeticSearchAllResponse
	err := c.get(ctx, "/v2/cosmetics/br/search/all", params, &result)
	return result, err
}

func (c *Client) GetBRCosmeticByIDs(ctx context.Context, ids []string, params BRCosmeticsByIDsParams) (BRCosmeticsByIDsResponse, error) {
	var result BRCosmeticsByIDsResponse

	if len(ids) == 0 {
		return result, ErrEmptyID
	}

	err := c.fetch(ctx, "POST", "/v2/cosmetics/br/search/ids", params, ids, &result)
	return result, err
}

func (c *Client) GetCreatorCode(ctx context.Context, params CreatorCodeParams) (CreatorCodeResponse, error) {
	var result CreatorCodeResponse
	err := c.get(ctx, "/v2/creatorcode", params, &result)
	return result, err
}

func (c *Client) GetBRMap(ctx context.Context, params BRMapParams) (BRMapResponse, error) {
	var result BRMapResponse
	err := c.get(ctx, "/v1/map", params, &result)
	return result, err
}

func (c *Client) GetAllNews(ctx context.Context, params AllNewsParams) (AllNewsResponse, error) {
	var result AllNewsResponse
	err := c.get(ctx, "/v2/news", params, &result)
	return result, err
}

func (c *Client) GetBRNews(ctx context.Context, params BRNewsParams) (BRNewsResponse, error) {
	var result BRNewsResponse
	err := c.get(ctx, "/v2/news/br", params, &result)
	return result, err
}

func (c *Client) GetSTWNews(ctx context.Context, params STWNewsParams) (STWNewsResponse, error) {
	var result STWNewsResponse
	err := c.get(ctx, "/v2/news/stw", params, &result)
	return result, err
}

func (c *Client) GetCreativeNews(ctx context.Context, params CreativeNewsParams) (CreativeNewsResponse, error) {
	var result CreativeNewsResponse
	err := c.get(ctx, "/v2/news/creative", params, &result)
	return result, err
}

func (c *Client) GetPlaylists(ctx context.Context, params PlaylistsParams) (PlaylistsResponse, error) {
	var result PlaylistsResponse
	err := c.get(ctx, "/v1/playlists", params, &result)
	return result, err
}

func (c *Client) GetPlaylistByID(ctx context.Context, playlistID string, params PlaylistByIDParams) (PlaylistByIDResponse, error) {
	var result PlaylistByIDResponse
	if playlistID == "" {
		return result, ErrEmptyID
	}

	err := c.get(ctx, fmt.Sprintf("/v1/playlists/%s", playlistID), params, &result)
	return result, err
}

func (c *Client) GetShop(ctx context.Context, params ShopParams) (ShopResponse, error) {
	var result ShopResponse
	err := c.get(ctx, "/v2/shop", params, &result)
	return result, err
}

func (c *Client) GetBRStatsByName(ctx context.Context, params BRStatsByNameParams) (BRStatsByNameResponse, error) {
	if err := c.checkAPIKey(); err != nil {
		return BRStatsByNameResponse{}, err
	}

	var result BRStatsByNameResponse
	err := c.get(ctx, "/v2/stats/br/v2", params, &result)
	return result, err
}

func (c *Client) GetBRStatsByAccountID(ctx context.Context, accountID string, params BRStatsByIDParams) (BRStatsByIDResponse, error) {
	if err := c.checkAPIKey(); err != nil {
		return BRStatsByIDResponse{}, err
	}

	var result BRStatsByIDResponse
	if accountID == "" {
		return result, ErrEmptyID
	}

	err := c.get(ctx, fmt.Sprintf("/v2/stats/br/v2/%s", accountID), params, &result)
	return result, err
}

func structToQuery(query any, defaultLanguage Language) url.Values {
	values := url.Values{}
	if query == nil {
		return values
	}

	v := reflect.ValueOf(query)
	t := reflect.TypeOf(query)

	for i := range t.NumField() {
		field := t.Field(i)
		tag := field.Tag.Get("url")
		if tag == "" {
			tag = field.Name
		} else {
			tag = strings.Split(tag, ",")[0]
		}

		val := v.Field(i)
		if val.IsZero() {
			if tag == "language" {
				values.Set(tag, string(defaultLanguage))
			}

			continue
		}

		values.Set(tag, fmt.Sprint(val.Interface()))
	}

	return values
}
