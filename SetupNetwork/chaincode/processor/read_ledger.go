/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ============================================================================================================================
// Get All Abattoir Received
// ============================================================================================================================
func getAllProcessorReceived(stub  shim.ChaincodeStubInterface, option string, value string) pb.Response {
	fmt.Println("getAllProcessorReceived:Looking for All Abattoir Received");

	//get the AllAbattoirReceived index
	allBAsBytes, err := stub.GetState("allProcessorReceivedIds")
	if err != nil {
		return shim.Error("Failed to get all Abattoir Received")
	}

	var res AllProcessorReceivedIds
	err = json.Unmarshal(allBAsBytes, &res)
	//fmt.Println(allBAsBytes);
	if err != nil {
		fmt.Println("Printing Unmarshal error:-");
		fmt.Println(err);
		return shim.Error("Failed to Unmarshal all Processing Company Received Ids")
	}

	var allIds AllProcessorReceivedIds
	var allDetails AllProcessorReceivedDetails
	var sb ProcessorReceived
	if strings.ToLower(option) == "id" && value != "" {
		sbAsBytes, err := stub.GetState(value)
		if err != nil {
			return shim.Error("Failed to get Processing Company Receipt Number ")
		}
		json.Unmarshal(sbAsBytes, &sb)
		if sb.ProcessorReceiptNumber != "" {
			allDetails.ProcessorReceived = append(allDetails.ProcessorReceived,sb);	
		}
		rabAsBytes, _ := json.Marshal(allDetails)
		return shim.Success(rabAsBytes)	
	}

	for i := range res.ProcessorReceiptNumbers{
		sbAsBytes, err := stub.GetState(res.ProcessorReceiptNumbers[i])
		if err != nil {
			return shim.Error("Failed to get Processing Company Receipt Number ")
		}		
		json.Unmarshal(sbAsBytes, &sb)

		if strings.ToLower(option) == "ids" {
			allIds.ProcessorReceiptNumbers = append(allIds.ProcessorReceiptNumbers,sb.ProcessorReceiptNumber);	
		} else if strings.ToLower(option) == "details" {
			allDetails.ProcessorReceived = append(allDetails.ProcessorReceived,sb);	
		}
	}
	
	if strings.ToLower(option) == "ids" {
		rabAsBytes, _ := json.Marshal(allIds)		
		return shim.Success(rabAsBytes)	
	} else if strings.ToLower(option) == "details" {
		rabAsBytes, _ := json.Marshal(allDetails)
		return shim.Success(rabAsBytes)	
	}
	
	return shim.Success(nil)
}

// ============================================================================================================================
// Get All Processing Company Transactions
// ============================================================================================================================
func getAllProcessingTransactions(stub  shim.ChaincodeStubInterface, option string, value string) pb.Response {
	fmt.Println("getAllProcessingTransactions: Looking for All Processing Company Transactions");

	//get the All processing company batch codes index
	allBAsBytes, err := stub.GetState("allProcessingTransactionIds")
	if err != nil {
		return shim.Error("Failed to get all Processing company Batch Codes")
	}

	var res AllProcessingTransactionIds
	err = json.Unmarshal(allBAsBytes, &res)
	//fmt.Println(allBAsBytes);
	if err != nil {
		fmt.Println("Printing Unmarshal error:-");
		fmt.Println(err);
		return shim.Error("Failed to Unmarshal all Processing company Batch Codes records")
	}
	var allIds AllProcessingTransactionIds
	var allDetails AllProcessingTransactionDetails
	var sb ProcessingTransaction
	if strings.ToLower(option) == "id" && value != "" {
		sbAsBytes, err := stub.GetState(value)
		if err != nil {
			return shim.Error("Failed to get processing company batch code record.")
		}
		json.Unmarshal(sbAsBytes, &sb)
		if sb.ProcessorBatchCode != "" {
			allDetails.ProcessingTransaction = append(allDetails.ProcessingTransaction,sb);	
		}
		rabAsBytes, _ := json.Marshal(allDetails)
		return shim.Success(rabAsBytes)	
	}
	
	for i := range res.ProcessorBatchCodes{
		sbAsBytes, err := stub.GetState(res.ProcessorBatchCodes[i])
		if err != nil {
			return shim.Error("Failed to get processing company batch code record.")
		}
		json.Unmarshal(sbAsBytes, &sb)
		if strings.ToLower(option) == "ids" {
			allIds.ProcessorBatchCodes = append(allIds.ProcessorBatchCodes,sb.ProcessorBatchCode);	
		} else if strings.ToLower(option) == "details" {
			allDetails.ProcessingTransaction = append(allDetails.ProcessingTransaction,sb);	
		}
	}

	if strings.ToLower(option) == "ids" {
		rabAsBytes, _ := json.Marshal(allIds)		
		return shim.Success(rabAsBytes)	
	} else if strings.ToLower(option) == "details" {
		rabAsBytes, _ := json.Marshal(allDetails)
		return shim.Success(rabAsBytes)	
	}
	
	return shim.Success(nil)
}


