import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class BackupService {
  constructor(private http: HttpClient) {}
  getBackups(): Observable<any> {
    return this.http.get<any>(`/api/backups`);
  }

  getBackupFull(): Observable<any> {
    return this.http.get<any>(`/api/backups/full`);
  }

  createBackup(databaseId: string): Observable<any> {
    return this.http.post<any>(`/api/backup`, {
      database_id: databaseId,
    });
  }
  getBackupsBy(paramsObj: any): Observable<any> {
    let params = new HttpParams();

    // Ajouter dynamiquement les paramètres à l'URL
    for (const key in paramsObj) {
      if (paramsObj.hasOwnProperty(key) && paramsObj[key]) {
        params = params.append(key, paramsObj[key]);
      }
    }

    return this.http.get<any>(`/api/backups/options`, { params });
  }

  deleteBackup(backupId: string): Observable<any> {
    return this.http.delete<any>(`/api/backup/${backupId}`);
  }
}
