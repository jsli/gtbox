package main

import (
	"fmt"
	"github.com/jsli/gtbox/file"
	"github.com/jsli/gtbox/ota"
)

func main() {
	testGenerateImage()
	//	testGenerateOtaPackage()
	//	testMd5sum()
	//	testRecordMd5()
}

func testGenerateImage() {
	dest_path := "/home/manson/temp/test/tmp/radio.img"
	comp_list := make([]file.Component, 4)
	comp_list[0] = file.Component{"/home/manson/temp/test/tmp/HL_TD_CP.bin", 0}
	comp_list[1] = file.Component{"/home/manson/temp/test/tmp/HL_TD_M08_AI_A0_Flash.bin", 8388608}
	comp_list[2] = file.Component{"/home/manson/temp/test/tmp/HL_TD_DSDS_CP.bin", 10485760}
	comp_list[3] = file.Component{"/home/manson/temp/test/tmp/HL_TD_M08_AI_A0_DSDS_Flash.bin", 18874368}

	err := file.GenerateImage(comp_list, dest_path, 0)
	if err == nil {
		fmt.Println("test GenerateImage [PASS]")
	} else {
		fmt.Println("test GenerateImage [FAIL]")
		fmt.Println("error: ", err)
	}
}

func testGenerateOtaPackage() {
	cmd_params := make([]string, 5)
	cmd_params[0] = "--platform=jb-4.2"
	cmd_params[1] = "--product=pxa1t88ff_def"
	cmd_params[2] = "--oem=marvell"
	cmd_params[3] = "--output=/home/manson/temp/test/tmp/update.zip"
	cmd_params[4] = "--zipfile=/home/manson/temp/test/tmp/update_pkg.zip"

	err := ota.GenerateOtaPackage("/home/manson/server/ota/new/radio/updatetool/updatemk", cmd_params)
	if err == nil {
		fmt.Println("test GenerateOtaPackage [PASS]")
	} else {
		fmt.Println("test GenerateOtaPackage [FAIL]")
		fmt.Println("error: ", err)
	}
}

func testMd5sum() {
	md5, err := file.Md5Sum("/home/manson/temp/test/tmp/update.zip")
	if err == nil {
		fmt.Println("test GenerateOtaPackage [PASS]")
		fmt.Println(md5)
	} else {
		fmt.Println("test GenerateOtaPackage [FAIL]")
		fmt.Println("error: ", err)
	}
}

func testRecordMd5() {
	err := ota.RecordMd5("/home/manson/temp/test/tmp/update.zip", "/home/manson/temp/test/tmp/checksum.txt")
	if err == nil {
		fmt.Println("test RecordMd5 [PASS]")
	} else {
		fmt.Println("test RecordMd5 [FAIL]")
		fmt.Println("error: ", err)
	}
}
