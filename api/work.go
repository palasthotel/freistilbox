package api

import (
	"github.com/palasthotel/freistilbox/credentials"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sort"
)

func Perform(path string) []byte {
	var cred=credentials.Load();
	var req,_=http.NewRequest("GET","https://dashboard.freistilbox.com/api/"+path,nil)
	req.Header.Add("X-User-Email",cred.User);
	req.Header.Add("X-User-Token",cred.Token);
	var client=http.Client{}

	var resp,_=client.Do(req);
	defer resp.Body.Close()
	var data,_=ioutil.ReadAll(resp.Body);
	return data;
}

func Clusters() []Cluster {
	var input=Perform("clusters");
	var parsed ClusterCall=ClusterCall{}
	json.Unmarshal(input,&parsed);
	sort.Sort(ClusterList(parsed.Data));
	return parsed.Data;

}