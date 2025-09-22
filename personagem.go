// // personagem.go - Funções para movimentação e ações do personagem
// package main

// import (
// 	"fmt"
// 	"time"
// 	// Adicione esta linha
// )

// // Uma função auxiliar para obter a direcao (dx, dy) a partir de uma tecla
// func obterDirecao(tecla rune) (int, int) {
// 	var dx, dy int
// 	switch tecla {
// 	case 'w':
// 		dy = -1 // Move para cima
// 	case 'a':
// 		dx = -1 // Move para a esquerda
// 	case 's':
// 		dy = 1 // Move para baixo
// 	case 'd':
// 		dx = 1 // Move para a direita
// 	}
// 	return dx, dy
// }

// // Atualiza a posição do personagem com base na tecla pressionada (WASD)
// func personagemMover(tecla rune, jogo *Jogo) {
// 	dx, dy := obterDirecao(tecla)

// 	nx, ny := jogo.PosX+dx, jogo.PosY+dy
// 	// Verifica se o movimento é permitido e realiza a movimentação
// 	if jogoPodeMoverPara(jogo, nx, ny) {
// 		jogoMoverElemento(jogo, jogo.PosX, jogo.PosY, dx, dy)
// 		jogo.PosX, jogo.PosY = nx, ny
// 	}
// }

// // Define o que ocorre quando o jogador pressiona a tecla de interacao
// // Neste exemplo, apenas exibe uma mensagem de status
// // Você pode expandir essa função para incluir lógica de interacao com objetos
// func personagemInteragir(jogo *Jogo) {
// 	// Atualmente apenas exibe uma mensagem de status
// 	jogo.StatusMsg += fmt.Sprintf("\n -Interagindo em (%d, %d)", jogo.PosX, jogo.PosY)
// 	interfaceDesenharElemento(jogo.PosX+1, jogo.PosY, Inimigo)
// 	time.Sleep(time.Second * 5)

// }
// func personagemDefinirDirecao(tecla rune, jogo *Jogo) {
// 	dx, dy := obterDirecao(tecla)

// 	jogo.UltimaDirX = dx
// 	jogo.UltimaDirY = dy

// 	// Atualiza a mensagem de status com base na direcao
// 	switch tecla {
// 	case 'w':
// 		jogo.StatusMsg = "Direcao de interacao: Cima"
// 	case 's':
// 		jogo.StatusMsg = "Direcao de interacao: Baixo"
// 	case 'a':
// 		jogo.StatusMsg = "Direcao de interacao: Esquerda"
// 	case 'd':
// 		jogo.StatusMsg = "Direcao de interacao: Direita"
// 	}
// }

// // Processa o evento do teclado e executa a ação correspondente
// func personagemExecutarAcao(ev EventoTeclado, jogo *Jogo) bool {
// 	switch ev.Tipo {

//		case "direcao":
//			personagemDefinirDirecao(ev.Tecla, jogo) // Nova função para a lógica das setas
//		case "sair":
//			// Retorna false para indicar que o jogo deve terminar
//			return false
//		case "interagir":
//			// Executa a ação de interacao
//			personagemInteragir(jogo)
//		case "mover":
//			// Move o personagem com base na tecla
//			personagemMover(ev.Tecla, jogo)
//		}
//		return true // Continua o jogo
//	}
//
// =======
// // personagem.go - Funções para movimentação e ações do personagem
package main

import (
	"fmt"
	// Adicione esta linha
)

// Uma função auxiliar para obter a direcao (dx, dy) a partir de uma tecla
func obterDirecao(tecla rune) (int, int) {
	var dx, dy int
	switch tecla {
	case 'w':
		dy = -1 // Move para cima
	case 'a':
		dx = -1 // Move para a esquerda
	case 's':
		dy = 1 // Move para baixo
	case 'd':
		dx = 1 // Move para a direita
	}
	return dx, dy
}

// Atualiza a posição do personagem com base na tecla pressionada (WASD)
func personagemMover(tecla rune, jogo *Jogo) {
	dx, dy := obterDirecao(tecla)

	nx, ny := jogo.PosX+dx, jogo.PosY+dy
	// Verifica se o movimento é permitido e realiza a movimentação
	if jogoPodeMoverPara(jogo, nx, ny) {
		jogoMoverElemento(jogo, jogo.PosX, jogo.PosY, dx, dy)
		jogo.PosX, jogo.PosY = nx, ny
	}
}

// Define o que ocorre quando o jogador pressiona a tecla de interacao
// Neste exemplo, apenas exibe uma mensagem de status
// Você pode expandir essa função para incluir lógica de interacao com objetos
func personagemInteragir(jogo *Jogo) {
	//Coloca qual elemento esta interagindo
	alvoX := jogo.PosX + jogo.UltimaDirX
	alvoY := jogo.PosY + jogo.UltimaDirY

	jogo.StatusMsg += fmt.Sprintf("\n -Interagindo em (%d, %d)", alvoX, alvoY)

	//time.Sleep(time.Second * 5)

}
func personagemDefinirDirecao(tecla rune, jogo *Jogo) {
	dx, dy := obterDirecao(tecla)

	jogo.UltimaDirX = dx
	jogo.UltimaDirY = dy

	// Atualiza a mensagem de status com base na direcao
	switch tecla {
	case 'w':
		jogo.StatusMsg = "Direcao de interacao: Cima"
	case 's':
		jogo.StatusMsg = "Direcao de interacao: Baixo"
	case 'a':
		jogo.StatusMsg = "Direcao de interacao: Esquerda"
	case 'd':
		jogo.StatusMsg = "Direcao de interacao: Direita"
	}
}

// Processa o evento do teclado e executa a ação correspondente
func personagemExecutarAcao(ev EventoTeclado, jogo *Jogo) bool {
	switch ev.Tipo {

	case "direcao":
		personagemDefinirDirecao(ev.Tecla, jogo) // Nova função para a lógica das setas
	case "sair":
		// Retorna false para indicar que o jogo deve terminar
		return false
	case "interagir":
		// Executa a ação de interacao
		personagemInteragir(jogo)
	case "mover":
		// Move o personagem com base na tecla
		personagemMover(ev.Tecla, jogo)
	}
	return true // Continua o jogo
}
