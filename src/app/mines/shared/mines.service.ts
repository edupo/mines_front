import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { MinesGameStatus, MinesCommand, MinesGameSettings } from './mines.models';
import { Observable, of, tap, map, catchError } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class MinesService {

  private url = "http://localhost:4201/mines";

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private httpClient: HttpClient,
  ) { }

  newGame(settings: MinesGameSettings): Observable<MinesGameStatus> {
    return this.httpClient.post<MinesGameStatus>(this.url, settings, this.httpOptions).pipe(
      tap(_ => this.log('newGame')),
      catchError(this.handleError<MinesGameStatus>('newGame')),
    )
  }

  getGame(status: MinesGameStatus): Observable<MinesGameStatus> {
    return this.httpClient.get<MinesGameStatus>(`${this.url}/${status.id}`)
      .pipe(
        tap(_ => this.log('getGame')),
        catchError(this.handleError<MinesGameStatus>('getBoard')),
      );
  }

  sendCommand(status: MinesGameStatus, command: MinesCommand): Observable<MinesGameStatus> {
    return this.httpClient.post<MinesGameStatus>(`${this.url}/${status.id}`, command, this.httpOptions).pipe(
      tap(_ => this.log(`sendCommand ${command.id}:${command.action}`)),
      catchError(this.handleError<MinesGameStatus>('sendCommand')),
    )
  }

  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      console.error(`${operation} failes: ${error.message}`);
      return of(result as T);
    }
  }

  private log(message: string) {
    console.log(`MinesService: ${message}`);
  }

}
