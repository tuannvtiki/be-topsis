package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"gorm.io/gorm"
	"topsis/internal/domain/model"
	"topsis/internal/domain/repository"
)

type StatisticalDomain struct {
	statisticalRepo repository.StatisticalRepositoryInterface
}

func NewStatisticalDomain(
	statisticalRepo repository.StatisticalRepositoryInterface,
) *StatisticalDomain {
	return &StatisticalDomain{
		statisticalRepo: statisticalRepo,
	}
}

func (s *StatisticalDomain) UpsertStatistical(records []*model.ChartInfo) error {
	chartsInfo, timeChart := make([]*model.ChartInfo, 0), strings.Join([]string{time.Now().Month().String(), fmt.Sprintf("%v", time.Now().Year())}, "-")
	statistical, err := s.statisticalRepo.GetByQueries(map[string]interface{}{
		"time_chart": timeChart,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			metadataCreate, errMarshal := json.Marshal(records)
			if errMarshal != nil {
				return errMarshal
			}
			return s.statisticalRepo.Create(&model.Statistical{
				TimeChart: timeChart,
				Metadata:  string(metadataCreate),
			})
		}
		return err
	}
	err = json.Unmarshal([]byte(statistical.Metadata), &chartsInfo)
	if err != nil {
		return err
	}
	chartsInfo = append(chartsInfo, records...)
	metadataUpdate, errMarshal := json.Marshal(chartsInfo)
	if errMarshal != nil {
		return errMarshal
	}
	return s.statisticalRepo.UpdateWithMap(statistical, map[string]interface{}{
		"metadata": string(metadataUpdate),
	})
}

func (s *StatisticalDomain) GetNameFilePath(queries map[string]interface{}) (string, error) {
	statistical, err := s.statisticalRepo.GetByQueries(queries)
	if err != nil {
		return "", err
	}

	chartInfo := make([]*model.ChartInfo, 0)
	err = json.Unmarshal([]byte(statistical.Metadata), &chartInfo)
	if err != nil {
		return "", err
	}

	// Generate chart
	graph := chart.BarChart{
		Width:  1280,
		Height: 720,
		Title:  "Statistical Chart",
	}

	bars := make([]chart.Value, 0)
	for _, value := range chartInfo {
		bars = append(bars, chart.Value{
			Label: value.Day,
			Value: value.Value,
		})
	}
	graph.Bars = bars

	nameFile := strings.Join([]string{"./statistical/chart-", fmt.Sprintf("%v", time.Now().UnixNano()), ".png"}, "")
	f, _ := os.Create(nameFile)
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			return
		}
	}(f)
	err = graph.Render(chart.PNG, f)
	if err != nil {
		return "", err
	}
	return nameFile, nil
}