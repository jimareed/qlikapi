package main

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
)

type ItemResourceAttributes struct {
	Id               string `json:"id"`
	CreatedDate      string `json:"createdDate"`
	Description      string `json:"description"`
	DynamicColor     string `json:"dynamicColor"`
	HasSectionAccess bool   `json:"hasSectionAccess"`
	LastReloadTime   string `json:"lastReloadTime"`
	ModifiedDate     string `json:"modifiedDate"`
	Name             string `json:"name"`
	Owner            string `json:"owner"`
	PublishTime      string `json:"publishTime"`
	Published        bool   `json:"published"`
	SpaceId          string `json:"spaceId"`
	Thumbnail        string `json:"thumbnail"`
	ResourceType     string `json:"_resourceType"`
}

type ItemResourceCustomAttributes struct {
	Dummy string `json:",omitempty"`
}


type Item struct {
	Id                       string                       `json:"id"`
	Name                     string                       `json:"name"`
	Description              string                       `json:"description"`
	ResourceAttributes       ItemResourceAttributes       `json:"resourceAttributes"`
	ResourceCreatedAt        string                       `json:"resourceCreatedAt"`
	ResourceCreatedBySubject string                       `json:"resourceCreatedBySubject"`
	ResourceCustomAttributes ItemResourceCustomAttributes `json:"resourceCustomAttributes"`
	ResourceId               string                       `json:"resourceId"`
	ResourceType             string                       `json:"resourceType"`
	SpaceId                  string                       `json:"spaceId,omitempty"`
}

type Items struct {
	Data []Item
}

func getItems(apiKey string, tenantUrl string) (Items, error) {

	var items Items

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	req, err := http.NewRequest("GET", tenantUrl+"/api/v1/items", nil)
	req.Header.Add("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return items, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		return items, err
	}

	return items, nil
}


func getItem(apiKey string, tenantUrl string, name string) (Items, error) {

	var items Items

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	req, err := http.NewRequest("GET", tenantUrl+"/api/v1/items?limit=10&query="+name+"&sort=-createdAt", nil)
	req.Header.Add("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return items, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		return items, err
	}

	return items, nil
}