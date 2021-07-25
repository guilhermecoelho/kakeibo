import { Component, OnInit } from '@angular/core';
import { EMPTY, Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Group } from '../models/group.model';
import { GroupServiceService } from './services/group-service.service';

@Component({
  selector: 'app-groups',
  templateUrl: './groups.component.html',
  styleUrls: ['./groups.component.css']
})

export class GroupsComponent implements OnInit {

  groups$: Observable<Group[]>;

  constructor(private service: GroupServiceService) { }

  ngOnInit(): void {
    this.groups$ = this.service.getGroups()
  }
}
