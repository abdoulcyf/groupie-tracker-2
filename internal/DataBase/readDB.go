package DataBase

import "fmt"

//

func ReadDB(id int) {
	fmt.Println("\nID:", id)
	fmt.Println("\nName:", Mydb.TableOfArtists[id].Name)

	fmt.Println("\nMembers:", Mydb.TableOfMembers[id])

	fmt.Println("\nImageUrl:", Mydb.TableOfArtists[id].ImageUrl)
	fmt.Println("\nFirstAlbum:", Mydb.TableOfArtists[id].FirstAlbum)
	fmt.Println("\nCreationDate:", Mydb.TableOfArtists[id].CreationDate)
	fmt.Println()
	//fmt.Println(Mydb.TableOfDates[id])
	fmt.Println("\nlocations:", Mydb.TableOfLocations[id])
	fmt.Println()
	fmt.Println("Relations:")
	fmt.Println()
	for key, value := range Mydb.TableOfRelation[id] {
		fmt.Println(key, "=")
		for _, d := range value {
			fmt.Println(d)
		}
	}
	fmt.Println()
	//fmt.Println(Mydb.Tabler[id])
}
