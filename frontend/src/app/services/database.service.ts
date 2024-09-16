import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class DatabaseService {
  // In a real-world application, this URL should be configurable
  // Since in private network in docker, it is not possible to use localhost but the name of the service
  private apiURL = 'http://backend:8080';
  constructor(private http: HttpClient) {}

  getDatabases(): Observable<any> {
    return this.http.get<any>(`${this.apiURL}/databases`);
  }

  addDatabase(database: any): Observable<any> {
    return this.http.post<any>(`http://localhost:8080/database`, database);
  }

  testConnection(database: any): Observable<any> {
    const params = new HttpParams()
      .set('host', database.host)
      .set('port', database.port)
      .set('username', database.username)
      .set('password', database.password)
      .set('dbName', database.database_name)
      .set('dbType', database.type);

    return this.http.get<any>(`http://localhost:8080/database/test`, {
      params,
    });
  }
}
