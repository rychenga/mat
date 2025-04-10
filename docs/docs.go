// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/mat/v1/wms/PairCids": {
            "post": {
                "description": "回傳 憏單明細查詢 結果",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "憑單明細查詢 API",
                "parameters": [
                    {
                        "description": "查詢條件",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PairCidsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PairCidsResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.PairCidsItemResponse": {
            "type": "object",
            "properties": {
                "ONB_APPLY_QTY": {
                    "type": "integer"
                },
                "ONB_APPLY_TOT": {
                    "type": "number"
                },
                "ONB_APV_QTY": {
                    "type": "integer"
                },
                "ONB_APV_TOT": {
                    "type": "number"
                },
                "ONB_INV_ID": {
                    "type": "integer"
                },
                "ONB_ITEM_ID": {
                    "type": "integer"
                },
                "ONB_ITEM_ON": {
                    "type": "integer"
                },
                "ONB_PURCH_ON": {},
                "ONB_RITEM_ON": {
                    "type": "integer"
                },
                "OVC_BELONG": {
                    "type": "string"
                },
                "OVC_CID": {
                    "type": "string"
                },
                "OVC_INTG_PIN": {
                    "type": "string"
                },
                "OVC_IN_STO_DATE": {
                    "type": "string"
                },
                "OVC_ITEM_CNAME": {
                    "type": "string"
                },
                "OVC_LOC": {
                    "type": "string"
                },
                "OVC_MRB": {
                    "type": "string"
                },
                "OVC_PART_NO": {
                    "type": "string"
                },
                "OVC_PROD_YEAR": {
                    "type": "string"
                },
                "OVC_SPEC": {
                    "type": "string"
                },
                "OVC_TEMP_PICK_NO": {
                    "type": "string"
                },
                "OVC_TMATST": {
                    "type": "string"
                },
                "OVC_UN_TYPE": {
                    "type": "string"
                }
            }
        },
        "models.PairCidsRequest": {
            "type": "object",
            "properties": {
                "OVC_END_DATE_TIME": {
                    "type": "string"
                },
                "OVC_START_DATE_TIME": {
                    "type": "string"
                },
                "PairType": {
                    "type": "string"
                }
            }
        },
        "models.PairCidsResponse": {
            "type": "object",
            "properties": {
                "Items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PairCidsItemResponse"
                    }
                },
                "ONB_SHIP_TIMES": {},
                "OVC_3L_MEMO": {},
                "OVC_ANLY_CODE": {
                    "type": "string"
                },
                "OVC_APPLY_DATE": {
                    "type": "string"
                },
                "OVC_APPLY_ID": {
                    "type": "string"
                },
                "OVC_APPLY_NAME": {
                    "type": "string"
                },
                "OVC_APV_DATE": {
                    "type": "string"
                },
                "OVC_CID": {
                    "type": "string"
                },
                "OVC_CREATE_NAME": {
                    "type": "string"
                },
                "OVC_CTRL_UNIT": {
                    "type": "string"
                },
                "OVC_DEPT_ID": {
                    "type": "string"
                },
                "OVC_DOC_ON": {},
                "OVC_JOB_CODE": {
                    "type": "string"
                },
                "OVC_LEDGER_CATE": {
                    "type": "string"
                },
                "OVC_LOGISTICS_DATE": {},
                "OVC_LOGISTICS_ID": {},
                "OVC_LOGISTICS_KIND": {},
                "OVC_LOGISTICS_MEMO": {},
                "OVC_LOGISTICS_NAME": {},
                "OVC_LOGISTICS_PHON": {},
                "OVC_LOGISTICS_WH": {},
                "OVC_MEMO": {
                    "type": "string"
                },
                "OVC_MINUS": {
                    "type": "string"
                },
                "OVC_POJ_CODE": {
                    "type": "string"
                },
                "OVC_PURCH": {},
                "OVC_PURCH_SUB": {},
                "OVC_RANLY_CODE": {
                    "type": "string"
                },
                "OVC_RCID": {
                    "type": "string"
                },
                "OVC_RDEPT_ID": {
                    "type": "string"
                },
                "OVC_REQ_DATE": {
                    "type": "string"
                },
                "OVC_RM_DEM_NO": {
                    "type": "string"
                },
                "OVC_RPOJ_CODE": {
                    "type": "string"
                },
                "OVC_RUSER_ID": {
                    "type": "string"
                },
                "OVC_RUSER_NAME": {
                    "type": "string"
                },
                "OVC_STATUS": {
                    "type": "string"
                },
                "OVC_WBS_NO": {
                    "type": "string"
                },
                "OVC_WH": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MAT API",
	Description:      "這是 MAT API 的 Swagger API 文件範例。",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
