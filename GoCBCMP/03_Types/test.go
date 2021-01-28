package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	body := `
	<?xml version="1.0" encoding="utf-8"?>
		<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
		<soap:Body>
			<ZBM0_INT21_GLS_EXT1>
			<IDOC>
				<EDI_DC40>
				<TABNAM>TEST</TABNAM>
				<DIRECT>1</DIRECT>
				</EDI_DC40>
				<E1EDT20>
				<E1EDT22 />
				<E1EDT47>
					<E1EDT56 />
				</E1EDT47>
				<E1ETD01 />
				<E1EDT57>
					<E1EDT58 />
				</E1EDT57>
				</E1EDT20>
			</IDOC>
			</ZBM0_INT21_GLS_EXT1>
		</soap:Body>
		</soap:Envelope>
	`

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://ws.nmbs.deblock.be:7443/NMBS.svc?wsdl", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", "http://sap.com/xi/WebService/soap1.1")
	req.Header.Add("UserAgent", "Alex2Agent")
	req.Header.Add("Accept-Encoding", "gzip,deflate,br")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "keep-alive")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
