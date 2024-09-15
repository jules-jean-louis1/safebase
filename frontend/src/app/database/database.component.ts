import { Component, OnInit } from '@angular/core';
import { DatabaseService } from '../services/database.service';
import { CommonModule } from '@angular/common';
import { DatabaseListComponent } from '../components/database-list/database-list.component';

@Component({
  selector: 'app-database',
  standalone: true,
  imports: [CommonModule, DatabaseListComponent],
  templateUrl: './database.component.html',
  styleUrl: './database.component.css',
})
export class DatabaseComponent {}
