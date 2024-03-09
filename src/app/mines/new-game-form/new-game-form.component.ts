import { Component, Output, EventEmitter } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { MinesGameSettings } from '../shared/mines.models';

@Component({
  selector: 'mines-new-game-form',
  standalone: true,
  templateUrl: './new-game-form.component.html',
  styleUrl: './new-game-form.component.scss',
  imports: [ReactiveFormsModule],
})
export class NewGameFormComponent {
  @Output() onSubmit = new EventEmitter<MinesGameSettings>();

  newGameForm = this.formBuilder.group({
    width: [10, [Validators.min(1), Validators.max(30)]],
    height: [10, [Validators.min(1), Validators.max(30)]],
    mines: [10, [Validators.min(1), Validators.max(100)]],
  })

  settings = {
    easy: { width: 8, height: 8, mines: 10 } as MinesGameSettings,
    medium: { width: 13, height: 15, mines: 40 } as MinesGameSettings,
    hard: { width: 16, height: 30, mines: 99 } as MinesGameSettings,
  }

  constructor(private formBuilder: FormBuilder) { }

  submit(level: string = '') {
    var _settings: MinesGameSettings
    switch (level) {
      case 'custom':
        var value = this.newGameForm.value
        _settings = { width: value.width, height: value.height, mines: value.mines } as MinesGameSettings
        break;
      case 'easy':
      case 'medium':
      case 'hard':
        _settings = this.settings[level]
        break;
      default:
        _settings = this.settings['easy']
        break;
    }
    this.onSubmit.emit(_settings);
  }
}
