import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Income } from '../models/income.model';
import { IncomeService } from './income.service';

@Component({
  selector: 'app-incomes',
  templateUrl: './incomes.component.html',
  styleUrls: ['./incomes.component.css']
})
export class IncomesComponent implements OnInit {
  @Input() columns = ['id', 'name','amount', 'date', 'note'];

  dataSource: Income[];

  constructor(
    private service: IncomeService,
    private router: Router
  ) { }

  ngOnInit(): void {

    this.service.getIncomes()
    .subscribe(data => this.dataSource = data)
  }

  clickedRows(row: any){
    this.router.navigate(['/incomes/insert', row.id]);
    
  }

  createNew(){
    this.router.navigate(['/incomes/insert']);
  }
}
