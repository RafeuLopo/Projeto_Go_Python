package treatment

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

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
