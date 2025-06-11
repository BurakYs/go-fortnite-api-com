package fortniteapi

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var testClient *Client

const (
	testCosmeticName = "Peely"
	testCosmeticID1  = "CID_349_Athena_Commando_M_Banana"
	testCosmeticID2  = "CID_049_Athena_Commando_M_HolidayGingerbread"

	testStatsName = "BurakYhs"
	testStatsID   = "05006cb489c347beaad83551a1b9544e"

	testPlaylistID  = "Playlist_DefaultSolo"
	testCreatorCode = "Ninja"
)

func requireAPIKey(t *testing.T) {
	if testClient.APIKey == "" {
		t.Skip("API_KEY is not set")
	}
}

func TestMain(m *testing.M) {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Println("API_KEY is not set in .env file, skipping tests that require it")
	}

	testClient = NewClient(LanguageEnglish, apiKey)

	os.Exit(m.Run())
}

func TestGetAESKey(t *testing.T) {
	_, err := testClient.GetAESKey(context.Background(), AESKeyParams{})
	assert.NoError(t, err)
}

func TestGetBanners(t *testing.T) {
	_, err := testClient.GetBanners(context.Background(), BannersParams{})
	assert.NoError(t, err)
}

func TestGetBannerColors(t *testing.T) {
	_, err := testClient.GetBannerColors(nil)
	assert.NoError(t, err)
}

func TestGetAllCosmetics(t *testing.T) {
	_, err := testClient.GetAllCosmetics(context.Background(), AllCosmeticsParams{})
	assert.NoError(t, err)
}

func TestGetNewCosmetics(t *testing.T) {
	_, err := testClient.GetNewCosmetics(context.Background(), NewCosmeticsParams{})
	assert.NoError(t, err)
}

func TestGetBRCosmeticsList(t *testing.T) {
	_, err := testClient.GetBRCosmeticsList(context.Background(), BRCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetTrackCosmeticsList(t *testing.T) {
	_, err := testClient.GetTrackCosmeticsList(context.Background(), TrackCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetInstrumentCosmeticsList(t *testing.T) {
	_, err := testClient.GetInstrumentCosmeticsList(context.Background(), InstrumentCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetCarCosmeticsList(t *testing.T) {
	_, err := testClient.GetCarCosmeticsList(context.Background(), CarCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetLegoCosmeticsList(t *testing.T) {
	_, err := testClient.GetLegoCosmeticsList(context.Background(), LegoCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetLegoKitCosmeticsList(t *testing.T) {
	_, err := testClient.GetLegoKitCosmeticsList(context.Background(), LegoKitCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetBeanCosmeticsList(t *testing.T) {
	_, err := testClient.GetBeanCosmeticsList(context.Background(), BeanCosmeticsListParams{})
	assert.NoError(t, err)
}

func TestGetBRCosmeticByID(t *testing.T) {
	resp, err := testClient.GetBRCosmeticByID(context.Background(), testCosmeticID1, BRCosmeticByIDParams{})
	assert.NoError(t, err)
	assert.Equal(t, testCosmeticID1, resp.ID)
}

func TestSearchBRCosmetic(t *testing.T) {
	resp, err := testClient.SearchBRCosmetic(context.Background(), BRCosmeticSearchParams{Name: testCosmeticName})
	assert.NoError(t, err)
	assert.Equal(t, testCosmeticName, resp.Name)
}

func TestSearchBRCosmetics(t *testing.T) {
	_, err := testClient.SearchBRCosmetics(context.Background(), BRCosmeticSearchAllParams{Name: testCosmeticName})
	assert.NoError(t, err)
}

func SearchBRCosmeticByIDs(t *testing.T) {
	ids := []string{testCosmeticID1, testCosmeticID2}
	resp, err := testClient.SearchBRCosmeticByIDs(context.Background(), ids, BRCosmeticsByIDsParams{})

	assert.NoError(t, err)
	assert.Len(t, resp, 2)

	for _, cosmetic := range resp {
		assert.Contains(t, ids, cosmetic.ID)
	}
}

func TestGetCreatorCode(t *testing.T) {
	_, err := testClient.GetCreatorCode(context.Background(), CreatorCodeParams{Name: testCreatorCode})
	assert.NoError(t, err)
}

func TestGetBRMap(t *testing.T) {
	_, err := testClient.GetBRMap(context.Background(), BRMapParams{})
	assert.NoError(t, err)
}

func TestGetNews(t *testing.T) {
	_, err := testClient.GetNews(context.Background(), AllNewsParams{})
	assert.NoError(t, err)
}

func TestGetBRNews(t *testing.T) {
	_, err := testClient.GetBRNews(context.Background(), BRNewsParams{})
	assert.NoError(t, err)
}

func TestGetSTWNews(t *testing.T) {
	_, err := testClient.GetSTWNews(context.Background(), STWNewsParams{})
	assert.NoError(t, err)
}

/*func TestGetCreativeNews(t *testing.T) {
	_, err := testClient.GetCreativeNews(context.Background(), CreativeNewsParams{})
	assert.NoError(t, err)
}*/

func TestGetPlaylists(t *testing.T) {
	_, err := testClient.GetPlaylists(context.Background(), PlaylistsParams{})
	assert.NoError(t, err)
}

func TestGetPlaylistByID(t *testing.T) {
	resp, err := testClient.GetPlaylistByID(context.Background(), testPlaylistID, PlaylistByIDParams{})
	assert.NoError(t, err)
	assert.Equal(t, testPlaylistID, resp.ID)
}

func TestGetShop(t *testing.T) {
	_, err := testClient.GetShop(context.Background(), ShopParams{})
	assert.NoError(t, err)
}

func TestGetBRStatsByName(t *testing.T) {
	requireAPIKey(t)

	_, err := testClient.GetBRStatsByName(context.Background(), BRStatsByNameParams{Name: testStatsName})
	assert.NoError(t, err)
}

func TestGetBRStatsByAccountID(t *testing.T) {
	requireAPIKey(t)

	_, err := testClient.GetBRStatsByAccountID(context.Background(), testStatsID, BRStatsByIDParams{})
	assert.NoError(t, err)
}
