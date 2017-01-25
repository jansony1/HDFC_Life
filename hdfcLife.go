package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)



 
// HDFC is a high level smart contract that HDFCs together business artifact based smart contracts
type HDFC struct {

}

// Init initializes the smart contracts
func (t *HDFC) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("ApplicationTable")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("ApplicationTable", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "applicationId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "age", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "martialStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "fatherName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "motherName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nationality", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "residentialStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "placeOfBirth", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "panNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "aadharNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "educationalQualification", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "politicallyExposed", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "disablePersonPolicy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "anyCriminalProceeding", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}

	return nil, nil
}


func (t *HDFC) getNumApplications(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0.")
	}

	var columns []shim.Column

	contractCounter := 0

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	for row := range rows {
		if len(row.Columns) != 0 {
			contractCounter++
		}
	}

	type count struct {
		NumContracts int
	}

	var c count
	c.NumContracts = contractCounter

	return json.Marshal(c)
}




// Invoke invokes the chaincode
func (t *HDFC) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "submitApplication" {
		if len(args) != 19 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 19. Got: %d.", len(args))
		}

		applicationId := args[0]
		status := args[1]
		title := args[2]
		firstName := args[3]
		lastName := args[4]
		gender := args[5]
		dob := args[6]
		age := args[7]
		martialStatus := args[8]
		motherName := args[9]
		nationality := args[10]
		residentialStatus := args[11]
		placeOfBirth := args[12]
		panNumber := args[13]
		aadharNumber := args[14]
		educationalQualification := args[15]
		politicallyExposed := args[16]
		disablePersonPolicy := args[17]
		anyCriminalProceeding := args[18]

		// Insert a row
		ok, err := stub.InsertRow("ApplicationTable", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: motherName}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: residentialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: placeOfBirth}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: educationalQualification}},
				&shim.Column{Value: &shim.Column_String_{String_: politicallyExposed}},
				&shim.Column{Value: &shim.Column_String_{String_: disablePersonPolicy}},
				&shim.Column{Value: &shim.Column_String_{String_: anyCriminalProceeding}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}

		return nil, err
	} else if function == "updateApplicationStatus" { 
		//return UpdateStatus(stub, args)
	} 

	return nil, errors.New("Invalid invoke function name.")

}

func (t *HDFC) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	type Status struct {
		Status string
	}
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(HDFC))
	if err != nil {
		fmt.Printf("Error starting HDFC: %s", err)
	}
} 
