import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

interface Backup {
  id: string;
  database_id: string;
  status: string;
  backup_type: string;
  filename: string;
  size?: string;
  error_msg?: string;
  log?: string;
  created_at: string;
  updated_at: string;
}

@Injectable({
  providedIn: 'root',
})
export class BackupService {
  constructor(private http: HttpClient) {}
  getBackups(): Observable<Backup> {
    return this.http.get<any>(`http://localhost:8080/backups`);
  }

  createBackup(databaseId: string): Observable<Backup> {
    return this.http.post<any>(`http://localhost:8080/backup`, {
      database_id: databaseId,
    });
  }
}
