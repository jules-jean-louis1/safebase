import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  private refreshListSubject = new Subject<void>();

  // Observable que les composants vont écouter
  get refreshList$() {
    return this.refreshListSubject.asObservable();
  }

  // Méthode pour notifier les composants
  notifyRefreshList() {
    console.log('Notification triggered in service'); // Ajout d'un log pour vérifier
    this.refreshListSubject.next();
  }
}
