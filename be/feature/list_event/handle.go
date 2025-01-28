package listevent

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func Handle(ctx *Ctx) *Ctx {
	fmt.Println(ctx.Input)
	day := fmt.Sprintf("%02d", ctx.Input.Day)
	month := fmt.Sprintf("%02d", ctx.Input.Month)

	data, err := http.Get("https://api.wikimedia.org/feed/v1/wikipedia/en/onthisday/all/" + month + "/" + day)
	if err != nil {
		slog.Error(err.Error())
		return ctx.InternalServerError()
	}
	defer data.Body.Close()

	var result struct {
		Events []struct {
			Year  int    `json:"year"`
			Text  string `json:"text"`
			Pages []struct {
				Title string `json:"title"`
			} `json:"pages"`
		} `json:"events"`
	}
	err = json.NewDecoder(data.Body).Decode(&result)
	if err != nil {
		ctx.WithError("ERROR_PARSE_JSON", err.Error())
		return ctx
	}

	ctx.WithOutput(&Output{Events: make([]string, 0, len(result.Events))})
	for _, event := range result.Events {
		ctx.Output.Events = append(ctx.Output.Events, fmt.Sprintf("%d - %s", event.Year, event.Text))
	}

	return ctx
}
