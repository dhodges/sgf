package parse

import "testing"

func TestListingZipArchive(t *testing.T) {
	sgfFileList, err := ListZipSGFfiles(zip_fixture_fpath("go4go_collection-20150118.zip"))
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
	sgfFileList, err := ListZipSGFfiles(zip_fixture_fpath("3_shusaku_games.zip"))
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
