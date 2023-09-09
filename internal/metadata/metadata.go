package metadata

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/srepio/sdk/types"
)

func Get() (*types.Metadata, error) {
	md := &types.Metadata{}

	mdJson, err := getJson()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(mdJson), md); err != nil {
		return nil, err
	}

	return md, nil
}

func Find(name string) (*types.Scenario, error) {
	md, err := Get()
	if err != nil {
		return nil, err
	}

	for _, s := range *md {
		if s.Name == name {
			return &s, nil
		}
	}

	return nil, fmt.Errorf("unknown senario %s", name)
}

func getJson() (string, error) {
	out, err := http.Get("https://raw.githubusercontent.com/srepio/containers/main/metadata.json")
	if err != nil {
		return "", err
	}
	defer out.Body.Close()

	body, err := io.ReadAll(out.Body)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}
