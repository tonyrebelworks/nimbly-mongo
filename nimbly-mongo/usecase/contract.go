package usecase

import (
	"strconv"
	"time"

	"chi-rest/bootstrap"
	"chi-rest/usecase/viewmodel"

	"github.com/andboson/carbon"
)

// UC default usecase dependencies
type UC struct {
	*bootstrap.App
}

func today() time.Time {
	// loc, _ := time.LoadLocation("Asia/Jakarta")
	// now := time.Now().In(loc)

	// return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 00, time.UTC)
	cb := carbon.Now()
	cb.SetTZ("UTC")

	return cb.Time
}

// SimplePaginationRes ...
func SimplePaginationRes(types string, maxID, firstID, lastID, limit int) viewmodel.SimplePaginationVM {
	var (
		prevPage string
		nextPage string
	)

	if maxID != 0 {
		prevPage = "?types=prev&max_id=" + strconv.Itoa(firstID) + "&limit=" + strconv.Itoa(limit)
	}
	nextPage = "?types=next&max_id=" + strconv.Itoa(lastID) + "&limit=" + strconv.Itoa(limit)

	pagination := viewmodel.SimplePaginationVM{
		PrevPage: prevPage,
		NextPage: nextPage,
	}

	return pagination
}
