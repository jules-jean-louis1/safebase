import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { DatabaseComponent } from './pages/database/database.component';
import { BackupComponent } from './pages/backup/backup.component';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'backups', component: BackupComponent },
  { path: 'databases', component: DatabaseComponent },
];
