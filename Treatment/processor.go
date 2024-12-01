package treatment

import (
	"bytes"
	"fmt"
	"os"
)

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
		VlrInvestimento:                  parseFloat64(record[55]),
		VlrRepasse:                       parseFloat64(record[56]),
		VlrContrapartida:                 parseFloat64(record[57]),
		VlrEmpenhado:                     parseFloat64(record[58]),
		VlrDesembolsado:                  parseFloat64(record[59]),
		VlrDesbloqueado:                  parseFloat64(record[60]),
		VlrPago:                          parseFloat64(record[61]),
		VlrTaxaAdm:                       parseFloat64(record[62]),
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

// Preprocessa o CSV para corrigir problemas comuns
func preprocessCSVFile(inputFilePath string) ([]byte, error) {
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	processedData := bytes.ReplaceAll(data, []byte(`""`), []byte(`"`))
	return processedData, nil
}
