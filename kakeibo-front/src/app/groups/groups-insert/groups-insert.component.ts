import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Group } from 'src/app/models/group.model';
import { GroupServiceService } from '../services/group-service.service';

@Component({
  selector: 'app-groups-insert',
  templateUrl: './groups-insert.component.html',
  styleUrls: ['./groups-insert.component.css']
})
export class GroupsInsertComponent implements OnInit {

  groupForm: FormGroup
  group: Group

  constructor(
    private formBuilder: FormBuilder, 
    private service: GroupServiceService, 
    private router: Router,
    private route: ActivatedRoute,
    ) {}
   
  ngOnInit(): void {

    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

    if(id !== null && !isNaN(id)){
      this.service.getGroupById(id)
      .subscribe(data => {
        this.group = data
        this.groupForm = this.formBuilder.group({
          name: [this.group.name]
        })
      })
    }
    
    this.groupForm = this.formBuilder.group({
      name: [``]
    })
  }

  onSubmit(): void{
    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);
    
    if(id !== null && !isNaN(id)) {
      this.group.name = this.groupForm.value.name
      this.group.id = id
      this.service.putGroup(this.group)
      .subscribe((resp) => {
        this.router.navigate(['groups']);
      })
      return
    }

    this.service.postGroups(this.groupForm.value)
    .subscribe((resp) => {
      this.router.navigate(['groups']);
    })
  }
  
  onDelete():void{
    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

      if(id !== null && !isNaN(id)) {

        this.service.deleteGroup(id)
        .subscribe((resp) => {
          this.router.navigate(['groups']);
        })
    }
  }
}
