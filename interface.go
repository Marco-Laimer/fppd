// interface.go - Interface gráfica do jogo usando termbox
// O código abaixo implementa a interface gráfica do jogo usando a biblioteca termbox-go.
// A biblioteca termbox-go é uma biblioteca de interface de terminal que permite desenhar
// elementos na tela, capturar eventos do teclado e gerenciar a aparência do terminal.

package main

import (
	"strings" // Adicionado para lidar com as strings da barra de status para mostrar 2 ou mais linhas /Marco

	"github.com/nsf/termbox-go"
)

// Define um tipo Cor para encapsuladar as cores do termbox
type Cor = termbox.Attribute

// Definições de cores utilizadas no jogo
const (
	CorPadrao      Cor = termbox.ColorDefault
	CorCinzaEscuro     = termbox.ColorDarkGray
	CorVermelho        = termbox.ColorRed
	CorVerde           = termbox.ColorGreen
	CorParede          = termbox.ColorWhite | termbox.AttrBold | termbox.AttrDim
	CorFundoParede     = termbox.ColorBlack
	CorTexto           = termbox.ColorDarkGray
)

// EventoTeclado representa uma ação detectada do teclado (como mover, sair ou interagir)
type EventoTeclado struct {
	Tipo  string // "sair", "interagir", "mover", "direcao"
	Tecla rune   // Tecla pressionada, usada no caso de movimento
}

// Inicializa a interface gráfica usando termbox
func interfaceIniciar() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
}

// Encerra o uso da interface termbox
func interfaceFinalizar() {
	termbox.Close()
}

// Lê um evento do teclado e o traduz para um EventoTeclado
func interfaceLerEventoTeclado() EventoTeclado {
	ev := termbox.PollEvent()
	if ev.Type != termbox.EventKey {
		return EventoTeclado{}
	}
	if ev.Key == termbox.KeyEsc {
		return EventoTeclado{Tipo: "sair"}
	}
	if ev.Ch == 'e' {
		return EventoTeclado{Tipo: "interagir"}
	}
	// Identifica as setas do teclado para a direção de interação
	if ev.Key == termbox.KeyArrowUp {
		return EventoTeclado{Tipo: "direcao", Tecla: 'w'}
	}
	if ev.Key == termbox.KeyArrowDown {
		return EventoTeclado{Tipo: "direcao", Tecla: 's'}
	}
	if ev.Key == termbox.KeyArrowLeft {
		return EventoTeclado{Tipo: "direcao", Tecla: 'a'}
	}
	if ev.Key == termbox.KeyArrowRight {
		return EventoTeclado{Tipo: "direcao", Tecla: 'd'}
	}

	return EventoTeclado{Tipo: "mover", Tecla: ev.Ch}
}

// Renderiza todo o estado atual do jogo na tela
func interfaceDesenharJogo(jogo *Jogo) {
	interfaceLimparTela()

	// Desenha todos os elementos do mapa
	for y, linha := range jogo.Mapa {
		for x, elem := range linha {
			interfaceDesenharElemento(x, y, elem)
		}
	}

	// Desenha o personagem sobre o mapa
	interfaceDesenharElemento(jogo.PosX, jogo.PosY, Personagem)

	// Desenha a barra de status
	interfaceDesenharBarraDeStatus(jogo)

	// Força a atualização do terminal
	interfaceAtualizarTela()
}

// Limpa a tela do terminal
func interfaceLimparTela() {
	termbox.Clear(CorPadrao, CorPadrao)
}

// Força a atualização da tela do terminal com os dados desenhados
func interfaceAtualizarTela() {
	termbox.Flush()
}

// Desenha um elemento na posição (x, y)
func interfaceDesenharElemento(x, y int, elem Elemento) {
	termbox.SetCell(x, y, elem.simbolo, elem.cor, elem.corFundo)
}

// Exibe uma barra de status com informações úteis ao jogador
func interfaceDesenharBarraDeStatus(jogo *Jogo) {
	// Divide a mensagem de status em linhas usando o caractere de nova linha
	linhas := strings.Split(jogo.StatusMsg, "\n")

	// Itera sobre cada linha e desenha na tela
	for i, linha := range linhas {
		for j, c := range linha {
			// A posição Y é calculada com base no índice da linha
			termbox.SetCell(j, len(jogo.Mapa)+1+i, c, CorTexto, CorPadrao)
		}
	}

	// Instruções fixas
	msg := "Use WASD para mover e E para interagir,Setas para escolher o locar de interagir e ESC para sair."
	for i, c := range msg {
		termbox.SetCell(i, len(jogo.Mapa)+3, c, CorTexto, CorPadrao)
	}
}
