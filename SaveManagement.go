package RPG

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"encoding/json"
	"io/ioutil"
)

type Save struct {
	CanLoad bool   `json:"CanLoad"`
	Chapter string `json:"Chapter"`
	PlayerX int    `json:"PlayerX"`
	PlayerY int    `json:"PlayerY"`
	PV      int    `json:"PV"`
	PA      int    `json:"PA"`
	PD      int    `json:"PD"`
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

func CreateSave(save *Save) {
	_, err := os.Open(Path)
	if os.IsNotExist(err) {
		os.Mkdir(FixPath(GetHomeDir()+"/AppData/Local/YLock's"), 0777)
		os.Create(Path)
		fmt.Println("Folder created")
	} else {
		savecontent, err := json.Marshal(save)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(Path, savecontent, 0777)
		if err != nil {
			panic(err)
		}
		fmt.Println("Save file created")
	}
}

func LoadSave(save *Save) {
	File, err := os.Open(Path)
	if os.IsNotExist(err) {
		CreateSave(save)
	} else if err != nil {
		log.Fatal(err)
	} else {
		u, _ := ioutil.ReadAll(File)
		json.Unmarshal(u, &save)
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

func CanLoad(save Save) bool {
	LoadSave(&save)
	return save.CanLoad
}

func UpdateSave(save *Save) {
	savecontent, err := json.Marshal(save)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(Path, savecontent, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println("Save file updated")
}
