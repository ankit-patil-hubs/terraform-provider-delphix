package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// Session
	url := "http://enginetf.dlpxdc.co/resources/json/delphix/session"
	jsonStr := []byte(`{"version":{"minor":11,"major":1,"micro":16,"type":"APIVersion"},"type":"APISession"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	x := resp.Header.Get("Set-Cookie")
	fmt.Println("response xxxxxxx:", x)

	// Authenticate
	url = "http://enginetf.dlpxdc.co/resources/json/delphix/login"
	jsonStr = []byte(`{"password":"delphix","type":"LoginRequest","target":"DOMAIN","username":"admin"}`)
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", resp.Header.Get("Set-Cookie"))
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	y := resp.Header.Get("Set-Cookie")
	fmt.Println("response yyyyyy:", y)

	// Create AppDataStagedSourceConfig

	// url = "http://enginetf.dlpxdc.co/resources/json/delphix/sourceconfig"
	// jsonStr = []byte(`{"type": "AppDataStagedSourceConfig","name": "CREATED_VIA_API_po","linkingEnabled": true,"repository": "APPDATA_REPOSITORY-4","parameters": {"name": "CREATED_VIA_API_po"}}`)
	// req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", resp.Header.Get("Set-Cookie"))
	// resp, err = client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ = ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	url = "http://enginetf.dlpxdc.co/resources/json/delphix/sourceconfig"
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", y)
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	//  GET Source Configs (List all currently available databases)
	//curl -i -b cookies.txt -X GET -H "Content-Type:application/json" http://ce1.dlpxdc.co/resources/json/delphix/sourceconfig
}
