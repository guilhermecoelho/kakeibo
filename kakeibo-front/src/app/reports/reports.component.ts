import { Component, Input, OnInit } from '@angular/core';
import { Report } from '../models/report.model';
import { ReportRequest } from '../models/reportRequest.model';
import { ReportsService } from './reports.service';
import { Moviment, ReportView } from './reportView.model';

@Component({
  selector: 'app-reports',
  templateUrl: './reports.component.html',
  styleUrls: ['./reports.component.css']
})
export class ReportsComponent implements OnInit {

  @Input() columns = ['date', 'name', 'amount'];

  reportView: ReportView;
  reportRequest: ReportRequest;

  constructor
  (
    private service: ReportsService,
  ) 
  {
    this.reportView = new ReportView();
    this.reportView.moviments = new Array()
  }

  ngOnInit(): void {
    
    this.reportRequest = new ReportRequest();

    let actualMonth = new Date().getMonth();
    let actualYear = new Date().getFullYear();

    this.reportRequest.dateStart = new Date(actualYear,actualMonth,1);
    this.reportRequest.dateFinish = new Date(actualYear,actualMonth+1,0);

    this.service.postReport(this.reportRequest)
    .subscribe(data => {
       this.reportView = new ReportView();
       this.reportView.moviments = new Array()

      let mergeMoviments = [...data.Expenses, ...data.Incomes]

      data.Expenses.forEach(item => {
        this.reportView.moviments.push(this.populateMoviment(item, true));
      });
      data.Incomes.forEach(item => {
        this.reportView.moviments.push(this.populateMoviment(item, false));
      });

      this.reportView.totalIncome = data.totalIncome
      this.reportView.totalExpense = data.totalExpense
      this.reportView.balance = data.balance
      this.reportView.dateStart = data.dateStart
      this.reportView.dateFinish = data.dateFinish


      console.log(this.reportView)
      //this.dataSource = data
    })
  }

  populateMoviment(item: any, isExpense:boolean): any{
    let moviment = new Moviment();
    moviment.name = item.name
    moviment.date = item.date
    moviment.amount = item.amount
    moviment.isExpense = isExpense

    return moviment
  }
}
