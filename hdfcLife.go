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


func (t *HDFC) UpdateStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	applicationId := args[0]
	newStatus := args[1]

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error %s", applicationId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}


	//currStatus := row.Columns[1].GetString_()



	//End- Check that the currentStatus to newStatus transition is accurate

	err = stub.DeleteRow(
		"ApplicationTable",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	//applicationId := row.Columns[0].GetString_()
	status := newStatus
	title := row.Columns[2].GetString_()
	firstName := row.Columns[3].GetString_()
	lastName := row.Columns[4].GetString_()
	gender := row.Columns[5].GetString_()
	dob := row.Columns[6].GetString_()
	age := row.Columns[7].GetString_()
	martialStatus := row.Columns[8].GetString_()
	fatherName := row.Columns[9].GetString_()
	motherName := row.Columns[10].GetString_()
	nationality := row.Columns[11].GetString_()
	residentialStatus := row.Columns[12].GetString_()
	placeOfBirth := row.Columns[13].GetString_()
	panNumber := row.Columns[14].GetString_()
	aadharNumber := row.Columns[15].GetString_()
	educationalQualification := row.Columns[16].GetString_()
	politicallyExposed := row.Columns[17].GetString_()
	disablePersonPolicy := row.Columns[18].GetString_()
	anyCriminalProceeding := row.Columns[19].GetString_()
	
	
	_, err = stub.InsertRow(
		"ApplicationTable",
		shim.Row{
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
				&shim.Column{Value: &shim.Column_String_{String_: fatherName}},
				&shim.Column{Value: &shim.Column_String_{String_: motherName}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: residentialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: placeOfBirth}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: educationalQualification}},
				&shim.Column{Value: &shim.Column_String_{String_: politicallyExposed}},
				&shim.Column{Value: &shim.Column_String_{String_: disablePersonPolicy}},
				&shim.Column{Value: &shim.Column_String_{String_: anyCriminalProceeding}}},
		})
	if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

	return nil, nil

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
		fatherName := args[9]
		motherName := args[10]
		nationality := args[11]
		residentialStatus := args[12]
		placeOfBirth := args[13]
		panNumber := args[14]
		aadharNumber := args[15]
		educationalQualification := args[16]
		politicallyExposed := args[17]
		disablePersonPolicy := args[18]
		anyCriminalProceeding := args[19]

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
				&shim.Column{Value: &shim.Column_String_{String_: fatherName}},
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
		t := HDFC{}
		return t.UpdateStatus(stub, args)
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
