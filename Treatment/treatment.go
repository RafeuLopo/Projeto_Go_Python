package treatment

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

func Treatment() {
	processedData, err := preprocessCSVFile("D:/Projeto_Go_Python/Uncompressed_Files/carteira_investimento_mcid.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := csv.NewReader(bytes.NewReader(processedData))
	reader.FieldsPerRecord = -1
	reader.Comma = ';'
	reader.LazyQuotes = true

	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading header:", err)
		return
	}

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	// Criar slice para guardar ContractData
	var contractDataList []ContractData

	for i, record := range records {
		fmt.Printf("Processing row %d\n", i+1)
		data := parseRow(record, i+1)

		// Pular linhas vazias
		if (data == ContractData{}) {
			continue
		}

		// Concatenar dados
		contractDataList = append(contractDataList, data)
	}

	// Salvar arquivo CSV
	err = saveToCSV(contractDataList, "D:/Projeto_Go_Python/Extracted_Parquet/processed_data.csv")
	if err != nil {
		fmt.Println("Error saving to CSV:", err)
	}
}
