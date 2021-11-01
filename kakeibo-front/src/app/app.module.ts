import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { GroupsComponent } from './groups/groups.component';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { GroupsInsertComponent } from './groups/groups-insert/groups-insert.component';
import { GroupServiceService } from './groups/services/group-service.service';
import { ReactiveFormsModule } from '@angular/forms';
import { HandleInterceptor } from './handle.interceptor';
import { LocalStorageService } from './local-storage.service';
import { LoginComponent } from './login/login.component';
import { CategoriesComponent } from './categories/categories.component';
import { CategoriesInsertComponent } from './categories/categories-insert/categories-insert.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatListModule } from '@angular/material/list';
import { MatTableModule } from '@angular/material/table';
import { ExpensesComponent } from './expenses/expenses.component';
import { ExpensesInsertComponent } from './expenses/expenses-insert/expenses-insert.component';
import { IncomesComponent } from './incomes/incomes.component';
import { IncomeInsertComponent } from './incomes/income-insert/income-insert.component';
import { Router } from '@angular/router';
import { UsersComponent } from './users/users.component';
import { UsersInsertComponent } from './users/users-insert/users-insert.component';
import { MatSelectModule } from '@angular/material/select';
import { ReportsComponent } from './reports/reports.component';
import {MatToolbarModule} from '@angular/material/toolbar'


@NgModule({
  declarations: [
    AppComponent,
    GroupsComponent,
    GroupsInsertComponent,
    LoginComponent,
    CategoriesComponent,
    CategoriesInsertComponent,
    ExpensesComponent,
    ExpensesInsertComponent,
    IncomesComponent,
    IncomeInsertComponent,
    UsersComponent,
    UsersInsertComponent,
    ReportsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatListModule,
    MatTableModule,
    MatSelectModule,
    MatToolbarModule
  ],
  providers: [
    GroupServiceService,
    { provide: HTTP_INTERCEPTORS, useClass: HandleInterceptor, multi: true, deps:[LocalStorageService, Router]}
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
