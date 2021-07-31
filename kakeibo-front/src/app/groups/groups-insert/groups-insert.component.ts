import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { GroupServiceService } from '../services/group-service.service';

@Component({
  selector: 'app-groups-insert',
  templateUrl: './groups-insert.component.html',
  styleUrls: ['./groups-insert.component.css']
})
export class GroupsInsertComponent implements OnInit {

  groupForm: FormGroup

  constructor(formBuilder: FormBuilder, private service: GroupServiceService, private router: Router) {
    this.groupForm = formBuilder.group({
      name: ['']
    })
   }
   
  ngOnInit(): void {
  }

  onSubmit(): void{

    this.service.postGroups(this.groupForm.value)
    .subscribe((resp) => {
      this.router.navigate(['groups']);
    })
  }

}
