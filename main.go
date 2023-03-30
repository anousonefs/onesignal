package main

import "fmt"

func main() {
	// SendNotiDemo()
	println(GenSha("", ""))
}

func SendNotiDemo() {
	noti := NotiInfo{
		AppID:      "",
		ApiKey:     "",
		ExternalID: []string{""},
		Segment:    []string{""},
		Title:      "laopost",
		Message:    fmt.Sprintf("ເຄື່ອງມາຮອດແລ້ວເດີ້  ລະຫັດເຄື່ອງ:%s", ""),
		AdditionalData: map[string]interface{}{
			"id":     "",
			"status": "",
		},
	}
	if err := SendNoti(noti); err != nil {
		panic(err)
	}
	println("done")
}
