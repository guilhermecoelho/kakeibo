import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { GroupServiceService } from 'src/app/groups/services/group-service.service';
import { Category } from 'src/app/models/category.model';
import { Group } from 'src/app/models/group.model';
import { CategoriesService } from '../categories.service';

@Component({
  selector: 'app-categories-insert',
  templateUrl: './categories-insert.component.html',
  styleUrls: ['./categories-insert.component.css']
})
export class CategoriesInsertComponent implements OnInit {

  categoryForm: FormGroup
  category: Category

  groups: Group[]

  constructor(
    private formBuilder: FormBuilder, 
    private service: CategoriesService, 
    private router: Router,
    private route: ActivatedRoute,
    private groupService: GroupServiceService
    ) {}
   
  ngOnInit(): void {

    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

    this.groupService.getGroups()
    .subscribe(
      data => this.groups = data
    );

    this.categoryForm = this.formBuilder.group({
      name: [``],
      group: [``]
    });

    if(id !== null && !isNaN(id)){
      this.service.getCategoryById(id)
      .subscribe(data => {
        this.category = data
        this.categoryForm.controls[`name`].setValue(this.category.name)
        this.categoryForm.controls[`group`].setValue(this.groups.find(x => x.id === this.category.groupId), {onlySelf:true})
      });
    }
  }

  onSubmit(): void{
    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

    if(id !== null && !isNaN(id)) {
      this.category.id = id
      this.category.name = this.categoryForm.value.name
      this.category.groupId = this.categoryForm.value.group.id

      this.service.putCategory(this.category)
      .subscribe((resp) => {
        this.router.navigate(['categories']);
      })
      return
    }

    this.service.postCategory(this.categoryForm.value)
    .subscribe((resp) => {
      this.router.navigate(['categories']);
    })
  }
  
  onDelete():void{
    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

      if(id !== null && !isNaN(id)) {

        this.service.deleteCategory(id)
        .subscribe((resp) => {
          this.router.navigate(['categories']);
        })
    }
  }

}
