import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { DatabaseComponent } from './pages/database/database.component';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'databases', component: DatabaseComponent },
];
