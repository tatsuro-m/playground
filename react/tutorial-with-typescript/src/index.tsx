import {ReactElement, useState, VFC} from 'react';
import ReactDOM from 'react-dom';
import './index.css';

type FillSquare = 'X' | 'O' | null;

interface squareProps {
  value: FillSquare
  onClick: () => void;
}

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

const Board: VFC = () => {
  const [squares, setSquares] = useState<FillSquare[]>(Array(9).fill(null))

  const handleClick = (i: number): void => {
    const squaresSlice = squares.slice();
    squaresSlice[i] = 'X';
    setSquares(squaresSlice);
  };

  const renderSquare = (i: number): ReactElement => {
    return <Square value={squares[i]} onClick={() => handleClick(i)} />
  }

  const status = 'Next player: X';

  return (
    <div>
      <div className="status">{status}</div>
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

const Game: VFC = () => (
  <div className="game">
    <div className="game-board">
      <Board/>
    </div>
    <div className="game-info">
      <div>{/* status */}</div>
      <ol>{/* TODO */}</ol>
    </div>
  </div>
);

// ========================================

ReactDOM.render(<Game/>, document.getElementById('root'));
