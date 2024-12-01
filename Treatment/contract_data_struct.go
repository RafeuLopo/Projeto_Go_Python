package treatment

import "time"

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