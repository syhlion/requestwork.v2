# requestwork

> ⚠️ **已封存（ARCHIVED）／不再維護**
>
> 此專案為早期作品，現已封存，僅作為開發歷程的紀錄留存，請勿用於新專案。
>
> **當初的考量**：想用一個固定大小的 worker pool 限制「同時進行的 HTTP 請求數」，避免下游被打爆；以 N 個 goroutine ＋ channel 實作 job 佇列。
>
> **為什麼不成熟（不建議再用）**：
> - 「限制併發數」用標準庫的 `http.Transport.MaxConnsPerHost`，或一個 `chan struct{}` semaphore 就能達成，worker pool 屬多此一舉；
> - 每個請求額外多開一個 goroutine，且 timeout 路徑與 handler 之間存在 data race；
> - 預設 `DisableKeepAlives: true` 關閉連線重用，使內建連線池形同虛設；
> - 全 workspace 無任何獨立使用者，僅作為 [greq](https://github.com/syhlion/greq) 的注入參數存在。
>
> **現在該用什麼**：直接用標準庫 `net/http`（必要時加 `golang.org/x/sync/semaphore` 做全域上限），或成熟的 [resty](https://resty.dev) / [imroc/req](https://github.com/imroc/req)。

[![Go Report Card](https://goreportcard.com/badge/github.com/syhlion/requestwork.v2)](https://goreportcard.com/report/github.com/syhlion/requestwork.v2)
[![Build Status](https://drone.syhlion.tw/api/badges/syhlion/requestwork.v2/status.svg)](https://drone.syhlion.tw/syhlion/requestwork.v2)

a lib for go to batch processing send web request

## Install

`go get github.com/syhlion/requestwork.v2`

### Usage

```

func main() {

    // Init request
	req, err := http.NewRequest("GET", "http://tw.yahoo.com", nil)
	if err != nil {
		t.Error("request error: ", err)
	}

	// Init worker
	a := New(5)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = a.Execute(ctx, req, func(resp *http.Response, err error) error {

		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil

	})
}

```
