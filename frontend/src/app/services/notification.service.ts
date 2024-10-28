import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  private refreshListSubject = new Subject<void>();

  notifyRefreshList() {
    console.log('Notification sent from NotificationService'); // AJOUTER CECI
    this.refreshListSubject.next();
  }

  getRefreshListObservable() {
    return this.refreshListSubject.asObservable();
  }
}
