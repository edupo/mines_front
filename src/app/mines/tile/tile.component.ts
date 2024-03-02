import { Component, EventEmitter, Input, Output } from '@angular/core';
import { NgIf, NgClass } from '@angular/common';

import { MinesService } from '../shared/mines.service';
import { MinesTile, MinesCommand, MinesBoard } from '../shared/mines.models';

@Component({
  selector: 'mines-tile',
  standalone: true,
  imports: [NgIf, NgClass],
  templateUrl: './tile.component.html',
  styleUrl: './tile.component.scss'
})
export class TileComponent {
  @Input() tile: MinesTile = {
    id: 0,
    around: 0,
    mines: 0,
    flags: 0,
    uncovered: false
  }
  @Output() onUpdate = new EventEmitter<MinesBoard>();

  constructor(private minesService: MinesService) { }

  sendCommand(action: "nothing" | "flag" | "uncover"): void {
    console.log(action)
    this.minesService.sendCommand({ id: this.tile.id, action: action } as MinesCommand)
      .subscribe(
        board => this.onUpdate.emit(board)
      )

  }

  calculateClasses() {
    return {
      'mines__tile': true,
      'mines__tile__covered': !this.tile.uncovered,
      'mines__tile__uncovered': this.tile.uncovered,
      'mines__tile__mine': this.tile.mines,
      'mines__tile__color__r1': !this.tile.mines && this.tile.around == 1,
      'mines__tile__color__r2': !this.tile.mines && this.tile.around == 2,
      'mines__tile__color__r3': !this.tile.mines && this.tile.around == 3,
      'mines__tile__color__r4': !this.tile.mines && this.tile.around == 4,
      'mines__tile__color__r5': !this.tile.mines && this.tile.around == 5,
      'mines__tile__color__r6': !this.tile.mines && this.tile.around == 6,
      'mines__tile__color__r7': !this.tile.mines && this.tile.around == 7,
      'mines__tile__color__r8': !this.tile.mines && this.tile.around == 8,
      'mines__tile__color__r9': !this.tile.mines && this.tile.around > 9,
    }
  }
}
