import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class DatabaseService {
  constructor(private http: HttpClient) {}

  getDatabases(): Observable<any> {
    if (environment.useMockData) {
      return of([{ id: 1, name: 'MockDB' }]);
    }
    return this.http.get<any>(`/api/databases`);
  }

  addDatabase(database: any): Observable<any> {
    return this.http.post<any>(`/api/database`, database);
  }

  testConnection(database: any): Observable<any> {
    const params = new HttpParams()
      .set('host', database.host)
      .set('port', database.port)
      .set('username', database.username)
      .set('password', database.password)
      .set('dbName', database.database_name)
      .set('dbType', database.type);

    return this.http.get<any>(`/api/database/test`, {
      params,
    });
  }

  updateDatabase(database: any): Observable<any> {
    return this.http.put<any>(`/api/database`, database);
  }

  deleteDatabase(id: string): Observable<any> {
    return this.http.delete<any>(`/api/database/${id}`);
  }
}