// ============================================================================================================================
// Get All Processing Company Dispatch
// ============================================================================================================================
func getAllProcessorDispatch(stub  shim.ChaincodeStubInterface, option string, value string) pb.Response {
	fmt.Println("getAllProcessorDispatch: Looking for All Processing Company Dispatch records");

	//get the All processing company batch codes index
	allBAsBytes, err := stub.GetState("allProcessorDispatchIds")
	if err != nil {
		return shim.Error("Failed to get all Processing company dispatch consignment numbers")
	}

	var res AllProcessorDispatchIds
	err = json.Unmarshal(allBAsBytes, &res)
	//fmt.Println(allBAsBytes);
	if err != nil {
		fmt.Println("Printing Unmarshal error:-");
		fmt.Println(err);
		return shim.Error("Failed to Unmarshal all Processing company dispatch records")
	}
	var allIds AllProcessorDispatchIds
	var allDetails AllProcessorDispatchDetails
	var sb ProcessorDispatch
	if strings.ToLower(option) == "id" && value != "" {
		sbAsBytes, err := stub.GetState(value)
		if err != nil {
			return shim.Error("Failed to get processing dispatch record.")
		}
		json.Unmarshal(sbAsBytes, &sb)
		if sb.ConsignmentNumber != "" {
			allDetails.ProcessorDispatch = append(allDetails.ProcessorDispatch,sb);	
		}
		rabAsBytes, _ := json.Marshal(allDetails)
		return shim.Success(rabAsBytes)	
	}
	
	for i := range res.ConsignmentNumbers{

		sbAsBytes, err := stub.GetState(res.ConsignmentNumbers[i])
		if err != nil {
			return shim.Error("Failed to get processing company batch code record.")
		}		
		json.Unmarshal(sbAsBytes, &sb)
		if strings.ToLower(option) == "ids" {
			allIds.ConsignmentNumbers = append(allIds.ConsignmentNumbers,sb.ConsignmentNumber);	
		} else if strings.ToLower(option) == "details" {
			allDetails.ProcessorDispatch = append(allDetails.ProcessorDispatch,sb);	
		}
	}

	if strings.ToLower(option) == "ids" {
		rabAsBytes, _ := json.Marshal(allIds)		
		return shim.Success(rabAsBytes)	
	} else if strings.ToLower(option) == "details" {
		rabAsBytes, _ := json.Marshal(allDetails)
		return shim.Success(rabAsBytes)	
	}
	
	return shim.Success(nil)
}

// ============================================================================================================================
// Get All Logistic Transactions
// ============================================================================================================================
func getAllLogisticTransactions(stub  shim.ChaincodeStubInterface, user string) pb.Response {
	fmt.Println("getAllLogisticTransactions: Looking for All Logistic Transactions");

	//get the LogisticTransactions index
	allBAsBytes, err := stub.GetState("allLogisticTransactions")
	if err != nil {
		return shim.Error("Failed to get all Abattoir Received")
	}

	var res AllLogisticTransactions
	err = json.Unmarshal(allBAsBytes, &res)
	//fmt.Println(allBAsBytes);
	if err != nil {
		fmt.Println("Printing Unmarshal error:-");
		fmt.Println(err);
		return shim.Error("Failed to Unmarshal all Logistic Transactions records")
	}

	var rab AllLogisticTransactions

	for i := range res.LogisticTransactionList{

		sbAsBytes, err := stub.GetState(res.LogisticTransactionList[i].ConsignmentNumber)
		if err != nil {
			return shim.Error("Failed to get Logistic Transaction record.")
		}
		var sb LogisticTransaction
		json.Unmarshal(sbAsBytes, &sb)

		// append all transactions to list
		rab.LogisticTransactionList = append(rab.LogisticTransactionList,sb);
	}

	rabAsBytes, _ := json.Marshal(rab)

	return shim.Success(rabAsBytes)
}