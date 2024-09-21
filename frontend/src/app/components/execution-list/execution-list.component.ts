import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common'; // Importer CommonModule
import { ExecutionService } from '../../services/execution.service';
import { MessageService } from 'primeng/api';
import { TableModule } from 'primeng/table';
import { TagModule } from 'primeng/tag';

@Component({
  selector: 'app-execution-list',
  standalone: true,
  imports: [CommonModule, TableModule, TagModule],
  templateUrl: './execution-list.component.html',
  styleUrl: './execution-list.component.css',
  providers: [ExecutionService, MessageService],
})
export class ExecutionListComponent implements OnInit {
  executions: any[] = [];
  constructor(private executionService: ExecutionService) {}

  ngOnInit(): void {
    this.executionService.getExecutions().subscribe({
      next: (executions) => {
        this.executions = executions.items;
      },
      error: (error) => {
        console.error(error);
      },
    });
  }
  getExecutionType(type: string) {
    switch (type) {
      case 'restore':
        return 'warning';
      case 'backup':
        return 'success';
      default:
        return 'danger';
    }
  }
}
