package usecase

import (
	"context"
	"encoding/json"
	"sync"

	"topsis/internal/domain/model"
	"topsis/internal/domain/repository"
)

type ConsultDomain struct {
	userRepo        repository.UserRepositoryInterface
	standardRepo    repository.StandardRepositoryInterface
	scoreRatingRepo repository.ScoreRatingRepositoryInterface
}

func NewConsultDomain(
	userRepo repository.UserRepositoryInterface,
	standardRepo repository.StandardRepositoryInterface,
	scoreRatingRepo repository.ScoreRatingRepositoryInterface,
) *ConsultDomain {
	return &ConsultDomain{
		userRepo:        userRepo,
		standardRepo:    standardRepo,
		scoreRatingRepo: scoreRatingRepo,
	}
}

func (c *ConsultDomain) Consult(ctx context.Context, userId string) ([]*model.ConsultResult, error) {
	// Step 1: Create an evaluation matrix consisting of m alternatives and n criteria
	var (
		wg       sync.WaitGroup
		weights  []float32
		matrix   [][]float32
		listName []string
		err      error
	)

	// Get list standards
	wg.Add(1)
	go func() {
		defer wg.Done()
		standards, scopedErr := c.standardRepo.GetStandardByQueries(ctx, map[string]interface{}{
			"user_id": userId,
		})
		if scopedErr != nil {
			err = scopedErr
			return
		}
		// Parse weight
		weights = ParseWeight(standards)
	}()

	// Get list score ratings
	wg.Add(1)
	go func() {
		defer wg.Done()
		scoreRatings, scopedErr := c.scoreRatingRepo.GetScoreRatingByListQueries(ctx, map[string]interface{}{
			"user_id": userId,
		})
		if scopedErr != nil {
			err = scopedErr
			return
		}
		// Parse metadata score rating
		var metadataStruct []*model.MetadataScoreRating
		for _, value := range scoreRatings {
			scopedErr = json.Unmarshal([]byte(value.Metadata), &metadataStruct)
			if scopedErr != nil {
				err = scopedErr
				return
			}
			var score []float32
			for _, v := range metadataStruct {
				score = append(score, float32(v.Score))
			}
			matrix = append(matrix, score)
			if len(metadataStruct) > 0 {
				listName = append(listName, metadataStruct[0].Name)
			}
		}
	}()

	wg.Wait()
	if err != nil {
		return nil, err
	}

	// Step 2: The matrix is then normalised to form the matrix
	matrix = Normalize(matrix)

	// Step 3: Calculate the weighted normalised decision matrix
	matrix = CalculateTheWeighted(matrix, weights)

	// Step 4: Determine the worst alternative and the best alternative

	return nil, nil
}

func ParseWeight(standards []*model.Standard) (weights []float32) {
	for _, value := range standards {
		weights = append(weights, float32(value.Weight))
	}
	return weights
}

func Normalize(matrix [][]float32) [][]float32 {
	for col := 0; col < len(matrix[0]); col++ {
		var sumSquare float32
		for row := 0; row < len(matrix); row++ {
			sumSquare += matrix[row][col] * matrix[row][col]
		}
		// Set data is normalized
		for row := 0; row < len(matrix); row++ {
			matrix[row][col] /= sumSquare
		}
	}
	return matrix
}

func CalculateTheWeighted(matrix [][]float32, weights []float32) [][]float32 {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			matrix[row][col] *= weights[col]
		}
	}
	return matrix
}
