import { Component } from "@angular/core";
import { ExecutionListComponent } from "../../components/execution-list/execution-list.component";

@Component({
  selector: "app-execution",
  standalone: true,
  imports: [ExecutionListComponent],
  templateUrl: "./execution.component.html",
  styleUrl: "./execution.component.css",
})
export class ExecutionComponent {

}
