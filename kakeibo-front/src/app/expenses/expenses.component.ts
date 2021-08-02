import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Expense } from '../models/expense.model';
import { ExpensesService } from './expenses.service';

@Component({
  selector: 'app-expenses',
  templateUrl: './expenses.component.html',
  styleUrls: ['./expenses.component.css']
})
export class ExpensesComponent implements OnInit {
  @Input() columns = ['id', 'name'];

  dataSource: Expense[];

  constructor(
    private service: ExpensesService,
    private router: Router
  ) { }

  ngOnInit(): void {

    this.service.getExpenses()
    .subscribe(data => this.dataSource = data)
  }

  clickedRows(row: any){
    this.router.navigate(['/expenses/insert', row.id]);
    
  }

  createNew(){
    this.router.navigate(['/expenses/insert']);
  }

}
