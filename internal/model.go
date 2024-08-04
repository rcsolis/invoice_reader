package internal

import (
	"encoding/xml"
)

type InvoiceFile struct {
	XmlName            xml.Name `xml:"Comprobante" json:"-"`
	Xmlns_cfdi         string   `xml:"xmlns cfdi,attr,omitempty"`
	Xmlns_xsi          string   `xml:"xmlns xsi,attr,omitempty"`
	Xsi_schemaLocation string   `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr,omitempty"`
	Version            string   `xml:"Version,attr" json:"version"`
	LocationExpedition string   `xml:"LugarExpedicion,attr" json:"location_expedition"`
	PaymentMethod      string   `xml:"MetodoPago,attr" json:"payment_method"`
	Confirmation       string   `xml:"Confirmacion,attr" json:"confirmation"`
	Currency           string   `xml:"Moneda,attr" json:"currency"`
	Discount           string   `xml:"Descuento,attr" json:"discount"`
	Folio              string   `xml:"Folio,attr" json:"folio"`
	ExchangeRate       string   `xml:"TipoCambio,attr" json:"exchange_rate"`
	Serie              string   `xml:"Serie,attr" json:"serie"`
	Export             string   `xml:"Exportacion,attr" json:"export"`
	InvoiceType        string   `xml:"TipoDeComprobante,attr" json:"invoice_type"`
	PaymentForm        string   `xml:"FormaPago,attr" json:"payment_form"`
	PaymentConditions  string   `xml:"CondicionesDePago,attr" json:"payment_conditions"`
	Date               string   `xml:"Fecha,attr" json:"data_of"`
	SubTotal           string   `xml:"SubTotal,attr" json:"sub_total"`
	Total              string   `xml:"Total,attr" json:"total"`
	CertificateNumber  string   `xml:"NoCertificado,attr" json:"certificate_number"`
	Certificate        string   `xml:"Certificado,attr" json:"certificate"`
	Seal               string   `xml:"Sello,attr" json:"seal"`

	GlobalInformation GlobalInformationType `xml:"InformacionGlobal" json:"global_information"`
	RelatedInvoices   RelatedInvoicesType   `xml:"CfdiRelacionados" json:"related_invoices"`
	Issuer            IssuerType            `xml:"Emisor" json:"issuer"`
	Receiver          ReceiverType          `xml:"Receptor" json:"receiver"`
	Concepts          ConceptsType          `xml:"Conceptos" json:"concepts"`
	Taxes             TotalTaxesType        `xml:"Impuestos" json:"taxes"`
}

type GlobalInformationType struct {
	XMLName     xml.Name `xml:"InformacionGlobal" json:"-"`
	Months      string   `xml:"Meses,attr" json:"months"`
	Year        string   `xml:"Año,attr" json:"year"`
	Periodicity string   `xml:"Periodicidad,attr" json:"periodicity"`
}

type RelatedInvoiceType struct {
	XMLName xml.Name `xml:"CfdiRelacionado" json:"-"`
	UUID    string   `xml:"UUID,attr" json:"uuid"`
}
type RelatedInvoicesType struct {
	XMLName         xml.Name             `xml:"CfdiRelacionados" json:"-"`
	RelationType    string               `xml:"TipoRelacion,attr" json:"type"`
	RelatedInvoices []RelatedInvoiceType `xml:"CfdiRelacionado" json:"invoices"`
}

type IssuerType struct {
	XMLName        xml.Name `xml:"Emisor" json:"-"`
	Regime         string   `xml:"RegimenFiscal,attr" json:"regime"`
	FacAtrAdquired string   `xml:"FacAtrAdquirente,attr" json:"fac_atr_adquired"`
	Name           string   `xml:"Nombre,attr" json:"name"`
	Rfc            string   `xml:"Rfc,attr" json:"rfc"`
}

