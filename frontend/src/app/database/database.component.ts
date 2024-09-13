import { Component, OnInit } from '@angular/core';
import { DatabaseService } from '../services/database.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-database',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './database.component.html',
  styleUrl: './database.component.css',
})
export class DatabaseComponent implements OnInit {
  databases: any[] = [];
  errorMessage: string = '';

  constructor(private databaseService: DatabaseService) {}

  ngOnInit(): void {
    this.databaseService.getDatabases().subscribe({
      next: (data) => {
        this.databases = data; // Assigne les données reçues à la variable databases
      },
      error: (error) => {
        this.errorMessage =
          'Erreur lors de la récupération des bases de données.';
        console.error(error); // Affiche l'erreur dans la console
      },
    });
  }
}
