import {ReactElement, useState, VFC} from 'react';
import ReactDOM from 'react-dom';
import './index.css';

type FillSquare = 'X' | 'O' | null;

interface squareProps {
  value: FillSquare
  onClick: () => void;
}
type BoardProps = {
  squares: FillSquare[];
  onClick: (i: number) => void;
};

const calculateWinner = (squares: FillSquare[]) => {
  const lines = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6],
  ];
  for (let i = 0; i < lines.length; i += 1) {
    const [a, b, c] = lines[i];
    if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
      return squares[a];
    }
  }

  return null;
};

const Square: VFC<squareProps> = (props) => {
  const {value, onClick} = props
  return (
    <button type="button" className="square">
      <button type="button" className="square" onClick={onClick}>
        {value}
      </button>
    </button>
  )
}

const Board: VFC<BoardProps> = (props) => {
  const {squares, onClick} = props;

  const renderSquare = (i: number): ReactElement => {
    return <Square value={squares[i]} onClick={() => onClick(i)} />
  }

  return (
    <div>
      <div className="board-row">
        {renderSquare(0)}
        {renderSquare(1)}
        {renderSquare(2)}
      </div>
      <div className="board-row">
        {renderSquare(3)}
        {renderSquare(4)}
        {renderSquare(5)}
      </div>
      <div className="board-row">
        {renderSquare(6)}
        {renderSquare(7)}
        {renderSquare(8)}
      </div>
    </div>
  );
};

const Game: VFC = () => {
  const [history, setHistory] = useState([{squares: Array(9).fill(null)}]);
  const [xIsNext, setXIsNext] = useState(true);

  const jumpTo = (step: number) => {
    setStepNumber(step);
    setXIsNext(step % 2 === 0);
  };

  const moves = history.map((step, move) => {
    const desc = move ? `Go to move #${move}` : 'Go to game start';

    return (
      <li key={move.toString()}>
        <button type="button" onClick={() => jumpTo(move)}>
          {desc}
        </button>
      </li>
    )});

    const [stepNumber, setStepNumber] = useState(0);
  const handleClick = (i: number): void => {
    const historySlice = history.slice(0, stepNumber + 1);
    const current = historySlice[historySlice.length - 1];
    const squares = current.squares.slice();

    // 勝者確定かマスが埋まっていたら、クリックしてもマスが変化しないようにする
    if (calculateWinner(squares) || squares[i]) {
      return;
    }

    squares[i] = xIsNext ? 'X' : 'O';
    setHistory([...historySlice, { squares }]);
    setStepNumber(historySlice.length);
    setXIsNext(!xIsNext);
  };

  const current = history[history.length - 1];
  const winner = calculateWinner(current.squares);
  const status = winner
    ? `Winner: ${winner}`
    : `Next player: ${xIsNext ? 'X' : 'O'}`;

  return (
    <div className="game">
      <div className="game-board">
        <Board squares={current.squares} onClick={(i) => handleClick(i)} />
      </div>
      <div className="game-info">
        <div>{status}</div>
        <ol>{/* TODO */}</ol>
      </div>
    </div>
  );
};


// ========================================

ReactDOM.render(<Game/>, document.getElementById('root'));
