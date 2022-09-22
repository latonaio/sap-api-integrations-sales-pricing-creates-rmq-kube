package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-pricing-creates-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSalesPricingConditionValidity(raw []byte, l *logger.Logger) (*SalesPricingConditionValidity, error) {
	pm := &responses.SalesPricingConditionValidity{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesPricingConditionValidity. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D

	salesPricingConditionValidity := &SalesPricingConditionValidity{
		ConditionRecord:            data.ConditionRecord,
		ConditionValidityEndDate:   data.ConditionValidityEndDate,
		ConditionValidityStartDate: data.ConditionValidityStartDate,
		ConditionApplication:       data.ConditionApplication,
		ConditionType:              data.ConditionType,
		ConditionReleaseStatus:     data.ConditionReleaseStatus,
		SalesDocument:              data.SalesDocument,
		SalesDocumentItem:          data.SalesDocumentItem,
		ConditionContract:          data.ConditionContract,
		CustomerGroup:              data.CustomerGroup,
		CustomerPriceGroup:         data.CustomerPriceGroup,
		MaterialPricingGroup:       data.MaterialPricingGroup,
		SoldToParty:                data.SoldToParty,
		BPForSoldToParty:           data.BPForSoldToParty,
		Customer:                   data.Customer,
		BPForCustomer:              data.BPForCustomer,
		PayerParty:                 data.PayerParty,
		BPForPayerParty:            data.BPForPayerParty,
		ShipToParty:                data.ShipToParty,
		BPForShipToParty:           data.BPForShipToParty,
		Supplier:                   data.Supplier,
		BPForSupplier:              data.BPForSupplier,
		MaterialGroup:              data.MaterialGroup,
		Material:                   data.Material,
		PriceListType:              data.PriceListType,
		CustomerTaxClassification1: data.CustomerTaxClassification1,
		ProductTaxClassification1:  data.ProductTaxClassification1,
		SDDocument:                 data.SDDocument,
		ReferenceSDDocument:        data.ReferenceSDDocument,
		ReferenceSDDocumentItem:    data.ReferenceSDDocumentItem,
		SalesOffice:                data.SalesOffice,
		SalesGroup:                 data.SalesGroup,
		SalesOrganization:          data.SalesOrganization,
		DistributionChannel:        data.DistributionChannel,
		TransactionCurrency:        data.TransactionCurrency,
		ConditionProcessingStatus:  data.ConditionProcessingStatus,
		PricingDate:                data.PricingDate,
		ConditionScaleBasisValue:   data.ConditionScaleBasisValue,
		TaxCode:                    data.TaxCode,
		ServiceDocument:            data.ServiceDocument,
		ServiceDocumentItem:        data.ServiceDocumentItem,
		CustomerConditionGroup:     data.CustomerConditionGroup,
	}

	return salesPricingConditionValidity, nil
}

func ConvertToSalesPricingConditionRecord(raw []byte, l *logger.Logger) (*SalesPricingConditionRecord, error) {
	p := &responses.SalesPricingConditionRecord{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesPricingConditionRecord. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	salesPricingConditionRecord := &SalesPricingConditionRecord{
		ConditionRecord:              data.ConditionRecord,
		ConditionSequentialNumber:    data.ConditionSequentialNumber,
		ConditionTable:               data.ConditionTable,
		ConditionApplication:         data.ConditionApplication,
		ConditionType:                data.ConditionType,
		ConditionValidityEndDate:     data.ConditionValidityEndDate,
		ConditionValidityStartDate:   data.ConditionValidityStartDate,
		CreationDate:                 data.CreationDate,
		PricingScaleType:             data.PricingScaleType,
		PricingScaleBasis:            data.PricingScaleBasis,
		ConditionScaleQuantity:       data.ConditionScaleQuantity,
		ConditionScaleQuantityUnit:   data.ConditionScaleQuantityUnit,
		ConditionScaleAmount:         data.ConditionScaleAmount,
		ConditionScaleAmountCurrency: data.ConditionScaleAmountCurrency,
		ConditionCalculationType:     data.ConditionCalculationType,
		ConditionRateValue:           data.ConditionRateValue,
		ConditionRateValueUnit:       data.ConditionRateValueUnit,
		ConditionQuantity:            data.ConditionQuantity,
		ConditionQuantityUnit:        data.ConditionQuantityUnit,
		BaseUnit:                     data.BaseUnit,
		ConditionIsDeleted:           data.ConditionIsDeleted,
		PaymentTerms:                 data.PaymentTerms,
		IncrementalScale:             data.IncrementalScale,
		PricingScaleLine:             data.PricingScaleLine,
		ConditionReleaseStatus:       data.ConditionReleaseStatus,
	}

	return salesPricingConditionRecord, nil
}
