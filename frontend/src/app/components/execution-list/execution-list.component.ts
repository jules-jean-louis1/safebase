import { Component, OnInit } from '@angular/core';
import { ExecutionService } from '../../services/execution.service';
import { MessageService } from 'primeng/api';
import { TableModule } from 'primeng/table';

@Component({
  selector: 'app-execution-list',
  standalone: true,
  imports: [],
  templateUrl: './execution-list.component.html',
  styleUrl: './execution-list.component.css',
  providers: [ExecutionService, MessageService, TableModule],
})
export class ExecutionListComponent implements OnInit {
  executions: any[] = [];
  constructor(private executionService: ExecutionService) {}

  ngOnInit(): void {
    this.executionService.getExecutions().subscribe({
      next: (executions) => {
        this.executions = executions;
      },
      error: (error) => {
        console.error(error);
      },
    });
  }
}
