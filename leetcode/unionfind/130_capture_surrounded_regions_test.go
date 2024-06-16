package unionfind

// 130. Surrounded Regions
// https://leetcode.com/problems/surrounded-regions/description/
func captureSurroundedRegions(board [][]byte) {
	if board == nil || len(board) <= 1 || len(board[0]) <= 1 {
		return
	}

	const (
		X = 'X'
		O = 'O'
		T = 'T'
	)

	var dfs func(board [][]byte, i, j int)
	dfs = func(board [][]byte, i, j int) {
		board[i][j] = T
		if i > 0 && board[i-1][j] == O {
			dfs(board, i-1, j)
		}
		if i < len(board)-1 && board[i+1][j] == O {
			dfs(board, i+1, j)
		}
		if j > 0 && board[i][j-1] == O {
			dfs(board, i, j-1)
		}
		if j < len(board[0])-1 && board[i][j+1] == O {
			dfs(board, i, j+1)
		}
	}

	h := len(board)
	w := len(board[0])

	for i := 0; i < w; i++ {
		if board[0][i] == O {
			dfs(board, 0, i)
		}
		if board[h-1][i] == O {
			dfs(board, h-1, i)
		}
	}
	for i := 1; i < h-1; i++ {
		if board[i][0] == O {
			dfs(board, i, 0)
		}
		if board[i][w-1] == O {
			dfs(board, i, w-1)
		}
	}

	for i, row := range board {
		for j, ch := range row {
			if ch == T {
				board[i][j] = O
			} else if ch == O {
				board[i][j] = X
			}
		}
	}
}
