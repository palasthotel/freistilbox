package api

import "strconv"

type Cluster struct {
	Id int
	Name string
	Description string
	Esla bool
}

type ClusterList []Cluster

func (data ClusterList) Len() int {
	return len(data);
}

func (data ClusterList) Less(i, j int) bool {
	a := data[i];
	b := data[j];
	var valA,_ = strconv.Atoi(a.Name[1:]);
	var valB,_ = strconv.Atoi(b.Name[1:]);
	return valA < valB;
}

func (data ClusterList) Swap(i,j int) {
	temp := data[i];
	data[i]=data[j];
	data[j]=temp;
}