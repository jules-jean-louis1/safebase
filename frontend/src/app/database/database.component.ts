import { Component, OnInit } from '@angular/core';
import { DatabaseService } from '../services/database.service';

@Component({
  selector: 'app-database',
  standalone: true,
  imports: [],
  templateUrl: './database.component.html',
  styleUrl: './database.component.css',
})
export class DatabaseComponent implements OnInit {
  database: any;

  constructor(private databaseService: DatabaseService) {}

  ngOnInit(): void {}
}
