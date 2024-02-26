import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MinesGameComponent } from './mines/game.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, MinesGameComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'Mines';
}
