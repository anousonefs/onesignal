package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api"
	"github.com/sirupsen/logrus"
)

type NotiInfo struct {
	AppID          string
	ApiKey         string
	ExternalID     []string
	Segment        []string
	Title          string
	Message        string
	AdditionalData map[string]interface{}
}

func SendNoti(data NotiInfo) (err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("helper.SendNoti(): %v\n", err)
		}
	}()
	ctx := context.WithValue(context.Background(), onesignal.AppAuth, data.ApiKey)
	notification := *onesignal.NewNotification(data.AppID)
	notification.SetIncludedSegments(data.Segment)
	notification.SetIncludeExternalUserIds(data.ExternalID)
	notification.SetData(data.AdditionalData)
	notification.SetHeadings(onesignal.StringMap{
		En: &data.Title,
	})
	notification.SetContents(onesignal.StringMap{
		En: &data.Message,
	})
	notification.UnsetIsIos()
	configuration := onesignal.NewConfiguration()
	oneCli := onesignal.NewAPIClient(configuration)
	res, r, err := oneCli.DefaultApi.CreateNotification(ctx).
		Notification(notification).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNotification`: %+v\n", res)
	return nil
}

func GenSha(userID, onesignalApiKey string) string {
	s := hmac.New(sha256.New, []byte(onesignalApiKey))
	s.Write([]byte(userID))
	sha := hex.EncodeToString(s.Sum(nil))
	return sha
}
