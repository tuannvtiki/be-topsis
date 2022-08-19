package model

type Field struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type Attachment struct {
	Fields   []Field `json:"fields,omitempty"`
	ImageUrl string  `json:"image_url,omitempty"`
}

type SlackMessage struct {
	Text        string        `json:"text"`
	IconEmoji   string        `json:"icon_emoji"`
	Attachments []*Attachment `json:"attachments"`
}

type TextMessageNotifyRun struct {
	CurrentTime string `json:"currentTime"`
	Location    string `json:"location"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
	IsRunning   string `json:"isRunning"`
	Note        string `json:"note"`
}

type TextMessageNotifySummary struct {
	CurrentTime  string `json:"currentTime"`
	Distance     string `json:"distance"`
	MovingTime   string `json:"moving_time"`
	AverageSpeed string `json:"average_speed"`
	MaxSpeed     string `json:"max_speed"`
	Note         string `json:"note"`
}

type TextMessageNotifyStatistical struct {
	ImageUrl string `json:"image_url"`
}

func (s *TextMessageNotifyRun) ToAttachment() []*Attachment {
	fields := make([]Field, 0)
	fields = append(fields, Field{
		Short: true,
		Title: "Current Time",
		Value: s.CurrentTime,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Location",
		Value: s.Location,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Temperature",
		Value: s.Temperature,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Weather",
		Value: s.Weather,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Running",
		Value: s.IsRunning,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Note",
		Value: s.Note,
	})

	return []*Attachment{{Fields: fields}}
}

func (s *TextMessageNotifySummary) ToAttachment() []*Attachment {
	fields := make([]Field, 0)
	fields = append(fields, Field{
		Short: true,
		Title: "Current Time",
		Value: s.CurrentTime,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Kilometers",
		Value: s.Distance,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Moving Time",
		Value: s.MovingTime,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Average Speed",
		Value: s.AverageSpeed,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Max Speed",
		Value: s.MaxSpeed,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Note",
		Value: s.Note,
	})

	return []*Attachment{{Fields: fields}}
}

func (s *TextMessageNotifyStatistical) ToAttachment() []*Attachment {
	return []*Attachment{{ImageUrl: s.ImageUrl}}
}
