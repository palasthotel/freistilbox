package api

import "strconv"

type ClusterAttributes struct {
	Description string
	Name string
	Esla bool
}

type ClusterWebsiteRelation struct {
	Id int
}

type ClusterWebsitesRelationCollection struct {
	Data []ClusterWebsiteRelation
}

type ClusterRelations struct {
	Websites ClusterWebsitesRelationCollection
}

type Cluster struct {
	Id int
	Attributes ClusterAttributes
}

type ClusterList []Cluster

type ClusterCall struct {
	Data ClusterList;
}

func (data ClusterList) Len() int {
	return len(data);
}

func (data ClusterList) Less(i, j int) bool {
	a := data[i];
	b := data[j];
	var valA,_ = strconv.Atoi(a.Attributes.Name[1:]);
	var valB,_ = strconv.Atoi(b.Attributes.Name[1:]);
	return valA < valB;
}

func (data ClusterList) Swap(i,j int) {
	temp := data[i];
	data[i]=data[j];
	data[j]=temp;
}