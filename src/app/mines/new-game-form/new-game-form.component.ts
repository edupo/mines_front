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
    width: [10, [Validators.min(1), Validators.max(20)]],
    height: [10, [Validators.min(1), Validators.max(20)]],
    mines: [10, [Validators.min(1), Validators.max(100)]],
  })

  constructor(private formBuilder: FormBuilder) { }

  submit() {
    var value = this.newGameForm.value
    this.onSubmit.emit({ width: value.width, height: value.height, mines: value.mines } as MinesGameSettings);
  }
}
