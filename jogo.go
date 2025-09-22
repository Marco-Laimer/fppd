// // jogo.go - Funções para manipular os elementos do jogo, como carregar o mapa e mover o personagem
// package main

// import (
// 	"bufio"
// 	"os"
// )

// // Elemento representa qualquer objeto do mapa (parede, personagem, vegetação, etc)
// type Elemento struct {
// 	simbolo  rune
// 	cor      Cor
// 	corFundo Cor
// 	tangivel bool // Indica se o elemento bloqueia passagem
// }

// // Jogo contém o estado atual do jogo
// type Jogo struct {
// 	Mapa                   [][]Elemento // grade 2D representando o mapa
// 	PosX, PosY             int          // posição atual do personagem
// 	UltimoVisitado         Elemento     // elemento que estava na posição do personagem antes de mover
// 	StatusMsg              string       // mensagem para a barra de status
// 	UltimaDirX, UltimaDirY int          // A última direção de interação selecionada pelas setas
// 	trapX                  int          // posição X da armadilha
// 	trapY                  int          // posição Y da armadilha
// 	sentido                int          //Sentido de movimento da armadilha (1 = descendo, -1 = subindo)
// 	limiteSuperior         int          //Até onde a trap pode subir
// 	limiteInferior         int          //Até onde a trap pode descer
// }

// // Elementos visuais do jogo
// var (
// 	Personagem = Elemento{'☺', CorCinzaEscuro, CorPadrao, true}
// 	Inimigo    = Elemento{'☠', CorVermelho, CorPadrao, true}
// 	Parede     = Elemento{'▓', CorParede, CorFundoParede, true}
// 	Vegetacao  = Elemento{'♣', CorVerde, CorPadrao, false}
// 	Vazio      = Elemento{' ', CorPadrao, CorPadrao, false}
// 	Trap       = Elemento{'▲', CorVermelho, CorPadrao, true}
// )

// // Cria e retorna uma nova instância do jogo
// func jogoNovo() Jogo {
// 	// O ultimo elemento visitado é inicializado como vazio
// 	// pois o jogo começa com o personagem em uma posição vazia
// 	return Jogo{
// 		UltimoVisitado: Vazio,
// 		UltimaDirX:     1, // Define a direção inicial para a direita
// 		UltimaDirY:     0,
// 		StatusMsg:      "Direção de interação: Direita", // Define a mensagem inicial
// 		trapX:          15,                              // coluna
// 		trapY:          3,                               // linha
// 		sentido:        1,                               // Sentido inicial (1 = descendo, -1 = subindo)
// 		limiteSuperior: 3,
// 		limiteInferior: 9,
// 	}
// }

// // Lê um arquivo texto linha por linha e constrói o mapa do jogo
// func jogoCarregarMapa(nome string, jogo *Jogo) error {
// 	arq, err := os.Open(nome)
// 	if err != nil {
// 		return err
// 	}
// 	defer arq.Close()

// 	scanner := bufio.NewScanner(arq)
// 	y := 0
// 	for scanner.Scan() {
// 		linha := scanner.Text()
// 		var linhaElems []Elemento
// 		for x, ch := range linha {
// 			e := Vazio
// 			switch ch {
// 			case Parede.simbolo:
// 				e = Parede
// 			case Inimigo.simbolo:
// 				e = Inimigo
// 			case Vegetacao.simbolo:
// 				e = Vegetacao
// 			case Personagem.simbolo:
// 				jogo.PosX, jogo.PosY = x, y // registra a posição inicial do personagem
// 			case Trap.simbolo:
// 				jogo.trapX, jogo.trapY = x, y //usa o valor do mapa.txt, dai nao precisa inicializar valor de posição no código(não funcionou ):)
// 				//e = Trap //apagar depois de testar
// 			}
// 			linhaElems = append(linhaElems, e)
// 		}
// 		jogo.Mapa = append(jogo.Mapa, linhaElems)
// 		y++
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Verifica se o personagem pode se mover para a posição (x, y)
// func jogoPodeMoverPara(jogo *Jogo, x, y int) bool {
// 	// Verifica se a coordenada Y está dentro dos limites verticais do mapa
// 	if y < 0 || y >= len(jogo.Mapa) {
// 		return false
// 	}

