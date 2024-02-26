import { Component, OnInit } from '@angular/core';
import { NgStyle, NgClass, NgFor } from '@angular/common';

import { TileComponent } from './tile/tile.component';
import { MinesService } from './shared/mines.service';
import { MinesBoard } from './shared/mines.models';

@Component({
  selector: 'mines-game',
  standalone: true,
  imports: [TileComponent, NgStyle, NgClass, NgFor],
  templateUrl: './game.component.html',
  styleUrl: './game.component.scss',
})
export class MinesGameComponent implements OnInit {
  board: MinesBoard = {
    id: 0,
    mines: 0,
    width: 0,
    height: 0,
    tiles: [],
  };

  constructor(private minesService: MinesService) { }

  ngOnInit(): void {
    this.getBoard();
  }

  getBoard(): void {
    this.minesService.getBoard().subscribe(board => this.board = board);
  }

  calculateStyle() {
    return {
      'grid-template-columns': `repeat(${this.board.width}, minmax(0, 1fr))`,
    };
  }
}
