package notify

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"topsis/client"
	"topsis/client/model"
	"topsis/client/strava"
	"topsis/client/upload_images"
	"topsis/client/weather"
	"topsis/config"
	internalModel "topsis/internal/domain/model"
	"topsis/internal/domain/usecase"
)

type ProcessNotifyBotInterface interface {
	ProcessNotifyRun() error
	ProcessNotifySummary() error
	ProcessNotifyStatistical() error
}

const (
	Location       = "Ha Noi, Viet Nam"
	Text           = "Weather information"
	Celsius        = "°C"
	FormatDateTime = "02-01-2006 15:04:05"
	FormatDate     = "02-01-2006"
	Latitude       = "21.0245"
	Longitude      = "105.8412"
	Units          = "metric"
	Exclude        = "minutely,hourly,daily,alerts"
	DistanceGoal   = 3000
)

type BotNotify struct {
	cfg               *config.Config
	statisticalDomain *usecase.StatisticalDomain
}

func NewBotNotify(cfg *config.Config, statisticalDomain *usecase.StatisticalDomain) *BotNotify {
	return &BotNotify{
		cfg:               cfg,
		statisticalDomain: statisticalDomain,
	}
}

func (b *BotNotify) ProcessNotifyRun() error {
	weatherInfo, err := weather.GetWeatherInfo("https://openweathermap.org/data/2.5/onecall", &model.ParamOpenWeather{
		Lat:     Latitude,
		Lon:     Longitude,
		Units:   Units,
		Appid:   b.cfg.AppID,
		Exclude: Exclude,
	})
	if err != nil {
		logrus.Errorf("get weather information from open-weather fail %v", err)
		return err
	}
	textMessage := &model.TextMessageNotifyRun{
		CurrentTime: time.Now().Format(FormatDateTime),
		Location:    Location,
		Temperature: fmt.Sprintf("%v %v feels like %v %v", weatherInfo.Current.Temp, Celsius, weatherInfo.Current.FeelsLike, Celsius),
		Weather:     weatherInfo.Current.Weather[0].Description,
		IsRunning:   "Yes",
		Note:        fmt.Sprintf("Hôm nay %v, trời không mưa nên chạy đi nhé", time.Now().Format(FormatDate)),
	}
	if weatherInfo.Current.Weather[0].Main == "Rain" {
		textMessage.IsRunning = "No"
		textMessage.Note = fmt.Sprintf("Hôm nay %v, trời mưa rồi nên ở nhà đi nhé", time.Now().Format(FormatDate))
	}
	message := &model.SlackMessage{
		Text:        Text,
		IconEmoji:   client.DefaultEmoji,
		Attachments: textMessage.ToAttachment(),
	}
	return client.SendMessageSlack(b.cfg.WebhookSlack, message)
}

func (b *BotNotify) ProcessNotifySummary() error {
	stravaActivities, err := strava.GetStravaActivityInfo(&model.ParamStrava{
		ClientId:     b.cfg.ClientId,
		ClientSecret: b.cfg.ClientSecret,
		RefreshToken: b.cfg.RefreshToken,
		GrantType:    b.cfg.GrantType,
	})
	if err != nil || len(stravaActivities) == 0 {
		return errors.New("strava activities information empty")
	}

	timeChart, messageActives, chartsInfo := time.Now().AddDate(0, 0, -1).Format(FormatDate), make([]*model.TextMessageNotifySummary, 0), make([]*internalModel.ChartInfo, 0)
	for _, activity := range stravaActivities {
		if timeChart != activity.StartDateLocal.Format(FormatDate) {
			continue
		}
		textMessage := &model.TextMessageNotifySummary{
			CurrentTime:  time.Now().Format(FormatDateTime),
			Distance:     fmt.Sprintf("%v km", math.Round((activity.Distance/float64(1000))*100)/100),
			MovingTime:   time.Duration(activity.MovingTime * 1000000000).String(),
			AverageSpeed: fmt.Sprintf("%.2f km/h", activity.AverageSpeed*60*60/float64(1000)),
			MaxSpeed:     fmt.Sprintf("%.2f km/h", activity.MaxSpeed*60*60/float64(1000)),
			Note:         fmt.Sprintf("Chúc mừng bạn đã hoàn thành mục tiêu chạy %v km ngày hôm qua %v nhé", DistanceGoal/1000, time.Now().AddDate(0, 0, -1).Format(FormatDate)),
		}
		if activity.Distance < DistanceGoal {
			textMessage.Note = fmt.Sprintf("Bạn đã không hoàn thành mục tiêu chạy %v km ngày hôm qua %v rồi :sleepy: ", DistanceGoal/1000, time.Now().AddDate(0, 0, -1).Format(FormatDate))
		}
		messageActives = append(messageActives, textMessage)
		chartsInfo = append(chartsInfo, &internalModel.ChartInfo{
			Day:   timeChart,
			Value: math.Round((activity.Distance/float64(1000))*100) / 100,
		})
	}

	if len(messageActives) == 0 {
		messageActives = append(messageActives, &model.TextMessageNotifySummary{
			CurrentTime:  time.Now().Format(FormatDateTime),
			Distance:     fmt.Sprintf("%v km", 0),
			MovingTime:   "0h0m0s",
			AverageSpeed: fmt.Sprintf("%v km/h", 0),
			MaxSpeed:     fmt.Sprintf("%v km/h", 0),
			Note:         fmt.Sprintf("Hôm qua %v, bạn không chạy à :rage:", time.Now().AddDate(0, 0, -1).Format(FormatDate)),
		})
	}

	err = b.statisticalDomain.UpsertStatistical(chartsInfo)
	for _, messageActivity := range messageActives {
		message := &model.SlackMessage{
			Text:        "Activity Information",
			IconEmoji:   client.DefaultEmoji,
			Attachments: messageActivity.ToAttachment(),
		}
		err = client.SendMessageSlack(b.cfg.WebhookSlack, message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BotNotify) ProcessNotifyStatistical() error {
	nameFile, err := b.statisticalDomain.GetNameFilePath(map[string]interface{}{
		"time_chart": strings.Join([]string{time.Now().Month().String(), fmt.Sprintf("%v", time.Now().Year())}, "-"),
	})
	if err != nil {
		return err
	}

	imageInfo, err := upload_images.UploadImage(&model.ParamUploadImage{
		Path:   nameFile,
		ApiKey: b.cfg.ApiKeyUploadImage,
	})
	if err != nil {
		logrus.Errorf("get weather information from open-weather fail %v", err)
		return err
	}

	// Send message to slack
	textMessage := &model.TextMessageNotifyStatistical{
		ImageUrl: imageInfo.Image.Url,
	}

	message := &model.SlackMessage{
		Text:        "Statistical Chart",
		IconEmoji:   client.DefaultEmoji,
		Attachments: textMessage.ToAttachment(),
	}
	return client.SendMessageSlack(b.cfg.WebhookSlack, message)
}