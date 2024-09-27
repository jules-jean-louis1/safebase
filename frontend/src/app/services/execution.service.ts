import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class ExecutionService {
  constructor(private http: HttpClient) {}

  getExecutions(): Observable<any> {
    if (environment.useMockData) {
      return of([{ id: 1, name: 'MockDB' }]);
    }
    return this.http.get<any>(`/api/executions`);
  }
}
