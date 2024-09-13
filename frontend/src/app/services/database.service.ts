import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class DatabaseService {
  private apiURL = 'http://localhost:8080';
  constructor(private http: HttpClient) {}

  getDatabases(): Observable<any> {
    return this.http.get(`${this.apiURL}/databases`);
  }
}
