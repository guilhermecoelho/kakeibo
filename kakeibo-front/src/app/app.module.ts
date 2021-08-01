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
import { LoginService } from './login/login.service';
import { CategoriesComponent } from './categories/categories.component';
import { CategoriesInsertComponent } from './categories/categories-insert/categories-insert.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatListModule } from '@angular/material/list';
import { MatTableModule } from '@angular/material/table';

@NgModule({
  declarations: [
    AppComponent,
    GroupsComponent,
    GroupsInsertComponent,
    LoginComponent,
    CategoriesComponent,
    CategoriesInsertComponent
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
    MatTableModule
  ],
  providers: [
    GroupServiceService,
    { provide: HTTP_INTERCEPTORS, useClass: HandleInterceptor, multi: true, deps:[LocalStorageService, LoginService]}
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
