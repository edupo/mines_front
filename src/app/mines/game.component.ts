import { Component, OnInit } from '@angular/core';
import { NgStyle, NgClass, NgFor } from '@angular/common';

import { TileComponent } from './tile/tile.component';
import { NewGameFormComponent } from './new-game-form/new-game-form.component';
import { MinesService } from './shared/mines.service';
import { MinesGameStatus, MinesGameSettings, MinesCommand } from './shared/mines.models';

@Component({
  selector: 'mines-game',
  standalone: true,
  imports: [TileComponent, NewGameFormComponent, NgStyle, NgClass, NgFor],
  templateUrl: './game.component.html',
  styleUrl: './game.component.scss',
})
export class MinesGameComponent implements OnInit {
  status: MinesGameStatus = {
    id: 0,
    mines: 0,
    width: 0,
    height: 0,
    tiles: [],
  };

  constructor(private service: MinesService) { }

  ngOnInit(): void {
    this.newGame({ width: 10, height: 10, mines: 10 } as MinesGameSettings)
  }

  newGame(settings: MinesGameSettings): void {
    this.service
      .newGame(settings)
      .subscribe(new_status => this.status = new_status);
  }

  sendCommand(command: MinesCommand): void {
    this.service
      .sendCommand(this.status, command)
      .subscribe(new_status => this.status = new_status)
  }

  calculateStyle() {
    return {
      'grid-template-columns': `repeat(${this.status.width}, minmax(0, 1fr))`,
    };
  }
}
