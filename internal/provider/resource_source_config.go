package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AppDataStagedSourceConfig struct {
	Type            string      `json:"type"`
	Reference       string      `json:"reference"`
	Namespace       interface{} `json:"namespace"`
	Name            string      `json:"name"`
	LinkingEnabled  bool        `json:"linkingEnabled"`
	Discovered      bool        `json:"discovered"`
	EnvironmentUser string      `json:"environmentUser"`
	Repository      string      `json:"repository"`
	Toolkit         string      `json:"toolkit"`
	Parameters      struct {
		Name string `json:"name"`
	} `json:"parameters"`
}

func resourceSourceConfig() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for Oracle dSource creation.",

		CreateContext: resourceSourceConfigCreate,
		ReadContext:   resourceSourceConfigRead,
		UpdateContext: resourceSourceConfigUpdate,
		DeleteContext: resourceSourceConfigDelete,

		Schema: map[string]*schema.Schema{
			"engine_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Output
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"environment_user": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSourceConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	//GET /resources/json/delphix/action/{ref}
	var diags diag.Diagnostics
	repository_id := d.Id()

	var engine string
	var user string
	var password string

	if v, has_v := d.GetOk("engine_ip"); has_v {
		engine = v.(string)
	}
	if v, has_v := d.GetOk("user_name"); has_v {
		user = v.(string)
	}
	if v, has_v := d.GetOk("password"); has_v {
		password = v.(string)
	}

	session_cookie := session(engine)
	auth_cookie := Auth(engine, session_cookie, user, password)
	res := GET_SOURCE_API(engine, repository_id, auth_cookie)

	var response AppDataStagedSourceConfig
	// Unmarshal the JSON string into the struct
	err1 := json.Unmarshal([]byte(string(res)), &response)
	if err1 != nil {
		fmt.Println("Error unmarshalling JSON:", err1)
		//return
	}

	d.Set("id", d.Id())
	d.Set("type", response.Type)
	d.Set("reference", response.Reference)
	d.Set("name", response.Name)
	d.Set("repository", response.Repository)
	d.Set("environment_user", response.EnvironmentUser)

	return diags
}

func resourceSourceConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("Action update not implemented for resource : dSource")
}

func resourceSourceConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	repository_id := d.Id()
	var engine string
	var user string
	var password string

	if v, has_v := d.GetOk("engine_ip"); has_v {
		engine = v.(string)
	}
	if v, has_v := d.GetOk("user_name"); has_v {
		user = v.(string)
	}
	if v, has_v := d.GetOk("password"); has_v {
		password = v.(string)
	}

	session_cookie := session(engine)
	auth_cookie := Auth(engine, session_cookie, user, password)
	DELETE_SOURCE_API(engine, repository_id, auth_cookie)
	//return diag.Errorf("Action update not implemented for resource : dSource")
	return diags
}

func resourceSourceConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	var engine string
	var user string
	var password string

	if v, has_v := d.GetOk("engine_ip"); has_v {
		engine = v.(string)
	}
	if v, has_v := d.GetOk("user_name"); has_v {
		user = v.(string)
	}
	if v, has_v := d.GetOk("password"); has_v {
		password = v.(string)
	}

	session_cookie := session(engine)
	auth_cookie := Auth(engine, session_cookie, user, password)
	res := SOURCE_CREATE_API(engine, auth_cookie)
	//	d.SetId()

	// Create an empty interface to store the JSON data
	var jsonData map[string]interface{}

	// Unmarshal the JSON string into the interface
	err1 := json.Unmarshal([]byte(res), &jsonData)

	if err1 != nil {
		fmt.Println("Error unmarshalling JSON:", err1)
		//return
	}

	// Access the value of "action" from the map
	actionValue, ok := jsonData["action"].(string)
	if !ok {
		fmt.Println("Error accessing 'action' field")
		//return
	}
	// Print the value of "action"
	fmt.Println("Action Value:", actionValue)
	id, ok := jsonData["result"].(string)
	d.SetId(id)

	fmt.Println("id Value:", id)

	PollForStatusCodeEngine(engine, actionValue, auth_cookie)

	// repository_id, ok := jsonData["result"].(string)
	// GET_SOURCE_API(engine, repository_id, auth_cookie)

	// var response Response
	// // Unmarshal the JSON string into the struct
	// err1 := json.Unmarshal([]byte(string(res)), &response)
	// if err1 != nil {
	// 	fmt.Println("Error unmarshalling JSON:", err1)
	// 	//return
	// }

	// d.Set("type", response.Type)
	// d.Set("status", response.Type)
	// d.Set("result", response.Type)
	// d.Set("job", response.Type)
	// d.Set("action", response.Type)
	//d.SetId(res)

	readDiags := resourceSourceConfigRead(ctx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}
	return diags
}

