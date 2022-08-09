package models

import "encoding/xml"

// NameSiloRecordModel namesilo 接口返回xml
type NameSiloRecordModel struct {
	XMLName xml.Name `xml:"namesilo"`
	Text    string   `xml:",chardata"`
	Request struct {
		Text      string `xml:",chardata"`
		Operation struct {
			Text string `xml:",chardata"`
		} `xml:"operation"`
		Ip struct {
			Text string `xml:",chardata"`
		} `xml:"ip"`
	} `xml:"request"`
	Reply struct {
		Text string `xml:",chardata"`
		Code struct {
			Text string `xml:",chardata"`
		} `xml:"code"`
		Detail struct {
			Text string `xml:",chardata"`
		} `xml:"detail"`
		ResourceRecord []struct {
			Text     string `xml:",chardata"`
			RecordID struct {
				Text string `xml:",chardata"`
			} `xml:"record_id"`
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"type"`
			Host struct {
				Text string `xml:",chardata"`
			} `xml:"host"`
			Value struct {
				Text string `xml:",chardata"`
			} `xml:"value"`
			Ttl struct {
				Text string `xml:",chardata"`
			} `xml:"ttl"`
			Distance struct {
				Text string `xml:",chardata"`
			} `xml:"distance"`
		} `xml:"resource_record"`
	} `xml:"reply"`
}

// UpDateDnsRecordRespModel Update an existing DNS resource record.
type UpDateDnsRecordRespModel struct {
	Namesilo xml.Name `xml:"namesilo"`
	Text     string   `xml:",chardata"`
	Request  struct {
		Text      string `xml:",chardata"`
		Operation struct {
			Text string `xml:",chardata"`
		} `xml:"operation"`
		Ip struct {
			Text string `xml:",chardata"`
		} `xml:"ip"`
	} `xml:"request"`
	Reply struct {
		Text string `xml:",chardata"`
		Code struct {
			Text string `xml:",chardata"`
		} `xml:"code"`
		Detail struct {
			Text string `xml:",chardata"`
		} `xml:"detail"`
		RecordID struct {
			Text string `xml:",chardata"`
		} `xml:"record_id"`
	} `xml:"reply"`
}
