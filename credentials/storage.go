package credentials

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"path"
)

type Credentials struct {
	User string;
	Token string;
}

func (c Credentials) Save() {
	var user,_=user.Current()
	var encoded,_=json.Marshal(c);
	var filename=path.Join(user.HomeDir,".freistilbox");
	ioutil.WriteFile(filename,encoded,0777)
}

func Load() Credentials {
	var c Credentials;
	var user,_=user.Current()
	var filename=path.Join(user.HomeDir,".freistilbox");
	var data,_=ioutil.ReadFile(filename);

	json.Unmarshal(data,&c);

	return c;
}
