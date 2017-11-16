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

// creating new part in blockchain
func createPart(stub  shim.ChaincodeStubInterface, args []string) pb.Response {	
	var err error
	fmt.Println("Running createPart")

	if len(args) != 9 {
		fmt.Println("Incorrect number of arguments. Expecting 9 - PartId, Part Code, Manufacture Date, User, Part Type, Part Name, Description, Batch Code, QR Code")
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	fmt.Println("Arguments :"+args[0]+","+args[1]+","+args[2]+","+args[3]+","+args[4]+","+args[5]+","+args[6]);

	var bt Part
	bt.PartId 			= args[0]
	bt.PartCode			= args[1]
	bt.PartType			= args[4]
	bt.PartName			= args[5]
	bt.Description			= args[6]
	bt.BatchCode			= args[7]
	bt.QRCode			= args[8]
	var tx Transaction
	tx.DateOfManufacture		= args[2]
	tx.TType 			= "CREATE"
	tx.User 			= args[3]
	bt.Transactions = append(bt.Transactions, tx)

	//Commit part to ledger
	fmt.Println("createPart Commit Part To Ledger");
	btAsBytes, _ := json.Marshal(bt)
	err = stub.PutState(bt.PartId, btAsBytes)
	if err != nil {		
		return shim.Error(err.Error())
	}

	//Update All Parts Array
	allBAsBytes, err := stub.GetState("allParts")
	if err != nil {
		return shim.Error("Failed to get all Parts")
	}
	var allb AllParts
	err = json.Unmarshal(allBAsBytes, &allb)
	if err != nil {
		return shim.Error("Failed to Unmarshal all Parts")
	}
	allb.Parts = append(allb.Parts,bt.PartId)

	allBuAsBytes, _ := json.Marshal(allb)
	err = stub.PutState("allParts", allBuAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Updating existing part in blockchain
func updatePart(stub  shim.ChaincodeStubInterface, args []string) pb.Response {	
	var err error
	fmt.Println("Running updatePart")

	if len(args) != 9 {
		fmt.Println("Incorrect number of arguments. Expecting 9 - PartId, Vehicle Id, Delivery Date, Installation Date, User, Warranty Start Date, Warranty End Date, Type, vin")
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}
	fmt.Println("Arguments :"+args[0]+","+args[1]+","+args[2]+","+args[3]+","+args[4]+","+args[5]+","+args[6]+","+args[7]);

	//Get and Update Part data
	bAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get Part #" + args[0])
	}
	var bch Part
	err = json.Unmarshal(bAsBytes, &bch)
	if err != nil {
		return shim.Error("Failed to Unmarshal Part #" + args[0])
	}

	var tx Transaction
	tx.TType 	= args[7];

	tx.VehicleId		= args[1]
	tx.DateOfDelivery	= args[2]
	tx.DateOfInstallation	= args[3]
	tx.User  		= args[4]
	tx.WarrantyStartDate	= args[5]
	tx.WarrantyEndDate	= args[6]
	tx.Vin	= args[8]

	bch.Transactions = append(bch.Transactions, tx)

	//Commit updates part to ledger
	fmt.Println("updatePart Commit Updates To Ledger");
	btAsBytes, _ := json.Marshal(bch)
	err = stub.PutState(bch.PartId, btAsBytes)
	if err != nil {		
		fmt.Println("error");
		return shim.Error(err.Error())
	}
	fmt.Println("success");
	return shim.Success(nil)
}

//Create AbattoirInward block
func saveAbattoirReceived(stub  shim.ChaincodeStubInterface, args []string) pb.Response {	
	var err error
	fmt.Println("Running saveAbattoirReceived..")

	if len(args) != 11 {
		fmt.Println("Incorrect number of arguments. Expecting 9 - AbattoirId..")
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	fmt.Println("Arguments :"+args[0]+","+args[1]+","+args[2]+","+args[3]+","+args[4]+","+args[5]+","+args[6]+","+args[7]+","+args[8]+","+args[9]+","+args[10]);

	var bt AbattoirMaterialReceived
	bt.AbattoirId				= args[0]
	bt.PurchaseOrderReferenceNumber			= args[1]
	bt.RawMaterialBatchNumber			= args[2]	
	bt.FarmerId					= args[3]
	bt.GUIDNumber					= args[4]
	bt.MaterialName				= args[5]
	bt.MaterialGrade			= args[6]
	bt.UseByDate				= args[7]
	bt.Quantity					= args[8]
	bt.QuantityUnit					= args[9]	

	var cert FarmersCertificate
	
	if args[10] != "" {
		p := strings.Split(args[10], ",")
		for i := range p {
			c := strings.Split(p[i], "^")
			cert.Id = c[0]
			cert.Name = c[1]
			bt.Certificates = append(bt.Certificates, cert)
		}
	}

	//Commit Inward entry to ledger
	fmt.Println("saveAbattoirReceived - Commit AbattoirInward To Ledger");
	btAsBytes, _ := json.Marshal(bt)
	err = stub.PutState(bt.PurchaseOrderReferenceNumber, btAsBytes)
	if err != nil {		
		return shim.Error(err.Error())
	}

	//Update All Abattoirs Array
	allBAsBytes, err := stub.GetState("allAbattoirReceivedIds")
	if err != nil {
		return shim.Error("Failed to get all Abattoir Inward Ids")
	}
	var allb AllAbattoirReceivedIds
	err = json.Unmarshal(allBAsBytes, &allb)
	if err != nil {
		return shim.Error("Failed to Unmarshal all Received")
	}
	allb.PurchaseOrderReferenceNumbers = append(allb.PurchaseOrderReferenceNumbers, bt.PurchaseOrderReferenceNumber)

	allBuAsBytes, _ := json.Marshal(allb)
	err = stub.PutState("allAbattoirReceivedIds", allBuAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//Create AbattoirDispatch block
func saveAbattoirDispatch(stub  shim.ChaincodeStubInterface, args []string) pb.Response {	
	var err error
	fmt.Println("Running saveAbattoirDispatch..")

	if len(args) != 13 {
		fmt.Println("Incorrect number of arguments. Expecting 12 - AbattoirId..")
		return shim.Error("Incorrect number of arguments. Expecting 12")
	}

	fmt.Println("Arguments :"+args[0]+","+args[1]+","+args[2]+","+args[3]+","+args[4]+","+args[5]+","+args[6]+","+args[7]+","+args[8]+","+args[9]+","+args[10]+","+args[11]+","+args[12]);

	var bt AbattoirDispatch	
	bt.AbattoirId				= args[0]
	bt.ConsignmentNumber		= args[1]
	bt.PurchaseOrderReferenceNumber			= args[2]
	bt.RawMaterialBatchNumber			= args[3]
	bt.GUIDNumber				= args[4]
	bt.MaterialName				= args[5]
	bt.MaterialGrade			= args[6]
	bt.TemperatureStorageMin	= args[7]
	bt.TemperatureStorageMax	= args[8]
	bt.ProductionDate			= args[9]
	bt.UseByDate				= args[10]	
	bt.Quantity					= args[11]
	bt.QuantityUnit				= args[12]
		
	//Commit Inward entry to ledger
	fmt.Println("saveAbattoirDispatch - Commit AbattoirDispatch To Ledger");
	btAsBytes, _ := json.Marshal(bt)
	err = stub.PutState(bt.ConsignmentNumber, btAsBytes)
	if err != nil {		
		return shim.Error(err.Error())
	}

	//Update All AbattoirDispatch Array
	allBAsBytes, err := stub.GetState("allAbattoirDispatchIds")
	if err != nil {
		return shim.Error("Failed to get all Abattoir Dispatch")
	}
	var allb AllAbattoirDispatchIds
	err = json.Unmarshal(allBAsBytes, &allb)
	if err != nil {
		return shim.Error("Failed to Unmarshal all dispatch")
	}
	allb.ConsignmentNumbers = append(allb.ConsignmentNumbers,bt.ConsignmentNumber)

	allBuAsBytes, _ := json.Marshal(allb)
	err = stub.PutState("allAbattoirDispatchIds", allBuAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//Create LogisticTransaction block
func createLogisticTransaction(stub  shim.ChaincodeStubInterface, args []string) pb.Response {	
	var err error
	fmt.Println("Running createLogisticTransaction..")

	if len(args) != 15 {
		fmt.Println("Incorrect number of arguments. Expecting 15")
		return shim.Error("Incorrect number of arguments. Expecting 15")
	}

	fmt.Println("Arguments :"+args[0]+","+args[1]+","+args[2]+","+args[3]+","+args[4]+","+args[5]+","+args[6]+","+args[7]+","+args[8]+","+args[9]+","+args[10]+","+args[11]+","+args[12]+","+args[13]+","+args[14]);

	var bt LogisticTransaction
	bt.LogisticProviderId				= args[0]
	bt.ConsignmentNumber				= args[1]
	bt.RouteId							= args[2]
	bt.AbattoirConsignmentId			= args[3]
	bt.VehicleId						= args[4]
	bt.VehicleType						= args[5]
	bt.PickupDateTime					= args[6]
	bt.ExpectedDeliveryDateTime			= args[7]
	bt.ActualDeliveryDateTime			= args[8]
	bt.TemperatureStorageMin			= args[9]
	bt.TemperatureStorageMax			= args[10]
	bt.Quantity							= args[11]
	bt.HandlingInstruction				= args[12]
	
	
	var st ShipmentStatusTransaction
	st.ShipmentStatus		= args[13]		// Default shipment status should be PickedUp
	st.ShipmentDate 		= args[14]	
	bt.ShipmentStatus = append(bt.ShipmentStatus, st)

	//Commit Inward entry to ledger
	fmt.Println("createLogisticTransaction - Commit LogisticTransaction To Ledger");
	btAsBytes, _ := json.Marshal(bt)
	err = stub.PutState(bt.ConsignmentNumber, btAsBytes)
	if err != nil {		
		return shim.Error(err.Error())
	}

	//Update All AbattoirDispatch Array
	allBAsBytes, err := stub.GetState("allLogisticTransactions")
	if err != nil {
		return shim.Error("Failed to get all Abattoir Dispatch")
	}
	var allb AllLogisticTransactions
	err = json.Unmarshal(allBAsBytes, &allb)
	if err != nil {
		return shim.Error("Failed to Unmarshal all dispatch")
	}
	allb.LogisticTransactionList = append(allb.LogisticTransactionList,bt)

	allBuAsBytes, _ := json.Marshal(allb)
	err = stub.PutState("allLogisticTransactions", allBuAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}



// **********************************************************************
//		Updating Logistics transation status in blockchain
// **********************************************************************
func updateLogisticTransactionStatus(stub  shim.ChaincodeStubInterface, args []string) pb.Response {	
	var err error
	fmt.Println("Running updateLogisticTransactionStatus..")

	if len(args) != 4 {
		fmt.Println("Incorrect number of arguments. Expecting 4 - ConsignmentNumber, LogisticProviderId, ShipmentStatus, ShipmentDate.")
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	fmt.Println("Arguments :"+args[0]+","+args[1]+","+args[2]+","+args[3]);

	//Get and Update LogisticTransaction data
	bAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get LogisticTransaction # " + args[0])
	}
	var bch LogisticTransaction
	err = json.Unmarshal(bAsBytes, &bch)
	if err != nil {
		return shim.Error("Failed to Unmarshal LogisticTransaction # " + args[0])
	}

	var tx ShipmentStatusTransaction
	tx.ShipmentStatus 	= args[2];
	tx.ShipmentDate		= args[3];

	bch.ShipmentStatus = append(bch.ShipmentStatus, tx)

	//Commit updates LogisticTransaction status to ledger
	fmt.Println("updateLogisticTransactionStatus Commit Updates To Ledger");
	btAsBytes, _ := json.Marshal(bch)
	err = stub.PutState(bch.ConsignmentNumber, btAsBytes)
	if err != nil {		
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// **********************************************************************
//		Updating Logistics transation status in blockchain
// **********************************************************************
func pushIotDetailsToLogisticTransaction(stub  shim.ChaincodeStubInterface, args []string) pb.Response {	
	var err error
	fmt.Println("Running pushIotDetailsToLogisticTransaction..")

	if len(args) != 4 {
		fmt.Println("Incorrect number of arguments. Expecting 4 - ConsignmentNumber, LogisticProviderId, ShipmentStatus, ShipmentDate.")
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	fmt.Println("Arguments :"+args[0]+","+args[1]+","+args[2]+","+args[3]);

	//Get and Update LogisticTransaction data
	bAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get LogisticTransaction # " + args[0])
	}
	var bch LogisticTransaction
	err = json.Unmarshal(bAsBytes, &bch)
	if err != nil {
		return shim.Error("Failed to Unmarshal LogisticTransaction # " + args[0])
	}

	var tx IotHistory
	tx.Temperature 	= args[2];
	tx.Location		= args[3];

	bch.IotTemperatureHistory = append(bch.IotTemperatureHistory, tx)

	//Commit updates LogisticTransaction status to ledger
	fmt.Println("pushIotDetailsToLogisticTransaction Commit Updates To Ledger");
	btAsBytes, _ := json.Marshal(bch)
	err = stub.PutState(bch.ConsignmentNumber, btAsBytes)
	if err != nil {		
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

