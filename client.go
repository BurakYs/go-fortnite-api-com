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
	"time"
)

const (
	Version = "1.2.0"
	BaseURL = "https://fortnite-api.com"
)

var ErrNoAPIKey = fmt.Errorf("an API key is required for this request")

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

type Client struct {
	HTTPClient *http.Client
	Language
	APIKey string
}

func NewClient(language Language, apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		Language:   language,
		APIKey:     apiKey,
	}
}

func (c *Client) Fetch(ctx context.Context, method, endpoint string, query, body, result any) error {
	u, err := url.Parse(BaseURL + endpoint)
	if err != nil {
		return fmt.Errorf("parse URL: %w", err)
	}

	params, err := structToQuery(query, c.Language)
	if err != nil {
		return fmt.Errorf("query conversion: %w", err)
	}

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
		return fmt.Errorf("unmarshal response: %w", err)
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

func (c *Client) Get(ctx context.Context, endpoint string, params, result any) error {
	return c.Fetch(ctx, "GET", endpoint, params, nil, result)
}

func (c *Client) checkAPIKey() error {
	if c.APIKey == "" {
		return ErrNoAPIKey
	}

	return nil
}

func (c *Client) GetAESKey(ctx context.Context, params AESKeyParams) (AESKeyResponse, error) {
	var result AESKeyResponse
	err := c.Get(ctx, "/v2/aes", params, &result)
	return result, err
}

func (c *Client) GetBanners(ctx context.Context, params BannersParams) (BannersResponse, error) {
	var result BannersResponse
	err := c.Get(ctx, "/v1/banners", params, &result)
	return result, err
}

func (c *Client) GetBannerColors(ctx context.Context) (BannerColorsResponse, error) {
	var result BannerColorsResponse
	err := c.Get(ctx, "/v1/banners/colors", nil, &result)
	return result, err
}

func (c *Client) GetAllCosmetics(ctx context.Context, params AllCosmeticsParams) (AllCosmeticsResponse, error) {
	var result AllCosmeticsResponse
	err := c.Get(ctx, "/v2/cosmetics", params, &result)
	return result, err
}

func (c *Client) GetNewCosmetics(ctx context.Context, params NewCosmeticsParams) (NewCosmeticsResponse, error) {
	var result NewCosmeticsResponse
	err := c.Get(ctx, "/v2/cosmetics/new", params, &result)
	return result, err
}

func (c *Client) GetBRCosmeticsList(ctx context.Context, params BRCosmeticsListParams) (BRCosmeticsListResponse, error) {
	var result BRCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/br", params, &result)
	return result, err
}

func (c *Client) GetTrackCosmeticsList(ctx context.Context, params TrackCosmeticsListParams) (TrackCosmeticsListResponse, error) {
	var result TrackCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/tracks", params, &result)
	return result, err
}

func (c *Client) GetInstrumentCosmeticsList(ctx context.Context, params InstrumentCosmeticsListParams) (InstrumentCosmeticsListResponse, error) {
	var result InstrumentCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/instruments", params, &result)
	return result, err
}

func (c *Client) GetCarCosmeticsList(ctx context.Context, params CarCosmeticsListParams) (CarCosmeticsListResponse, error) {
	var result CarCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/cars", params, &result)
	return result, err
}

func (c *Client) GetLegoCosmeticsList(ctx context.Context, params LegoCosmeticsListParams) (LegoCosmeticsListResponse, error) {
	var result LegoCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/lego", params, &result)
	return result, err
}

func (c *Client) GetLegoKitCosmeticsList(ctx context.Context, params LegoKitCosmeticsListParams) (LegoKitCosmeticsListResponse, error) {
	var result LegoKitCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/lego/kits", params, &result)
	return result, err
}

func (c *Client) GetBeanCosmeticsList(ctx context.Context, params BeanCosmeticsListParams) (BeanCosmeticsListResponse, error) {
	var result BeanCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/beans", params, &result)
	return result, err
}

func (c *Client) GetBRCosmeticByID(ctx context.Context, cosmeticID string, params BRCosmeticByIDParams) (BRCosmeticByIDResponse, error) {
	var result BRCosmeticByIDResponse

	if cosmeticID == "" {
		return result, fmt.Errorf("ID parameter cannot be empty")
	}

	err := c.Get(ctx, fmt.Sprintf("/v2/cosmetics/br/%s", cosmeticID), params, &result)
	return result, err
}

func (c *Client) SearchBRCosmetic(ctx context.Context, params SearchBRCosmeticParams) (SearchBRCosmeticResponse, error) {
	var result SearchBRCosmeticResponse
	err := c.Get(ctx, "/v2/cosmetics/br/search", params, &result)
	return result, err
}

func (c *Client) SearchBRCosmetics(ctx context.Context, params SearchBRCosmeticsParams) (SearchBRCosmeticsResponse, error) {
	var result SearchBRCosmeticsResponse
	err := c.Get(ctx, "/v2/cosmetics/br/search/all", params, &result)
	return result, err
}

func (c *Client) SearchBRCosmeticsByIDs(ctx context.Context, ids []string, params BRCosmeticsByIDsParams) (BRCosmeticsByIDsResponse, error) {
	var result BRCosmeticsByIDsResponse

	if len(ids) == 0 {
		return result, fmt.Errorf("IDs parameter cannot be empty")
	}

	err := c.Fetch(ctx, "POST", "/v2/cosmetics/br/search/ids", params, ids, &result)
	return result, err
}

func (c *Client) GetCreatorCode(ctx context.Context, params CreatorCodeParams) (CreatorCodeResponse, error) {
	var result CreatorCodeResponse

	if params.Name == "" {
		return result, fmt.Errorf("name parameter cannot be empty")
	}

	err := c.Get(ctx, "/v2/creatorcode", params, &result)
	return result, err
}

func (c *Client) GetBRMap(ctx context.Context, params BRMapParams) (BRMapResponse, error) {
	var result BRMapResponse
	err := c.Get(ctx, "/v1/map", params, &result)
	return result, err
}

func (c *Client) GetNews(ctx context.Context, params AllNewsParams) (AllNewsResponse, error) {
	var result AllNewsResponse
	err := c.Get(ctx, "/v2/news", params, &result)
	return result, err
}

func (c *Client) GetBRNews(ctx context.Context, params BRNewsParams) (BRNewsResponse, error) {
	var result BRNewsResponse
	err := c.Get(ctx, "/v2/news/br", params, &result)
	return result, err
}

func (c *Client) GetSTWNews(ctx context.Context, params STWNewsParams) (STWNewsResponse, error) {
	var result STWNewsResponse
	err := c.Get(ctx, "/v2/news/stw", params, &result)
	return result, err
}

func (c *Client) GetCreativeNews(ctx context.Context, params CreativeNewsParams) (CreativeNewsResponse, error) {
	var result CreativeNewsResponse
	err := c.Get(ctx, "/v2/news/creative", params, &result)
	return result, err
}

func (c *Client) GetPlaylists(ctx context.Context, params PlaylistsParams) (PlaylistsResponse, error) {
	var result PlaylistsResponse
	err := c.Get(ctx, "/v1/playlists", params, &result)
	return result, err
}

func (c *Client) GetPlaylistByID(ctx context.Context, playlistID string, params PlaylistByIDParams) (PlaylistByIDResponse, error) {
	var result PlaylistByIDResponse

	if playlistID == "" {
		return result, fmt.Errorf("ID parameter cannot be empty")
	}

	err := c.Get(ctx, fmt.Sprintf("/v1/playlists/%s", playlistID), params, &result)
	return result, err
}

func (c *Client) GetShop(ctx context.Context, params ShopParams) (ShopResponse, error) {
	var result ShopResponse
	err := c.Get(ctx, "/v2/shop", params, &result)
	return result, err
}

func (c *Client) GetBRStatsByName(ctx context.Context, params BRStatsByNameParams) (BRStatsByNameResponse, error) {
	var result BRStatsByNameResponse

	if err := c.checkAPIKey(); err != nil {
		return result, err
	}

	if params.Name == "" {
		return result, fmt.Errorf("name parameter cannot be empty")
	}

	err := c.Get(ctx, "/v2/stats/br/v2", params, &result)
	return result, err
}

func (c *Client) GetBRStatsByAccountID(ctx context.Context, accountID string, params BRStatsByIDParams) (BRStatsByIDResponse, error) {
	var result BRStatsByIDResponse

	if err := c.checkAPIKey(); err != nil {
		return result, err
	}

	if accountID == "" {
		return result, fmt.Errorf("ID parameter cannot be empty")
	}

	err := c.Get(ctx, fmt.Sprintf("/v2/stats/br/v2/%s", accountID), params, &result)
	return result, err
}
