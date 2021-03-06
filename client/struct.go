package client

type ClientConf struct {
	WebAddr    string `json:"web_addr"`
	LatestIP   string `json:"latest_ip"`
	IsIPv6     bool   `json:"is_ipv6"`
	EnableDdns bool   `json:"enable_ddns"`
	DNSPod     bool   `json:"dnspod"`
	Aliyun     bool   `json:"aliyun"`
}

type DNSPodConf struct {
	Id           string `json:"id"`
	Token        string `json:"token"`
	Domain       string `json:"domain"`
	SubDomain    string `json:"sub_domain"`
	RecordId     string `json:"record_id"`
	RecordType   string `json:"record_type"`
	RecordLineId string `json:"record_line_id"`
}

type AliyunConf struct {
	AccessKeyId     string `json:"accesskey_id"`
	AccessKeySecret string `json:"accesskey_secret"`
	Domain          string `json:"domain"`
	SubDomain       string `json:"sub_domain"`
	RecordId        string `json:"record_id"`
	RecordType      string `json:"record_type"`
}
