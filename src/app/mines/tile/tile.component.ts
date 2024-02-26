import { Component, Input } from '@angular/core';
import { NgIf, NgClass } from '@angular/common';

import { MinesTile } from '../shared/mines.models';

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
