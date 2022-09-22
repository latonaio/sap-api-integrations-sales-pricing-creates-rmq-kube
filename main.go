package main

import (
	"fmt"
	sap_api_caller "sap-api-integrations-sales-pricing-creates-rmq-kube/SAP_API_Caller"
	"sap-api-integrations-sales-pricing-creates-rmq-kube/SAP_API_Caller/requests"
	sap_api_input_reader "sap-api-integrations-sales-pricing-creates-rmq-kube/SAP_API_Input_Reader"
	"sap-api-integrations-sales-pricing-creates-rmq-kube/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	sap_api_time_value_converter "github.com/latonaio/sap-api-time-value-converter"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	pc := sap_api_request_client_header_setup.NewSAPRequestClientWithOption(conf.SAP)
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), conf.RMQ.QueueTo())
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		"100",
		pc,
		conf.RMQ.QueueTo(),
		rmq,
		l,
	)

	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	for msg := range iter {
		err = callProcess(caller, msg)
		if err != nil {
			msg.Fail()
			l.Error(err)
			continue
		}
		msg.Success()
	}
}

func callProcess(caller *sap_api_caller.SAPAPICaller, msg rabbitmq.RabbitmqMessage) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("error occurred: %w", e)
			return
		}
	}()
	salesPricingConditionValidity, salesPricingConditionRecord := extractData(msg.Data())
	accepter := getAccepter(msg.Data())
	caller.AsyncPostSalesPricing(salesPricingConditionValidity, salesPricingConditionRecord, accepter)
	return nil
}

func extractData(data map[string]interface{}) (
	salesPricingConditionValidity *requests.SalesPricingConditionValidity,
	salesPricingConditionRecord *requests.SalesPricingConditionRecord,
) {

	sdc := sap_api_input_reader.ConvertToSDC(data)
	sap_api_time_value_converter.ChangeTimeFormatToSAPFormatStruct(&sdc)

	salesPricingConditionValidity = sdc.ConvertToSalesPricingConditionValidity()
	salesPricingConditionRecord = sdc.ConvertToSalesPricingConditionRecord()
	return
}

func getAccepter(data map[string]interface{}) []string {
	sdc := sap_api_input_reader.ConvertToSDC(data)
	accepter := sdc.Accepter
	if len(sdc.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"SalesPricingConditionValidity", "SalesPricingConditionRecord",
		}
	}
	return accepter
}