package keys

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	keysBucket = os.Getenv("KEYS_BUCKET")
)

func Load(keyType KeyType) {
	keyFileData, err := getKeyFileData(keyType)
	if err != nil {
		log.Printf("error occurred reading keyFileData for '%s': %s", keyType, err)
		return
	}

	keyTypeValue := reflect.ValueOf(keyType).Elem()

	keys := strings.Split(string(keyFileData), "\n")

	for _, key := range keys {
		if len(key) == 0 {
			continue
		}

		keyData := strings.Split(key, "=")
		keyTypeValue.FieldByName(keyData[0]).SetString(keyData[1])
	}
}

func getKeyFileData(keyType KeyType) (data []byte, err error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	keyObjectName := fmt.Sprintf("%s.keys", keyType.GetKeyFilePrefix())

	rc, err := client.Bucket(keysBucket).Object(keyObjectName).NewReader(ctx)
	if err != nil {
		return
	}

	defer rc.Close()

	return ioutil.ReadAll(rc)
}
