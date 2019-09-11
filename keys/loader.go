package keys

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

func Load(keyType KeyType, env string) {
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

	env := "stage"
	keyObjectName := fmt.Sprintf("%s.%s.keys", keyType.GetKeyFilePrefix(), env)

	rc, err := client.Bucket("hellodoctor-staging-keys").Object(keyObjectName).NewReader(ctx)
	if err != nil {
		return
	}

	defer rc.Close()

	return ioutil.ReadAll(rc)
}
