import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Category } from '../models/category.model';
import { CategoriesService } from './categories.service';

@Component({
  selector: 'app-categories',
  templateUrl: './categories.component.html',
  styleUrls: ['./categories.component.css']
})
export class CategoriesComponent implements OnInit {
  @Input() columns = ['id', 'name'];

  dataSource: Category[];

  constructor(
    private service: CategoriesService,
    private router: Router
  ) { }

  ngOnInit(): void {

    this.service.getCategories()
    .subscribe(data => this.dataSource = data)
  }

  clickedRows(row: any){
    this.router.navigate(['/categories/insert', row.id]);
    
  }

  createNew(){
    this.router.navigate(['/categories/insert']);
  }

}
