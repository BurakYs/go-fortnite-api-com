package fortniteapi

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testClient *Client

const (
	testCosmetic    = "Peely"
	testCreatorCode = "Ninja"
	testStats       = "BurakYhs"
)

func TestMain(m *testing.M) {
	apiKey := ""
	testClient = NewClient(LanguageEnglish, apiKey)
	testClient.SetContext(context.TODO())

	os.Exit(m.Run())
}

func TestGetAESKey(t *testing.T) {
	_, err := testClient.GetAESKey(nil, AESKeyParams{})
	assert.NoError(t, err)
}

func TestGetBanners(t *testing.T) {
	_, err := testClient.GetBanners(nil, BannersParams{})
	assert.NoError(t, err)
}

func TestGetBannerColors(t *testing.T) {
	_, err := testClient.GetBannerColors(nil)
	assert.NoError(t, err)
}

func TestGetAllCosmetics(t *testing.T) {
	_, err := testClient.GetAllCosmetics(nil, AllCosmeticsParams{})
	assert.NoError(t, err)
}

func TestGetNewCosmetics(t *testing.T) {
	_, err := testClient.GetNewCosmetics(nil, NewCosmeticsParams{})
	assert.NoError(t, err)
}

func TestGetBRCosmeticsList(t *testing.T) {
	_, err := testClient.GetBRCosmeticsList(nil, BRCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetTrackCosmeticsList(t *testing.T) {
	_, err := testClient.GetTrackCosmeticsList(nil, TrackCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetInstrumentCosmeticsList(t *testing.T) {
	_, err := testClient.GetInstrumentCosmeticsList(nil, InstrumentCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetCarCosmeticsList(t *testing.T) {
	_, err := testClient.GetCarCosmeticsList(nil, CarCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetLegoCosmeticsList(t *testing.T) {
	_, err := testClient.GetLegoCosmeticsList(nil, LegoCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetLegoKitCosmeticsList(t *testing.T) {
	_, err := testClient.GetLegoKitCosmeticsList(nil, LegoKitCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetBeanCosmeticsList(t *testing.T) {
	_, err := testClient.GetBeanCosmeticsList(nil, BeanCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestSearchBRCosmetic(t *testing.T) {
	_, err := testClient.SearchBRCosmetic(nil, BRCosmeticSearchParams{Name: testCosmetic})
	assert.NoError(t, err)
}

func TestSearchBRCosmetics(t *testing.T) {
	_, err := testClient.SearchBRCosmetics(nil, BRCosmeticSearchAllParams{Name: testCosmetic})
	assert.NoError(t, err)
}

func TestGetCreatorCode(t *testing.T) {
	_, err := testClient.GetCreatorCode(nil, CreatorCodeParams{Name: testCreatorCode})
	assert.NoError(t, err)
}

func TestGetBRMap(t *testing.T) {
	_, err := testClient.GetBRMap(nil, BRMapParams{})
	assert.NoError(t, err)
}

func TestGetNews(t *testing.T) {
	_, err := testClient.GetAllNews(nil, AllNewsParams{})
	assert.NoError(t, err)
}

func TestGetBRNews(t *testing.T) {
	_, err := testClient.GetBRNews(nil, BRNewsParams{})
	assert.NoError(t, err)
}

func TestGetSTWNews(t *testing.T) {
	_, err := testClient.GetSTWNews(nil, STWNewsParams{})
	assert.NoError(t, err)
}

/*func TestGetCreativeNews(t *testing.T) {
	_, err := testClient.GetCreativeNews(nil, CreativeNewsParams{})
	assert.NoError(t, err)
}*/

func TestGetPlaylists(t *testing.T) {
	_, err := testClient.GetPlaylists(nil, PlaylistsParams{})
	assert.NoError(t, err)
}

func TestGetPlaylistByID(t *testing.T) {
	playlists, err := testClient.GetPlaylists(nil, PlaylistsParams{})
	assert.NoError(t, err)

	_, err = testClient.GetPlaylistByID(nil, playlists[0].ID, PlaylistByIDParams{})
	assert.NoError(t, err)
}

func TestGetShop(t *testing.T) {
	_, err := testClient.GetShop(nil, ShopParams{})
	assert.NoError(t, err)
}

/*func TestGetBRStats(t *testing.T) {
	_, err := testClient.GetBRStats(nil, BRStatsByNameParams{Name: testStats})
	assert.NoError(t, err)
}

func TestGetBRStatsByAccountID(t *testing.T) {
	stats, err := testClient.GetBRStats(nil, BRStatsByNameParams{Name: testStats})
	noErr := assert.NoError(t, err)

	if !noErr {
		return
	}

	id := stats.Account.ID
	_, err = testClient.GetBRStatsByAccountID(nil, id, BRStatsByIDParams{})
	assert.NoError(t, err)
}*/

func TestGetBRCosmeticByID(t *testing.T) {
	list, err := testClient.GetBRCosmeticsList(nil, BRCosmeticsListParams{})
	assert.NoError(t, err)

	_, err = testClient.GetBRCosmeticByID(nil, list[0].ID, BRCosmeticByIDParams{})
	assert.NoError(t, err)
}

func TestGetBRCosmeticByIDs(t *testing.T) {
	list, err := testClient.GetBRCosmeticsList(nil, BRCosmeticsListParams{})
	assert.NoError(t, err)

	ids := []string{list[0].ID, list[1].ID}
	resp, err := testClient.GetBRCosmeticByIDs(nil, ids, BRCosmeticsByIDsParams{})
	assert.NoError(t, err)
	assert.Len(t, resp, 2)
}
