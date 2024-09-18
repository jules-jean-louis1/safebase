import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

interface CreateRestore {
  database_id: string;
  backup_id: string;
}

@Injectable({
  providedIn: 'root'
})
export class RestoreService {

  constructor(private http: HttpClient) { 
  }
  insertRestore(restore: CreateRestore): Observable<any> {
    return this.http.post<any>(`http://localhost:8080/restore`, restore);
  }
}
