package main 

import "fmt"


type veiculo struct{
	portas int
	cor string
}

type caminhonete struct{
	veic veiculo
	tracao bool
}

type sedan struct{
	veic veiculo
	modeloLuxo bool
}



func main() {


	caminhonete1 := caminhonete{veiculo{4, "cinza"}, true}
	sedan1 := sedan{veiculo{4, "azul marinho"}, true}


	fmt.Println(caminhonete1.veic)
	fmt.Println(sedan1.veic.cor)
}