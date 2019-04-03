package structs

import (
	"encoding/json"
)

type Project struct {
	Id                int    `json:"id"`                //key
	Name              string `json:"name"`              //專案名稱
	Ip                string `json:"ip"`                //連線 ip 位置
	PushAddress       string `json:"pushAddress"`       //adb push 位置
	DebugApkPath      string `json:"debugApkPath"`      //debug apk 位置
	DebugSignedName   string `json:"debugSignedName"`   //debug 加簽後 apk 名稱
	ReleaseApkPath    string `json:"releaseApkPath"`    //release apk 位置
	ReleaseSignedName string `json:"releaseSignedName"` //release 加簽後 apk 名稱
	SignedId          int    `json:"signedId"`          //加簽檔案 id
}

func toStruct(jsonStr string) (*Project, error) {
	project := Project{}
	err := json.Unmarshal([]byte(jsonStr), project)
	if err != nil {
		return &project, err
	}
	return nil, err
}

func toString(project *Project) (jsonStr string, result error) {
	data, result := json.Marshal(project)
	if result != nil {
		return "", result
	}
	jsonStr = string(data)
	return jsonStr, result
}
