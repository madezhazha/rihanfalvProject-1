import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { FeedbackComponent } from './components/feedback/feedback.component';
import { PaperComponent } from './components/paper/paper.component';
import { PaperwebComponent } from './components/paper/paperweb/paperweb.component';
import { SearchComponent } from './components/search/search.component';
import { CaseComponent } from './components/case/case.component';
import { PersonpageComponent } from './components/personpage/personpage.component';
import { ChatComponent } from './components/chat/chat.component';
import { HeadComponent } from './components/head/head.component';

//头部组件
import { LangingComponent } from './components/head/langing/langing.component';
import { LandComponent } from './components/head/langing/land/land.component';
import { RegisterComponent } from './components/head/langing/register/register.component';
import { ForgetpasswordComponent } from './components/head/langing/forgetpassword/forgetpassword.component';
import { WebheadComponent } from './components/head/webhead/webhead.component';


import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';

//

@NgModule({
  declarations: [
    AppComponent,
    FeedbackComponent,
    PaperComponent,
    PaperwebComponent,
    SearchComponent,
    CaseComponent,
    PersonpageComponent,
    ChatComponent,
    HeadComponent,

    //头部组件
    LangingComponent,
    LandComponent,
    RegisterComponent,
    ForgetpasswordComponent,
    WebheadComponent,


  ],
  imports: [
    BrowserModule,
    AppRoutingModule,

    ReactiveFormsModule,
    HttpClientModule,
    FormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
