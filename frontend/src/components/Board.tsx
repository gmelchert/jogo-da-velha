interface Props {
    board: number[];
    onPlay: (pos: number) => void;
}

export const Board = ({
    board,
    onPlay,
}: Props) => {

    return (
        <div className="grid grid-cols-3 gap-2 w-48 m-auto mt-4">
            {board.map((v, i) => (
                <button
                    key={i}
                    className="w-16 h-16 text-2xl border"
                    onClick={() => onPlay(i)}
                    disabled={v !== 0}
                >
                    {v === 1 ? "X" : v === 2 ? "O" : ""}
                </button>
            ))}
        </div>
    )
}