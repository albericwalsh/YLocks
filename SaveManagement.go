package RPG

import (
	"log"
	"os"
	"os/user"

	"encoding/json"
	"io/ioutil"
)

type Save struct {
	Chapter int `json:"Chapter"`
	PlayerX int `json:"PlayerX"`
	PlayerY int `json:"PlayerY"`
	PV      int `json:"PV"`
	PA      int `json:"PA"`
	PD      int `json:"PD"`
}

// return usr home directory
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

var (
	Path string = FixPath(GetHomeDir() + "/AppData/Local/YLock's/savefile.json")
)

func FixPath(Path string) string {
	//convert all \ to / in path
	var newpath = []rune(Path)
	for i, r := range newpath {
		if r == 92 {
			newpath[i] = '/'
		}
	}
	return string(newpath)
}

func CreateSave(save Save) {
	File, err := os.Open(Path)
	if os.IsNotExist(err) {
		os.Mkdir(FixPath(GetHomeDir()+"/AppData/Local/YLock's"), 0777)
		os.Create(Path)
		// MarshalSave(save)
		u, err := json.Marshal(File)
    	if err != nil {
			log.Fatal(err)
    	}
    	ioutil.WriteFile(Path, u, 0777)
	} else {
		log.Fatal(err)
	}
	File.Close()
}

func LoadSave(save Save) {
	File, err := os.Open(Path)
	if os.IsNotExist(err) {
		CreateSave(save)
	} else if err != nil {
		log.Fatal(err)
	} else {
		u := json.NewDecoder(File)
		data := []byte{}
		u.Decode(data)
		err := json.Unmarshal(data, &save)
		if err != nil {
			log.Fatal(err)
		}
	}
	File.Close()
}

func CheckSave() bool {
	File, err := os.Open(Path)
	if os.IsNotExist(err) {
		return false
	} else {
		log.Fatal(err)
	}
	File.Close()
	return true
}
