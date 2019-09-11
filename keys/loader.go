package keys

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

func Load(keyType KeyType, env string) {
	appRoot, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	appRoot = "/keys"

	keysRoot := fmt.Sprintf("%s/.keys", appRoot)
	keyFile := fmt.Sprintf("%s/%s.%s.keys", keysRoot, keyType.GetKeyFilePrefix(), env)

	v := reflect.ValueOf(keyType).Elem()

	keyFileData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Printf("error occurred reading keyFile %s: %s", keyFile, err)
		return
	}

	keys := strings.Split(string(keyFileData), "\n")
	for _, key := range keys {
		if len(key) == 0 {
			continue
		}

		keyData := strings.Split(key, "=")
		v.FieldByName(keyData[0]).SetString(keyData[1])
	}
}
