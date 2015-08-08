package parse

import (
	"testing"

	"github.com/dhodges/sgfinfo/sgf"
	"github.com/dhodges/sgfinfo/fixtures"
)

func TestListingZipArchive(t *testing.T) {
	sgfFileList, err := ListZipSGFfiles(fixtures.Zip_fixture_fpath("go4go_collection-20150118.zip"))
	if err != nil {
		t.Error(err)
		return
	}

	if len(sgfFileList) != 36 {
		t.Errorf("expected 36 files, but found %d", len(sgfFileList))
	}

	expectedName := "__go4go_20150111_Li-Ziqi_Hu-Aohua.sgf"
	foundName := sgfFileList[8]
	if foundName != expectedName {
		t.Errorf("wrong name: sgfFileList[8], \nexpected  %q \nbut found %q", expectedName, foundName)
	}
}

func TestListingZipArchiveOnlySGFfiles(t *testing.T) {
	sgfFileList, err := ListZipSGFfiles(fixtures.Zip_fixture_fpath("3_shusaku_games.zip"))
	if err != nil {
		t.Error(err)
		return
	}

	if len(sgfFileList) != 3 {
		t.Errorf("expected 3 files, but found %d", len(sgfFileList))
	}

	expected := "1842/Ota_Yuzo-Kuwahara_Shusaku.sgf"
	if sgfFileList[1] != expected {
		t.Errorf("zip filelist is wrong: \nfound    %q \nexpected %q", sgfFileList[1], expected)
	}
}

func TestParsingZipArchiveSGFfile(t *testing.T) {
	zipArchive := fixtures.Zip_fixture_fpath("3_shusaku_games.zip")
	games, err := ParseZipSGFfile(zipArchive, "1840/Ito_Showa-Kuwahara_Shusaku.sgf")
	if err != nil {
		t.Error(err)
		return
	}
	game := games[0]

	foundName, _ := game.GetInfo(sgf.WhitePlayerName)
	if foundName != "Ito Showa" {
		t.Errorf("wrong white player name found: %q expected: Ito Showa", foundName)
	}
	foundName, _ = game.GetInfo(sgf.BlackPlayerName)
	if foundName != "Kuwahara Shusaku" {
		t.Errorf("wrong black player name found: %q expected: %q", foundName, "Kuwahara Shusaku")
	}

	if game.NodeCount() != 202 {
		t.Errorf("wrong node count, found: %d, expected: 202", game.NodeCount())
	}
}

func TestParsingZipArchiveAllSGFfiles(t *testing.T) {
	zipArchive := fixtures.Zip_fixture_fpath("3_shusaku_games.zip")
	games, err := ParseZipAllSGFfiles(zipArchive)
	if err != nil {
		t.Error(err)
		return
	}

	if len(games) != 3 {
		t.Errorf("wrong number of games, found: %d, expected 3", len(games))
		return
	}

	foundDate, _ := games[0].GetInfo(sgf.Date)
	if foundDate != "1840-03-14" {
		t.Errorf("first game is incorrect, found %q, expected '1840-03-14'", foundDate)
	}

	foundName, _ := games[1].GetInfo(sgf.WhitePlayerName)
	if foundName != "Ota Yuzo" {
		t.Errorf("wrong white player name \nfound:   %q \nexpected: Ota Yuzo", foundName)
	}

	foundName, _ = games[2].GetInfo(sgf.WhitePlayerName)
	if foundName != "Kadono Tadazaemon" {
		t.Errorf("wrong white player name \nfound:   %q \nexpected: Kadono Tadazaemon", foundName)
	}
}
