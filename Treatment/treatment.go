package treatment

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type ContractData struct {
	CodTCI                           string
	NumConvenio                      string
	CodProposta                      string
	NumProposta                      string
	CodOperacao                      string
	CodDV                            string
	CodContrato                      string
	CodSaci                          *float64
	CodAgOperador                    string
	CodAgFinanceiro                  string
	CodMdrAntigo                     string
	BlnCargaLegadoTCI                string
	CodCipiProjetoInvest             string
	CodCipiIntervencao               string
	IdGoverna                        string
	NumGenericoContrato              string
	TxtOrigem                        string
	TxtUF                            string
	TxtRegiao                        string
	CodIbge6dig                      *float64
	CodIbge7dig                      *float64
	TxtMunicipio                     string
	TxtTipoInstrumento               string
	DscObjetoInstrumento             string
	DscConcedente                    string
	TxtSiglaSecretaria               string
	TxtNomeSecretaria                string
	BlnPac                           string
	DscFasePac                       string
	BlnEmenda                        string
	NumEmendas                       string
	QtdEmendas                       *float64
	CodTomador                       string
	TxtTomador                       string
	CnpjAgenteFinanceiro             string
	TxtAgenteFinanceiro              string
	NumAnoInstrumento                *float64
	DteAssinaturaContrato            *time.Time
	DteFimContrato                   *time.Time
	DteInicioObra                    *time.Time
	DteFimObra                       *time.Time
	BlnCarteiraMcid                  string
	BlnCarteiraAtivaMcid             string
	BlnCarteiraAndamento             string
	DscSituacaoContratoMcid          string
	DscSituacaoObjetoMcid            string
	TxtMotivoParalisacaoMcid         string
	TxtPrincipalMotivoParalisacao    string
	DscDetalhamentoMotivoParalisacao string
	DscMotivoParalisacao             string
	DteParalisacao                   *time.Time
	DscFonte                         string
	TxtFonte                         string
	DscSubFonte                      string
	CodAcaoUltimoEmpenho             string
	VlrInvestimento                  string
	VlrRepasse                       string
	VlrContrapartida                 string
	VlrEmpenhado                     string
	VlrDesembolsado                  string
	VlrDesbloqueado                  string
	VlrPago                          string
	VlrTaxaAdm                       string
	PrcExecucaoFisica                *float64
	QtdUh                            int
	QtdEntregues                     *float64
	QtdUhDistratadas                 int
	QtdVigentes                      int
	DteControle                      *time.Time
	DteCarga                         *time.Time
	DscSituacaoAtual                 string
	DteAtualizacaoSituacaoAtual      *time.Time
}

// Funções auxiliares para conversões
func parseDate(dateStr string) *time.Time {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" {
		return nil
	}
	t, err := time.Parse("2006/01/02 15:04:05", dateStr)
	if err != nil {
		return nil
	}
	return &t
}

func parseFloat64(val string) *float64 {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil
	}
	val = strings.ReplaceAll(val, ",", ".") // Normaliza números com vírgulas
	parsed, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil
	}
	return &parsed
}

func parseInt(val string) int {
	val = strings.TrimSpace(val)
	parsed, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return parsed
}

// Preprocessa o CSV para corrigir problemas comuns
func preprocessCSVFile(inputFilePath string) ([]byte, error) {
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	processedData := bytes.ReplaceAll(data, []byte(`""`), []byte(`"`))
	return processedData, nil
}

