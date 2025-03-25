package models

type PairCidsRequest struct {
	OVC_START_DATE_TIME string `json:"OVC_START_DATE_TIME"`
	OVC_END_DATE_TIME   string `json:"OVC_END_DATE_TIME"`
	PairType            string `json:"PairType"`
}

type PairCidsResponse struct {
	OVC_CID            string                 `json:"OVC_CID"`
	OVC_DEPT_ID        string                 `json:"OVC_DEPT_ID"`
	OVC_CTRL_UNIT      string                 `json:"OVC_CTRL_UNIT"`
	OVC_WH             string                 `json:"OVC_WH"`
	OVC_POJ_CODE       string                 `json:"OVC_POJ_CODE"`
	OVC_WBS_NO         string                 `json:"OVC_WBS_NO,omitempty"`
	OVC_ANLY_CODE      string                 `json:"OVC_ANLY_CODE,omitempty"`
	OVC_REQ_DATE       string                 `json:"OVC_REQ_DATE,omitempty"`
	OVC_LEDGER_CATE    string                 `json:"OVC_LEDGER_CATE,omitempty"`
	OVC_MEMO           string                 `json:"OVC_MEMO,omitempty"`
	OVC_APPLY_ID       string                 `json:"OVC_APPLY_ID,omitempty"`
	OVC_APPLY_DATE     string                 `json:"OVC_APPLY_DATE,omitempty"`
	OVC_STATUS         string                 `json:"OVC_STATUS,omitempty"`
	OVC_APV_DATE       string                 `json:"OVC_APV_DATE,omitempty"`
	OVC_MINUS          string                 `json:"OVC_MINUS,omitempty"`
	OVC_JOB_CODE       *string                `json:"OVC_JOB_CODE,omitempty"`
	OVC_RM_DEM_NO      *string                `json:"OVC_RM_DEM_NO,omitempty"`
	OVC_PURCH          interface{}            `json:"OVC_PURCH,omitempty"`
	OVC_PURCH_SUB      interface{}            `json:"OVC_PURCH_SUB,omitempty"`
	ONB_SHIP_TIMES     interface{}            `json:"ONB_SHIP_TIMES,omitempty"`
	OVC_RCID           string                 `json:"OVC_RCID"`
	OVC_RUSER_ID       string                 `json:"OVC_RUSER_ID"`
	OVC_RPOJ_CODE      string                 `json:"OVC_RPOJ_CODE"`
	OVC_RDEPT_ID       string                 `json:"OVC_RDEPT_ID"`
	OVC_RANLY_CODE     string                 `json:"OVC_RANLY_CODE,omitempty"`
	OVC_3L_MEMO        interface{}            `json:"OVC_3L_MEMO,omitempty"`
	OVC_APPLY_NAME     string                 `json:"OVC_APPLY_NAME"`
	OVC_RUSER_NAME     string                 `json:"OVC_RUSER_NAME"`
	OVC_CREATE_NAME    string                 `json:"OVC_CREATE_NAME"`
	OVC_DOC_ON         interface{}            `json:"OVC_DOC_ON,omitempty"`
	OVC_LOGISTICS_KIND interface{}            `json:"OVC_LOGISTICS_KIND,omitempty"`
	OVC_LOGISTICS_NAME interface{}            `json:"OVC_LOGISTICS_NAME,omitempty"`
	OVC_LOGISTICS_ID   interface{}            `json:"OVC_LOGISTICS_ID,omitempty"`
	OVC_LOGISTICS_PHON interface{}            `json:"OVC_LOGISTICS_PHON,omitempty"`
	OVC_LOGISTICS_DATE interface{}            `json:"OVC_LOGISTICS_DATE,omitempty"`
	OVC_LOGISTICS_WH   interface{}            `json:"OVC_LOGISTICS_WH,omitempty"`
	OVC_LOGISTICS_MEMO interface{}            `json:"OVC_LOGISTICS_MEMO,omitempty"`
	Items              []PairCidsItemResponse `json:"Items"`
}

type PairCidsItemResponse struct {
	OVC_CID          string      `json:"OVC_CID"`
	ONB_ITEM_ON      int         `json:"ONB_ITEM_ON"`
	ONB_ITEM_ID      int64       `json:"ONB_ITEM_ID"`
	ONB_INV_ID       int64       `json:"ONB_INV_ID"`
	ONB_RITEM_ON     int         `json:"ONB_RITEM_ON"`
	ONB_PURCH_ON     interface{} `json:"ONB_PURCH_ON,omitempty"`
	OVC_LOC          string      `json:"OVC_LOC"`
	OVC_IN_STO_DATE  string      `json:"OVC_IN_STO_DATE"`
	OVC_PROD_YEAR    string      `json:"OVC_PROD_YEAR"`
	OVC_TEMP_PICK_NO string      `json:"OVC_TEMP_PICK_NO"`
	OVC_TMATST       string      `json:"OVC_TMATST"`
	OVC_MRB          string      `json:"OVC_MRB"`
	OVC_BELONG       string      `json:"OVC_BELONG"`
	ONB_APPLY_QTY    int         `json:"ONB_APPLY_QTY"`
	ONB_APPLY_TOT    float64     `json:"ONB_APPLY_TOT"`
	ONB_APV_QTY      int         `json:"ONB_APV_QTY"`
	ONB_APV_TOT      float64     `json:"ONB_APV_TOT"`
	OVC_PART_NO      string      `json:"OVC_PART_NO"`
	OVC_INTG_PIN     string      `json:"OVC_INTG_PIN"`
	OVC_ITEM_CNAME   string      `json:"OVC_ITEM_CNAME"`
	OVC_SPEC         string      `json:"OVC_SPEC"`
	OVC_UN_TYPE      string      `json:"OVC_UN_TYPE"`
}
