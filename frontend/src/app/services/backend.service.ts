import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, map, Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BackendService {

  constructor(private http: HttpClient) { }
    
  isBackendReachable(): Observable<boolean> {
    return this.http.get('/api/health-check').pipe(
      map(() => true),
      catchError(() => of(false))
    );
  }
}
