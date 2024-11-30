package main

import (
	"fmt"
	"os"
	"projeto-go-python/Extraction"
	"projeto-go-python/Treatment"
)

func main() {
	url := "https://dadosabertos.cidades.gov.br/dataset/ec20859d-48a1-403f-893a-0de6f41ad99c/resource/6e634ccb-62b2-466a-9894-74c92168490c/download/carteira_investimento_mcid.rar"
	destFolder := "RAR_Files"
	destFileName := "carteira_investimento.rar"

	err := extraction.DownloadRARFile(url, destFolder, destFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Resource downloaded successfully!")
	rarFile := destFolder + "/" + destFileName
	outputDir := "Uncompressed_Files"
	err = extraction.ExtractRAR(rarFile, outputDir)
	if err != nil {
		fmt.Println("Error extracting RAR file:", err)
		return
	}

	fmt.Println("RAR file extracted successfully!")
	dir := "Extracted_Parquet"

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	treatment.Treatment()
}
