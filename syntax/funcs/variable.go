package main

func YourName(name string, aliases ...string) {

}
func CallYourName() {
	YourName("cgy")
	YourName("cgy", "Cgy")
	YourName("cgy", "Cgy", "CGy")
	YourName("cgy", "Cgy", "CGy", "CGY")
	aliases := []string{"cgy", "Cgy"}
	YourName("cgy", aliases...)
}
