package api

import (
	"github.com/palasthotel/freistilbox/credentials"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"sort"
	"fmt"
	"bytes"
)

func createClient() http.Client {
	/*proxy,_ := url.Parse("http://localhost:8888");
	transport := &http.Transport {
		Proxy: http.ProxyURL(proxy),
	};*/
	var client=http.Client{}//Transport:transport}
	return client;
}

func Perform(path string) []byte {
	var cred=credentials.Load();
	var req,_=http.NewRequest("GET","https://dashboard.freistilbox.com/api/"+path,nil)
	req.Header.Add("X-User-Email",cred.User);
	req.Header.Add("X-User-Token",cred.Token);
	client := createClient();

	var resp,_=client.Do(req);
	defer resp.Body.Close()
	var data,_=ioutil.ReadAll(resp.Body);
	return data;
}

func Post(path string, payload []byte) []byte {
	var cred=credentials.Load();
	var reader = bytes.NewReader(payload);
	var req,_=http.NewRequest("POST","https://dashboard.freistilbox.com/api/"+path,reader)
	req.Header.Add("X-User-Email",cred.User);
	req.Header.Add("X-User-Token",cred.Token);
	req.Header.Add("Content-Type","application/json");

	client := createClient();
	var resp,_=client.Do(req);
	defer resp.Body.Close()
	var data,_=ioutil.ReadAll(resp.Body);
	return data;
}

func Put(path string, payload []byte) []byte {
	var cred=credentials.Load();
	var reader = bytes.NewReader(payload);
	var req,_=http.NewRequest("PUT","https://dashboard.freistilbox.com/api/"+path,reader);
	req.Header.Add("X-User-Email",cred.User);
	req.Header.Add("X-User-Token",cred.Token);
	req.Header.Add("Content-Type","application/json");

	client := createClient();
	var resp,_=client.Do(req);
	defer resp.Body.Close();
	var data,_=ioutil.ReadAll(resp.Body);
	return data;
}

func FetchClusters() []Cluster {
	var input=Perform("clusters");
	var parsed ClusterCall=ClusterCall{}
	json.Unmarshal(input,&parsed);
	sort.Sort(ClusterList(parsed.Data));
	return parsed.Data;

}

func FetchWebsite(Id string) Website {
	var input=Perform(fmt.Sprintf("websites/%s",Id));
	var parsed WebsiteCall=WebsiteCall{}
	json.Unmarshal(input,&parsed);
	return parsed.Data;
}