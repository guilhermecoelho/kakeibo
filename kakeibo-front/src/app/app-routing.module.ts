import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { GroupsInsertComponent } from './groups/groups-insert/groups-insert.component';
import { GroupsComponent } from './groups/groups.component';
import { LoginComponent } from './login/login.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'groups', component: GroupsComponent },
  { path: 'groups/insert', component: GroupsInsertComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
