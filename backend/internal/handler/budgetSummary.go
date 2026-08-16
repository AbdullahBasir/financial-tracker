package handler

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (cfg *apiConfig) GetBudgetSummary(ctx context.Context, userID uuid.UUID, month string) (BudgetSummaryResponse, error) {
	budgets, err := cfg.dbQueries.GetBudgetsForMonth(ctx, sqlc.GetBudgetsForMonthParams{
		UserID: userID,
		Month:  month,
	})
	if err != nil {
		slog.Error("failed to retrieve budgets from database",
			"error", err,
			"user_id", userID,
		)
		return BudgetSummaryResponse{}, err
	}

	spending, err := cfg.dbQueries.GetMonthlySpending(ctx, sqlc.GetMonthlySpendingParams{
		UserID:  userID,
		Column2: month,
	})
	if err != nil {
		slog.Error("failed to retrieve spending budget from database",
			"error", err,
			"user_id", userID,
		)
		return BudgetSummaryResponse{}, err
	}

	spendingMap := make(map[uuid.UUID]decimal.Decimal)
	for _, spent := range spending {
		spendingMap[spent.CategoryID.UUID] = spent.TotalSpent
	}

	summary := BudgetSummaryResponse{
		Month:       month,
		Items:       make([]BudgetSummaryItem, 0, len(budgets)),
		TotalBudget: decimal.Zero,
		TotalSpent:  decimal.Zero,
	}

	for _, budget := range budgets {
		category, err := cfg.dbQueries.GetCategory(ctx, budget.CategoryID)
		if err != nil {
			slog.Error("failed to retrieve category from database",
				"error", err,
				"category_id", budget.CategoryID,
			)
			return BudgetSummaryResponse{}, err
		}

		if category.Type != "expenses" {
			return BudgetSummaryResponse{}, fmt.Errorf("budgets can only be created for expenses categories")
		}

		spent, ok := spendingMap[budget.CategoryID]
		if !ok {
			spent = decimal.Zero
		}

		remaining := budget.MonthlyLimit.Sub(spent)

		item := BudgetSummaryItem{
			CategoryID:   budget.CategoryID,
			CategoryName: category.Name,
			MonthlyLimit: budget.MonthlyLimit,
			TotalSpent:   spent,
			Remaining:    remaining,
			IsOverBudget: spent.GreaterThan(budget.MonthlyLimit),
		}

		summary.Items = append(summary.Items, item)
		summary.TotalBudget = summary.TotalBudget.Add(budget.MonthlyLimit)
		summary.TotalSpent = summary.TotalSpent.Add(spent)
	}

	summary.TotalRemaining = summary.TotalBudget.Sub(summary.TotalSpent)
	return summary, nil
}
