package RPG

import (
	"log"
	"os"
	"os/user"
)

// return user home directory
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

var Path string = FixPath(GetHomeDir() + "/AppData/Local/YLock's/savefile.json")

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

//check if save file exists
//if it does, load it
//if it doesn't, call NewGame()

func CheckSave() bool {
	//check if save file exists
	if _, err := os.Stat(Path); err == nil {
		return true
	} else {
		return false
	}
}

func LoadGame() {
	//load the save file
	os.Open(Path)
}

func CreateSaveFile() {
	// Creating an empty file
    // Using Create() function
	os.Mkdir(GetHomeDir() + "/AppData/Local/YLock's", os.ModePerm)
    savefile, err := os.Create(Path)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(savefile)
    savefile.Close()
}
