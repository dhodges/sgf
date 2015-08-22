package tests

import (
	"testing"

	"github.com/dhodges/sgfinfo/sgf"
	"github.com/dhodges/sgfinfo/parse"
	"github.com/dhodges/sgfinfo/fixtures"
  "github.com/stretchr/testify/assert"
)

func TestListingZipArchive(t *testing.T) {
	sgfFileList, err := parse.ListZipSGFfiles(fixtures.Zip_fpath("go4go_collection-20150118.zip"))
	assert.Equal(t, err, nil, "problem listing zip fixtures")

	assert.Equal(t, len(sgfFileList), 36, "wrong file count")
	assert.Equal(t, sgfFileList[8], "__go4go_20150111_Li-Ziqi_Hu-Aohua.sgf", "wrong name: sgfFileList[8]")
}

func TestListingZipArchiveOnlySGFfiles(t *testing.T) {
	sgfFileList, err := parse.ListZipSGFfiles(fixtures.Zip_fpath("3_shusaku_games.zip"))
	assert.Equal(t, err, nil, "problem listing zip fixtures")

	assert.Equal(t, len(sgfFileList), 3, "wrong file list length")
	assert.Equal(t, sgfFileList[1],   "1842/Ota_Yuzo-Kuwahara_Shusaku.sgf", "zip filelist is wrong")
}

func TestParsingZipArchiveSGFfile(t *testing.T) {
	zipArchive := fixtures.Zip_fpath("3_shusaku_games.zip")
	games, err := parse.ParseZipSGFfile(zipArchive, "1840/Ito_Showa-Kuwahara_Shusaku.sgf")
	assert.Equal(t, err, nil, "problem parsing zip archive")

	game := games[0]
	assert.Equal(t, game.GameInfo[sgf.PlayerWhiteName], "Ito Showa", "wrong white player name found")
	assert.Equal(t, game.GameInfo[sgf.PlayerBlackName], "Kuwahara Shusaku", "wrong black player name found")
	assert.Equal(t, game.NodeCount(), 202,              "wrong node count")
}

func TestParsingZipArchiveAllSGFfiles(t *testing.T) {
	zipArchive := fixtures.Zip_fpath("3_shusaku_games.zip")
	games, err := parse.ParseZipAllSGFfiles(zipArchive)
	assert.Equal(t, err, nil, "problem parsing zip archive")

	assert.Equal(t, len(games), 3, "wrong number of games")
	assert.Equal(t, games[0].GameInfo[sgf.Date],            "1840-03-14",        "first game is incorrect")
	assert.Equal(t, games[1].GameInfo[sgf.PlayerWhiteName], "Ota Yuzo",          "wrong white player name")
	assert.Equal(t, games[2].GameInfo[sgf.PlayerWhiteName], "Kadono Tadazaemon", "wrong white player name")
}