func parseRow(record []string, rowIndex int) ContractData {
	// Preencher valores padrão se a linha não tiver o número esperado de colunas
	if len(record) < 72 {
		fmt.Printf("Warning: Row %d does not have enough columns. Expected 72, got %d\n", rowIndex, len(record))
		
		// Preenche a linha com valores vazios até atingir 72 colunas
		for len(record) < 72 {
			record = append(record, "")
		}
	}

	// Se a linha estiver vazia , retorna vazio
	if isRowEmpty(record) {
		fmt.Printf("Skipping empty row %d\n", rowIndex)
		return ContractData{}
	}

	return ContractData{
		CodTCI:                           record[0],
		NumConvenio:                      record[1],
		CodProposta:                      record[2],
		NumProposta:                      record[3],
		CodOperacao:                      record[4],
		CodDV:                            record[5],
		CodContrato:                      record[6],
		CodSaci:                          parseFloat64(record[7]),
		CodAgOperador:                    record[8],
		CodAgFinanceiro:                  record[9],
		CodMdrAntigo:                     record[10],
		BlnCargaLegadoTCI:                record[11],
		CodCipiProjetoInvest:             record[12],
		CodCipiIntervencao:               record[13],
		IdGoverna:                        record[14],
		NumGenericoContrato:              record[15],
		TxtOrigem:                        record[16],
		TxtUF:                            record[17],
		TxtRegiao:                        record[18],
		CodIbge6dig:                      parseFloat64(record[19]),
		CodIbge7dig:                      parseFloat64(record[20]),
		TxtMunicipio:                     record[21],
		TxtTipoInstrumento:               record[22],
		DscObjetoInstrumento:             record[23],
		DscConcedente:                    record[24],
		TxtSiglaSecretaria:               record[25],
		TxtNomeSecretaria:                record[26],
		BlnPac:                           record[27],
		DscFasePac:                       record[28],
		BlnEmenda:                        record[29],
		NumEmendas:                       record[30],
		QtdEmendas:                       parseFloat64(record[31]),
		CodTomador:                       record[32],
		TxtTomador:                       record[33],
		CnpjAgenteFinanceiro:             record[34],
		TxtAgenteFinanceiro:              record[35],
		NumAnoInstrumento:                parseFloat64(record[36]),
		DteAssinaturaContrato:            parseDate(record[37]),
		DteFimContrato:                   parseDate(record[38]),
		DteInicioObra:                    parseDate(record[39]),
		DteFimObra:                       parseDate(record[40]),
		BlnCarteiraMcid:                  record[41],
		BlnCarteiraAtivaMcid:             record[42],
		BlnCarteiraAndamento:             record[43],
		DscSituacaoContratoMcid:          record[44],
		DscSituacaoObjetoMcid:            record[45],
		TxtMotivoParalisacaoMcid:         record[46],
		TxtPrincipalMotivoParalisacao:    record[47],
		DscDetalhamentoMotivoParalisacao: record[48],
		DscMotivoParalisacao:             record[49],
		DteParalisacao:                   parseDate(record[50]),
		DscFonte:                         record[51],
		TxtFonte:                         record[52],
		DscSubFonte:                      record[53],
		CodAcaoUltimoEmpenho:             record[54],
		VlrInvestimento:                  record[55],
		VlrRepasse:                       record[56],
		VlrContrapartida:                 record[57],
		VlrEmpenhado:                     record[58],
		VlrDesembolsado:                  record[59],
		VlrDesbloqueado:                  record[60],
		VlrPago:                          record[61],
		VlrTaxaAdm:                       record[62],
		PrcExecucaoFisica:                parseFloat64(record[63]),
		QtdUh:                            parseInt(record[64]),
		QtdEntregues:                     parseFloat64(record[65]),
		QtdUhDistratadas:                 parseInt(record[66]),
		QtdVigentes:                      parseInt(record[67]),
		DteControle:                      parseDate(record[68]),
		DteCarga:                         parseDate(record[69]),
		DscSituacaoAtual:                 record[70],
		DteAtualizacaoSituacaoAtual:      parseDate(record[71]),
	}
}

func isRowEmpty(record []string) bool {
	for _, value := range record {
		if value != "" {
			return false
		}
	}
	return true
}

