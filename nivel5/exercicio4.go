package main 

import "fmt"

func main() {
	x := struct{
		nome string
		sabor string
		ondetem []string // slice
		vaibemcom map[string]string // lembrar que é junto os types key:value
	}{ // como é anonimo, ele precisa da atricução assim que acaba a declaração dos campos
		nome: "Stroopwafel",
		sabor: "doce",
		ondetem: []string{"na Holanda", "na casa de Tia Beta por que ela é rica"},
		vaibemcom: map[string]string{ // fazer uma declaração de inicialização do mapa novamente
			"no café da manhã" : "café",
			"no almoço": "café",
			"no jantar" : "talvez um chá, mas com café é melhor",
		},
	}

	fmt.Println(x)
}