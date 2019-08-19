package controllers

import (
	"bytes"
	"encoding/json"
	"golangapi/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InquiryBalance(c *gin.Context) {
	endpoint := "http://10.243.131.177:3333/gateway/CoreBankBalanceInquiryCASA/1.0/balanceInquiryCASA"
	var obj models.Soa
	var request string = `{
		"balanceInquiryCASARequest": {
		  "soaHeader": {
			"messageVersion": "1.0",
			"messageType": "JSON",
			"messageSubType": "balanceInquiryCASA",
			"messageSender": "Mandol",
			"senderDomain": "Omni",
			"messageTimeStamp": "2018-11-09 01:48:42.578",
			"initiatedTimeStamp": "2018-11-09 01:48:42.578",
			"trackingID": "123456789012"
		  },
		  "messageHeader": {
			"property": [
			  {
				"propertyKey": "tellerId",
				"propertyValue": "9980032"
			  },
			  {
				"propertyKey": "journalSequence",
				"propertyValue": "63522"
			  },
			  {
				"propertyKey": "transactionCode",
				"propertyValue": "2622"
			  }
			]
		  },
		  "payload": {
			"channelId": "32",
			"accountNumber": "1520012684110",
			"accountType": "S"
		  }
		}
	  }`

	var requestBody = []byte(request)

	response, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	response.Header.Set("Authorization", "Basic TWFuZG9sMjpNYW5kb2xAMTIz")
	response.Header.Set("Content-Type", "application/json")
	if err != nil {
		error(c, err)
	}

	client := &http.Client{}
	responseData, err := client.Do(response)
	if err != nil {
		error(c, err)
	}
	if responseData.StatusCode == 401 {
		error(c, "Unauthorized")
	} else {
		body, _ := ioutil.ReadAll(responseData.Body)
		json.Unmarshal(body, &obj)
		data := gin.H{
			"balance":               obj.BalanceInquiryCASAResponse.Payload.BalanceInfo,
			"statementCycleReplica": obj.BalanceInquiryCASAResponse.Payload.StatementCycle,
		}
		success(c, data)
	}
}
