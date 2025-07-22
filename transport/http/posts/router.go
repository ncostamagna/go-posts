package posts

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/ncostamagna/go-http-utils/response"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ctxKey string

const (
	ctxParam  ctxKey = "params"
	ctxHeader ctxKey = "header"
	ctxQuery  ctxKey = "query"
)

func NewHTTPServer(_ context.Context, endpoints Endpoints) http.Handler {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}
	r.Use(ginDecode())

	r.GET("/posts", gin.WrapH(httptransport.NewServer(endpoints.GetAll, decodeGetAllHandler, encodeResponse, opts...)))
	r.POST("/posts", gin.WrapH(httptransport.NewServer(endpoints.Store, decodeStoreHandler, encodeResponse, opts...)))

	r.GET("/posts/:id", gin.WrapH(httptransport.NewServer(endpoints.Get, decodeGetHandler, encodeResponse, opts...)))
	r.PATCH("/posts/:id", gin.WrapH(httptransport.NewServer(endpoints.Update, decodeUpdateHandler, encodeResponse, opts...)))
	r.DELETE("/posts/:id", gin.WrapH(httptransport.NewServer(endpoints.Delete, decodeDeleteHandler, encodeResponse, opts...)))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}

func ginDecode() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ctxParam, c.Params)
		ctx = context.WithValue(ctx, ctxHeader, c.Request.Header)
		ctx = context.WithValue(ctx, ctxQuery, c.Request.URL.Query())
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func decodeGetHandler(ctx context.Context, _ *http.Request) (interface{}, error) {
	params := ctx.Value(ctxParam).(gin.Params)

	return GetReq{
		ID: uuid.MustParse(params.ByName("id")),
	}, nil
}

func decodeGetAllHandler(ctx context.Context, _ *http.Request) (interface{}, error) {
	query := ctx.Value(ctxQuery).(url.Values)
	page, err := strconv.ParseInt(query.Get("page"), 10, 32)
	if err != nil {
		page = 0
	}
	limit, err := strconv.ParseInt(query.Get("limit"), 10, 32)
	if err != nil {
		limit = 15
	}

	return GetAllReq{
		Page:  int32(page),
		Limit: int32(limit),
	}, nil
}

func decodeStoreHandler(_ context.Context, r *http.Request) (interface{}, error) {
	var req StoreReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(err.Error())
	}

	return req, nil
}

func decodeUpdateHandler(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateReq
	params := ctx.Value(ctxParam).(gin.Params)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(err.Error())
	}

	req.ID = uuid.MustParse(params.ByName("id"))

	return req, nil
}

func decodeDeleteHandler(ctx context.Context, _ *http.Request) (interface{}, error) {
	params := ctx.Value(ctxParam).(gin.Params)
	return DeleteReq{
		ID: uuid.MustParse(params.ByName("id")),
	}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(response.Response)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := err.(response.Response)
	w.WriteHeader(resp.StatusCode())
	_ = json.NewEncoder(w).Encode(resp)
}
