package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 点击发牌按钮，抽排
func sendCards(g *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		x1, x2, y1, y2 := GetXYRangeInt(SendButton)
		if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
			g.Shuffle()
			g.DrawCard(5)
		}
	}
}

// 悬停判断， 判断是否悬停的同时，还需要返回index，表示卡牌的位置
func isMouseOverCard(g *Game) bool {
	x, y := ebiten.CursorPosition()
	return x >= int(g.cardX) && x <= int(g.cardX)+imageWidth && y >= int(g.cardY) && y <= int(g.cardY)+imageHeight
}

// 核心拖拽函数
/*
1、鼠标进来，判断鼠标位置，判断鼠标动作，判断选择了哪张图片，获取图片的index ※
2、
*/
func checkCardDrag(g *Game) {
	length := len(g.HandCards)

	index := getExpandIndex(length)
	if index == -1 {
		return
	}
	//开始拖动图片的判断
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.isDragging = true
	}

	//拖动过程中
	if g.isDragging {
		x, y := ebiten.CursorPosition()
		g.cardX = float64(x) - imageWidth/2
		g.cardY = float64(y) - imageHeight/2

		//松开鼠标的那一刹那
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			g.isDragging = false
			if isMouseOverEnemy() {
				g.showCard = false // 牌消失
			} else {
				g.cardX = g.cardOriginalX
				g.cardY = g.cardOriginalY
			}
		}
	}
}

// 判断鼠标出现在了卡牌消失的地方
func isMouseOverEnemy() bool {
	x, y := ebiten.CursorPosition()
	enemyX, enemyY := GetXY(EnemyPos)
	return x >= int(enemyX) && x <= int(enemyX)+imageWidth && y >= int(enemyY) && y <= int(enemyY)+imageHeight
}

// // 点击“发牌”的按钮
// func checkRefreshButtonClick(g *Game) {
// 	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
// 		x, y := ebiten.CursorPosition()
// 		if x >= ScreenWidth-50 && y <= 50 {
// 			g.RefreshCard()
// 		}
// 	}
// }
