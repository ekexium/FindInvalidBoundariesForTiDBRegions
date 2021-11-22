package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pingcap/tidb/util/codec"
)

func main() {
	jsonFile, err := os.Open("region.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	regions := result["regions"].([]interface{})
	for _, region := range regions {
		regionMap := region.(map[string]interface{})
		startKeyString := regionMap["start_key"].(string)
		startKey, err := hex.DecodeString(startKeyString)
		if err != nil {
			fmt.Println(err)
		}
		endKeyString := regionMap["end_key"].(string)
		endKey, err := hex.DecodeString(endKeyString)
		if err != nil {
			fmt.Println(err)
		}

		_, _, err = codec.DecodeBytes(startKey, nil)
		if err != nil && len(startKey) > 0 {
			json, _ := json.MarshalIndent(&regionMap, "", "  ")
			fmt.Printf("\n%v\n=========================\ninvalid start key:\n%s\n", err, string(json))
		}
		_, _, err = codec.DecodeBytes(endKey, nil)
		if err != nil && len(endKey) > 0 {
			json, _ := json.MarshalIndent(&regionMap, "", "  ")
			fmt.Printf("\n%v\n=========================\ninvalid end key:\n%s\n", err, string(json))
		}
	}
}
