import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { DatabaseComponent } from './database/database.component';
import { BackupComponent } from './backup/backup.component';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'databases', component: DatabaseComponent },
  { path: 'backups', component: BackupComponent },
];
