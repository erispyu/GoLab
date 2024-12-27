package main

import (
	"encoding/json"
	"fmt"
)

type BeneficiaryType string

const (
	BeneficiaryTypeIndividual BeneficiaryType = "individual"
	BeneficiaryTypeCompany    BeneficiaryType = "company"
)

type Config struct {
	BeneficiaryType interface{} `json:"beneficiary_type"`
}

func parse(jsonStr string) {
	config := Config{}
	err := json.Unmarshal([]byte(jsonStr), &config)
	if err != nil {
		fmt.Println(err)
	}
	switch config.BeneficiaryType.(type) {
	case string:
		fmt.Println("config.BeneficiaryType is string")
		if config.BeneficiaryType.(string) == string(BeneficiaryTypeIndividual) {
			fmt.Println("individual match")
		}
	case []interface{}:
		fmt.Println("config.BeneficiaryType is list")
		for _, t := range config.BeneficiaryType.([]interface{}) {
			if t.(string) == string(BeneficiaryTypeIndividual) {
				fmt.Println("individual match")
			}
			fmt.Println("-", t, ", type:", fmt.Sprintf("%T", t))
		}
	}
	fmt.Println("====================================")
}

func main() {
	singleJson := "{\"beneficiary_type\": \"individual\"}"
	listJson := "{\"beneficiary_type\": [\"individual\", \"company\"]}"
	invalidListJson := "{\"beneficiary_type\": 200}"
	parse(singleJson)
	parse(listJson)
	parse(invalidListJson)
}