// 	// Verifica se a coordenada X está dentro dos limites horizontais do mapa
// 	if x < 0 || x >= len(jogo.Mapa[y]) {
// 		return false
// 	}

// 	// Verifica se o elemento de destino é tangível (bloqueia passagem)
// 	if jogo.Mapa[y][x].tangivel {
// 		return false
// 	}

// 	// Pode mover para a posição
// 	return true
// }

// // Move um elemento para a nova posição
// func jogoMoverElemento(jogo *Jogo, x, y, dx, dy int) {
// 	nx, ny := x+dx, y+dy

// 	// Obtem elemento atual na posição
// 	elemento := jogo.Mapa[y][x] // guarda o conteúdo atual da posição

// 	jogo.Mapa[y][x] = jogo.UltimoVisitado   // restaura o conteúdo anterior
// 	jogo.UltimoVisitado = jogo.Mapa[ny][nx] // guarda o conteúdo atual da nova posição
// 	jogo.Mapa[ny][nx] = elemento            // move o elemento
// }

// func jogoAtualizarTrap(jogo *Jogo) {

// 	// Limpa a posição atual
// 	jogo.Mapa[jogo.trapY][jogo.trapX] = Vazio

// 	// Próxima posição
// 	novoY := jogo.trapY + jogo.sentido

// 	// Atualiza posição
// 	jogo.trapY = novoY
// 	jogo.Mapa[jogo.trapY][jogo.trapX] = Trap

//		// Se bater em parede ou limite, inverte o sentido (valores foram ajustados para o mapa atual, algoritmo não vai funcionar para outros mapas com valores X e Y diferentes)
//		if novoY <= jogo.limiteSuperior || novoY >= jogo.limiteInferior {
//			jogo.sentido *= -1
//			novoY = jogo.trapY + jogo.sentido
//		}
//	}
//
// =======
// jogo.go - Funções para manipular os elementos do jogo, como carregar o mapa e mover o personagem
package main

import (
	"bufio"
	"os"
	"time"
)

// Elemento representa qualquer objeto do mapa (parede, personagem, vegetação, etc)
type Elemento struct {
	simbolo  rune
	cor      Cor
	corFundo Cor
	tangivel bool // Indica se o elemento bloqueia passagem
}

// Jogo contém o estado atual do jogo
type Jogo struct {
	Mapa                   [][]Elemento // grade 2D representando o mapa
	PosX, PosY             int          // posição atual do personagem
	UltimoVisitado         Elemento     // elemento que estava na posição do personagem antes de mover
	StatusMsg              string       // mensagem para a barra de status
	UltimaDirX, UltimaDirY int          // A última direção de interação selecionada pelas setas
	trapX                  int          // posição X da armadilha
	trapY                  int          // posição Y da armadilha
	sentido                int          //Sentido de movimento da armadilha (1 = descendo, -1 = subindo)
	limiteSuperior         int          //Até onde a trap pode subir
	limiteInferior         int          //Até onde a trap pode descer
}

// Elementos visuais do jogo
var (
	Personagem    = Elemento{'☺', CorCinzaEscuro, CorPadrao, true}
	Inimigo       = Elemento{'☠', CorVermelho, CorPadrao, true}
	Parede        = Elemento{'▓', CorParede, CorFundoParede, true}
	Vegetacao     = Elemento{'♣', CorVerde, CorPadrao, false}
	Vazio         = Elemento{' ', CorPadrao, CorPadrao, false}
	Botao         = Elemento{'■', CorVerde, CorVerde, true}
	BotaoVermelho = Elemento{'■', CorVermelho, CorVermelho, true}
	Trap          = Elemento{'▲', CorVermelho, CorPadrao, true}
)

