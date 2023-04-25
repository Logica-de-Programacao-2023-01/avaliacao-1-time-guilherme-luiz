package q5

import "strings"

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {

	newS := ""
	vogais := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, ranVogais := range vogais {

		s = strings.ReplaceAll(s, ranVogais, "")

	}

	splitS := strings.Split(s, "")

	for i := 0; i < len(s); i++ {

		splitS[i] = "." + splitS[i]

	}

	for _, ranSplitS := range splitS {

		newS += ranSplitS

	}

	newS = strings.ToLower(newS)

	return newS

}
