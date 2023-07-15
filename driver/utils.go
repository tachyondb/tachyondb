// package driver

// import (
// 	"log"
// 	"os"
// )

// // TODO: update all this
// func CreateBinDirIfNotExist() (error) {
// 	if _, err := os.Stat("bin"); os.IsNotExist(err) {
// 		if err != nil {
// 			return err
// 		}

// 		if err := os.Mkdir("bin", 0755); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func SaveObject(resource string, T proto.Message) (error) {
// 	out, err := proto.Marshal(T)
// 	if err != nil {
// 		return err
// 	}

// 	if err := CreateBinDirIfNotExist(); err != nil {
// 		return err
// 	}

// 	filename := "bin/" + resource + ".bin"
// 	if err := os.WriteFile(filename, out, 0644); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func GetObject(resource string, data proto.Message) error {
// 	filename := "bin/" + resource + ".bin"

// 	in, err := os.ReadFile(filename)
// 	if err != nil {
// 		log.Fatalln("error reading file:", err)
// 		return err
// 	}

// 	if err := proto.Unmarshal(in, data); err != nil {
// 		return err
// 	}

// 	return nil
// }