// Cria e retorna uma nova instância do jogo
func jogoNovo() Jogo {
	// O ultimo elemento visitado é inicializado como vazio
	// pois o jogo começa com o personagem em uma posição vazia
	return Jogo{
		UltimoVisitado: Vazio,
		UltimaDirX:     1, // Define a direção inicial para a direita
		UltimaDirY:     0,
		StatusMsg:      "Direção de interacao: Direita", // Define a mensagem inicial
		trapX:          15,                              // coluna
		trapY:          3,                               // linha
		sentido:        1,                               // Sentido inicial (1 = descendo, -1 = subindo)
		limiteSuperior: 3,
		limiteInferior: 9,
	}
}

// Lê um arquivo texto linha por linha e constrói o mapa do jogo
func jogoCarregarMapa(nome string, jogo *Jogo) error {
	arq, err := os.Open(nome)
	if err != nil {
		return err
	}
	defer arq.Close()

	scanner := bufio.NewScanner(arq)
	y := 0
	for scanner.Scan() {
		linha := scanner.Text()
		var linhaElems []Elemento
		for x, ch := range linha {
			e := Vazio
			switch ch {
			case Parede.simbolo:
				e = Parede
			case Inimigo.simbolo:
				e = Inimigo
			case Vegetacao.simbolo:
				e = Vegetacao
			case Botao.simbolo:
				e = Botao
			case Personagem.simbolo:
				jogo.PosX, jogo.PosY = x, y // registra a posição inicial do personagem
			case Botao.simbolo:
				e = Botao
			case Trap.simbolo:
				jogo.trapX, jogo.trapY = x, y //usa o valor do mapa.txt, dai nao precisa inicializar valor de posição no código(não funcionou ):)
				//e = Trap //apagar depois de testar
			}
			linhaElems = append(linhaElems, e)
		}
		jogo.Mapa = append(jogo.Mapa, linhaElems)
		y++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// Verifica se o personagem pode se mover para a posição (x, y)
func jogoPodeMoverPara(jogo *Jogo, x, y int) bool {
	// Verifica se a coordenada Y está dentro dos limites verticais do mapa
	if y < 0 || y >= len(jogo.Mapa) {
		return false
	}

	// Verifica se a coordenada X está dentro dos limites horizontais do mapa
	if x < 0 || x >= len(jogo.Mapa[y]) {
		return false
	}

	// Verifica se o elemento de destino é tangível (bloqueia passagem)
	if jogo.Mapa[y][x].tangivel {
		return false
	}

	// Pode mover para a posição
	return true
}

// Move um elemento para a nova posição
func jogoMoverElemento(jogo *Jogo, x, y, dx, dy int) {
	nx, ny := x+dx, y+dy

	// Obtem elemento atual na posição
	elemento := jogo.Mapa[y][x] // guarda o conteúdo atual da posição

	jogo.Mapa[y][x] = jogo.UltimoVisitado   // restaura o conteúdo anterior
	jogo.UltimoVisitado = jogo.Mapa[ny][nx] // guarda o conteúdo atual da nova posição
	jogo.Mapa[ny][nx] = elemento            // move o elemento
}

func jogoAtualizarTrap(jogo *Jogo) {

	// Limpa a posição atual
	jogo.Mapa[jogo.trapY][jogo.trapX] = Vazio

	// Próxima posição
	novoY := jogo.trapY + jogo.sentido

	// Atualiza posição
	jogo.trapY = novoY
	jogo.Mapa[jogo.trapY][jogo.trapX] = Trap

	// Se bater em parede ou limite, inverte o sentido (valores foram ajustados para o mapa atual, algoritmo não vai funcionar para outros mapas com valores X e Y diferentes)
	if novoY <= jogo.limiteSuperior || novoY >= jogo.limiteInferior {
		jogo.sentido *= -1
		novoY = jogo.trapY + jogo.sentido
	}
}

func jogoPiscarBotao(jogo *Jogo, x, y int, intervaloMs int) {
	visivel := true
	for {
		if visivel {
			jogo.Mapa[y][x] = BotaoVermelho // atualiza o estado do mapa
		} else {
			jogo.Mapa[y][x] = Botao
		}
		visivel = !visivel

		// Apenas redesenha essa célula
		interfaceDesenharElemento(x, y, jogo.Mapa[y][x])
		interfaceAtualizarTela()

		time.Sleep(time.Millisecond * time.Duration(intervaloMs))
	}
}