func saveToCSV(data []ContractData, outputFilePath string) error {
	// Saída do arquivo CSV
	file, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Colunas do CSV
	header := []string{
		"CodTCI", "NumConvenio", "CodProposta", "NumProposta", "CodOperacao", "CodDV", "CodContrato", "CodSaci", 
		"CodAgOperador", "CodAgFinanceiro", "CodMdrAntigo", "BlnCargaLegadoTCI", "CodCipiProjetoInvest", 
		"CodCipiIntervencao", "IdGoverna", "NumGenericoContrato", "TxtOrigem", "TxtUF", "TxtRegiao", 
		"CodIbge6dig", "CodIbge7dig", "TxtMunicipio", "TxtTipoInstrumento", "DscObjetoInstrumento", 
		"DscConcedente", "TxtSiglaSecretaria", "TxtNomeSecretaria", "BlnPac", "DscFasePac", "BlnEmenda", 
		"NumEmendas", "QtdEmendas", "CodTomador", "TxtTomador", "CnpjAgenteFinanceiro", "TxtAgenteFinanceiro", 
		"NumAnoInstrumento", "DteAssinaturaContrato", "DteFimContrato", "DteInicioObra", "DteFimObra", 
		"BlnCarteiraMcid", "BlnCarteiraAtivaMcid", "BlnCarteiraAndamento", "DscSituacaoContratoMcid", 
		"DscSituacaoObjetoMcid", "TxtMotivoParalisacaoMcid", "TxtPrincipalMotivoParalisacao", 
		"DscDetalhamentoMotivoParalisacao", "DscMotivoParalisacao", "DteParalisacao", "DscFonte", "TxtFonte", 
		"DscSubFonte", "CodAcaoUltimoEmpenho", "VlrInvestimento", "VlrRepasse", "VlrContrapartida", 
		"VlrEmpenhado", "VlrDesembolsado", "VlrDesbloqueado", "VlrPago", "VlrTaxaAdm", "PrcExecucaoFisica", 
		"QtdUh", "QtdEntregues", "QtdUhDistratadas", "QtdVigentes", "DteControle", "DteCarga", 
		"DscSituacaoAtual", "DteAtualizacaoSituacaoAtual",
	}

	// Escrever colunas no CSV
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing header to CSV: %w", err)
	}

	// Iterar sobre os dados e salvar cada linha no CSV
	for _, record := range data {
		recordRow := []string{
			record.CodTCI,
			record.NumConvenio,
			record.CodProposta,
			record.NumProposta,
			record.CodOperacao,
			record.CodDV,
			record.CodContrato,
			nullableToString(record.CodSaci),
			record.CodAgOperador,
			record.CodAgFinanceiro,
			record.CodMdrAntigo,
			record.BlnCargaLegadoTCI,
			record.CodCipiProjetoInvest,
			record.CodCipiIntervencao,
			record.IdGoverna,
			record.NumGenericoContrato,
			record.TxtOrigem,
			record.TxtUF,
			record.TxtRegiao,
			nullableToString(record.CodIbge6dig),
			nullableToString(record.CodIbge7dig),
			record.TxtMunicipio,
			record.TxtTipoInstrumento,
			record.DscObjetoInstrumento,
			record.DscConcedente,
			record.TxtSiglaSecretaria,
			record.TxtNomeSecretaria,
			record.BlnPac,
			record.DscFasePac,
			record.BlnEmenda,
			record.NumEmendas,
			nullableToString(record.QtdEmendas),
			record.CodTomador,
			record.TxtTomador,
			record.CnpjAgenteFinanceiro,
			record.TxtAgenteFinanceiro,
			nullableToString(record.NumAnoInstrumento),
			nullableDateToString(record.DteAssinaturaContrato),
			nullableDateToString(record.DteFimContrato),
			nullableDateToString(record.DteInicioObra),
			nullableDateToString(record.DteFimObra),
			record.BlnCarteiraMcid,
			record.BlnCarteiraAtivaMcid,
			record.BlnCarteiraAndamento,
			record.DscSituacaoContratoMcid,
			record.DscSituacaoObjetoMcid,
			record.TxtMotivoParalisacaoMcid,
			record.TxtPrincipalMotivoParalisacao,
			record.DscDetalhamentoMotivoParalisacao,
			record.DscMotivoParalisacao,
			nullableDateToString(record.DteParalisacao),
			record.DscFonte,
			record.TxtFonte,
			record.DscSubFonte,
			record.CodAcaoUltimoEmpenho,
			record.VlrInvestimento,
			record.VlrRepasse,
			record.VlrContrapartida,
			record.VlrEmpenhado,
			record.VlrDesembolsado,
			record.VlrDesbloqueado,
			record.VlrPago,
			record.VlrTaxaAdm,
			nullableToString(record.PrcExecucaoFisica),
			strconv.Itoa(record.QtdUh),
			nullableToString(record.QtdEntregues),
			strconv.Itoa(record.QtdUhDistratadas),
			strconv.Itoa(record.QtdVigentes),
			nullableDateToString(record.DteControle),
			nullableDateToString(record.DteCarga),
			record.DscSituacaoAtual,
			nullableDateToString(record.DteAtualizacaoSituacaoAtual),
		}

		if err := writer.Write(recordRow); err != nil {
			return fmt.Errorf("error writing record to CSV: %w", err)
		}
	}

	return nil
}

func nullableToString(val interface{}) string {
	if val == nil {
		return "0"
	}
	switch v := val.(type) {
	case *float64:
		if v == nil {
			return "0"
		}
		return fmt.Sprintf("%f", *v)
	case *time.Time:
		if v == nil {
			return "0"
		}
		return v.Format("2006/01/02 15:04:05")
	}
	return fmt.Sprintf("%v", val)
}

func nullableDateToString(val *time.Time) string {
	if val == nil {
		return "0"
	}
	return val.Format("2006/01/02 15:04:05")
}

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