export interface MinesTile {
    id: number,
    around: number,
    mines: number,
    flags: number,
    uncovered: boolean,
}

export interface MinesBoard {
    id: number;
    width: number;
    height: number;
    mines: number;
    tiles: MinesTile[];
}

export interface MinesCommand {
    id: number;
    action: "nothing" | "flag" | "uncover";
}