package utils

import "strconv"

// parsePaginationParams converts string pagination parameters to integers with defaults
func ParsePaginationParams(limit, page string) (limitInt, pageInt int) {
    limitInt, err := strconv.Atoi(limit)
    if err != nil || limitInt <= 0 {
        limitInt = 10 // Default to 10 if invalid
    }

    pageInt, err = strconv.Atoi(page)
    if err != nil || pageInt <= 0 {
        pageInt = 1 // Default to 1 if invalid
    }

    return limitInt, pageInt
}

// CalculateLastPage calculates the last page number based on total items and limit per page
func CalculateLastPage(total, limit int) int {
    return (total + limit - 1) / limit
}