func session(engine string) string {
	url := engine + "/session"
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
	return resp.Header.Get("Set-Cookie")
}

func Auth(engine string, session_cookie string, user string, password string) string {
	url := engine + "/login"
	jsonStr := []byte(`{"password":"delphix","type":"LoginRequest","target":"DOMAIN","username":"admin"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session_cookie)
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
	return resp.Header.Get("Set-Cookie")
}

func SOURCE_CREATE_API(engine string, auth_cookie string) string {
	url := engine + "/sourceconfig"
	jsonStr := []byte(`{"type": "AppDataStagedSourceConfig","name": "CREATED_VIA_API_po","linkingEnabled": true,"repository": "APPDATA_REPOSITORY-4","parameters": {"name": "CREATED_VIA_API_po"}}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", auth_cookie)
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

	return string(body)

}

func GET_SOURCE_API(engine string, id string, auth_cookie string) string {
	url := engine + "/sourceconfig/" + id
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", auth_cookie)
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

	return string(body)

}

func GET_ACTION_API(engine string, action_id string, auth_cookie string) string {
	url := engine + "/action/" + action_id
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", auth_cookie)
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

	return string(body)

}

func DELETE_SOURCE_API(engine string, id string, auth_cookie string) string {
	url := engine + "/sourceconfig/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", auth_cookie)
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

	return string(body)

}

func APITEST(engine string, auth_cookie string) string {
	url := engine + "/sourceconfig"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", auth_cookie)
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
	return string(body)
}

// poll counter is the retry counter for which an api call should be retried.
func PollForStatusCodeEngine(engine string, action_id string, auth_cookie string) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	res := GET_ACTION_API(engine, action_id, auth_cookie)

	// Create an empty interface to store the JSON data
	var jsonData map[string]interface{}
	// Unmarshal the JSON string into the interface
	err1 := json.Unmarshal([]byte(string(res)), &jsonData)
	if err1 != nil {
		fmt.Println("Error unmarshalling JSON:", err1)
	}
	// Access the value of "state" from the nested map
	result, ok := jsonData["result"].(map[string]interface{})
	if !ok {
		fmt.Println("Error accessing 'result' field")
		//return
	}

	state, ok := result["state"].(string)
	if !ok {
		fmt.Println("Error accessing 'state' field")
		//return
	}

	// Print the value of "state"
	fmt.Println("State:", state)

	// var i = 0
	// for state != Completed {
	// 	time.Sleep(time.Duration(JOB_STATUS_SLEEP_TIME) * time.Second)
	// 	res := GET_ACTION_API(engine, action_id, auth_cookie)
	// 	// Create an empty interface to store the JSON data
	// 	var jsonData map[string]interface{}
	// 	// Unmarshal the JSON string into the interface
	// 	err1 := json.Unmarshal([]byte(string(res)), &jsonData)
	// 	if err1 != nil {
	// 		fmt.Println("Error unmarshalling JSON:", err1)
	// 	}
	// 	state, ok := jsonData["state"].(string)
	// 	if !ok {
	// 		fmt.Println("Error accessing 'state' field")
	// 	}
	// 	i++
	// 	fmt.Printf("Engine-ACTION has Status:%s", state)
	// }

	if state != Completed {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(JOB_STATUS_SLEEP_TIME) * time.Second)
			res := GET_ACTION_API(engine, action_id, auth_cookie)
			var jsonData map[string]interface{}
			err1 := json.Unmarshal([]byte(string(res)), &jsonData)
			if err1 != nil {
				fmt.Println("Error unmarshalling JSON:", err1)
			}
			// Access the value of "state" from the nested map
			result, ok := jsonData["result"].(map[string]interface{})
			if !ok {
				fmt.Println("Error accessing 'result' field")
				//return
			}

			state, ok := result["state"].(string)
			if !ok {
				fmt.Println("Error accessing 'state' field")
				//return
			}

			// Print the value of "state"
			fmt.Println("State:", state)

			if state == Completed {
				break
			}
		}
	}

	return nil, diags
}