type ReceiverType struct {
	XMLName         xml.Name `xml:"Receptor" json:"-"`
	Name            string   `xml:"Nombre,attr" json:"name"`
	Rfc             string   `xml:"Rfc,attr" json:"rfc"`
	Use             string   `xml:"UsoCFDI,attr" json:"use"`
	FiscalRegime    string   `xml:"RegimenFiscalReceptor" json:"fiscal_regime"`
	FiscalAddress   string   `xml:"DomicilioFiscalReceptor" json:"fiscal_address"`
	FiscalResidence string   `xml:"ResidenciaFiscal" json:"fiscal_residence"`
	NumRegIdTrib    string   `xml:"NumRegIdTrib" json:"num_reg_id_trib"`
}
type ConceptsType struct {
	XMLName  xml.Name      `xml:"Conceptos" json:"-"`
	Concepts []ConceptType `xml:"Concepto" json:"elements"`
}
type ConceptType struct {
	XMLName           xml.Name               `xml:"Concepto" json:"-"`
	ObjectImp         string                 `xml:"ObjetoImp,attr" json:"object_imp"`
	ProdServiceKey    string                 `xml:"ClaveProdServ,attr" json:"prod_service_key"`
	UnitKey           string                 `xml:"ClaveUnidad,attr" json:"unit_key"`
	IdentificationNum string                 `xml:"NoIdentificacion,attr" json:"identification_num"`
	Quantity          string                 `xml:"Cantidad,attr" json:"quantity"`
	Unit              string                 `xml:"Unidad,attr" json:"unit"`
	Description       string                 `xml:"Descripcion,attr" json:"description"`
	UnitValue         string                 `xml:"ValorUnitario,attr" json:"unit_value"`
	Amount            string                 `xml:"Importe,attr" json:"amount"`
	Taxes             TaxesType              `xml:"Impuestos" json:"taxes"`
	Predial           PredialType            `xml:"CuentaPredial,attr" json:"predial"`
	CustomInformation CustomsInformationType `xml:"InformacionAduanera" json:"custom_information"`
	ThirdParty        ThirdPartyAccountType  `xml:"ACuentaTerceros,attr" json:"third_party"`
	PartType          PartType               `xml:"Parte" json:"part"`
}

type TotalTaxesType struct {
	XMLName        xml.Name       `xml:"Impuestos" json:"-"`
	TotalTransfer  string         `xml:"TotalImpuestosTrasladados,attr" json:"total_transfers"`
	TotalRetention string         `xml:"TotalImpuestosRetenidos,attr" json:"total_retentions"`
	Retentions     RetentionsType `xml:"Retenciones" json:"retentions"`
	Transfers      TransfersType  `xml:"Traslados" json:"transfers"`
}
type TaxesType struct {
	XMLName    xml.Name       `xml:"Impuestos" json:"-"`
	Tranfers   TransfersType  `xml:"Traslados" json:"transfers"`
	Retentions RetentionsType `xml:"Retenciones" json:"retentions"`
}

type TransfersType struct {
	XMLName   xml.Name       `xml:"Traslados" json:"-"`
	Transfers []TransferType `xml:"Traslado" json:"elements"`
}
type TransferType struct {
	XMLName xml.Name `xml:"Traslado" json:"-"`
	Base    string   `xml:"Base,attr" json:"base"`
	Tax     string   `xml:"Impuesto,attr" json:"tax"`
	Type    string   `xml:"TipoFactor,attr" json:"type"`
	Rate    string   `xml:"TasaOCuota,attr" json:"rate"`
	Amount  string   `xml:"Importe,attr" json:"amount"`
}
type RetentionsType struct {
	XMLName    xml.Name        `xml:"Retenciones" json:"-"`
	Retentions []RetentionType `xml:"Retencion" json:"elements"`
}
type RetentionType struct {
	XMLName xml.Name `xml:"Retencion" json:"-"`
	Base    string   `xml:"Base,attr" json:"base"`
	Tax     string   `xml:"Impuesto,attr" json:"tax"`
	Type    string   `xml:"TipoFactor,attr" json:"type"`
	Rate    string   `xml:"TasaOCuota,attr" json:"rate"`
	Amount  string   `xml:"Importe,attr" json:"amount"`
}

type PredialType struct {
	XMLName xml.Name `xml:"CuentaPredial" json:"-"`
	Number  string   `xml:"Numero,attr" json:"number"`
}
type ThirdPartyAccountType struct {
	XMLName xml.Name `xml:"ACuentaTerceros" json:"-"`
	Regime  string   `xml:"RegimenFiscalACuentaTerceros,attr" json:"regime"`
	Name    string   `xml:"NombreACuentaTerceros,attr" json:"name"`
	Rfc     string   `xml:"RfcACuentaTerceros,attr" json:"rfc"`
	Address string   `xml:"DomicilioFiscalACuentaTerceros,attr" json:"address"`
}

type CustomsInformationType struct {
	XMLName xml.Name `xml:"InformacionAduanera" json:"-"`
	Number  string   `xml:"Numero,attr" json:"number"`
}

type PartType struct {
	XMLName            xml.Name                 `xml:"Parte" json:"-"`
	Quantity           string                   `xml:"Cantidad,attr" json:"quantity"`
	Description        string                   `xml:"Descripcion,attr" json:"description"`
	UnitValue          string                   `xml:"ValorUnitario,attr" json:"unit_value"`
	Unit               string                   `xml:"Unidad,attr" json:"unit"`
	Amount             string                   `xml:"Importe,attr" json:"amount"`
	IdentificationNum  string                   `xml:"NoIdentificacion,attr" json:"identification_num"`
	ProdServiceKey     string                   `xml:"ClaveProdServ,attr" json:"prod_service_key"`
	CustomsInformation []CustomsInformationType `xml:"InformacionAduanera" json:"customs_information"`
}
