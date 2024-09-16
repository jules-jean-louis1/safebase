import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  private refreshListSubject = new Subject<void>();

  get refreshList$() {
    return this.refreshListSubject.asObservable();
  }

  notifyRefreshList() {
    this.refreshListSubject.next();
  }
}