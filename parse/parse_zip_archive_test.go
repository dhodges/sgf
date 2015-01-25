package parse

import "testing"

func TestListingZipArchive(t *testing.T) {
	sgfFileList, err := ListZipSGFfiles(zip_fixture_fpath("go4go_collection-20150118.zip"))
	if err != nil {
		t.Error(err)
		return
	}

	if len(sgfFileList) != 36 {
		t.Errorf("expected %d files, but found %d", 36, len(sgfFileList))
	}

	expectedName := "__go4go_20150111_Zhao-Wei_Wang-Chen.sgf"
	foundName := sgfFileList[8]
	if foundName != expectedName {
		t.Errorf("wrong name: sgfFileList[8], expected %q but found %q", expectedName, foundName)
	}
}
