import { Component, OnInit } from '@angular/core';
import { TableModule } from 'primeng/table';
import { CommonModule } from '@angular/common';
import { DatabaseService } from '../../services/database.service';

@Component({
  selector: 'app-database-list',
  standalone: true,
  imports: [TableModule, CommonModule],
  templateUrl: './database-list.component.html',
  styleUrl: './database-list.component.css',
  providers: [DatabaseService],
})
export class DatabaseListComponent implements OnInit {
  databases: any[] = [];
  errorMessage: string = '';

  constructor(private databaseService: DatabaseService) {}

  ngOnInit(): void {
    this.databaseService.getDatabases().subscribe({
      next: (data) => {
        this.databases = data;
      },
      error: (error) => {
        this.errorMessage =
          'Erreur lors de la récupération des bases de données.';
        console.error(error);
      },
    });
  }
}
