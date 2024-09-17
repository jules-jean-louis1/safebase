import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class BackupService {
  constructor(private http: HttpClient) {}
  getBackups(): Observable<any> {
    return this.http.get<any>(`http://localhost:8080/backups`);
  }

  getBackupFull(): Observable<any> {
    return this.http.get<any>(`http://localhost:8080/backups/full`);
  }

  createBackup(databaseId: string): Observable<any> {
    return this.http.post<any>(`http://localhost:8080/backup`, {
      database_id: databaseId,
    });
  }
}
