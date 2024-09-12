import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class DatabaseService {
  private apiURL = 'http://localhost:8080';
  constructor(private http: HttpClient) {}

  getDatabase() {
    return this.http.get(`${this.apiURL}/database`);
  }
}
