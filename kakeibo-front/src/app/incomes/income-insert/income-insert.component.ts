import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { CategoriesService } from 'src/app/categories/categories.service';
import { Category } from 'src/app/models/category.model';
import { Income } from 'src/app/models/income.model';
import { IncomeService } from '../income.service';

@Component({
  selector: 'app-income-insert',
  templateUrl: './income-insert.component.html',
  styleUrls: ['./income-insert.component.css']
})
export class IncomeInsertComponent implements OnInit {

  incomeForm: FormGroup
  income: Income

  categories: Category[]

  constructor(
    private formBuilder: FormBuilder, 
    private service: IncomeService, 
    private router: Router,
    private route: ActivatedRoute,
    private categoryService: CategoriesService
    ) {}
   
  ngOnInit(): void {

    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

    this.categoryService.getCategories()
    .subscribe(
      data => this.categories = data
    );

    this.incomeForm = this.formBuilder.group({
      name: [``],
      category: [``],
      amount: [``],
      date:[``],
      note:[``]
    });

    if(id !== null && !isNaN(id)){
      this.service.getIncomeById(id)
      .subscribe(data => {
        this.income = data
        this.incomeForm.controls[`name`].setValue(this.income.name)
        this.incomeForm.controls[`category`].setValue(this.categories.find(x => x.id === this.income.categoryId), {onlySelf:true})
        this.incomeForm.controls[`amount`].setValue(this.income.amount)
        this.incomeForm.controls[`date`].setValue(this.income.date)
        this.incomeForm.controls[`note`].setValue(this.income.note)
      });
    }
  }

  onSubmit(): void{
    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

    if(id !== null && !isNaN(id)) {
      this.income.id = id
      this.income.name = this.incomeForm.value.name
      this.income.categoryId = this.incomeForm.value.group.id

      this.service.putIncome(this.income)
      .subscribe((resp) => {
        this.router.navigate(['incomes']);
      })
      return
    }

    this.service.postIncome(this.incomeForm.value)
    .subscribe((resp) => {
      this.router.navigate(['incomes']);
    })
  }
  
  onDelete():void{
    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

      if(id !== null && !isNaN(id)) {

        this.service.deleteIncome(id)
        .subscribe((resp) => {
          this.router.navigate(['incomes']);
        })
    }
  }

}
