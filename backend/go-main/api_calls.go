package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func changeUsernamePayments(oldLogin, newLogin string) bool {

	addr := CONFIG_MAP["payments api addr"]
	resp, err := client.Post(addr+"/post/change-login/"+oldLogin+"/"+newLogin, "", nil)

	if err != nil {
		fmt.Println("changeUsernamePayments", err)
		return false
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("changeUsernamePayments ReadAll", err)
		return false
	}

	var res map[string]interface{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println("changeUsernamePayments Unmarshal", err)
		return false
	}
	if res["err_code"].(int) != 0 {
		fmt.Println("changeUsernamePayments respnose:", res)
		return false
	}
	return true
}

func addUser(login string) bool {
	addr := CONFIG_MAP["payments api addr"]
	resp, err := client.Post(addr+"/post/add-user/"+login, "", nil)
	fmt.Println("addUser(", login, ")", resp)
	if err != nil {
		fmt.Println("addUser", err)
		return false
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("addUser ReadAll", err)
		return false
	}

	var res map[string]interface{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println("addUser Unmarshal", err)
		return false
	}
	if res["err_code"].(int) != 0 {
		fmt.Println("addUser respnose:", res)
		return false
	}
	return true
}
