package main

func main() {
	CreateConfigDirAndFile()
	config := GetConfig()
	config.AddFavorite("columbus-blue-jackets", nhl)
	config.WriteToFile()
}
