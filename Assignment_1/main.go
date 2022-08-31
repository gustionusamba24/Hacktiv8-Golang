package main

import "fmt"

type students struct {
	id int
	name string
	address string
	profession string
	reason string
}

func findStudent(students []students, id int) {
	for i, student := range students {
		if student.id == id {
			fmt.Println("******************** Students Data ********************")
			fmt.Println("Nama \t\t:", students[i].name)
			fmt.Println("Address \t:", students[i].address)
			fmt.Println("Profession \t:", students[i].profession)
			fmt.Println("Reason \t\t:", students[i].reason)
			fmt.Println("********************* End of Data *********************")
			return
		} else if i == len(students) - 1 {
			fmt.Println("id :", id, "Not Found")
		}
	}	
}

func main(){
	var input int
	fmt.Print("Please input ID student :", input)
	fmt.Scan(&input)

	var studentsData = []students {
		{
			id: 1,
			name: "Gustio Nusamba",
			address: "Kalasan, Yogyakarta",
			profession: "Mobile Developer",
			reason: "want to explore about back end developer",
		},
		{
			id: 2,
			name: "Hatmaja Narotama",
			address: "Depok, Yogyakarta",
			profession: "Front End Developer",
			reason: "want to explore about Go languange",
		},
		{
			id: 3,
			name: "Rizki Tri Wahyudi",
			address: "Prambanan, Yogyakarta",
			profession: "Back End Developer",
			reason: "want to explore about Go programming",
		},
		{
			id: 4,
			name: "Arsya Diva Maharani",
			address: "Gamping, Yogyakarta",
			profession: "IT Support",
			reason: "want to learn back end",
		},
		{
			id: 5,
			name: "Dinda Kusuma Dewi",
			address: "Pakem, Yogyakarta",
			profession: "Software Engineer",
			reason: "want to focus on a career in back end engineer",
		},
	}

	findStudent(studentsData, input)
}