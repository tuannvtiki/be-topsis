package repository

import "context"

type RuleRepositoryInterface interface {
	CreateRule(ctx context.Context, ruleDto *dto.RuleDto) (*dto.RuleDto, error)
}
