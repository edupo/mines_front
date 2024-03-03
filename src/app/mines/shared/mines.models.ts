export interface MinesTileStatus {
    id: number,
    around: number,
    mines: number,
    flags: number,
    uncovered: boolean,
}

export interface MinesGameStatus {
    id: number;
    width: number;
    height: number;
    mines: number;
    tiles: MinesTileStatus[];
}

export interface MinesCommand {
    id: number;
    action: "nothing" | "flag" | "uncover";
}

export interface MinesGameSettings {
    width: number;
    height: number;
    mines: number;
}