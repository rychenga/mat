package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mat/logger"     // 匯入 logger
	"mat/middleware" // 匯入 middleware
	"mat/models"     // 匯入資料結構

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"go.uber.org/zap"

	// swagger 套件引入
	_ "mat/docs" // 引入 docs 產生的 swagger 文件

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// PairCidsHandler godoc
// @Summary 憑單明細查詢 API
// @Description 回傳 憏單明細查詢 結果
// @Tags example
// @Accept json
// @Produce json
// @Param request body models.PairCidsRequest true "查詢條件"
// @Success 200 {object} models.PairCidsResponse
// @Router /api/mat/v1/wms/PairCids [post]
func PairCidsHandler(c *gin.Context) {
	log, exists := c.Get("logger")
	if !exists {
		panic("Logger 未設定")
	}
	logger := log.(*zap.Logger)

	var request models.PairCidsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("解析请求体失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	logger.Info("Received request with parameters",
		zap.String("OVC_START_DATE_TIME", request.OVC_START_DATE_TIME),
		zap.String("OVC_END_DATE_TIME", request.OVC_END_DATE_TIME),
		zap.String("PairType", request.PairType))

	response := models.PairCidsResponse{
		OVC_CID:            "3C65001110825000035",
		OVC_DEPT_ID:        "6500",
		OVC_CTRL_UNIT:      "6514",
		OVC_WH:             "電子庫(221B)",
		OVC_POJ_CODE:       "VQCADS",
		OVC_WBS_NO:         "A140000700",
		OVC_ANLY_CODE:      "3110",
		OVC_REQ_DATE:       "2022/08/25",
		OVC_LEDGER_CATE:    "T",
		OVC_MEMO:           "系統自動調撥",
		OVC_APPLY_ID:       "1085399",
		OVC_APPLY_DATE:     "2022/08/25",
		OVC_STATUS:         "C02",
		OVC_APV_DATE:       "2022/08/25 13:58:23",
		OVC_MINUS:          "N",
		OVC_JOB_CODE:       nil,
		OVC_RM_DEM_NO:      nil,
		OVC_PURCH:          nil,
		OVC_PURCH_SUB:      nil,
		ONB_SHIP_TIMES:     nil,
		OVC_RCID:           "2C65001110825000035",
		OVC_RUSER_ID:       "1085399",
		OVC_RPOJ_CODE:      "VQCADS",
		OVC_RDEPT_ID:       "6500",
		OVC_RANLY_CODE:     "3110",
		OVC_3L_MEMO:        nil,
		OVC_APPLY_NAME:     "鍾穎琪",
		OVC_RUSER_NAME:     "鍾穎琪",
		OVC_CREATE_NAME:    "鍾穎琪",
		OVC_DOC_ON:         nil,
		OVC_LOGISTICS_KIND: nil,
		OVC_LOGISTICS_NAME: nil,
		OVC_LOGISTICS_ID:   nil,
		OVC_LOGISTICS_PHON: nil,
		OVC_LOGISTICS_DATE: nil,
		OVC_LOGISTICS_WH:   nil,
		OVC_LOGISTICS_MEMO: nil,
		Items: []models.PairCidsItemResponse{
			models.PairCidsItemResponse{
				OVC_CID:          "3C65001110825000035",
				ONB_ITEM_ON:      1,
				ONB_ITEM_ID:      2108290048531,
				ONB_INV_ID:       2201140114528,
				ONB_RITEM_ON:     1,
				ONB_PURCH_ON:     nil,
				OVC_LOC:          "A",
				OVC_IN_STO_DATE:  "200701080022212",
				OVC_PROD_YEAR:    "2006/05/29",
				OVC_TEMP_PICK_NO: "0",
				OVC_TMATST:       "951206-6",
				OVC_MRB:          "合格",
				OVC_BELONG:       "-",
				ONB_APPLY_QTY:    8,
				ONB_APPLY_TOT:    8,
				ONB_APV_QTY:      8,
				ONB_APV_TOT:      1560,
				OVC_PART_NO:      "5905M55342E06B29D4R",
				OVC_INTG_PIN:     "5905YETT53633",
				OVC_ITEM_CNAME:   "電阻晶片",
				OVC_SPEC:         "029.40 Ω 1% 1/20W，規範說明欄：M55342E06B29D4R",
				OVC_UN_TYPE:      "EA",
			},
		},
	}

	c.JSON(http.StatusOK, response)

}

// @title MAT API
// @version 1.0
// @description 這是 MAT API 的 Swagger API 文件範例。
// @host localhost:8080
// @BasePath /
func main() {
	fmt.Println("Hello, world GO MAT Sevice!")
	// 初始化 zap logger
	log := logger.InitLogger()
	defer log.Sync() // 確保 log buffer 被寫入

	// 使用 dig 進行依賴注入
	container := dig.New()
	if err := container.Provide(logger.InitLogger); err != nil {
		fmt.Println("Failed to provide logger:", err)
		return
	}

	// 設定 Gin 伺服器
	r := gin.New()

	// 使用 gin-zap 記錄 middleware
	r.Use(middleware.GinZapLoggerMiddleware(log))

	// 將 logger 注入到 gin context 中
	r.Use(func(c *gin.Context) {
		c.Set("logger", log)
		c.Next()
	})
	r.POST("/api/mat/v1/wms/PairCids", PairCidsHandler)

	// Swagger API 文件路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 啟動伺服器
	// Graceful_Shutdown 功能
	// 透過 http.Server 啟動 Gin
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	// 啟動 Gin 服務
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server start error", zap.Error(err))
		}
	}()
	fmt.Println("Server is running at http://localhost:8080")

	// 監聽系統訊號
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit // 阻塞，直到收到退出訊號
	log.Info("Shutting down server...")

	// 設定 5 秒的 timeout，讓請求完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	log.Info("Server exiting")

}
