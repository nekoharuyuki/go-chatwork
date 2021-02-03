package chatwork

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// CreateTask can create a task to chatwork
func CreateTask(message string, token string, ids []string, roomid string) {
	// URL : https://developer.chatwork.com/ja/endpoint_rooms.html#POST-rooms-room_id-tasks
	taskapi := "https://api.chatwork.com/v2/rooms" + roomid + "/tasks"
	data := url.Values{}
	data.Add("body", message)
	data.Add("to_ids", strings.Join(ids, ","))
	PostTo(taskapi, data, token)
}

// Post can post a message to chatwork
func Post(message string, token string, roomid string) {
	// URL : https://developer.chatwork.com/ja/endpoint_rooms.html#POST-rooms-room_id-messages
	messageapi := "https://api.chatwork.com/v2/rooms/" + roomid + "/messages"
	data := url.Values{}
	data.Add("body", message)
	PostTo(messageapi, data, token)
}

// PostTo is a post method to chatwork
func PostTo(url string, data url.Values, token string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	cliant := &http.Client{Transport: tr}
	requset, _ := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	requset.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	requset.Header.Add("X-ChatWorkToken", fmt.Sprintf("%v", token))

	response, err := cliant.Do(requset)
	if err != nil {
		fmt.Println("post error")
		fmt.Println(err)
		return
	}

	fmt.Println(response.Status)

	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("result: %s\n", contents)

}
