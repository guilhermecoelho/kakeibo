import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { CategoriesService } from 'src/app/categories/categories.service';
import { Category } from 'src/app/models/category.model';
import { Expense } from 'src/app/models/expense.model';
import { ExpensesService } from '../expenses.service';

@Component({
  selector: 'app-expenses-insert',
  templateUrl: './expenses-insert.component.html',
  styleUrls: ['./expenses-insert.component.css']
})
export class ExpensesInsertComponent implements OnInit {

  expenseForm: FormGroup
  expense: Expense
  selectedCategory: any;

  categories: Category[]

  constructor(
    private formBuilder: FormBuilder, 
    private service: ExpensesService, 
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

    this.expenseForm = this.formBuilder.group({
      name: [``],
      category: [``],
      amount: [``],
      date:[``],
      note:[``]
    });

    if(id !== null && !isNaN(id)){
      this.service.getExpenseById(id)
      .subscribe(data => {
        this.expense = data
        this.expenseForm.controls[`name`].setValue(this.expense.name)
        this.expenseForm.controls[`category`].setValue(this.categories.find(x => x.id === this.expense.categoryId), {onlySelf:true})
        this.expenseForm.controls[`amount`].setValue(this.expense.amount)
        this.expenseForm.controls[`date`].setValue(this.expense.date)
        this.expenseForm.controls[`note`].setValue(this.expense.note)

        this.selectedCategory = this.categories.find(x => x.id === this.expense.categoryId)
      });
    }
  }

  onSubmit(): void{
    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);
    this.expenseForm.value.category = this.selectedCategory

    if(id !== null && !isNaN(id)) {
      
      this.expense = {
        id: id,
        name:this.expenseForm.value.name,
        date: this.expenseForm.value.date,
        amount: this.expenseForm.value.amount,
        categoryId: this.expenseForm.value.category.id,
        category:this.selectedCategory,
        note: this.expenseForm.value.note
      }

      this.service.putExpense(this.expense)
      .subscribe((resp) => {
        this.router.navigate(['expenses']);
      })
      return
    }

    this.service.postExpense(this.expenseForm.value)
    .subscribe((resp) => {
      this.router.navigate(['expenses']);
    })
  }
  
  onDelete():void{
    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

      if(id !== null && !isNaN(id)) {

        this.service.deleteExpense(id)
        .subscribe((resp) => {
          this.router.navigate(['expenses']);
        })
    }
  }

}
