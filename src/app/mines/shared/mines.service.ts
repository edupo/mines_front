import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { MinesBoard, MinesCommand } from './mines.models';
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

  getBoard(): Observable<MinesBoard> {
    return this.httpClient.get<MinesBoard>(this.url)
      .pipe(
        tap(_ => this.log('fetched board')),
        catchError(this.handleError<MinesBoard>('getBoard')),
      );
  }

  sendCommand(command: MinesCommand): Observable<MinesBoard> {
    return this.httpClient.post<MinesBoard>(this.url, command, this.httpOptions).pipe(
      tap(_ => this.log('sendCommand: new board status')),
      catchError(this.handleError<MinesBoard>('sendCommand')),
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
