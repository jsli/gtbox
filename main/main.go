package main

import (
	"fmt"
	"github.com/jsli/gtbox/archive"
	"github.com/jsli/gtbox/file"
	"github.com/jsli/gtbox/ota"
	"github.com/jsli/gtbox/pathutil"
	"os"
)

func checkErr(err error, tag string) {
	if err == nil {
		fmt.Printf("test %s [PASS]\n", tag)
	} else {
		fmt.Printf("test %s [FAIL]\n", tag)
		fmt.Printf("error: %s\n", err)
	}
}

func main() {
	testGenerateImage()
	testGenerateOtaPackage()
	testMd5sum()
	testRecordMd5()
	testMkDir()
	testUnZip()
	testZip()
	testCopyFile()
	testCopyDir()
	testWriteReader2File()
}

func testWriteReader2File() {
	src, _ := os.Open("/home/manson/temp/test/tmp/radio.img")
	dest := "/home/manson/temp/test/tmp/test.img"
	err := file.WriteReader2File(src, dest)
	checkErr(err, "file.WriteReader2File")
}

func testCopyDir() {
	//	src := "/home/manson/temp/test/tmp/sss"
	src := "/home/manson/temp/test/tmp/sss/"
	dest := "/home/manson/temp/test/tmp/aaa"
	//	dest := "/home/manson/temp/test/tmp/aaa/"
	err := file.CopyDir(src, dest)
	checkErr(err, "file.CopyDir")
}

func testCopyFile() {
	src := "/home/manson/temp/test/tmp/radio.img"
	dest := "/home/manson/temp/test/tmp/cp_radio.img"
	//	dest := "/home/manson/temp/test/tmp/sss/"
	//	dest := "/home/manson/temp/test/tmp/sss/cp.img"
	//	dest := "/home/manson/temp/test/tmp/sss/cp/"
	//	dest := "/home/manson/temp/test/tmp/sss/cp/cp.img"
	_, err := file.CopyFile(src, dest)
	checkErr(err, "file.CopyFile")
}

func testZip() {
	err := archive.ArchiveZipFile("/home/manson/temp/test/tmp/unzip/", "/home/manson/temp/test/tmp/unzip.zip")
	checkErr(err, "archive.ArchiveZipFile")
}

func testUnZip() {
	err := archive.ExtractZipFile("/home/manson/temp/test/tmp/update_pkg.zip", "/home/manson/temp/test/tmp/unzip")
	checkErr(err, "archive.ExtractZipFile")
}

func testMkDir() {
	path := "/home/manson/temp/test/tmp/dir"
	err := pathutil.MkDir(path)
	checkErr(err, "pathutil.MkDir")
}

func testGenerateImage() {
	dest_path := "/home/manson/temp/test/tmp/radio.img"
	comp_list := make([]file.Component, 4)
	comp_list[0] = file.Component{"/home/manson/temp/test/tmp/HL_TD_CP.bin", 0}
	comp_list[1] = file.Component{"/home/manson/temp/test/tmp/HL_TD_M08_AI_A0_Flash.bin", 8388608}
	comp_list[2] = file.Component{"/home/manson/temp/test/tmp/HL_TD_DSDS_CP.bin", 10485760}
	comp_list[3] = file.Component{"/home/manson/temp/test/tmp/HL_TD_M08_AI_A0_DSDS_Flash.bin", 18874368}

	err := file.GenerateImage(comp_list, dest_path, 0)
	checkErr(err, "file.GenerateImage")
}

func testGenerateOtaPackage() {
	cmd_params := make([]string, 5)
	cmd_params[0] = "--platform=jb-4.2"
	cmd_params[1] = "--product=pxa1t88ff_def"
	cmd_params[2] = "--oem=marvell"
	cmd_params[3] = "--output=/home/manson/temp/test/tmp/update.zip"
	cmd_params[4] = "--zipfile=/home/manson/temp/test/tmp/update_pkg.zip"

	err := ota.GenerateOtaPackage("/home/manson/server/ota/new/radio/updatetool/updatemk", cmd_params)
	checkErr(err, "ota.GenerateOtaPackage")
}

func testMd5sum() {
	md5, err := file.Md5Sum("/home/manson/temp/test/tmp/update.zip")
	checkErr(err, "file.Md5Sum")
	fmt.Println(md5)
}

func testRecordMd5() {
	err := ota.RecordMd5("/home/manson/temp/test/tmp/update.zip", "/home/manson/temp/test/tmp/checksum.txt")
	checkErr(err, "ota.RecordMd5")
}
