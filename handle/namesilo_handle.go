// Package tools 工具库  域名 ddns 动态修改
// api :https://www.namesilo.com/api-reference#dns/dns-list-records
package handle

import (
	"encoding/xml"
	"errors"
	"go_ddns_namesilo/config"
	models "go_ddns_namesilo/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type MyIpMOdel struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	Cc      string `json:"cc"`
}

// DDnsByNameSilo ddns 动态绑定 ip至dns
func DDnsByNameSilo() {
	//1、获取 dns 记录id
	records, err := dnsListRecords()
	if err != nil {
		log.Println(err)
		return
	}
	//2、 匹配 需要 动态绑定的host
	rrId, err := matchDomainRecordId(records)
	if err != nil {
		log.Println(err)
		return
	}
	// 3、获取当前地址IP
	ip, err := myIp()
	if err != nil {
		log.Println(err)
		return
	}
	//4、更新dns
	err = updateDnsRecord(rrId, ip.IP)
	if err != nil {
		log.Println(err)
		return
	}
}

// MyIp 我的本地ip
func myIp() (*MyIpMOdel, error) {
	httpUrl := "http://checkip.amazonaws.com/"
	resp, err := http.Get(httpUrl)
	if err != nil {
		return nil, err
	}

	var myIp *MyIpMOdel
	// body 正确响应 json  格式 {"ip":"118.112.111.89","country":"China","cc":"CN"}
	if resp.StatusCode == http.StatusOK {
		ip, _ := ioutil.ReadAll(resp.Body) //把	body 内容读入字符串 s
		if ipStr := string(ip); ipStr != "" {
			myIp = &MyIpMOdel{IP: ipStr}
			return myIp, nil
		}
	}
	return nil, errors.New("当前地址ip查询失败")
}

// DnsListRecords 获取 namesilo列出当前 DNS 记录
func dnsListRecords() (*models.NameSiloRecordModel, error) {

	httpUrl := "https://www.namesilo.com/api/dnsListRecords"

	req, _ := http.NewRequest(http.MethodGet, httpUrl, nil)
	req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 6.0.1; VTR-AL00 Build/V417IR)")

	//设置查询参数
	query := req.URL.Query()
	query.Add("domain", config.MyConfig.Domain)
	query.Add("key", config.MyConfig.ApiKey)
	query.Add("type", "xml")
	query.Add("version", "1")

	req.URL.RawQuery = query.Encode()
	//发起 请求 响应
	resp, _ := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	// body 正确响应 xml 格式
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var record *models.NameSiloRecordModel

	//判断是否转换失败
	if err = xml.Unmarshal(body, &record); err != nil {
		return nil, err
	}
	return record, nil
}

// MatchDomain 匹配 地址  ddns 获取 rrid
func matchDomainRecordId(record *models.NameSiloRecordModel) (string, error) {
	resourceRecord := record.Reply.ResourceRecord

	for _, value := range resourceRecord {
		// 匹配成功 ddns 需要的 域名 返回
		if value.Host.Text == config.MyConfig.DDnsHost+"."+config.MyConfig.Domain {
			return value.RecordID.Text, nil
		}
	}
	return "", errors.New("no match ddns domain")
}

// UpdateDnsRecord Update an existing DNS resource record.
// https://www.namesilo.com/api/dnsUpdateRecord?version=1
// &type=xml
// &key=12345
// &domain=namesilo.com
// &rrid=1a2b3 rrid：资源记录的唯一 ID。您可以使用 dnsListRecords 获取此值。
// &rrhost=test
// &rrvalue=55.55.55.55
// &rrttl=7207
func updateDnsRecord(rrId, updateValue string) error {
	httpUrl := "https://www.namesilo.com/api/dnsUpdateRecord?version=1&type=xml"

	req, _ := http.NewRequest(http.MethodGet, httpUrl, nil)
	req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 6.0.1; VTR-AL00 Build/V417IR)")

	//设置查询参数
	query := req.URL.Query()
	query.Add("version", "1")
	query.Add("type", "xml")
	query.Add("key", config.MyConfig.ApiKey)
	query.Add("domain", config.MyConfig.Domain)
	query.Add("rrid", rrId)
	query.Add("rrhost", config.MyConfig.DDnsHost)
	query.Add("rrvalue", updateValue)
	query.Add("rrttl", "3603")

	req.URL.RawQuery = query.Encode()
	//发起 请求 响应
	resp, _ := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// body 正确响应 xml 格式
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var updateResp *models.UpDateDnsRecordRespModel
	if err := xml.Unmarshal(body, &updateResp); err != nil {
		return err
	}

	if "success" == updateResp.Reply.Detail.Text {
		log.Printf("name-silo ddns 成功 host：%v , IP：%v \n", config.MyConfig.DDnsHost+"."+config.MyConfig.Domain, updateValue)
		return nil
	}

	return errors.New("UpdateDnsRecord error :" + updateResp.Text)
}
